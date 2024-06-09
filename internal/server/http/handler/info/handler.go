package info

import (
	"errors"

	"github.com/SmartCityFlensburg/green-space-management/internal/entities/info"
	"github.com/SmartCityFlensburg/green-space-management/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type GetAppInfoResponse struct {
	Version   string         `json:"version"`
	BuildTime string         `json:"buildTime"`
	GoVersion string         `json:"goVersion"`
	Git       GitResponse    `json:"git"`
	Server    ServerResponse `json:"server"`
}

type GitResponse struct {
	Branch     string `json:"branch"`
	Commit     string `json:"commit"`
	Repository string `json:"repository"`
}

type ServerResponse struct {
	OS        string `json:"os"`
	Arch      string `json:"arch"`
	Hostname  string `json:"hostname"`
	Url       string `json:"url"`
	IP        string `json:"ip"`
	Port      int    `json:"port"`
	Interface string `json:"interface"`
  UpTime    string `json:"uptime"`
}

func GetAppInfo(svc service.InfoService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		appInfo, err := svc.GetAppInfo(c.Context())
		if err != nil {
			code := fiber.StatusInternalServerError
			var svcErr service.Error

			if errors.As(err, &svcErr) {
				switch svcErr.Code {
				case service.NotFound:
					code = fiber.StatusNotFound
				case service.BadRequest:
					code = fiber.StatusBadRequest
				case service.Forbidden:
					code = fiber.StatusForbidden
				case service.Unauthorized:
					code = fiber.StatusUnauthorized
				case service.InternalError:
					code = fiber.StatusInternalServerError
				default:
					log.Debugf("missing service error code %d", svcErr.Code)
				}
			}

			return fiber.NewError(code, err.Error())
		}

		return c.JSON(mapToDto(appInfo))
	}
}

func mapToDto(entity *info.App) *GetAppInfoResponse {
	return &GetAppInfoResponse{
		Version:   entity.Version,
		BuildTime: entity.BuildTime.String(),
		GoVersion: entity.GoVersion,
		Git: GitResponse{
			Branch:     entity.Git.Branch,
			Commit:     entity.Git.Commit,
			Repository: entity.Git.Repository.String(),
		},
		Server: ServerResponse{
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
