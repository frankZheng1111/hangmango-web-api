package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"hangmango-web-api/config"
	db "hangmango-web-api/model"
	"net/http"
	"strconv"
	"time"
)

var loginSecretKey string = "secret-hangmango-web-key" + config.Config.ENV

func ValidationErrorResponse(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"msg": "ParamsValidationError",
	})
	return
}

func GenerateLoginToken(userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"iat":    time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(loginSecretKey))

	return tokenString, err
}

func ParsePaginateFromQuery(c *gin.Context) *db.Paginate {
	paginate := new(db.Paginate)
	page := c.Query("page")
	pageSize := c.Query("page_size")
	paginate.Page, _ = strconv.Atoi(page)
	paginate.PageSize, _ = strconv.Atoi(pageSize)
	return paginate
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
