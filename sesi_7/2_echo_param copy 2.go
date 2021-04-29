package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func ExampleEchoParam() {
	r := echo.New()

	r.GET("/page1", func(ctx echo.Context) error {
		name := ctx.QueryParam("name")
		data := fmt.Sprintf("hello %s", name)
		return ctx.String(http.StatusOK, data)
	})

	r.GET("/page2/:name", func(ctx echo.Context) error {
		name := ctx.Param("name")
		data := fmt.Sprintf("hello %s", name)
		return ctx.String(http.StatusOK, data)
	})

	r.GET("/page3/:name/*", func(ctx echo.Context) error {
		name := ctx.Param("name")
		message := ctx.Param("*")
		data := fmt.Sprintf("hello %s, I have message for you: %s", name, message)
		return ctx.String(http.StatusOK, data)
	})

	r.POST("/page4", func(ctx echo.Context) error {
		name := ctx.FormValue("name")
		message := ctx.FormValue("message")
		data := fmt.Sprintf("hello %s, I have message for you: %s", name, message)
		return ctx.String(http.StatusOK, data)
	})

	r.Start(":9000")
}
