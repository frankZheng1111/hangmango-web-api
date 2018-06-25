package controller

import (
	"github.com/gin-gonic/gin"
	db "hangmango-web-api/model"
	"hangmango-web-api/serializer"
	"net/http"
	"strings"
)

type LoginInfo struct {
	LoginName string `json:"login_name" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

func bindUserLoginInfo(c *gin.Context) (loginInfo *LoginInfo, err error) {
	loginInfo = new(LoginInfo)
	if err = c.BindJSON(loginInfo); err != nil {
		if !strings.Contains(err.Error(), "validation") {
			panic(err)
		}
		ValidationErrorResponse(c)
	}
	return
}

func UserSignIn(c *gin.Context) {
	signInBody, err := bindUserLoginInfo(c)
	if err != nil {
		return
	}

	user, err := db.UserLogin(signInBody.LoginName, signInBody.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "LoginFail",
		})
		return
	}
	// SetSession(c, "userId", user.Id)
	token, expiredAt, err := GenerateLoginToken(user.Id)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "LoginFail",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id":    user.Id,
		"expired_at": expiredAt,
		"token":      token,
	})
	return
}

func UserSignUp(c *gin.Context) {
	signUpBody, err := bindUserLoginInfo(c)
	if err != nil {
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

func GetBestUsers(c *gin.Context) {
	count, users := db.GetBestUsers(ParsePaginateFromQuery(c))
	c.JSON(http.StatusOK, serializer.SerializeBaseUsers(count, users))
	return
}
