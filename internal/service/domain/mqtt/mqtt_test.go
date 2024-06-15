package mqtt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMqttService(t *testing.T) {
  t.Run("should create a new service", func(t *testing.T) {
    svc := NewMqttService()
    assert.NotNil(t, svc)
  })
}


// other test cases
