package models

import "gorm.io/gorm"

type UserEntity struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string `gorm:"unique"`
	Password  string
}
