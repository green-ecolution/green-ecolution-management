package mangodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type DBConfig struct {
	url     string
	Timeout time.Duration
}

func newDBCfg() *DBConfig {
	return &DBConfig{
		url:     "mongodb://localhost:27017",
		Timeout: 10 * time.Second,
	}
}

var (
	Database         = "green-space-management"
	SensorCollection = "sensor"
)

func NewMongoDBConnection() (*mongo.Client, error) {
	return connectMongoDB(newDBCfg())
}
func connectMongoDB(dbCfg *DBConfig) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(dbCfg.url)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), dbCfg.Timeout)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to MongoDB!")
	return client, nil
}
