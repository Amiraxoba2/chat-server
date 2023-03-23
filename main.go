package main

import (
	"github.com/Amiraxoba2/chat-server/internal/data"
	"github.com/Amiraxoba2/chat-server/internal/route"
	"github.com/gin-gonic/gin"
)

func main() {
	data.InitDB()
	engine := gin.Default()
	engine.POST("/register", route.Register)
	engine.POST("/login", route.Login)
	engine.GET("/messages", route.LoadMessages)
	engine.Run(":8080")
}
