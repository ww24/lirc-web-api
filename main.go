package main

import (
	"net/http"

	lirc "github.com/inando/go-lirc"
	"github.com/labstack/echo"
)

type response struct {
	Status string `json:"status"`
	List   []string
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		client, err := lirc.New()
		if err != nil {
			return
		}

		err = client.Send("%s %s %s", "LIST", "''", "''")
		if err != nil {
			return
		}

		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/", func(c echo.Context) (err error) {
		client, err := lirc.New()
		if err != nil {
			return
		}

		err = client.Send("%s %s %s", "SEND_ONCE", "aircon", "on")
		if err != nil {
			return
		}

		return c.JSON(http.StatusOK, &response{"ok"})
	})

	e.Logger.Fatal(e.Start(":1323"))
}
