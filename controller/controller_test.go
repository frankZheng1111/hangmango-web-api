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
