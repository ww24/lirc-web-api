package main

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/ww24/lirc-web-api/lirc"
)

type signal struct {
	Remote string `json:"remote"`
	Name   string `json:"name"`
}

func apiv1(g *echo.Group) {
	g.GET("/", func(c echo.Context) (err error) {
		signals, err := fetchSignals()
		if err != nil {
			return wrapError(err)
		}

		return &response{
			code:    http.StatusOK,
			Status:  "ok",
			Signals: signals,
		}
	})

	g.POST("/", func(c echo.Context) (err error) {
		sig := new(signal)
		if err = c.Bind(sig); err != nil {
			return wrapError(err)
		}

		client, err := lirc.New()
		if err != nil {
			return wrapError(err)
		}
		defer client.Close()

		replies, err := client.List(sig.Remote, sig.Name)
		if err != nil {
			return wrapError(err)
		}
		if len(replies) == 0 {
			return &response{
				code:    http.StatusBadRequest,
				Status:  "ng",
				Message: "invalid signal",
			}
		}

		err = client.SendOnce(sig.Remote, sig.Name)
		if err != nil {
			return wrapError(err)
		}

		return &response{
			code:   http.StatusOK,
			Status: "ok",
		}
	})
}

func fetchSignals() (signals []signal, err error) {
	client, err := lirc.New()
	if err != nil {
		return
	}
	defer client.Close()

	remotes, err := client.List("")
	if err != nil {
		return
	}

	signals = make([]signal, 0, len(remotes)*2)
	for _, remote := range remotes {
		var replies []string
		replies, err = client.List(remote)
		if err != nil {
			return
		}

		for _, reply := range replies {
			name := strings.Split(reply, " ")
			if len(name) == 2 {
				signals = append(signals, signal{
					Remote: remote,
					Name:   name[1],
				})
			}
		}
	}

	return
}
