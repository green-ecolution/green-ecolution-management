package mongodb

import (
	"context"
	"fmt"
	"log"

	"github.com/SmartCityFlensburg/green-space-management/config"
	"github.com/SmartCityFlensburg/green-space-management/internal/storage"
	"github.com/SmartCityFlensburg/green-space-management/internal/storage/mongodb/sensor"
	"github.com/SmartCityFlensburg/green-space-management/internal/storage/mongodb/tree"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient(ctx context.Context, cfg config.DatabaseConfig) (*mongo.Client, error) {
	mongoUri := fmt.Sprintf("mongodb://%s:%s@%s:%d", cfg.User, cfg.Password, cfg.Host, cfg.Port)

	clientOptions := options.Client().ApplyURI(mongoUri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, storage.ErrMongoCannotCreateClient
	}

	log.Println("Trying to connect to MongoDB...")

	ctx, cancel := context.WithTimeout(ctx, cfg.Timeout)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		log.Println(err)
		return nil, storage.ErrMongoCannotPingClient
	}

	log.Println("Connected to MongoDB!")

	return client, nil
}

func NewRepository(cfg *config.Config) (*storage.Repository, error) {
	ctx := context.TODO()
	mongoClient, err := NewMongoClient(ctx, cfg.Database)
	if err != nil {
		return nil, err
	}

	sensorCollection := mongoClient.Database(cfg.Database.Name).Collection("sensors")
	mongoSensorRepo := sensor.NewSensorRepository(mongoClient, sensorCollection)

	treeCollection := mongoClient.Database(cfg.Database.Name).Collection("trees")
	mongoTreeRepo := tree.NewTreeRepository(mongoClient, treeCollection)

	return &storage.Repository{
		Sensor: mongoSensorRepo,
		Tree:   mongoTreeRepo,
	}, nil
}
