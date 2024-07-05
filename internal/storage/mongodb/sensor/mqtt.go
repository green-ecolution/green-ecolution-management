package sensor

import (
	"context"
	"log"

	"github.com/SmartCityFlensburg/green-space-management/internal/entities/sensor"
	"github.com/SmartCityFlensburg/green-space-management/internal/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SensorRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewSensorRepository(client *mongo.Client, collection *mongo.Collection) *SensorRepository {
	return &SensorRepository{client: client, collection: collection}
}

func (r *SensorRepository) Upsert(ctx context.Context, data sensor.MqttData) error {
	filter := bson.M{"end_device_ids.device_id": data.Data.EndDeviceIDs.DeviceID}
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

func (r *SensorRepository) Get(ctx context.Context, id string) (*sensor.MqttData, error) {
	filter := bson.M{"end_device_ids.device_id": id}
	var data sensor.MqttData
	err := r.collection.FindOne(ctx, filter).Decode(&data)
	if err != nil {
		return nil, storage.ErrMongoDataNotFound
	}

	return &data, nil
}

func (r *SensorRepository) GetFirst(ctx context.Context) (*sensor.MqttData, error) {
	var data sensor.MqttData
	if err := r.collection.FindOne(ctx, bson.D{}).Decode(&data); err != nil {
		return nil, storage.ErrMongoDataNotFound
	}

	return &data, nil
}

func (r *SensorRepository) GetAllByTreeID(ctx context.Context, treeID string) ([]sensor.MqttData, error) {
	filter := bson.M{"tree_id": treeID}
	var data []sensor.MqttData
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		log.Println(err)
		return nil, storage.ErrMongoDataNotFound
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var d sensor.MqttData
		if err := cursor.Decode(&d); err != nil {
			log.Println(err)
			return nil, storage.ErrMongoDataNotFound
		}
		data = append(data, d)
	}

	return data, nil
}

func (r *SensorRepository) GetLastByTreeID(ctx context.Context, treeID string) (*sensor.MqttData, error) {
	filter := bson.M{"tree_id": treeID}
	opts := options.FindOne().SetSort(bson.D{{Key: "time", Value: -1}})
	var data sensor.MqttData
	err := r.collection.FindOne(ctx, filter, opts).Decode(&data)
	if err != nil {
		return nil, storage.ErrMongoDataNotFound
	}

	return &data, nil
}
