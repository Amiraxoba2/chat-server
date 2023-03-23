package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Author  uint
	Content string
	ChatId  uint
}
