package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/ww24/lirc-web-api/config"
	"github.com/ww24/lirc-web-api/lirc"
)

var (
	// -ldflags "-X main.version=$API_VERSION"
	version string

	outputAPIVersion bool
	apiPort          int
)

func wrapError(err error) error {
	return &response{
		code:    http.StatusInternalServerError,
		Status:  "ng",
		Message: err.Error(),
	}
}

func init() {
	flag.BoolVar(&outputAPIVersion, "v", false, "output version")
	flag.IntVar(&apiPort, "p", 3000, "set API port")
	flag.Parse()
}

func main() {
	if outputAPIVersion {
		fmt.Println(version)
		return
	}

	e := echo.New()

	e.Logger.SetLevel(log.INFO)
	e.Logger.Infof("API version: %s", version)
	e.Logger.Infof("Running mode: %s", config.Mode)

	e.Pre(middleware.RemoveTrailingSlash())

	e.GET("/status", func(c echo.Context) error {
		client, err := lirc.New()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &status{
				Status:  "ng",
				Message: err.Error(),
			})
		}
		defer client.Close()

		lircdVersion, err := client.Version()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &status{
				Status:  "ng",
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusOK, &status{
			Status:       "ok",
			Message:      "LIRC Web API works",
			Version:      version,
			LIRCDVersion: lircdVersion,
		})
	})

	// create api v1 group and set error handling middleware
	apiv1g := e.Group("/api/v1", func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			msg := "unknown"

			defer func() {
				cause := recover()
				if cause != nil {
					e.Logger.Errorf("Panic:%v", cause)
					if err, ok := cause.(error); ok && config.IsDev() {
						msg = err.Error()
					}
					c.JSON(http.StatusInternalServerError, &response{
						Status:  "ng",
						Message: msg,
					})
				}
			}()

			err := next(c)
			if err != nil {
				if res, ok := err.(*response); ok {
					if res.Status == "ok" {
						return c.JSON(res.code, res)
					}

					e.Logger.Errorf("InternalServerError:%s", res)
					if config.IsProd() {
						res.Message = msg
					}
					return c.JSON(http.StatusInternalServerError, res)
				}
			}
			return err
		}
	})
	apiv1(apiv1g)

	e.Static("/", "./frontend")

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(apiPort)))
}
