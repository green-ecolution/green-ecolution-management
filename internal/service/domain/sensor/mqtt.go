package sensor

import (
	"context"
	"encoding/json"
	"log"

	"github.com/SmartCityFlensburg/green-space-management/internal/entities/sensor"
	"github.com/SmartCityFlensburg/green-space-management/internal/storage"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type MqttService struct {
	sensorRepo    storage.SensorRepository
	isConnected bool
}

func NewMqttService(sensorRepository storage.SensorRepository) *MqttService {
	return &MqttService{sensorRepo: sensorRepository}
}

func (s *MqttService) HandleMessage(client MQTT.Client, msg MQTT.Message) {
	jsonStr := string(msg.Payload())
	log.Printf("Received message: %s\n", jsonStr)

	var sensorData sensor.MqttData
	if err := json.Unmarshal([]byte(jsonStr), &sensorData); err != nil {
    log.Printf("Error unmarshalling sensor data: %v\n", err)
    return
	}
	log.Printf("Sensor data: %v\n", sensorData)

  if err := s.sensorRepo.Upsert(context.TODO(), sensorData); err != nil {
    log.Printf("Error upserting sensor data: %v\n", err)
    return
  }
}

func (s *MqttService) SetConnected(ready bool) {
	s.isConnected = ready
}

func (s *MqttService) Ready() bool {
	return s.isConnected
}
