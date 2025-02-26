package test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"imooc.com/ccmouse/learngo/learngin/gincookie/initRouter"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router *gin.Engine

func init()  {
	gin.SetMode(gin.TestMode)
	router = initRouter.SetupRouter()
}

func TestIndexHtml(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "hello gin get methods", "返回的HTML页面中应该包含 hello gin get methods")
}
