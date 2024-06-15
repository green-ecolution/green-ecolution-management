package http

import (
	"context"
	"fmt"

	"github.com/SmartCityFlensburg/green-space-management/config"
	"github.com/SmartCityFlensburg/green-space-management/internal/service"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	cfg      *config.Config
	services *service.Service
}

func NewServer(cfg *config.Config, services *service.Service) *Server {
	return &Server{
		cfg:      cfg,
		services: services,
	}
}

func (s *Server) Run(ctx context.Context) error {
	app := fiber.New()
	app.Mount("/", s.router())

	go func() {
		<-ctx.Done()
		fmt.Println("Shutting down HTTP Server")
		app.Shutdown()
	}()

	return app.Listen(fmt.Sprintf("%s", s.cfg.Url))
}
