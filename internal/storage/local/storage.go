package local

import (
	"github.com/SmartCityFlensburg/green-space-management/config"
	"github.com/SmartCityFlensburg/green-space-management/internal/storage"
	"github.com/SmartCityFlensburg/green-space-management/internal/storage/local/info"
)

func NewRepository(cfg *config.Config) (*storage.Repository, error) {
	infoRepo, err := info.NewInfoRepository(cfg)
	if err != nil {
		return nil, err
	}

	return &storage.Repository{
		Info: infoRepo,
	}, nil
}
