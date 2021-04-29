package main

import (
	"github.com/labstack/echo"
)

func ExampleEchoStatic() {
	r := echo.New()

	r.Static("/static", "assets")

	r.Start(":9000")
}
