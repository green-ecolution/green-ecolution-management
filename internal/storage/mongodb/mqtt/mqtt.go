package mqtt

import (
	"context"
	"fmt"
	"log"

	"github.com/SmartCityFlensburg/green-space-management/config"
	"github.com/SmartCityFlensburg/green-space-management/internal/entities/sensor"
	"github.com/SmartCityFlensburg/green-space-management/internal/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MqttMongoRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoClient(ctx context.Context, cfg config.DatabaseConfig) (*mongo.Client, error) {
	mongoUri := fmt.Sprintf("mongodb://%s:%s@%s:%d", cfg.User, cfg.Password, cfg.Host, cfg.Port)

	clientOptions := options.Client().ApplyURI(mongoUri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, storage.ErrMongoCannotCreateClient
	}

	ctx, cancel := context.WithTimeout(ctx, cfg.Timeout)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		return nil, storage.ErrMongoCannotPingClient
	}

	log.Println("Connected to MongoDB!")
	return client, nil
}

func NewMqttMongoRepository(client *mongo.Client, collection *mongo.Collection) *MqttMongoRepository {
	return &MqttMongoRepository{client: client, collection: collection}
}

func (r *MqttMongoRepository) Upsert(ctx context.Context, data sensor.Data) error {
	filter := bson.M{"end_device_ids.device_id": data.EndDeviceIDs.DeviceID}
	update := bson.M{"$set": data}
	opts := options.Update().SetUpsert(true)
	result, err := r.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return storage.ErrMongoCannotUpsertData
	}

	if result.UpsertedCount > 0 {
		log.Printf("Inserted a new document with ID %v\n", result.UpsertedID)
	} else {
		log.Println("Updated an existing document")
	}

	return nil
}

func (r *MqttMongoRepository) Get(ctx context.Context, id string) (*sensor.Data, error) {
	filter := bson.M{"end_device_ids.device_id": id}
	var data sensor.Data
	err := r.collection.FindOne(ctx, filter).Decode(&data)
	if err != nil {
		return nil, storage.ErrMongoDataNotFound
	}

	return &data, nil
}

func (r *MqttMongoRepository) GetFirst(ctx context.Context) (*sensor.Data, error) {
	var data sensor.Data
	if err := r.collection.FindOne(ctx, bson.D{}).Decode(&data); err != nil {
		return nil, storage.ErrMongoDataNotFound
	}

	return &data, nil
}
