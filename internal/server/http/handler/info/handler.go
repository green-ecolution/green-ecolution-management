package info

import (
	internal "github.com/SmartCityFlensburg/green-space-management/internal/entities/info"
	"github.com/SmartCityFlensburg/green-space-management/internal/server/http/handler"
	"github.com/SmartCityFlensburg/green-space-management/internal/service"
	"github.com/SmartCityFlensburg/green-space-management/internal/service/entities/info"
	"github.com/gofiber/fiber/v2"
)

// @Summary		Get info about the app
// @Description	Get info about the app and the server
// @Id				get-app-info
// @Tags			Info
// @Produce		json
// @Success		200	{object}	info.GetAppInfoResponse
// @Failure		400	{object}	HTTPError
// @Failure		401	{object}	HTTPError
// @Failure		403	{object}	HTTPError
// @Failure		404	{object}	HTTPError
// @Failure		500	{object}	HTTPError
// @Router			/info [get]
func GetAppInfo(svc service.InfoService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		appInfo, err := svc.GetAppInfo(c.Context())
		if err != nil {
			return handler.HandleError(err)
		}

		return c.JSON(mapToDto(appInfo))
	}
}

func mapToDto(entity *internal.App) *info.GetAppInfoResponse {
	return &info.GetAppInfoResponse{
		Version:   entity.Version,
		BuildTime: entity.BuildTime.String(),
		GoVersion: entity.GoVersion,
		Git: info.GitResponse{
			Branch:     entity.Git.Branch,
			Commit:     entity.Git.Commit,
			Repository: entity.Git.Repository.String(),
		},
		Server: info.ServerResponse{
			OS:        entity.Server.OS,
			Arch:      entity.Server.Arch,
			Hostname:  entity.Server.Hostname,
			Url:       entity.Server.Url.String(),
			IP:        entity.Server.IP.String(),
			Port:      entity.Server.Port,
			Interface: entity.Server.Interface,
			UpTime:    entity.Server.Uptime.String(),
		},
	}
}
