package main

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Handlers struct {
}

func (h *Handlers) getPeople(c echo.Context) error {
	params := PeopleGet{
		Query: "",
		Limit: 5,
		Page:  1,
	}
	if err := c.Bind(&params); err != nil {
		log.Debugln(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	log.Infof("Getting people with params: %+v", params)
	people, err := getPeople(params)
	if err != nil {
		log.Debugln(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, people)
}

func (h *Handlers) createPerson(c echo.Context) error {
	var err error
	var person PersonCreate

	if err = c.Bind(&person); err != nil {
		log.Debugln(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(&person); err != nil {
		log.Debugln(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err = createPerson(person); err != nil {
		log.Debugln(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	log.Infof("Created new person")
	return c.NoContent(http.StatusCreated)
}

func (h *Handlers) updatePerson(c echo.Context) error {
	var data PersonUpdate
	err := c.Bind(&data)
	if err != nil {
		log.Debugln(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err = updatePerson(data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	log.Infof("Updated person with id %d", data.ID)
	return c.NoContent(http.StatusOK)
}

func (h *Handlers) deletePerson(c echo.Context) error {
	var data PersonUpdate
	err := c.Bind(&data)
	if err != nil {
		log.Debugln(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	err = deletePerson(data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	log.Infof("Deleted person with id %d", data.ID)
	return c.NoContent(http.StatusOK)
}
