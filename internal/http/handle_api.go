package http

import (
	"go-clean-project/internal/application"
	"go-clean-project/internal/controller"

	"github.com/gin-gonic/gin"
)

func registerAPIRoutes(r gin.IRouter, app *application.Application) {
	registerAPIUserRoutes(r.Group("user"), app)
}

func registerAPIUserRoutes(r gin.IRouter, app *application.Application) {
	uc := controller.NewUserController()
	r.POST("register", uc.Register(app.UseCase.User.Register))
}
