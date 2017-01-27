package main

import (
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/ww24/lirc-web-api/lirc"
)

var (
	// ErrBadSignal because signal not found in lircd.conf
	ErrBadSignal = errors.New("signal not found")
)

func apiv1(g *echo.Group) {
	g.GET("", func(c echo.Context) (err error) {
		signals, err := fetchSignals("")
		if err != nil {
			return wrapError(err)
		}

		return &response{
			code:    http.StatusOK,
			Status:  "ok",
			Signals: signals,
		}
	})

	g.GET("/list", func(c echo.Context) (err error) {
		signals, err := fetchSignals("")
		if err != nil {
			return wrapError(err)
		}

		return &response{
			code:    http.StatusOK,
			Status:  "ok",
			Signals: signals,
		}
	})

	g.GET("/:remote", func(c echo.Context) (err error) {
		signals, err := fetchSignals(c.Param("remote"))
		if err != nil {
			return wrapError(err)
		}

		return &response{
			code:    http.StatusOK,
			Status:  "ok",
			Signals: signals,
		}
	})

	g.POST("/:remote/:name", func(c echo.Context) (err error) {
		sig := &signal{
			Remote: c.Param("remote"),
			Name:   c.Param("name"),
		}

		err = sendSignal(sig)
		if err != nil {
			switch err {
			case ErrBadSignal:
				return &response{
					code:    http.StatusBadRequest,
					Status:  "ng",
					Message: "invalid signal",
				}
			default:
				return wrapError(err)
			}
		}

		return &response{
			code:   http.StatusOK,
			Status: "ok",
		}
	})

	g.POST("", func(c echo.Context) (err error) {
		sig := new(signal)
		if err = c.Bind(sig); err != nil {
			return wrapError(err)
		}

		err = sendSignal(sig)
		if err != nil {
			switch err {
			case ErrBadSignal:
				return &response{
					code:    http.StatusBadRequest,
					Status:  "ng",
					Message: "invalid signal",
				}
			default:
				return wrapError(err)
			}
		}

		return &response{
			code:   http.StatusOK,
			Status: "ok",
		}
	})
}

func fetchSignals(remote string) (signals []signal, err error) {
	client, err := lirc.New()
	if err != nil {
		return
	}
	defer client.Close()

	remotes, err := client.List(remote)
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

func sendSignal(sig *signal) (err error) {
	client, err := lirc.New()
	if err != nil {
		return
	}
	defer client.Close()

	replies, err := client.List(sig.Remote, sig.Name)
	if err != nil {
		return
	}
	if len(replies) == 0 {
		return ErrBadSignal
	}

	err = client.SendOnce(sig.Remote, sig.Name)
	if err != nil {
		return
	}

	return
}
