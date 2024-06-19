package info

import (
	"context"
	"errors"

	"github.com/SmartCityFlensburg/green-space-management/internal/entities/info"
	"github.com/SmartCityFlensburg/green-space-management/internal/service"
	"github.com/SmartCityFlensburg/green-space-management/internal/storage"
)

type InfoService struct {
	infoRepository storage.InfoRepository
}

func NewInfoService(infoRepository storage.InfoRepository) *InfoService {
	return &InfoService{
		infoRepository: infoRepository,
	}
}

func (s *InfoService) GetAppInfo(ctx context.Context) (*info.App, error) {
	appInfo, err := s.infoRepository.GetAppInfo(ctx)
	if err != nil {
		if errors.Is(err, storage.ErrIpNotFound) {
			return nil, service.NewError(service.InternalError, err.Error())
		}
		if errors.Is(err, storage.ErrIFacesNotFound) {
			return nil, service.NewError(service.InternalError, err.Error())
		}
		if errors.Is(err, storage.ErrIFacesAddressNotFound) {
			return nil, service.NewError(service.InternalError, err.Error())
		}
		if errors.Is(err, storage.ErrHostnameNotFound) {
			return nil, service.NewError(service.InternalError, err.Error())
		}
	}

	return appInfo, nil
}

func (s *InfoService) Ready() bool {
	return s.infoRepository != nil
}
