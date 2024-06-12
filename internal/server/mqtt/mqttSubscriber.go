package mqtt

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func handleTemperature(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Temperature: %s\n", msg.Payload())
}

func handleHumidity(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Humidity: %s\n", msg.Payload())
}

var connectHandler MQTT.OnConnectHandler = func(client MQTT.Client) {
	fmt.Println("Connected to MQTT Broker")
}

var connectLostHandler MQTT.ConnectionLostHandler = func(client MQTT.Client, err error) {
	fmt.Printf("Connection lost: %v\n", err)
}

func RunSubscriber(mqttBroker string) {
	fmt.Println("Brocker: " + mqttBroker)
	opts := MQTT.NewClientOptions().AddBroker(mqttBroker).SetClientID("smartphone")
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe("TEMPERATURE", 1, handleTemperature); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return
	}

	if token := client.Subscribe("HUMIDITY", 1, handleHumidity); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return
	}

	// Keep the client running indefinitely
	select {}
}
