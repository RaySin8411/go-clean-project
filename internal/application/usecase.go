package application

import (
	"go-clean-project/internal/usecase/api/user/login"
	"go-clean-project/internal/usecase/api/user/register"
)

type UseCase struct {
	User *UserUseCase
}

func NewUseCase(app *Application) *UseCase {
	return &UseCase{
		User: NewUserUseCase(app),
	}
}

type UserUseCase struct {
	Register *register.UseCase
	Login    *login.UseCase
}

func NewUserUseCase(app *Application) *UserUseCase {
	return &UserUseCase{
		// Register: register.NewUseCase(app.Database, app.Service.Password, app.Service.Token),
		Login: login.NewUseCase(app.Database, app.Service.Password, app.Service.Token),
	}

}
