package status

import (
	"testing"
	"time"

	mockClient "github.com/pexip/go-infinity-sdk/v38/internal/mock"
	"github.com/pexip/go-infinity-sdk/v38/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_GetSystemStatus(t *testing.T) {
	client := &mockClient.Client{}

	expectedStatus := &SystemStatus{
		Status:      "healthy",
		Version:     "29.0.0",
		Uptime:      3600,
		Timestamp:   util.InfinityTime{Time: time.Now()},
		HostName:    "pexip-mgmt",
		TotalMemory: 8589934592,
		UsedMemory:  4294967296,
		CPULoad:     25.5,
	}

	client.On("GetJSON", t.Context(), "status/v1/system_status/", mock.AnythingOfType("*status.SystemStatus")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*SystemStatus)
		*result = *expectedStatus
	})

	service := New(client)
	result, err := service.GetSystemStatus(t.Context())

	assert.NoError(t, err)
	assert.Equal(t, expectedStatus, result)
	client.AssertExpectations(t)
}
