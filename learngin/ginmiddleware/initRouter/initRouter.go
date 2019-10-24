package initRouter

import (
	"imooc.com/ccmouse/learngo/learngin/ginmiddleware/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"imooc.com/ccmouse/learngo/learngin/ginmiddleware/handler"
	"imooc.com/ccmouse/learngo/learngin/ginmiddleware/utils"
)

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(middleware.Logger(), gin.Recovery())

	router.LoadHTMLGlob("/Users/mac/go/src/imooc.com/ccmouse/learngo/learngin/ginmiddle/templates/*")
	//router.LoadHTMLGlob("templates/*")
	router.StaticFile("/favicon.ico", "/Users/mac/go/src/imooc.com/ccmouse/learngo/learngin/ginmiddleware/favicon.ico")
	router.Static("/statics", "/Users/mac/go/src/imooc.com/ccmouse/learngo/learngin/ginmiddleware/statics")
	router.StaticFS("/avatar", http.Dir(utils.RootPath()+"avatar/"))
	index := router.Group("/")
	{
		index.Any("", handler.Index)
	}

	userRouter := router.Group("/user")
	{
		userRouter.POST("/register", handler.UserRegister)
		userRouter.POST("/login", handler.UserLogin)
		userRouter.GET("/profile/", middleware.Auth(), handler.UserProfile)
		userRouter.POST("/update", middleware.Auth(), handler.UpdateUserProfile)
	}

	return router
}
