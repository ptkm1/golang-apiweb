package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ptkm1/golang-apiweb.git/use_cases/book"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/api/v1/")
	{
		books := main.Group("Books")
		{
			books.GET("/", book.ShowBooks)
		}
	}

	return router
}
