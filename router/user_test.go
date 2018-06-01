package router

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInitUserRouters(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/users/probe", nil)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()

	r := gin.Default()

	InitUserRouters(r.Group("/v1/users"))

	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)

	var respJson map[string]interface{}
	decoder := json.NewDecoder(w.Body)
	if err := decoder.Decode(&respJson); err != nil {
		panic(err)
	}
	assert.Equal(t, respJson["msg"], "success")
}
