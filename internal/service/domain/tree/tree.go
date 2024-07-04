package tree

import (
	"context"
	"errors"

	"github.com/SmartCityFlensburg/green-space-management/internal/entities/sensor"
	"github.com/SmartCityFlensburg/green-space-management/internal/entities/tree"
	"github.com/SmartCityFlensburg/green-space-management/internal/service"
	"github.com/SmartCityFlensburg/green-space-management/internal/storage"
)

type TreeService struct {
	treeRepo   storage.TreeRepository
	sensorRepo storage.SensorRepository
}

func NewTreeService(treeRepo storage.TreeRepository, sensorRepo storage.SensorRepository) *TreeService {
	return &TreeService{
		treeRepo:   treeRepo,
		sensorRepo: sensorRepo,
	}
}

func (s *TreeService) GetTreeByID(ctx context.Context, id string) (*tree.Tree, error) {
	tree, err := s.treeRepo.Get(ctx, id)
	if err != nil {
		if errors.Is(err, storage.ErrMongoDataNotFound) {
			return nil, service.NewError(service.NotFound, err.Error())
		}

		return nil, service.NewError(service.InternalError, err.Error())
	}

	return tree, nil
}

func (s *TreeService) GetAllTrees(ctx context.Context) ([]tree.Tree, error) {
	trees, err := s.treeRepo.GetAll(ctx)
	if err != nil {
		if errors.Is(err, storage.ErrMongoDataNotFound) {
			return nil, service.NewError(service.NotFound, err.Error())
		}

		return nil, service.NewError(service.InternalError, err.Error())
	}

	return trees, nil
}

func (s *TreeService) InsertTree(ctx context.Context, data tree.Tree) error {
	err := s.treeRepo.Insert(ctx, data)
	if err != nil {
		if errors.Is(err, storage.ErrMongoCannotUpsertData) {
			return service.NewError(service.InternalError, err.Error())
		}

		return service.NewError(service.InternalError, err.Error())
	}

	return nil
}

func (s *TreeService) Ready() bool {
	return s.treeRepo != nil
}

func (s *TreeService) GetSensorDataByTreeID(ctx context.Context, treeID string) ([]sensor.MqttData, error) {
	return s.sensorRepo.GetAllByTreeID(ctx, treeID)
}

func (s *TreeService) GetTreePrediction(ctx context.Context, id string) (*tree.SensorPrediction, error) {
	treeData, err := s.treeRepo.Get(ctx, id)
	if err != nil {
		if errors.Is(err, storage.ErrMongoDataNotFound) {
			return nil, service.NewError(service.NotFound, err.Error())
		}
		return nil, service.NewError(service.InternalError, err.Error())
	}

	lastSensorData, err := s.sensorRepo.GetLastByTreeID(ctx, id)
	if err != nil {
		if errors.Is(err, storage.ErrMongoDataNotFound) {
			return nil, service.NewError(service.NotFound, err.Error())
		}

		return nil, service.NewError(service.InternalError, err.Error())
	}

	humidity := lastSensorData.Data.UplinkMessage.DecodedPayload.Humidity

	return &tree.SensorPrediction{
		SensorID: lastSensorData.Data.EndDeviceIDs.DeviceID,
		Humidity: humidity,
		Health:   getHealth(humidity),
		Tree:     treeData,
	}, nil
}

func getHealth(humidity int) tree.PredictedHealth {
	if humidity < 40 {
		return tree.HealthBad
	} else if humidity < 70 {
		return tree.HealthModerate
	}

	return tree.HealthGood
}
