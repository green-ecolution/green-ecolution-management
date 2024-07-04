package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/SmartCityFlensburg/green-space-management/internal/entities/info"
	"github.com/SmartCityFlensburg/green-space-management/internal/entities/sensor"
	"github.com/SmartCityFlensburg/green-space-management/internal/entities/tree"
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
	Service
	GetAppInfo(context.Context) (*info.App, error)
}

type MqttService interface {
	Service
	HandleMessage(client MQTT.Client, msg MQTT.Message)
	SetConnected(bool)
}

type SenserService interface {
	Service
  GetHumidity(context.Context) (int, error)
  GetBattery(context.Context) (float64, error)
  GetMqttDataByTreeID(context.Context, string) ([]sensor.MqttData, error)
  GetMqttDataByTreeIDLast(context.Context, string) (*sensor.MqttData, error)
}

type TreeService interface {
  Service
  GetTreeByID(ctx context.Context, id string) (*tree.Tree, error)
  GetAllTrees(context.Context) ([]tree.Tree, error)
  InsertTree(ctx context.Context, data tree.Tree) error
  GetSensorDataByTreeID(ctx context.Context, treeID string) ([]sensor.MqttData, error)
  GetTreePrediction(ctx context.Context, treeID string) (*tree.SensorPrediction, error)
}

type Service interface {
	Ready() bool
}

type Services struct {
	InfoService InfoService
	MqttService MqttService
  SenserService SenserService
  TreeService TreeService
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
