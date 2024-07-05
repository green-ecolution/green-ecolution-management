package http

import (
	"context"
	"errors"
	"fmt"

	"github.com/SmartCityFlensburg/green-space-management/config"
	"github.com/SmartCityFlensburg/green-space-management/internal/service"
	"github.com/gofiber/fiber/v2"
)

type HTTPError struct {
	Error  string `json:"error"`
	Code   int    `json:"code"`
	Path   string `json:"path"`
	Method string `json:"method"`
} //@Name HTTPError

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
		return c.JSON(HTTPError{
			e.Message,
			e.Code,
			c.Path(),
			c.Method(),
		})
	}
	return nil
}
