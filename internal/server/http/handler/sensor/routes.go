package sensor

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterSensorRoutes() *fiber.App {
	app := fiber.New()

	app.Get("/", getSensorData())

	return app
}
