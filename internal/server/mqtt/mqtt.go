package mqtt

import (
	"context"
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
	opts := MQTT.NewClientOptions()
	opts.AddBroker(m.cfg.MQTT.Broker)
	opts.SetClientID(m.cfg.MQTT.ClientID)
	opts.SetUsername(m.cfg.MQTT.Username)
	opts.SetPassword(m.cfg.MQTT.Password)

	opts.OnConnect = func(client MQTT.Client) {
		log.Println("Connected to MQTT Broker")
		m.svc.MqttService.SetConnected(true)
	}
	opts.OnConnectionLost = func(client MQTT.Client, err error) {
		log.Printf("Connection lost to MQTT Broker: %v\n", err)
	}

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
		return
	}

	token := client.Subscribe(m.cfg.MQTT.Topic, 1, m.svc.MqttService.HandleMessage)
	go func(token MQTT.Token) {
		_ = token.Wait()
		if token.Error() != nil {
			log.Println(token.Error())
		}
	}(token)

	<-ctx.Done()
	log.Println("Shutting down MQTT Subscriber")
}
