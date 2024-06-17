package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/SmartCityFlensburg/green-space-management/internal/entities/info"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var (
	ErrIpNotFound            = errors.New("local ip not found")
	ErrIFacesNotFound        = errors.New("cant get interfaces")
	ErrIFacesAddressNotFound = errors.New("cant get interfaces address")
	ErrHostnameNotFound      = errors.New("cant get hostname")
)

type Error struct {
	Message string
	Code    ErrorCode
}

func NewError(code ErrorCode, msg string) Error {
	return Error{Code: code, Message: msg}
}

func (e Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

type ErrorCode int

const (
	BadRequest    ErrorCode = 400
	Unauthorized  ErrorCode = 401
	Forbidden     ErrorCode = 403
	NotFound      ErrorCode = 404
	InternalError ErrorCode = 500
)

type InfoService interface {
	GetAppInfo(context.Context) (*info.App, error)
	Ready() bool
}

type MqttService interface {
	HandleTemperature(client MQTT.Client, msg MQTT.Message)
	HandleHumidity(client MQTT.Client, msg MQTT.Message)
	SetConnected(bool)
	Ready() bool
}

type Service interface {
	Ready() bool
}

type Services struct {
	InfoService InfoService
	MqttService MqttService
}

func (s *Services) AllServicesReady() bool {
	v := reflect.ValueOf(s).Elem()
	for i := 0; i < v.NumField(); i++ {
		service := v.Field(i).Interface()
		if srv, ok := service.(Service); ok {
			if !srv.Ready() {
				return false
			}
		} else {
			log.Printf("Service %s does not implement the Service interface", v.Field(i).Type().Name())
			return false
		}
	}
	return true
}
