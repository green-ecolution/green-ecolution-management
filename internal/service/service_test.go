package service

import (
	"testing"

	serviceMock "github.com/SmartCityFlensburg/green-space-management/internal/service/_mock"
	"github.com/stretchr/testify/assert"
)

func TestAllServiceReady(t *testing.T) {
	t.Run("should return true if all service implemented the ServiceReady interface", func(t *testing.T) {
		// given
		infoSvc := serviceMock.NewMockInfoService(t)
		mqttSvc := serviceMock.NewMockMqttService(t)
		treeSvc := serviceMock.NewMockTreeService(t)
		svc := Services{
			InfoService:   infoSvc,
			MqttService:   mqttSvc,
			TreeService:   treeSvc,
		}

		// when
		infoSvc.EXPECT().Ready().Return(true)
		mqttSvc.EXPECT().Ready().Return(true)
		treeSvc.EXPECT().Ready().Return(true)

		ready := svc.AllServicesReady()

		// then
		assert.True(t, ready)
	})
}
