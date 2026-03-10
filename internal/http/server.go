package http

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"go-clean-project/internal/application"
	"go-clean-project/internal/http/middleware"

	"github.com/gin-gonic/gin"
)

var (
	gracefulShutdownPeriod             = 5 * time.Second
	defaultHTTPServerReadHeaderTimeout = 2 * time.Second
)

type Server struct {
	server *http.Server
	engine *gin.Engine
}

func NewServer(app *application.Application) (*Server, error) {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(middleware.NewRequestLogger(app.Logger).Execute())
	engine.Use(middleware.NewInjectLogger(app.Logger).Execute())

	registerRoutes(engine, app)

	server := &http.Server{
		Addr:        ":" + app.Config.Server.Port,
		Handler:     engine,
		ReadTimeout: defaultHTTPServerReadHeaderTimeout,
		IdleTimeout: 60 * time.Second,
	}

	return &Server{
		server: server,
		engine: engine,
	}, nil
}

func (s *Server) Start() {
	go func() {
		if err := s.server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Panic(err)
		}
	}()
}

func (s *Server) Shutdown() {
	log.Println("Server is shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), gracefulShutdownPeriod)
	defer cancel()

	err := s.server.Shutdown(ctx)

	if err != nil {
		log.Panic(err)
	}

	log.Println("Server exiting")
}