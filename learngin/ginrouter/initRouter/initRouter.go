package initRouter

import (
	"github.com/gin-gonic/gin"
	"imooc.com/ccmouse/learngo/learngin/ginrouter/handler"
	"net/http"
	"strings"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	index := router.Group("/")
	{
		//index.GET("", retHelloGinAndMethod)

		//index.POST("", retHelloGinAndMethod)

		//index.PUT("", retHelloGinAndMethod)

		//index.DELETE("", retHelloGinAndMethod)

		//index.PATCH("", retHelloGinAndMethod)

		//index.HEAD("", retHelloGinAndMethod)

		//index.OPTIONS("", retHelloGinAndMethod)

		index.Any("", retHelloGinAndMethod)
	}

	userRouter := router.Group("/user")
	{
		userRouter.GET("/:name", handler.UserSave)
		userRouter.GET("", handler.UserSaveByQuery)
	}

	return router
}

func retHelloGinAndMethod(context *gin.Context)  {
	context.String(http.StatusOK, "hello gin " + strings.ToLower(context.Request.Method) + " method")
}