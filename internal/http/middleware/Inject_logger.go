package middleware

import (
	"go-clean-project/internal/logger"

	"github.com/gin-gonic/gin"
)

type InjectLogger struct {
	logger logger.Logger
}

func NewInjectLogger(logger logger.Logger) *InjectLogger {
	return &InjectLogger{
		logger: logger,
	}
}

func (i *InjectLogger) Execute() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("logger", i.logger.With(c))
		c.Next()
	}
}
