package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/ww24/lirc-web-api/config"
	"github.com/ww24/lirc-web-api/lirc"
)

var (
	// -ldflags "-X main.version=$API_VERSION"
	version string
)

type response struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	List    []string `json:"list,omitempty"`
}

func main() {
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
					if err, ok := cause.(error); ok && config.IsDev() {
						msg = err.Error()
					}
					e.Logger.Errorf("Panic:%v", cause)
					c.JSON(http.StatusInternalServerError, &response{
						Status:  "ng",
						Message: msg,
					})
				}
			}()

			err := next(c)
			if err != nil {
				e.Logger.Errorf("InternalServerError:%s", err)
				if config.IsDev() {
					msg = err.Error()
				}
				return c.JSON(http.StatusInternalServerError, &response{
					Status:  "ng",
					Message: msg,
				})
			}
			return err
		}
	})

	e.GET("/", func(c echo.Context) (err error) {
		client, err := lirc.New()
		if err != nil {
			return
		}

		replies, err := client.List("")
		if err != nil {
			return
		}

		return c.JSON(http.StatusOK, &response{
			Status: "ok",
			List:   replies,
		})
	})

	// TODO: implement
	e.POST("/", func(c echo.Context) (err error) {
		client, err := lirc.New()
		if err != nil {
			return
		}

		err = client.SendOnce("aircon", "on")
		if err != nil {
			return
		}

		return c.JSON(http.StatusOK, &response{
			Status: "ok",
		})
	})

	port := strings.TrimSpace(os.Getenv("LIRC_WEB_API_PORT"))
	if port == "" {
		// default port
		port = "3000"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
