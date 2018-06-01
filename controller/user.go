package controller

import (
	"github.com/gin-gonic/gin"
	db "hangmango-web-api/model"
	"hangmango-web-api/serializer"
	"net/http"
	"strings"
)

func UserSignIn(c *gin.Context) {
	var signUpBody struct {
		LoginName string `json:"login_name" binding:"required"`
		Password  string `json:"password" binding:"required"`
	}

	if err := c.BindJSON(&signUpBody); err != nil {
		if !strings.Contains(err.Error(), "validation") {
			panic(err)
		}
		ValidationErrorResponse(c)
		return
	}

	user, err := db.UserLogin(signUpBody.LoginName, signUpBody.Password)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "LoginFail",
		})
		return
	}
	c.JSON(http.StatusOK, serializer.SerializeBaseUsers(1, []*db.User{user}))
	return
}

func UserSignUp(c *gin.Context) {
	var signUpBody struct {
		LoginName string `json:"login_name" binding:"required"`
		Password  string `json:"password" binding:"required"`
	}

	if err := c.BindJSON(&signUpBody); err != nil {
		if !strings.Contains(err.Error(), "validation") {
			panic(err)
		}
		ValidationErrorResponse(c)
		return
	}
	user, err := db.CreateUser(signUpBody.LoginName, signUpBody.Password)
	if err != nil {
		if !strings.Contains(err.Error(), "Duplicate entry") {
			panic(err)
		}
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "UserAlreadyExist",
		})
		return
	}
	c.JSON(http.StatusOK, serializer.SerializeBaseUsers(1, []*db.User{user}))
	return
}
