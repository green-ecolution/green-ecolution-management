package sensor

import (
	"fmt"
	"github.com/SmartCityFlensburg/green-space-management/internal/storage/mangodb"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func getSensorData() fiber.Handler {
	return func(c *fiber.Ctx) error {
		connection, err := mangodb.NewMongoDBConnection()
		if err != nil {

			log.Fatalf("Failed to connect to MongoDB: %v", err)
		}

		first, err := mangodb.GetFirst(connection)
		if err != nil {
			log.Fatalf("Failed to unmarshal JSON: %v", err)
		} else {
			fmt.Printf("Data updated...: \n")
		}
		return c.JSON(first)
	}
}
