package main

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Config struct {
	Port        int  `env:"PORT" envDefault:"8000"`
	Development bool `env:"DEVELOPMENT" envDefault:"false"`
}

var version = "develop"

func main() {
	godotenv.Load()

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(cfg)
	})

	fmt.Printf("Version: %s\n", version)

	if cfg.Development {
		fmt.Println("Running in dev mode")
	}

	log.Fatal(app.Listen(fmt.Sprintf(":%d", cfg.Port)))
}
