package sensor

import (
	"context"
	"log"

	"github.com/SmartCityFlensburg/green-space-management/internal/storage"
	"github.com/SmartCityFlensburg/green-space-management/internal/storage/entities/sensor"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r *SensorRepository) Insert(ctx context.Context, data sensor.MqttEntity) (*sensor.MqttEntity, error) {
	if data.ID == primitive.NilObjectID {
		objID := primitive.NewObjectID()
		data.ID = objID
	}
	_, err := r.collection.InsertOne(ctx, data)
	if err != nil {
		return nil, storage.ErrMongoCannotUpsertData
	}

	return &data, nil
}

func (r *SensorRepository) Get(ctx context.Context, id string) (*sensor.MqttEntity, error) {
	filter := bson.M{"end_device_ids.device_id": id}
	var data sensor.MqttEntity
	err := r.collection.FindOne(ctx, filter).Decode(&data)
	if err != nil {
		return nil, storage.ErrMongoDataNotFound
	}

	return &data, nil
}

func (r *SensorRepository) GetFirst(ctx context.Context) (*sensor.MqttEntity, error) {
	var data sensor.MqttEntity
	if err := r.collection.FindOne(ctx, bson.D{}).Decode(&data); err != nil {
		return nil, storage.ErrMongoDataNotFound
	}

	return &data, nil
}

func (r *SensorRepository) GetAllByTreeID(ctx context.Context, treeID string) ([]*sensor.MqttEntity, error) {
	filter := bson.M{"tree_id": treeID}
	var data []*sensor.MqttEntity
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		log.Println(err)
		return nil, storage.ErrMongoDataNotFound
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var d *sensor.MqttEntity
		if err := cursor.Decode(&d); err != nil {
			log.Println(err)
			return nil, storage.ErrMongoDataNotFound
		}
		data = append(data, d)
	}

	return data, nil
}

func (r *SensorRepository) GetLastByTreeID(ctx context.Context, treeID string) (*sensor.MqttEntity, error) {
	filter := bson.M{"tree_id": treeID}
	opts := options.FindOne().SetSort(bson.D{{Key: "time", Value: -1}})
	var data sensor.MqttEntity
	err := r.collection.FindOne(ctx, filter, opts).Decode(&data)
	if err != nil {
		return nil, storage.ErrMongoDataNotFound
	}

	return &data, nil
}
