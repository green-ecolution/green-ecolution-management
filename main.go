//go:generate mockery
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/SmartCityFlensburg/green-space-management/config"
	"github.com/SmartCityFlensburg/green-space-management/internal/server/http"
	"github.com/SmartCityFlensburg/green-space-management/internal/server/mqtt"
	"github.com/SmartCityFlensburg/green-space-management/internal/service/domain"
	"github.com/SmartCityFlensburg/green-space-management/internal/storage/local"
)

var version = "develop"

func main() {
	cfg, err := config.GetAppConfig()
	if err != nil {
		log.Fatal(err)
	}

	go mqtt.RunSubscriber(cfg.MQTTBroker)

	repositories, err := local.NewRepository(cfg)
	if err != nil {
		log.Fatal(err)
	}

	services := domain.NewService(cfg, repositories)
	httpServer := http.NewServer(cfg, services)

	fmt.Printf("Version: %s\n", version)
	if cfg.Development {
		fmt.Println("Running in dev mode")
	}

	log.Fatal(httpServer.Run(context.Background()))
}
