package route

import (
	"fmt"
	"github.com/Amiraxoba2/chat-server/internal/data"
	"github.com/Amiraxoba2/chat-server/internal/model"
	"github.com/Amiraxoba2/chat-server/internal/safety"
	"github.com/gin-gonic/gin"
)

func LoadMessages(context *gin.Context) {
	chat := context.Query("chat")
	var messages []model.Message
	data.DB.Raw("SELECT * FROM messages WHERE chat_id=?", chat).Scan(&messages)
	response := ""
	for _, message := range messages {
		var author model.User
		data.DB.Raw("SELECT * FROM users WHERE id=?", message.Author).Scan(&author)
		response += fmt.Sprintf("%s > %s\n", author.Name, safety.Decrypt(message.Content, *safety.LoadKey()))
	}
	context.Writer.WriteString(response)
}
