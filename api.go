package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ww24/lirc-web-api/service"
)

func apiv1(g *echo.Group) {
	g.GET("", func(c echo.Context) (err error) {
		signals, err := service.FetchSignals("")
		if err != nil {
			return wrapError(err)
		}

		return &response{
			code:    http.StatusOK,
			Status:  "ok",
			Signals: signals,
		}
	})

	g.GET("/signals", func(c echo.Context) (err error) {
		signals, err := service.FetchSignals("")
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
		signals, err := service.FetchSignals(c.Param("remote"))
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
		sig := &service.Signal{
			Remote: c.Param("remote"),
			Name:   c.Param("name"),
		}

		err = service.SendSignal(&service.SendSignalParam{Signal: sig})
		if err != nil {
			switch err {
			case service.ErrBadSignal:
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
		sendParam := new(service.SendSignalParam)
		if err = c.Bind(sendParam); err != nil {
			return wrapError(err)
		}

		err = service.SendSignal(sendParam)
		if err != nil {
			switch err {
			case service.ErrBadSignal:
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
