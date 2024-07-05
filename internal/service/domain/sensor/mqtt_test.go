package sensor

import (
	"testing"

	storageMock "github.com/SmartCityFlensburg/green-space-management/internal/storage/_mock"
	"github.com/stretchr/testify/assert"
)

func TestNewSensorService(t *testing.T) {
	repo := storageMock.NewMockSensorRepository(t)
	t.Run("should create a new service", func(t *testing.T) {
		svc := NewMqttService(repo)
		assert.NotNil(t, svc)
	})
}

// other test cases
