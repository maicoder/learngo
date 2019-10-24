package handler

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"imooc.com/ccmouse/learngo/learngin/ginupload/model"
	"imooc.com/ccmouse/learngo/learngin/ginupload/utils"
)

func UserRegister(context *gin.Context) {
	var user model.UserModel
	if err := context.ShouldBind(&user); err != nil {
		context.String(http.StatusBadRequest, "输入的数据不合法")
		log.Println("err ->", err.Error())
	}
	passwordAgain := context.PostForm("password-again")
	if passwordAgain != user.Password {
		context.String(http.StatusBadRequest, "密码校验无效，两次密码不一致")
		log.Panicln("密码校验无效，两次密码不一致")
	}
	id := user.Save()
	log.Println("id is ", id)
	context.Redirect(http.StatusMovedPermanently, "/")
}

func UserLogin(ctx *gin.Context) {
	var user model.UserModel
	if e := ctx.Bind(&user); e != nil {
		log.Panicln("login 绑定错误", e.Error())
	}

	u := user.QueryByEmail()
	if u.Password == user.Password {
		log.Println("登录成功", u.Email)
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"email": u.Email,
		})
	}
}

func UserProfile(ctx *gin.Context) {
	id := ctx.Query("id")
	var user model.UserModel
	i, err := strconv.Atoi(id)
	u, e := user.QueryById(i)
	if e != nil || err != nil {
		ctx.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
	}
	ctx.HTML(http.StatusOK, "user_profile.tmpl", gin.H{
		"user": u,
	})
}

func UpdateUserProfile(ctx *gin.Context) {
	var user model.UserModel
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.HTML(http.StatusOK, "error.tmpl", gin.H{
			"errpr": err,
		})
		log.Panicln("绑定发生错误", err.Error())
	}

	file, e := ctx.FormFile("avatar-file")
	if e != nil {
		ctx.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("文件上传错误", e.Error())
	}

	path := utils.RootPath()
	path = filepath.Join(path, "avatar")
	e = os.MkdirAll(path, os.ModePerm)
	if e != nil {
		ctx.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("无法创建文件夹", e.Error())
	}

	fileName := strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	e = ctx.SaveUploadedFile(file, path+fileName)
	if e != nil {
		ctx.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("无法保存文件", e.Error())
	}

	avatarUrl := "http://127.0.0.1:8080/avatar/" + fileName
	user.Avatar = sql.NullString{String: avatarUrl}
	e = user.Update(int(user.Id))
	if e != nil {
		ctx.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("数据无法更新", e.Error())
	}
	ctx.Redirect(http.StatusMovedPermanently, "/user/profile?id="+strconv.Itoa(int(user.Id)))
}
