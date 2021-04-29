package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type User struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}

type Person struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"gte=0,lte=80"`
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func ExampleEchoRequestPayload() {
	r := echo.New()
	r.Validator = &CustomValidator{validator: validator.New()}

	r.Any("/user", func(ctx echo.Context) error {
		u := new(User)
		if err := ctx.Bind(u); err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, u)
	})

	r.POST("/person", func(ctx echo.Context) error {
		p := new(Person)
		if err := ctx.Bind(p); err != nil {
			return err
		}

		if err := ctx.Validate(p); err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, p)
	})

	r.Start(":9000")
}

func CustomErrorHandler(err error, ctx echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	ctx.Logger().Error(report)
	ctx.JSON(report.Code, report)
}

func CustomReadableErrorHandler(err error, ctx echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				report.Message = fmt.Sprintf("%s is required", err.Field())
			case "email":
				report.Message = fmt.Sprintf("%s is not valid email", err.Field())
			case "gte":
				report.Message = fmt.Sprintf("%s value must be greater than %s", err.Field(), err.Param())
			case "lte":
				report.Message = fmt.Sprintf("%s value must be lower than %s", err.Field(), err.Param())
			}

			break
		}
	}

	ctx.Logger().Error(report)
	ctx.JSON(report.Code, report)
}

func CustomErrorPage(err error, ctx echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	errPage := fmt.Sprintf("%d.html", report.Code)
	if err := ctx.File(errPage); err != nil {
		ctx.HTML(report.Code, "Errorrr")
	}
}
