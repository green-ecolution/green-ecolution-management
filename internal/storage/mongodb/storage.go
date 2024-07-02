package mongodb

import (
	"context"

	"github.com/SmartCityFlensburg/green-space-management/config"
	"github.com/SmartCityFlensburg/green-space-management/internal/storage"
	"github.com/SmartCityFlensburg/green-space-management/internal/storage/mongodb/mqtt"
)

func NewRepository(cfg *config.Config) (*storage.Repository, error) {
	ctx := context.TODO()
	mongoClient, err := mqtt.NewMongoClient(ctx, cfg.Database)
	if err != nil {
		return nil, err
	}

	collection := mongoClient.Database(cfg.Database.Name).Collection(cfg.Database.Collection)
	mongoMqttRepo := mqtt.NewMqttMongoRepository(mongoClient, collection)

	return &storage.Repository{
		MqttMongo: mongoMqttRepo,
	}, nil
}
