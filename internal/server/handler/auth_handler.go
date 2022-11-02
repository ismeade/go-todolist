package handler

import (
	"go-todolist/internal/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func AuthHandler(c *gin.Context) {
	// 用户发送用户名和密码过来
	var login Login
	err := c.ShouldBind(&login)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "无效的参数",
		})
		return
	}
	// 校验用户名和密码是否正确
	if login.Username == "admin" && login.Password == "123" {
		// 生成Token
		tokenString, _ := token.GenToken(login.Username)
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": gin.H{"token": tokenString},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 2002,
		"msg":  "鉴权失败",
	})
	return
}
