package initRouter

import (
	"github.com/gin-gonic/gin"
	"imooc.com/ccmouse/learngo/learngin/ginmysql/handler"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("/Users/mac/go/src/imooc.com/ccmouse/learngo/learngin/ginmysql/templates/*")
	//router.LoadHTMLGlob("templates/*")
	router.Static("/statics", "/Users/mac/go/src/imooc.com/ccmouse/learngo/learngin/ginmysql/statics")
	router.StaticFile("/favicon.ico", "/Users/mac/go/src/imooc.com/ccmouse/learngo/learngin/ginmysql/favicon.ico")
	index := router.Group("/")
	{
		index.Any("", handler.Index)
	}

	userRouter := router.Group("/user")
	{
		userRouter.POST("/register", handler.UserRegister)
		userRouter.POST("/login", handler.UserLogin)
	}

	return router
}

//func retHelloGinAndMethod(context *gin.Context)  {
//	context.String(http.StatusOK, "hello gin " + strings.ToLower(context.Request.Method) + " method")
//}