package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/ww24/lirc-web-api/config"
	"github.com/ww24/lirc-web-api/lirc"
)

var (
	// -ldflags "-X main.version=$API_VERSION"
	version string

	isAPIVersion bool
	apiPort      int
)

type response struct {
	code    int
	Status  string   `json:"status"`
	Message string   `json:"message,omitempty"`
	List    []string `json:"list,omitempty"`
}

func (res *response) Error() string {
	return res.Message
}

func wrapError(err error) error {
	return &response{
		code:    http.StatusInternalServerError,
		Status:  "ng",
		Message: err.Error(),
	}
}

func init() {
	flag.BoolVar(&isAPIVersion, "v", false, "output version")
	flag.IntVar(&apiPort, "p", 3000, "set API port")
	flag.Parse()
}

func main() {
	if isAPIVersion {
		fmt.Println(version)
		return
	}

	e := echo.New()

	e.Logger.SetLevel(log.INFO)
	e.Logger.Infof("API version: %s", version)
	e.Logger.Infof("Running mode: %s", config.Mode)

	// error handling middleware
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
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

	e.GET("/", func(c echo.Context) (err error) {
		client, err := lirc.New()
		if err != nil {
			return wrapError(err)
		}

		replies, err := client.List("")
		if err != nil {
			return wrapError(err)
		}

		return &response{
			code:   http.StatusOK,
			Status: "ok",
			List:   replies,
		}
	})

	// TODO: implement
	e.POST("/", func(c echo.Context) (err error) {
		client, err := lirc.New()
		if err != nil {
			return wrapError(err)
		}

		err = client.SendOnce("aircon", "on")
		if err != nil {
			return wrapError(err)
		}

		return &response{
			code:   http.StatusOK,
			Status: "ok",
		}
	})

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(apiPort)))
}
