package router

import (
	"github.com/gin-gonic/gin"
	db "hangmango-web-api/model"
	"net/http"
	"strings"
)

func InitUserRouters(userGroup *gin.RouterGroup) {
	userGroup.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"total_count": 0,
			"data":        []int{},
		})
	})

	userGroup.POST("/", func(c *gin.Context) {
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
		var res struct {
			TotalCount int        `json:"total_count"`
			Data       []*db.User `json:"data"`
		}
		user, _ := db.CreateUser(signUpBody.LoginName, signUpBody.Password)
		res.TotalCount = 1
		res.Data = []*db.User{user}
		c.JSON(http.StatusOK, res)
		return
	})
}
