package routes

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.Group("/api", func(ctx *gin.Context) {})
}
