package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Token string
	Name  string
}
