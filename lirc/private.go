package lirc

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func (cli *Client) send(cmd, remote string, code ...string) (replies []Reply, err error) {
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
		if err != nil {
			return
		}

		err = cli.writer.Flush()
		if err != nil {
			return
		}

		reply := <-cli.reply
		if !reply.Success {
			err = errors.New(strings.Join(reply.Data, ","))
			return
		}

		replies = append(replies, reply)
	}

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

			sizeText := sc.Text()
			size, err := strconv.ParseUint(sizeText, 10, 64)
			if err != nil {
				if cli.Verbose {
					log.Println("illegal error:", sizeText)
				}
			}

			replies := make([]string, size)
			for i := uint64(0); i < size && sc.Scan(); i++ {
				replies[i] = sc.Text()
			}

			reply.Data = replies
		case "END":
			cli.reply <- *reply
		default:
			if cli.Verbose {
				log.Println("illegal error:", line)
			}
		}
	}

	return
}
