package lirc

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
)

const socket = "/var/run/lirc/lircd"

// Client for lircd
type Client struct {
	sync.Mutex

	conn    net.Conn
	writer  *bufio.Writer
	reply   chan Reply
	Verbose bool
}

// Reply from lircd
type Reply struct {
	Command string
	Success bool
	Data    []string
}

// NewClient construct Client instance
func NewClient(path ...string) (client *Client, err error) {
	socketPath := socket
	if len(path) > 0 {
		socketPath = path[0]
	}

	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		return
	}

	client = &Client{
		conn:   conn,
		writer: bufio.NewWriter(conn),
		reply:  make(chan Reply, 1),
	}

	go client.read()

	return
}

func (cli *Client) send(cmd string, remote string, code ...string) (replies []Reply, err error) {
	cli.Lock()
	defer cli.Unlock()

	if remote != "" {
		cmd = fmt.Sprintf("%s %s", cmd, remote)
	}

	if len(code) == 0 {
		code = append(code, "")
	}

	replies = make([]Reply, 0, len(code))

	for _, c := range code {
		cmdStr := fmt.Sprintf("%s %s\n", cmd, c)
		_, err = cli.writer.WriteString(cmdStr)
		cli.writer.Flush()
		if err != nil {
			return
		}
		replies = append(replies, <-cli.reply)
	}

	fmt.Println(replies)

	return
}

func (cli *Client) read() {
	sc := bufio.NewScanner(cli.conn)

	var reply *Reply
	for sc.Scan() {
		line := sc.Text()

		if cli.Verbose {
			log.Println("DEBUG:", line)
		}

		switch line {
		case "BEGIN":
			reply = &Reply{}

			if !sc.Scan() {
				continue
			}

			reply.Command = sc.Text()
			if cli.Verbose {
				log.Println("BEGIN:", reply.Command)
			}
		case "SUCCESS":
			reply.Success = true
		case "ERROR":
			reply.Success = false
		case "DATA":
			if !sc.Scan() {
				continue
			}

			var size uint64
			size, err := strconv.ParseUint(sc.Text(), 10, 64)
			replies := make([]string, size)
			if err != nil {
				if cli.Verbose {
					log.Println("illegal error")
				}
			}

			for i := uint64(0); i < size && sc.Scan(); i++ {
				replies[i] = sc.Text()
			}

			reply.Data = replies
		case "END":
			cli.reply <- *reply
		default:
			if cli.Verbose {
				log.Println("illegal error")
			}
		}
	}

	return
}

// List command
func (cli *Client) List(remote string, code ...string) (replies []string, err error) {
	reps, err := cli.send("LIST", remote, code...)
	if err != nil {
		return
	}

	replies = make([]string, 0, len(reps))
	for _, r := range reps {
		replies = append(replies, r.Data...)
		if !r.Success {
			err = errors.New("some errors occurred")
		}
	}
	return
}

// SendOnce command
func (cli *Client) SendOnce(remote string, code ...string) (err error) {
	_, err = cli.send("SEND_ONCE", remote, code...)
	if err != nil {
		return
	}
	return
}

// SendStart command
func (cli *Client) SendStart(remote string, code string) (err error) {
	_, err = cli.send("SEND_START", remote, code)
	return
}

// SendStop command
func (cli *Client) SendStop(remote string, code string) (err error) {
	_, err = cli.send("SEND_STOP", remote, code)
	return
}

// Close connection
func (cli *Client) Close() error {
	cli.Lock()
	defer cli.Unlock()

	close(cli.reply)
	return cli.conn.Close()
}
