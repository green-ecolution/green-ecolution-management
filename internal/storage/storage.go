package storage

import (
	"context"
	"errors"

	"github.com/SmartCityFlensburg/green-space-management/internal/entities/info"
	"github.com/SmartCityFlensburg/green-space-management/internal/entities/sensor"
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

type MqttRepository interface {
	Upsert(ctx context.Context, data sensor.Data) error
	Get(ctx context.Context, id string) (*sensor.Data, error)
	GetFirst(ctx context.Context) (*sensor.Data, error)
}

type Repository struct {
	Info InfoRepository
	Mqtt MqttRepository
}
