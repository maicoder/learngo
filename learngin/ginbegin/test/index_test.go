package test

import (
	"github.com/stretchr/testify/assert"
	"imooc.com/ccmouse/learngo/learngin/ginbegin/initRouter"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexGetRouter(t *testing.T)  {
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Hello gin", w.Body.String())
}