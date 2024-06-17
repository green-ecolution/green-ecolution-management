package mqtt

import (
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type MqttService struct {
	isConnected bool
}

func NewMqttService() *MqttService {
	return &MqttService{}
}

func (s *MqttService) HandleTemperature(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Temperature: %s\n", msg.Payload())
}

func (s *MqttService) HandleHumidity(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Humidity: %s\n", msg.Payload())
}

func (s *MqttService) SetConnected(ready bool) {
	s.isConnected = ready
}

func (s *MqttService) Ready() bool {
	return s.isConnected
}
