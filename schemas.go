package main

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Validator struct {
	validator *validator.Validate
}

func (cv *Validator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

type PeopleGet struct {
	Query string `query:"q"`
	Limit int    `query:"limit"`
	Page  int    `query:"page"`
}

type PersonCreate struct {
	Name       string `validate:"required"`
	Surname    string `validate:"required"`
	Patronymic string
}

type PersonUpdate struct {
	ID          uint `param:"id"`
	Name        string
	Surname     string
	Patronymic  string
	Age         int8
	Gender      Gender
	Nationality Nationality
}
