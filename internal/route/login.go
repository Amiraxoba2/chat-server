package route

import (
	"encoding/base64"
	"github.com/Amiraxoba2/chat-server/internal/data"
	"github.com/Amiraxoba2/chat-server/internal/model"
	"github.com/AureumApes/goutil"
	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")
	token := base64.StdEncoding.EncodeToString([]byte(goutil.Hash(username+password, "amiraxoba", "token")))
	if data.DB.First(&model.User{}, "token = ?", token).RowsAffected == 0 {
		context.Status(404)
		context.Writer.WriteString("An error occured\n")
	} else {
		context.Status(202)
		context.Writer.Write([]byte(token))
	}
}
