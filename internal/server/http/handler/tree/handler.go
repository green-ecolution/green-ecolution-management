package tree

import (
	"github.com/SmartCityFlensburg/green-space-management/internal/server/http/handler"
	"github.com/SmartCityFlensburg/green-space-management/internal/service"
	"github.com/gofiber/fiber/v2"
	_ "github.com/SmartCityFlensburg/green-space-management/internal/service/entities/tree"
)

// @Summary		Get all trees
// @Description	Get all trees
// @Id				get-all-trees
// @Tags			Trees
// @Produce		json
// @Param			sensor_data	query		boolean	false	"Get raw sensor data for each tree"
// @Success		200			{object}	tree.TreeSensorDataResponse
// @Failure		400			{object}	HTTPError
// @Failure		401			{object}	HTTPError
// @Failure		403			{object}	HTTPError
// @Failure		404			{object}	HTTPError
// @Failure		500			{object}	HTTPError
// @Router			/tree [get]
func GetAllTree(svc service.TreeService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		trees, err := svc.GetAllTreesResponse(c.Context(), c.QueryBool("sensor_data"))
		if err != nil {
			return handler.HandleError(err)
		}

		return c.JSON(trees)
	}
}

// @Summary		Get tree by ID
// @Description	Get tree by ID
// @Id				get-tree-by-id
// @Tags			Trees
// @Produce		json
// @Param			treeID		path		string	true	"Tree ID"
// @Param			sensor_data	query		boolean	false	"Get raw sensor data for each tree"
// @Success		200			{object}	tree.TreeSensorDataResponse
// @Failure		400			{object}	HTTPError
// @Failure		401			{object}	HTTPError
// @Failure		403			{object}	HTTPError
// @Failure		404			{object}	HTTPError
// @Failure		500			{object}	HTTPError
// @Router			/tree/{treeID} [get]
func GetTreeByID(svc service.TreeService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tree, err := svc.GetTreeByIDResponse(c.Context(), c.Params("id"), c.QueryBool("sensor_data"))
		if err != nil {
			return handler.HandleError(err)
		}

		return c.JSON(tree)
	}
}

// @Summary		Get tree prediction by tree ID
// @Description	Get tree prediction by tree ID
// @Id				get-tree-prediction-by-id
// @Tags			Trees
// @Produce		json
// @Param			treeID		path		string	true	"Tree ID"
// @Param			sensor_data	query		boolean	false	"Get raw sensor data for each tree"
// @Success		200			{object}	tree.TreeSensorPredictionResponse
// @Failure		400			{object}	HTTPError
// @Failure		401			{object}	HTTPError
// @Failure		403			{object}	HTTPError
// @Failure		404			{object}	HTTPError
// @Failure		500			{object}	HTTPError
// @Router			/tree/{treeID}/prediction [get]
func GetTreePredictions(svc service.TreeService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tree, err := svc.GetTreePredictionResponse(c.Context(), c.Params("id"), c.QueryBool("sensor_data"))
		if err != nil {
			return handler.HandleError(err)
		}

		return c.JSON(tree)
	}
}
