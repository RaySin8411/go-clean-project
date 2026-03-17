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
	// 將 app.UseCase.User.Register 這個已建立好的 UseCase 實例，
	// 作為參數傳遞給 uc.Register 方法。
	r.POST("register", uc.Register(app.UseCase.User.Register))
}
