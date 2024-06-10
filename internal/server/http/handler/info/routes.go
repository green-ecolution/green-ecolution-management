package info

import (
	"github.com/SmartCityFlensburg/green-space-management/internal/service"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(svc service.InfoService) *fiber.App {
	app := fiber.New()

	app.Get("/", GetAppInfo(svc))

	return app
}
