package main

import (
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name        string      `gorm:"type:varchar(32);not null"`
	Surname     string      `gorm:"type:varchar(32);not null"`
	Patronymic  string      `gorm:"type:varchar(32)"`
	Age         int8        `gorm:"not null"`
	Gender      Gender      `gorm:"type:varchar(32);not null"`
	Nationality Nationality `gorm:"type:varchar(2);not null"`
}
