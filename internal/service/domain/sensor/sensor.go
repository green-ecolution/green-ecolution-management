package sensor

import (
	"context"

	"github.com/SmartCityFlensburg/green-space-management/internal/entities/sensor"
	"github.com/SmartCityFlensburg/green-space-management/internal/storage"
)

type SensorService struct {
	sensorRepo    storage.SensorRepository
}

func NewSensorService(sensorRepository storage.SensorRepository) *SensorService {
	return &SensorService{sensorRepo: sensorRepository}
}

func (s *SensorService) GetHumidity(ctx context.Context) (int, error) {
  return 50, nil
}

func (s *SensorService) GetBattery(ctx context.Context) (float64, error) {
  return 20.0, nil
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

