package lirc

import (
	"bufio"
	"errors"
	"net"
	"sync"
)

const socket = "/var/run/lirc/lircd"

var (
	// ErrInvalidReplies because lircd replies invalid format data
	ErrInvalidReplies = errors.New("invalid replies")
)

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

// New construct Client instance
func New(path ...string) (client ClientAPI, err error) {
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
		reply:  make(chan Reply),
	}

	go client.read()

	return
}

// Version command
func (cli *Client) Version() (version string, err error) {
	reps, err := cli.send("VERSION", "")
	if err != nil {
		return
	}
	if len(reps) == 0 {
		err = ErrInvalidReplies
		return
	}
	if len(reps[0].Data) == 0 {
		err = ErrInvalidReplies
		return
	}
	version = reps[0].Data[0]
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
	}
	return
}

// SendOnce command
func (cli *Client) SendOnce(remote string, code ...string) (err error) {
	_, err = cli.send("SEND_ONCE", remote, code...)
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
	defer close(cli.reply)
	return cli.conn.Close()
}
