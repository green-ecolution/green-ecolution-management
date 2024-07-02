package handler

import (
	"errors"

	"github.com/SmartCityFlensburg/green-space-management/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func HandleError(err error) *fiber.Error {
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
