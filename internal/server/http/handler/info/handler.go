package info

import (
	"github.com/SmartCityFlensburg/green-space-management/internal/server/http/handler"
	"github.com/SmartCityFlensburg/green-space-management/internal/service"
	_ "github.com/SmartCityFlensburg/green-space-management/internal/service/entities/info"
	"github.com/gofiber/fiber/v2"
)

// @Summary		Get info about the app
// @Description	Get info about the app and the server
// @Id				get-app-info
// @Tags			Info
// @Produce		json
// @Success		200	{object}	info.AppInfoResponse
// @Failure		400	{object}	HTTPError
// @Failure		401	{object}	HTTPError
// @Failure		403	{object}	HTTPError
// @Failure		404	{object}	HTTPError
// @Failure		500	{object}	HTTPError
// @Router			/info [get]
func GetAppInfo(svc service.InfoService) fiber.Handler {
	return func(c *fiber.Ctx) error {
    info, err := svc.GetAppInfoResponse(c.Context())
    if err != nil {
      return handler.HandleError(err)
    }
    return c.JSON(info)
	}
}
