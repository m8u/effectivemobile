package main

import (
	"gorm.io/gorm"
)

type Gender string

const (
	GENDER_MALE   Gender = "male"
	GENDER_FEMALE Gender = "female"
)

type Nationality string

const (
	NATIONALITY_RU Nationality = "RU"
	NATIONALITY_UA Nationality = "UA"
)

type Person struct {
	gorm.Model
	Name        string      `gorm:"type:varchar(32);not null" validate:"required"`
	Surname     string      `gorm:"type:varchar(32);not null" validate:"required"`
	Patronymic  string      `gorm:"type:varchar(32)"`
	Age         int8        `gorm:"not null"`
	Gender      Gender      `gorm:"type:varchar(32);not null"`
	Nationality Nationality `gorm:"type:varchar(2);not null"`
}

func getPeople() []Person {
	var people []Person
	db.Find(&people)
	return people
}

func createPerson(person Person) error {
	var err error
	person.Age, err = getAge(person.Name)
	if err != nil {
		return err
	}
	person.Gender, err = getGender(person.Name)
	if err != nil {
		return err
	}
	person.Nationality, err = getNationality(person.Name)
	if err != nil {
		return err
	}
	db.Create(&person)
	return nil
}
