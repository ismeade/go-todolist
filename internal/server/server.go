package server

import (
	"go-todolist/internal/server/handler"

	"github.com/gin-gonic/gin"
)

func StartUp() {
	r := gin.Default()

	r.POST("/auth", handler.AuthHandler)
	r.GET("/home", handler.JWTAuthMiddleware(), handler.HomeHandler)

	r.Run()
}
