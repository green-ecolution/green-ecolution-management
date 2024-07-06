package tree

import "github.com/SmartCityFlensburg/green-space-management/internal/service/entities/sensor"

type TreeSensorPredictionResponse struct {
	Tree             *TreeResponse             `json:"tree,omitempty"`
	SensorPrediction *SensorPredictionResponse `json:"sensor_prediction,omitempty"`
	SensorData       []*sensor.MqttPayloadResponse `json:"sensor_data,omitempty"`
} //@Name TreeSensorPrediction

type TreeSensorDataResponse struct {
	Tree       *TreeResponse             `json:"tree,omitempty"`
	SensorData []*sensor.MqttPayloadResponse `json:"sensor_data,omitempty"`
} //@Name TreeSensorData
