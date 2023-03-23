package data

import (
	"github.com/Amiraxoba2/chat-server/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB, _ = gorm.Open(sqlite.Open("resource/chat.sqlite"))

func InitDB() {
	DB.AutoMigrate(model.User{})
	DB.AutoMigrate(model.Chat{})
	DB.AutoMigrate(model.Message{})
	DB.Migrator().CreateConstraint(&model.Message{}, "ChatId")
}
