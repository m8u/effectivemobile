package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type Handlers struct {
}

func (h *Handlers) getPeople(c echo.Context) error {
	return c.JSON(http.StatusOK, getPeople())
}

func (h *Handlers) createPerson(c echo.Context) error {
	var err error
	var person Person

	if err = c.Bind(&person); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	if err = c.Validate(&person); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	if err = createPerson(person); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.NoContent(http.StatusCreated)
}

func (h *Handlers) updatePerson(c echo.Context) error {
	var person Person
	err := c.Bind(&person)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	//updatePerson(person)
	return c.NoContent(http.StatusOK)
}

func (h *Handlers) deletePerson(c echo.Context) error {
	var person Person
	err := c.Bind(&person)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	//deletePerson(person)
	return c.NoContent(http.StatusOK)
}
