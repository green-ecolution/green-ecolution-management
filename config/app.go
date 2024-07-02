package config

import (
	"net/url"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Host       string        `env:"HOST" envDefault:"localhost"`
	Port       int           `env:"PORT" envDefault:"27017"`
	User       string        `env:"USER"`
	Password   string        `env:"PASSWORD"`
	Name       string        `env:"NAME"`
	Timeout    time.Duration `env:"TIMEOUT" envDefault:"30"`
	Database   string        `env:"DATABASE" envDefault:"green-space-management"`
	Collection string        `env:"COLLECTION" envDefault:"sensors"`
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
	godotenv.Load()

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
