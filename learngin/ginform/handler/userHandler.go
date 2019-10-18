package handler

import (
	"github.com/gin-gonic/gin"
	"imooc.com/ccmouse/learngo/learngin/ginform/model"
	"log"
	"net/http"
)

func UserSave(context *gin.Context) {
	username := context.Param("name")
	context.String(http.StatusOK, "用户" + username + "已经保存")
}

func UserSaveByQuery(context *gin.Context) {
	username := context.Query("name")
	age := context.DefaultQuery("age", "20")
	context.String(http.StatusOK, "用户:" + username + ", 年龄:" + age + "已经保存")
}

func UserRegister(context *gin.Context)  {
	var user model.UserModel
	if err := context.ShouldBind(&user); err != nil {
		log.Println("err ->", err.Error())
		context.String(http.StatusBadRequest, "输入的数据不合法")
	} else {
		log.Println("email", user.Email, "password", user.Password, "password again", user.PasswordAgain)
		context.Redirect(http.StatusMovedPermanently, "/")
	}
}
