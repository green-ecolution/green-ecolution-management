package config

import (
	"net/url"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Url         *url.URL `env:"APP_URL,expand" envDefault:"localhost:$PORT"`
	Port        int      `env:"PORT" envDefault:"8000"`
	Development bool     `env:"DEVELOPMENT" envDefault:"false"`
	MQTTBroker  string   `env:"MQTT_BROKER" envDefault:"tcp://mqtt.eclipseprojects.io:1883"`
}

func GetAppConfig() (*Config, error) {
  if err := godotenv.Load(); err != nil {
    return nil, err
  }

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
