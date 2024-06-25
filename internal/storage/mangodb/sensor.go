package mangodb

import (
	"context"
	"fmt"
	"github.com/SmartCityFlensburg/green-space-management/internal/entities/sensor"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Upsert(client *mongo.Client, data sensor.Data) error {
	collection := client.Database(Database).Collection(SensorCollection)
	filter := bson.M{"end_device_ids.device_id": data.EndDeviceIDs.DeviceID}
	update := bson.M{
		"$set": data,
	}
	opts := options.Update().SetUpsert(true)
	result, err := collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		return err
	}
	if result.UpsertedCount > 0 {
		fmt.Printf("Inserted a new document with ID %v\n", result.UpsertedID)
	} else {
		fmt.Println("Updated an existing document")
	}
	return nil
}

func Get(client *mongo.Client, deviceID string) (sensor.Data, error) {
	collection := client.Database(Database).Collection(SensorCollection)
	var data sensor.Data
	filter := bson.M{"end_device_ids.device_id": deviceID}
	err := collection.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func GetFirst(client *mongo.Client) (sensor.Data, error) {
	collection := client.Database(Database).Collection(SensorCollection)
	var data sensor.Data
	filter := bson.M{}
	opts := options.FindOne().SetSort(bson.D{{"_id", 1}})
	err := collection.FindOne(context.TODO(), filter, opts).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}
