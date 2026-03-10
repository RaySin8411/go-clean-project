package application

import (
	"context"

	"go-clean-project/internal/config"
	"go-clean-project/internal/logger"
)

type Application struct {
	Config *config.Config
	Logger logger.Logger

	Service *Service
	UseCase *UseCase
}

func New(cfg *config.Config) (*Application, error) {
	ctx := context.Background()
	logger := logger.NewSLogger(ctx, cfg.Logger)

	app := &Application{
		Config: cfg,
		Logger: logger,
	}

	app.Service = NewService(app)
	app.UseCase = NewUseCase(app)
	return app, nil
}
