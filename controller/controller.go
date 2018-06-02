package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ValidationErrorResponse(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"msg": "ParamsValidationError",
	})
	return
}

func SetSession(c *gin.Context, key string, value interface{}) {
	session := sessions.Default(c)
	session.Set(key, value)
	session.Save()
}

func GetSession(c *gin.Context, key string) interface{} {
	session := sessions.Default(c)
	return session.Get(key)
}
