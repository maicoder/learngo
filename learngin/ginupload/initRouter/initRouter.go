package initRouter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"imooc.com/ccmouse/learngo/learngin/ginupload/handler"
	"imooc.com/ccmouse/learngo/learngin/ginupload/utils"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("/Users/mac/go/src/imooc.com/ccmouse/learngo/learngin/ginupload/templates/*")
	//router.LoadHTMLGlob("templates/*")
	router.StaticFile("/favicon.ico", "/Users/mac/go/src/imooc.com/ccmouse/learngo/learngin/ginupload/favicon.ico")
	router.Static("/statics", "/Users/mac/go/src/imooc.com/ccmouse/learngo/learngin/ginupload/statics")
	router.StaticFS("/avatar", http.Dir(utils.RootPath()+"avatar/"))
	index := router.Group("/")
	{
		index.Any("", handler.Index)
	}

	userRouter := router.Group("/user")
	{
		userRouter.POST("/register", handler.UserRegister)
		userRouter.POST("/login", handler.UserLogin)
		userRouter.GET("/profile", handler.UserProfile)
		userRouter.POST("/update", handler.UpdateUserProfile)
	}

	return router
}
