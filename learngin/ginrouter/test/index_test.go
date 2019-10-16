package test

import (
	"github.com/stretchr/testify/assert"
	"imooc.com/ccmouse/learngo/learngin/ginrouter/initRouter"
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
	assert.Equal(t, "Hello gin get method", w.Body.String())
}

func TestIndexPostRouter(t *testing.T)  {
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "hello gin post method", w.Body.String())
}

func TestIndexPutRouter(t *testing.T)  {
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "hello gin put method", w.Body.String())
}

func TestIndexDeleteRouter(t *testing.T)  {
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "hrllo gin delete method", w.Body.String())
}

func TestIndexPatchRouter(t *testing.T)  {
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPatch, "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "hello gin patch method", w.Body.String())
}

func TestIndexHeadRouter(t *testing.T)  {
	router := initRouter.SetupRouter()
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodHead, "/", nil)
	router.ServeHTTP(recorder, request)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "hello gin head method", recorder.Body.String())
}

func TestIndexOptionsRouter(t *testing.T)  {
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodOptions, "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "hello gin options method", w.Body.String())
}