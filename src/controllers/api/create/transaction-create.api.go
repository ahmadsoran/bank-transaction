package apicretae

import "github.com/gin-gonic/gin"

func Create(ctx *gin.Context) {
	// Create a new transaction
	ctx.JSON(200, gin.H{
		"message": "Create transaction",
	})
}
