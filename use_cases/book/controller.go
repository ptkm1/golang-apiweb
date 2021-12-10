package book

import "github.com/gin-gonic/gin"

func ShowBooks(context *gin.Context) {
	context.JSON(200, gin.H{
		"value": "key",
	})
}
