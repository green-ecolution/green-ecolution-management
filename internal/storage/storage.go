package storage

import (
	"context"
	"errors"

	"github.com/SmartCityFlensburg/green-space-management/internal/entities/info"
	"github.com/SmartCityFlensburg/green-space-management/internal/entities/sensor"
	"github.com/SmartCityFlensburg/green-space-management/internal/entities/tree"
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
	GetAppInfo(context.Context) (*info.App, error)
}

type SensorRepository interface {
	Upsert(ctx context.Context, data sensor.MqttData) error
	Get(ctx context.Context, id string) (*sensor.MqttData, error)
	GetFirst(ctx context.Context) (*sensor.MqttData, error)
  GetAllByTreeID(ctx context.Context, treeID string) ([]sensor.MqttData, error)
  GetLastByTreeID(ctx context.Context, treeID string) (*sensor.MqttData, error)
}

type TreeRepository interface {
  Insert(ctx context.Context, data tree.Tree) error
  Get(ctx context.Context, id string) (*tree.Tree, error)
  GetAll(ctx context.Context) ([]tree.Tree, error)
}

type Repository struct {
	Info   InfoRepository
	Sensor SensorRepository
  Tree   TreeRepository
}
