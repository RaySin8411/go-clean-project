package http

import (
	"go-clean-project/internal/application"

	"github.com/gin-gonic/gin"
)

func registerAPIRoutes(r gin.IRouter, _ *application.Application) {
	// Register your API routes here
	r.GET("hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Hello, World!"})
	})
}
