package handler

import (
	"fmt"
	"go-todolist/internal/server/repository"
	"go-todolist/internal/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignUp struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
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
	err2 := repository.VerifyCredential(login.Username, login.Password)
	if err2 != nil {
		fmt.Println(err2)
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  err2,
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

func SignUpHandler(c *gin.Context) {
	// 用户发送用户名和密码过来
	var signUp SignUp
	err := c.ShouldBind(&signUp)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "无效的参数",
		})
		return
	}
	// 校验用户名和密码是否正确
	if signUp.Username == "admin" && signUp.Password == "123" {
		// 生成Token
		tokenString, _ := token.GenToken(signUp.Username)
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
