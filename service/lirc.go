package service

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/ww24/lirc-web-api/lirc"
)

var (
	// ErrBadSignal because signal not found in lircd.conf
	ErrBadSignal   = errors.New("signal not found")
	newLIRCService = lirc.New
)

// Signal .
type Signal struct {
	Remote string `json:"remote"`
	Name   string `json:"name"`
}

// SendSignalParam .
type SendSignalParam struct {
	*Signal
	Duration int64 `json:"duration,omitempty"`
}

// GetDuration convert int64[msec] to time.Duration
func (s *SendSignalParam) GetDuration() time.Duration {
	return time.Millisecond * time.Duration(s.Duration)
}

// FetchSignals .
func FetchSignals(remote string) (signals []Signal, err error) {
	client, err := newLIRCService()
	if err != nil {
		return
	}
	defer client.Close()

	remotes, err := client.List(remote)
	if err != nil {
		return
	}

	signals = make([]Signal, 0, len(remotes)*2)
	for _, remote := range remotes {
		var replies []string
		replies, err = client.List(remote)
		if err != nil {
			return
		}

		for _, reply := range replies {
			name := strings.Split(reply, " ")
			if len(name) == 2 {
				signals = append(signals, Signal{
					Remote: remote,
					Name:   name[1],
				})
			}
		}
	}

	return
}

// SendSignal .
func SendSignal(s *SendSignalParam) (err error) {
	client, err := newLIRCService()
	if err != nil {
		return
	}
	defer client.Close()

	replies, err := client.List(s.Remote, s.Name)
	if err != nil {
		return
	}
	if len(replies) == 0 {
		return ErrBadSignal
	}

	if s.Duration > 0 {
		log.Printf("send signal:%s:%s\tduration:%s\n", s.Remote, s.Name, s.GetDuration())
		err = client.SendStart(s.Remote, s.Name)
		if err != nil {
			return
		}
		defer client.SendStop(s.Remote, s.Name)
		time.Sleep(s.GetDuration())
	} else {
		log.Printf("send signal:%s:%s\n", s.Remote, s.Name)
		err = client.SendOnce(s.Remote, s.Name)
		if err != nil {
			return
		}
	}

	return
}
