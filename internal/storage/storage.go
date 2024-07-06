package storage

import (
	"context"
	"errors"

	"github.com/SmartCityFlensburg/green-space-management/internal/storage/entities/info"
	"github.com/SmartCityFlensburg/green-space-management/internal/storage/entities/sensor"
	"github.com/SmartCityFlensburg/green-space-management/internal/storage/entities/tree"
)

var (
	ErrIpNotFound            = errors.New("local ip not found")
	ErrIFacesNotFound        = errors.New("cant get interfaces")
	ErrIFacesAddressNotFound = errors.New("cant get interfaces address")
	ErrHostnameNotFound      = errors.New("cant get hostname")

	ErrMongoCannotCreateClient = errors.New("cannot create mongo client")
	ErrMongoCannotPingClient   = errors.New("cannot ping mongo client")
	ErrMongoCannotUpsertData   = errors.New("cannot upsert data")
	ErrMongoDataNotFound       = errors.New("data not found")
)

type InfoRepository interface {
	GetAppInfo(context.Context) (*info.AppEntity, error)
}

type SensorRepository interface {
	Insert(ctx context.Context, data sensor.MqttEntity) (*sensor.MqttEntity, error)
	Get(ctx context.Context, id string) (*sensor.MqttEntity, error)
	GetFirst(ctx context.Context) (*sensor.MqttEntity, error)
	GetAllByTreeID(ctx context.Context, treeID string) ([]*sensor.MqttEntity, error)
	GetLastByTreeID(ctx context.Context, treeID string) (*sensor.MqttEntity, error)
}

type TreeRepository interface {
	Insert(ctx context.Context, data *tree.TreeEntity) error
	Get(ctx context.Context, id string) (*tree.TreeEntity, error)
	GetAll(ctx context.Context) ([]*tree.TreeEntity, error)
}

type Repository struct {
	Info   InfoRepository
	Sensor SensorRepository
	Tree   TreeRepository
}
