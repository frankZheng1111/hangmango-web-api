package controller

import (
	"github.com/gin-gonic/gin"
	db "hangmango-web-api/model"
	"hangmango-web-api/serializer"
	"net/http"
	"strings"
)

func SignUpUser(c *gin.Context) {
	var signUpBody struct {
		LoginName string `json:"login_name" binding:"required"`
		Password  string `json:"password" binding:"required"`
	}

	if err := c.BindJSON(&signUpBody); err != nil {
		if !strings.Contains(err.Error(), "validation") {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "params validation error",
		})
		return
	}
	user, _ := db.CreateUser(signUpBody.LoginName, signUpBody.Password)
	c.JSON(http.StatusOK, serializer.SerializeBaseUsers(1, []*db.User{user}))
	return
}
