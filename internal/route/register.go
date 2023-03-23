package route

import (
	"encoding/base64"
	"github.com/Amiraxoba2/chat-server/internal/data"
	"github.com/Amiraxoba2/chat-server/internal/model"
	"github.com/AureumApes/goutil"
	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")
	token := base64.StdEncoding.EncodeToString([]byte(goutil.Hash(username+password, "amiraxoba", "token")))
	data.DB.Create(&model.User{
		Name:  username,
		Token: token,
	})
	context.Writer.WriteString("Successfully Created your Account!\n")
}
