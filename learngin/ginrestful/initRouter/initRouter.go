package initRouter

import (
	"imooc.com/ccmouse/learngo/learngin/ginrestful/handler/article"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	articleRouter := router.Group("")
	{
		articleRouter.GET("/article/:id", article.GetOne)
		articleRouter.GET("/articles", article.GetAll)
		articleRouter.POST("/article", article.Insert)
		articleRouter.DELETE("/article/:id", article.DeleteOne)
	}

	return router
}
