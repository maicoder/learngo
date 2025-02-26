package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"
	"imooc.com/ccmouse/learngo/learngin/ginrestful/initRouter"
	"imooc.com/ccmouse/learngo/learngin/ginrestful/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router *gin.Engine

func init()  {
	router = initRouter.SetupRouter()
}

func TestInsertArticle(t *testing.T)  {
	article := model.Article{
		Type:"go",
		Content:"hello gin",
	}
	marshal, _ := json.Marshal(article)
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/article", bytes.NewBufferString(string(marshal)))
	req.Header.Add("content-type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.NotEqual(t, "{id:-1}", w.Body.String())
}

func TestGetOneArticle(t *testing.T) {
	article := model.Article{
		Id:      1,
		Type:    "go",
		Content: "hello gin",
	}
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/article/1", nil)
	router.ServeHTTP(w, req)
	marshal, _ := json.Marshal(article)
	fmt.Println(w.Body.String())
	fmt.Println(string(marshal))
	assert.Equal(t, w.Code, 200)
	assert.Equal(t, w.Body.String(), string(marshal))
}

func TestGetAllArticle(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/articles", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusOK)
}

func TestDeleteOne(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/article/3", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusOK)
}