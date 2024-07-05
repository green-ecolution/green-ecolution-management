package tree

import (
	"github.com/SmartCityFlensburg/green-space-management/internal/service"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(svc service.TreeService) *fiber.App {
	app := fiber.New()

	app.Get("/", GetAllTree(svc))
	app.Get("/:id", GetTreeByID(svc))
	app.Get("/:id/prediction", GetTreePredictions(svc))

	return app
}
