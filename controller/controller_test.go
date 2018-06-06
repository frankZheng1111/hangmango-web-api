package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestValidationErrorResponse(t *testing.T) {
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("/test", ValidationErrorResponse)
	r.ServeHTTP(w, req)

	var respJson map[string]interface{}
	decoder := json.NewDecoder(w.Body)
	if err := decoder.Decode(&respJson); err != nil {
		panic(err)
	}

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "ParamsValidationError", respJson["msg"])
}

func TestCommonPanicHandle(t *testing.T) {
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("/test", CommonPanicHandle(func(c *gin.Context) {
		panic("testError")
	}))
	r.ServeHTTP(w, req)

	var respJson map[string]interface{}

	decoder := json.NewDecoder(w.Body)
	if err := decoder.Decode(&respJson); err != nil {
		panic(err)
	}

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, "SERVER_ERROR", respJson["msg"])
}

func TestValidAuthToken(t *testing.T) {
	var resultC *gin.Context
	r := gin.Default()
	r.GET("/test", ValidAuthToken, func(c *gin.Context) {
		resultC = c
	})

	// test valid token: missing token
	//
	reqWithoutToken, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqWithoutToken)
	var respJson map[string]interface{}
	decoder := json.NewDecoder(w.Body)
	if err := decoder.Decode(&respJson); err != nil {
		panic(err)
	}
	assert.Equal(t, 401, w.Code)
	assert.Equal(t, "NeedAuthToken", respJson["msg"])

	// test valid token: invalid token
	//
	reqWithInvalidToken, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		panic(err)
	}

	reqWithInvalidToken.Header.Add("hangmango-auth-token", "INVALIDTOKEN")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, reqWithInvalidToken)
	decoder = json.NewDecoder(w.Body)
	if err := decoder.Decode(&respJson); err != nil {
		panic(err)
	}
	assert.Equal(t, 401, w.Code)
	assert.Equal(t, "TokenAuthFail", respJson["msg"])
}
