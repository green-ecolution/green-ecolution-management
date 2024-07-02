package config

import (
	"net/url"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Host     string        `env:"HOST" envDefault:"localhost"`
	Port     int           `env:"PORT" envDefault:"27017"`
	User     string        `env:"USER" envDefault:"root"`
	Password string        `env:"PASSWORD" envDefault:"example"`
	Name     string        `env:"NAME" envDefault:"green-space-management"`
	Timeout  time.Duration `env:"TIMEOUT" envDefault:"10s"`
}

type MQTTConfig struct {
	Broker   string `env:"BROKER" envDefault:"eu1.cloud.thethings.network:1883"`
	ClientID string `env:"CLIENT_ID"`
	Username string `env:"USERNAME"`
	Password string `env:"PASSWORD"`
	Topic    string `env:"TOPIC"`
}

type Config struct {
	Url         *url.URL       `env:"APP_URL,expand" envDefault:"localhost:$PORT"`
	Port        int            `env:"PORT" envDefault:"8000"`
	Development bool           `env:"DEVELOPMENT" envDefault:"false"`
	MQTT        MQTTConfig     `envPrefix:"MQTT_"`
	Database    DatabaseConfig `envPrefix:"DB_"`
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
