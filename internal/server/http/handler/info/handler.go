package info

import (
	"github.com/SmartCityFlensburg/green-space-management/internal/entities/info"
	"github.com/SmartCityFlensburg/green-space-management/internal/server/http/handler"
	"github.com/SmartCityFlensburg/green-space-management/internal/service"
	"github.com/gofiber/fiber/v2"
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
      return handler.HandleError(err)
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
