package storage

import (
	"context"
	"errors"

	"github.com/SmartCityFlensburg/green-space-management/internal/entities/info"
)

var (
	ErrIpNotFound            = errors.New("local ip not found")
	ErrIFacesNotFound        = errors.New("cant get interfaces")
	ErrIFacesAddressNotFound = errors.New("cant get interfaces address")
	ErrHostnameNotFound      = errors.New("cant get hostname")
)

type InfoRepository interface {
	GetAppInfo(context.Context) (*info.App, error)
}

type Repository struct {
	Info InfoRepository
}
