package http

import (
	"go-clean-project/internal/application"

	"github.com/gin-gonic/gin"
)

func registerRoutes(engine *gin.Engine, app *application.Application) {
	registerRootRoutes(engine, app)
	registerAPIRoutes(engine.Group("api"), app)
}
