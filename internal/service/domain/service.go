package domain

import (
	"github.com/SmartCityFlensburg/green-space-management/config"
	"github.com/SmartCityFlensburg/green-space-management/internal/service"
	"github.com/SmartCityFlensburg/green-space-management/internal/service/domain/info"
	"github.com/SmartCityFlensburg/green-space-management/internal/service/domain/mqtt"
	"github.com/SmartCityFlensburg/green-space-management/internal/storage"
)

func NewService(cfg *config.Config, repositories *storage.Repository) *service.Services {
	return &service.Services{
		InfoService: info.NewInfoService(repositories.Info),
		MqttService: mqtt.NewMqttService(repositories.Mqtt),
	}
}
