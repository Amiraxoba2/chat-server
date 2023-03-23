package route

import (
	"github.com/Amiraxoba2/chat-server/internal/data"
	"github.com/Amiraxoba2/chat-server/internal/model"
	"github.com/gin-gonic/gin"
	"strings"
)

func SendMessage(context *gin.Context) {
	tokenRaw := context.PostForm("token")
	token := strings.ReplaceAll(tokenRaw, " ", "+")
	println(token)
	var user model.User
	r := data.DB.Raw("SELECT * FROM users WHERE token=?", token).Scan(&user)
	if r.RowsAffected == 0 {
		context.Status(404)
		context.Writer.WriteString("An error occured: Token invalid")
		return
	}
	var chat model.Chat
	if r = data.DB.Raw("SELECT * FROM chats WHERE id=?", context.Query("chat")).Scan(&chat); r.RowsAffected == 0 {
		context.Status(404)
		context.Writer.WriteString("An error occured: Chat not found")
		return
	}
	message := context.PostForm("message")
	data.DB.Create(&model.Message{
		Content: message,
		Author:  user.ID,
		ChatId:  chat.ID,
	})
}
