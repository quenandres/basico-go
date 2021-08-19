package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// go get para manejo de paquete de terceros.
	// Instanciar echo
	e := echo.New()

	// Ruta
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World from GO")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
