//go:generate mockery
package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"sync"
	"syscall"

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

	fmt.Printf("Version: %s\n", version)
	if cfg.Development {
		fmt.Println("Running in dev mode")
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	repositories, err := local.NewRepository(cfg)
	if err != nil {
		log.Fatal(err)
	}

	services := domain.NewService(cfg, repositories)
	httpServer := http.NewServer(cfg, services)
	mqttServer := mqtt.NewMqtt(cfg, services)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		mqttServer.RunSubscriber(ctx)
	}()

	go func() {
		defer wg.Done()
    if err := httpServer.Run(ctx); err != nil {
      log.Fatal(err)
    }
	}()

	wg.Wait()
}
