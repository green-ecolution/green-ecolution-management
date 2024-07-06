package mapper

import (
	domain "github.com/SmartCityFlensburg/green-space-management/internal/entities/sensor"
	response "github.com/SmartCityFlensburg/green-space-management/internal/service/entities/sensor"
	repo "github.com/SmartCityFlensburg/green-space-management/internal/storage/entities/sensor"
)

// goverter:converter
// goverter:extend TimeToTime
type MqttMapper interface {
	// goverter:autoMap Data
	FromEntity(src *repo.MqttEntity) *domain.MqttPayload
	FromEntityList(src []*repo.MqttEntity) []*domain.MqttPayload

	ToResponse(src *domain.MqttPayload) *response.MqttPayloadResponse
  ToResponseList(src []*domain.MqttPayload) []*response.MqttPayloadResponse
	FromResponse(src *response.MqttPayloadResponse) *domain.MqttPayload
}
