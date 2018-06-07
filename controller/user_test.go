package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	db "hangmango-web-api/model"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserSignUp(t *testing.T) {
	var respJson map[string]interface{}
	w := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/test", UserSignUp)
	db.InitTestDB()

	// test signup fail
	//
	reqFail, err := http.NewRequest("POST", "/test", strings.NewReader("{}"))
	if err != nil {
		panic(err)
	}
	r.ServeHTTP(w, reqFail)
	decoder := json.NewDecoder(w.Body)
	if err := decoder.Decode(&respJson); err != nil {
		panic(err)
	}
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "ParamsValidationError", respJson["msg"])

	// test signup success
	//
	reqSuccess, err := http.NewRequest("POST", "/test", strings.NewReader(`{"login_name": "nameUser", "password": "pass"}`))
	if err != nil {
		panic(err)
	}
	w = httptest.NewRecorder()
	r.ServeHTTP(w, reqSuccess)
	decoder = json.NewDecoder(w.Body)
	if err := decoder.Decode(&respJson); err != nil {
		panic(err)
	}
	assert.Equal(t, 1, int(respJson["total_count"].(float64)))
	assert.Equal(t, 200, w.Code)
}

func TestUserSignIn(t *testing.T) {
	db.InitTestDB()
	reqSuccess, err := http.NewRequest("POST", "/test", strings.NewReader(`{"login_name": "nameUser", "password": "pass"}`))
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()

	r := gin.Default()
	r.POST("/test", UserSignIn)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, reqSuccess)

	var respJson map[string]interface{}
	decoder := json.NewDecoder(w.Body)
	if err := decoder.Decode(&respJson); err != nil {
		panic(err)
	}

	assert.Equal(t, 401, w.Code)
}
