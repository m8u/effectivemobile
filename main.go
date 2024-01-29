package main

import (
	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB

func init() {
	var err error

	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	err = godotenv.Load()
	if err != nil {
		log.Fatalln("failed to load .env file")
	}

	db, err = gorm.Open(postgres.Open("postgres://postgres:postgres@db:5432/effectivemobile"), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed to connect to database")
	}
	err = db.AutoMigrate(&Person{})
	if err != nil {
		log.Fatalln("failed to migrate Person")
	}
}

func main() {
	e := echo.New()

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			log.WithFields(log.Fields{
				"URI":    values.URI,
				"status": values.Status,
			}).Info("request")

			return nil
		},
	}))

	e.Validator = &Validator{validator: validator.New()}
	h := &Handlers{}
	e.GET("/people", h.getPeople)
	e.POST("/person", h.createPerson)
	e.PATCH("/person/:id", h.updatePerson)
	e.DELETE("/person/:id", h.deletePerson)

	log.Fatalln(e.Start(":8000"))
}
