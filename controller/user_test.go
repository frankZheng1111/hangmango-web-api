package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserSignUp(t *testing.T) {
	reqFail, err := http.NewRequest("POST", "/test", strings.NewReader("{}"))
	reqSuccess, err := http.NewRequest("POST", "/test", strings.NewReader(`{"login_name": "nameUser", "password": "pass"}`))
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()

	r := gin.Default()
	r.POST("/test", UserSignUp)
	r.ServeHTTP(w, reqFail)

	var respJson map[string]interface{}
	decoder := json.NewDecoder(w.Body)
	if err := decoder.Decode(&respJson); err != nil {
		panic(err)
	}

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "ParamsValidationError", respJson["msg"])

	w = httptest.NewRecorder()
	r.ServeHTTP(w, reqSuccess)
	decoder = json.NewDecoder(w.Body)
	if err := decoder.Decode(&respJson); err != nil {
		panic(err)
	}

	assert.Equal(t, 200, w.Code)
}
