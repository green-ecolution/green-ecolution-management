package tree

import (
	"context"
	"log"

	"github.com/SmartCityFlensburg/green-space-management/internal/entities/sensor"
	"github.com/SmartCityFlensburg/green-space-management/internal/entities/tree"
	"github.com/SmartCityFlensburg/green-space-management/internal/server/http/handler"
	"github.com/SmartCityFlensburg/green-space-management/internal/service"
	"github.com/SmartCityFlensburg/green-space-management/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type TreeSensorDataResponse struct {
	Tree             *tree.Tree             `json:"tree,omitempty"`
	SensorPrediction *tree.SensorPrediction `json:"sensor_prediction,omitempty"`
	SensorData       []sensor.MqttData      `json:"sensor_data,omitempty"`
}

func GetAllTree(svc service.TreeService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		treeData, err := svc.GetAllTrees(c.Context())
		if err != nil {
			return handler.HandleError(err)
		}

		reponse := utils.Map(treeData, func(tree tree.Tree) TreeSensorDataResponse {
			var sensorData []sensor.MqttData
			if c.QueryBool("sensor_data") {
				data, err := GetSensorDataByTreeID(c.Context(), svc, tree.ID.Hex())
				if err != nil {
					log.Println(err)
				}
				sensorData = data
			}
			return TreeSensorDataResponse{
				Tree:       &tree,
				SensorData: sensorData,
			}
		})

		return c.JSON(reponse)
	}
}

func GetTreeByID(svc service.TreeService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var response TreeSensorDataResponse
		tree, err := svc.GetTreeByID(c.Context(), c.Params("id"))
		if err != nil {
			return handler.HandleError(err)
		}
		response.Tree = tree

		if c.QueryBool("sensor_data") {
			data, err := GetSensorDataByTreeID(c.Context(), svc, c.Params("id"))
			if err != nil {
				return handler.HandleError(err)
			}
			response.SensorData = data
		}

		return c.JSON(response)
	}
}

func GetTreePredictions(svc service.TreeService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var response TreeSensorDataResponse
		predict, err := svc.GetTreePrediction(c.Context(), c.Params("id"))
		if err != nil {
			return handler.HandleError(err)
		}

		response.SensorPrediction = predict

		if c.QueryBool("sensor_data") {
			data, err := GetSensorDataByTreeID(c.Context(), svc, c.Params("id"))
			if err != nil {
				return handler.HandleError(err)
			}

			response.SensorData = data
		}

		return c.JSON(response)
	}
}

func GetSensorDataByTreeID(ctx context.Context, svc service.TreeService, treeID string) ([]sensor.MqttData, error) {
	data, err := svc.GetSensorDataByTreeID(ctx, treeID)
	if err != nil {
		return nil, handler.HandleError(err)
	}

	return data, nil
}
