package main

import (
	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var db *gorm.DB

type Validator struct {
	validator *validator.Validate
}

func (cv *Validator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env file")
	}

	db, err = gorm.Open(postgres.Open("postgres://postgres:postgres@db:5432/effectivemobile"), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed to connect to database")
	}
	err = db.AutoMigrate(&Person{})
	if err != nil {
		log.Fatalln("failed to migrate Person")
	}

	e := echo.New()
	e.Validator = &Validator{validator: validator.New()}
	h := &Handlers{}
	e.GET("/people", h.getPeople)
	e.POST("/person", h.createPerson)
	e.Logger.Fatal(e.Start(":8000"))
}
