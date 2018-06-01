package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ValidationErrorResponse(c *gin.Context) {
	c.JSON(http.StatusUnprocessableEntity, gin.H{
		"msg": "ParamsValidationError",
	})
	return
}
