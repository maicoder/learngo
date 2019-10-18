package initRouter

import (
	"github.com/gin-gonic/gin"
	"imooc.com/ccmouse/learngo/learngin/ginform/handler"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("/Users/mac/go/src/imooc.com/ccmouse/learngo/learngin/ginform/templates/*")
	//router.LoadHTMLGlob("templates/*")
	router.Static("/statics", "/Users/mac/go/src/imooc.com/ccmouse/learngo/learngin/ginform/statics")
	router.StaticFile("/favicon.ico", "/Users/mac/go/src/imooc.com/ccmouse/learngo/learngin/ginform/favicon.ico")
	index := router.Group("/")
	{
		//index.GET("", handler.Index)

		//index.POST("", retHelloGinAndMethod)

		//index.PUT("", retHelloGinAndMethod)

		//index.DELETE("", retHelloGinAndMethod)

		//index.PATCH("", retHelloGinAndMethod)

		//index.HEAD("", retHelloGinAndMethod)

		//index.OPTIONS("", retHelloGinAndMethod)

		index.Any("", handler.Index)
	}

	userRouter := router.Group("/user")
	{
		userRouter.GET("/:name", handler.UserSave)
		userRouter.GET("", handler.UserSaveByQuery)
		userRouter.POST("/register", handler.UserRegister)
	}

	return router
}

//func retHelloGinAndMethod(context *gin.Context)  {
//	context.String(http.StatusOK, "hello gin " + strings.ToLower(context.Request.Method) + " method")
//}