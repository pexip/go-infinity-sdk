package config

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_GetGlobalConfiguration(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedConfig := &GlobalConfiguration{
		ID:                     1,
		EnableWebRTC:           true,
		EnableSIP:              true,
		EnableH323:             false,
		EnableRTMP:             true,
		CryptoMode:             "best_effort",
		MaxPixelsPerSecond:     "1920x1080@30",
		MediaPortsStart:        40000,
		MediaPortsEnd:          49999,
		SignallingPortsStart:   5060,
		SignallingPortsEnd:     5061,
		BurstingEnabled:        false,
		GuestsOnlyTimeout:      300,
		WaitingForChairTimeout: 600,
		EnableAnalytics:        true,
		EnableErrorReporting:   false,
		AdministratorEmail:     "admin@example.com",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/global/1/", mock.AnythingOfType("*config.GlobalConfiguration")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*GlobalConfiguration)
		*result = *expectedConfig
	})

	service := New(client)
	result, err := service.GetGlobalConfiguration(t.Context())

	assert.NoError(t, err)
	assert.Equal(t, expectedConfig, result)
	client.AssertExpectations(t)
}

func TestService_UpdateGlobalConfiguration(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	enableAnalytics := false
	enableWebRTC := true
	updateRequest := &GlobalConfigurationUpdateRequest{
		EnableWebRTC:           &enableWebRTC,
		EnableAnalytics:        &enableAnalytics,
		AdministratorEmail:     "newemail@example.com",
		GuestsOnlyTimeout:      intPtr(600),
		WaitingForChairTimeout: intPtr(900),
	}

	expectedConfig := &GlobalConfiguration{
		ID:                     1,
		EnableWebRTC:           true,
		EnableAnalytics:        false,
		AdministratorEmail:     "newemail@example.com",
		GuestsOnlyTimeout:      600,
		WaitingForChairTimeout: 900,
	}

	client.On("PutJSON", t.Context(), "configuration/v1/global/1/", updateRequest, mock.AnythingOfType("*config.GlobalConfiguration")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*GlobalConfiguration)
		*result = *expectedConfig
	})

	service := New(client)
	result, err := service.UpdateGlobalConfiguration(t.Context(), updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedConfig, result)
	client.AssertExpectations(t)
}

// Helper function for tests
func intPtr(i int) *int {
	return &i
}
