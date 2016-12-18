package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/lirc-web-api/lirc"
)

type response struct {
	Status string `json:"status"`
	List   []string
}

func main() {
	e := echo.New()

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

	e.Logger.Fatal(e.Start(":1323"))
}
