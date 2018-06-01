package router

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInitRouters(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/probe", nil)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()

	r := gin.Default()

	InitRouters(r.Group("/v1"))

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var respJson map[string]interface{}
	decoder := json.NewDecoder(w.Body)
	if err := decoder.Decode(&respJson); err != nil {
		panic(err)
	}
	assert.Equal(t, "success", respJson["msg"])
}
