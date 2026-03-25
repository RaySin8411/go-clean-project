package http

import (
	"go-clean-project/internal/application"

	"github.com/gin-gonic/gin"
)

func registerRootRoutes(r gin.IRouter, app *application.Application) {
	// Register your API routes here
	r.GET("health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})

	auth := gin.BasicAuth(gin.Accounts{
		app.Config.APIDocAuth.UserName: app.Config.APIDocAuth.Password,
	})

	apiGroup := r.Group("", auth)
	staticFile(apiGroup, app.Embed.APIDoc, "swagger", "swagger.html", "text/html", 86400)
	staticFile(apiGroup, app.Embed.APIDoc, "swagger.json", "swagger.json", "application/json", 0)
}
