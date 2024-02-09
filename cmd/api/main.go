package main

import (
	"net/http"

	conn "nuage/internal/db"

	"github.com/labstack/echo/v4"
)

func main() {
	conn.InitDB()
	//mustNot(err)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Yooo")
	})
	e.Logger.Fatal(e.Start(":3000"))
}

func mustNot(err error) {
	if err != nil {
		panic(err)
	}
}
