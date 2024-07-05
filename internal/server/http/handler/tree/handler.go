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

type TreeSensorPredictionResponse struct {
	Tree             *tree.Tree             `json:"tree,omitempty"`
	SensorPrediction *tree.SensorPrediction `json:"sensor_prediction,omitempty"`
	SensorData       []sensor.MqttData      `json:"sensor_data,omitempty"`
} //@Name TreeSensorPredictionResponse

type TreeSensorDataResponse struct {
	Tree       *tree.Tree        `json:"tree,omitempty"`
	SensorData []sensor.MqttData `json:"sensor_data,omitempty"`
} //@Name TreeSensorDataResponse

//	@Summary		Get all trees
//	@Description	Get all trees
//	@Id				get-all-trees
//	@Tags			Trees
//	@Produce		json
//	@Param			sensor_data	query		boolean	false	"Get raw sensor data for each tree"
//	@Success		200			{object}	tree.TreeSensorDataResponse
//	@Failure		400			{object}	HTTPError
//	@Failure		401			{object}	HTTPError
//	@Failure		403			{object}	HTTPError
//	@Failure		404			{object}	HTTPError
//	@Failure		500			{object}	HTTPError
//	@Router			/tree [get]
func GetAllTree(svc service.TreeService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		treeData, err := svc.GetAllTrees(c.Context())
		if err != nil {
			return handler.HandleError(err)
		}

		reponse := utils.Map(treeData, func(tree tree.Tree) TreeSensorDataResponse {
			var sensorData []sensor.MqttData
			if c.QueryBool("sensor_data") {
				data, err := getSensorDataByTreeID(c.Context(), svc, tree.ID.Hex())
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

//	@Summary		Get tree by ID
//	@Description	Get tree by ID
//	@Id				get-tree-by-id
//	@Tags			Trees
//	@Produce		json
//	@Param			treeID		path		string	true	"Tree ID"
//	@Param			sensor_data	query		boolean	false	"Get raw sensor data for each tree"
//	@Success		200			{object}	tree.TreeSensorDataResponse
//	@Failure		400			{object}	HTTPError
//	@Failure		401			{object}	HTTPError
//	@Failure		403			{object}	HTTPError
//	@Failure		404			{object}	HTTPError
//	@Failure		500			{object}	HTTPError
//	@Router			/tree/{treeID} [get]
func GetTreeByID(svc service.TreeService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var response TreeSensorDataResponse
		tree, err := svc.GetTreeByID(c.Context(), c.Params("id"))
		if err != nil {
			return handler.HandleError(err)
		}
		response.Tree = tree

		if c.QueryBool("sensor_data") {
			data, err := getSensorDataByTreeID(c.Context(), svc, c.Params("id"))
			if err != nil {
				return handler.HandleError(err)
			}
			response.SensorData = data
		}

		return c.JSON(response)
	}
}

//	@Summary		Get tree prediction by tree ID
//	@Description	Get tree prediction by tree ID
//	@Id				get-tree-prediction-by-id
//	@Tags			Trees
//	@Produce		json
//	@Param			treeID		path		string	true	"Tree ID"
//	@Param			sensor_data	query		boolean	false	"Get raw sensor data for each tree"
//	@Success		200			{object}	tree.TreeSensorPredictionResponse
//	@Failure		400			{object}	HTTPError
//	@Failure		401			{object}	HTTPError
//	@Failure		403			{object}	HTTPError
//	@Failure		404			{object}	HTTPError
//	@Failure		500			{object}	HTTPError
//	@Router			/tree/{treeID}/prediction [get]
func GetTreePredictions(svc service.TreeService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var response TreeSensorPredictionResponse
		predict, err := svc.GetTreePrediction(c.Context(), c.Params("id"))
		if err != nil {
			return handler.HandleError(err)
		}

		response.SensorPrediction = predict

		if c.QueryBool("sensor_data") {
			data, err := getSensorDataByTreeID(c.Context(), svc, c.Params("id"))
			if err != nil {
				return handler.HandleError(err)
			}

			response.SensorData = data
		}

		return c.JSON(response)
	}
}

func getSensorDataByTreeID(ctx context.Context, svc service.TreeService, treeID string) ([]sensor.MqttData, error) {
	data, err := svc.GetSensorDataByTreeID(ctx, treeID)
	if err != nil {
		return nil, handler.HandleError(err)
	}

	return data, nil
}
