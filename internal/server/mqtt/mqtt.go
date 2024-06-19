package mqtt

import (
	"context"
	"fmt"
	"log"

	"github.com/SmartCityFlensburg/green-space-management/config"
	"github.com/SmartCityFlensburg/green-space-management/internal/service"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type Mqtt struct {
	cfg *config.Config
	svc *service.Services
}

func NewMqtt(cfg *config.Config, services *service.Services) *Mqtt {
	return &Mqtt{
		cfg: cfg,
		svc: services,
	}
}

func (m *Mqtt) RunSubscriber(ctx context.Context) {
	fmt.Println("Brocker: " + m.cfg.MQTTBroker)
	opts := MQTT.NewClientOptions().AddBroker(m.cfg.MQTTBroker).SetClientID("smartphone")
	opts.OnConnect = func(client MQTT.Client) {
		fmt.Println("Connected to MQTT Broker")
		m.svc.MqttService.SetConnected(true)
	}
	opts.OnConnectionLost = func(client MQTT.Client, err error) {
		fmt.Printf("Connection lost: %v\n", err)
	}

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
		return
	}

	token := client.Subscribe("TEMPERATURE", 1, m.svc.MqttService.HandleTemperature)
	go func(token MQTT.Token) {
		_ = token.Wait()
		if token.Error() != nil {
			log.Println(token.Error())
		}
	}(token)

	token = client.Subscribe("HUMIDITY", 1, m.svc.MqttService.HandleHumidity)
	go func(token MQTT.Token) {
		_ = token.Wait()
		if token.Error() != nil {
			log.Println(token.Error())
		}
	}(token)

	<-ctx.Done()
	fmt.Println("Shutting down MQTT Subscriber")
}
