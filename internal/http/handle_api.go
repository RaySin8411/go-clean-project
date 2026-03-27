package http

import (
	"go-clean-project/internal/application"
	"go-clean-project/internal/controller"
	"go-clean-project/internal/http/middleware"

	"github.com/gin-gonic/gin"
)

func registerAPIRoutes(r gin.IRouter, app *application.Application) {
	registerAPIUserRoutes(r.Group("user"), app)
}

func registerAPIUserRoutes(r gin.IRouter, app *application.Application) {
	uc := controller.NewUserController()
	authMiddleware := middleware.NewAuth(app.Service.Token).Execute()

	r.POST("register", uc.Register(app.UseCase.User.Register))
	r.POST("login", uc.Login(app.UseCase.User.Login))

	r.PUT("password", authMiddleware, uc.ChangePassword(app.UseCase.User.ChangePassword))
}
