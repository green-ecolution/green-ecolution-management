package http

import (
	"context"
	"errors"
	"fmt"

	"github.com/SmartCityFlensburg/green-space-management/config"
	"github.com/SmartCityFlensburg/green-space-management/internal/service"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	cfg      *config.Config
	services *service.Services
}

func NewServer(cfg *config.Config, services *service.Services) *Server {
	return &Server{
		cfg:      cfg,
		services: services,
	}
}

func (s *Server) Run(ctx context.Context) error {
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})
	app.Use(s.healthCheck())
	app.Mount("/", s.router())

	go func() {
		<-ctx.Done()
		fmt.Println("Shutting down HTTP Server")
		if err := app.Shutdown(); err != nil {
			fmt.Println("Error while shutting down HTTP Server:", err)
		}
	}()

	return app.Listen(s.cfg.Url.String())
}

func errorHandler(c *fiber.Ctx, err error) error {
  c.Status(fiber.StatusInternalServerError)
  var e *fiber.Error
  if errors.As(err, &e) {
    return c.JSON(fiber.Map{
      "error": e.Message,
      "code": e.Code,
      "path": c.Path(),
      "method": c.Method(),
    })
  }
  return nil
}
