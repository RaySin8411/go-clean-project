package http

import (
	"go-clean-project/internal/application"

	"github.com/gin-gonic/gin"
)

func registerRootRoutes(r gin.IRouter, _ *application.Application) {
	// Register your API routes here
	r.GET("health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})
}
