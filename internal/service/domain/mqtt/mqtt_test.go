package mqtt

import (
	"testing"

	"github.com/stretchr/testify/assert"
	storageMock "github.com/SmartCityFlensburg/green-space-management/internal/storage/_mock"
)

func TestNewMqttService(t *testing.T) {
  repo := storageMock.NewMockMqttRepository(t)
	t.Run("should create a new service", func(t *testing.T) {
		svc := NewMqttService(repo)
		assert.NotNil(t, svc)
	})
}

// other test cases
