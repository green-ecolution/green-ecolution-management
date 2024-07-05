package sensor

import (
	"context"

	"github.com/SmartCityFlensburg/green-space-management/internal/entities/sensor"
	"github.com/SmartCityFlensburg/green-space-management/internal/storage"
)

type SensorService struct {
	sensorRepo storage.SensorRepository
}

func NewSensorService(sensorRepository storage.SensorRepository) *SensorService {
	return &SensorService{sensorRepo: sensorRepository}
}

func (s *SensorService) GetHumidityByTree(ctx context.Context, treeID string) (int, error) {
	data, err := s.sensorRepo.GetLastByTreeID(ctx, treeID)
	if err != nil {
		return 0, err
	}
	return data.Data.UplinkMessage.DecodedPayload.Humidity, nil
}

func (s *SensorService) GetBatteryByTree(ctx context.Context, treeID string) (float64, error) {
	data, err := s.sensorRepo.GetLastByTreeID(ctx, treeID)
	if err != nil {
		return 0, err
	}
	return data.Data.UplinkMessage.DecodedPayload.Battery, nil
}

func (s *SensorService) GetMqttDataByTreeID(ctx context.Context, treeID string) ([]sensor.MqttData, error) {
	return s.sensorRepo.GetAllByTreeID(ctx, treeID)
}

func (s *SensorService) GetMqttDataByTreeIDLast(ctx context.Context, treeID string) (*sensor.MqttData, error) {
	return s.sensorRepo.GetLastByTreeID(ctx, treeID)
}

func (s *SensorService) Ready() bool {
	return s.sensorRepo != nil
}
