package status

import (
	"testing"
	"time"

	mockClient "github.com/pexip/go-infinity-sdk/v38/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListConferenceSyncs(t *testing.T) {
	client := &mockClient.Client{}

	lastSync := time.Now().Add(-30 * time.Minute)

	expectedResponse := &ConferenceSyncListResponse{
		Objects: []ConferenceSync{
			{
				ID:           1,
				Name:         "Primary Sync",
				Status:       "active",
				LastSync:     &lastSync,
				SyncInterval: 300,
				ErrorMessage: "",
				ResourceURI:  "/api/admin/status/v1/conference_sync/1/",
			},
		},
	}

	client.On("GetJSON", t.Context(), "status/v1/conference_sync/", mock.AnythingOfType("*status.ConferenceSyncListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*ConferenceSyncListResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ListConferenceSyncs(t.Context(), nil)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "Primary Sync", result.Objects[0].Name)
	assert.Equal(t, "active", result.Objects[0].Status)
	client.AssertExpectations(t)
}

func TestService_ListConferenceSyncs_WithOptions(t *testing.T) {
	client := &mockClient.Client{}
	service := New(client)

	opts := &ListOptions{
		Limit:  5,
		Offset: 0,
	}

	expectedResponse := &ConferenceSyncListResponse{
		Objects: []ConferenceSync{
			{
				ID:           2,
				Name:         "Test Sync With Options",
				Status:       "active",
				SyncInterval: 300,
				ErrorMessage: "",
			},
		},
	}

	client.On("GetJSON", t.Context(), mock.MatchedBy(func(endpoint string) bool {
		return endpoint != "status/v1/conference_sync/"
	}), mock.AnythingOfType("*status.ConferenceSyncListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*ConferenceSyncListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListConferenceSyncs(t.Context(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "Test Sync With Options", result.Objects[0].Name)

	client.AssertExpectations(t)
}

func TestService_GetConferenceSync(t *testing.T) {
	client := &mockClient.Client{}

	lastSync := time.Now().Add(-15 * time.Minute)
	expectedSync := &ConferenceSync{
		ID:           1,
		Name:         "Test Sync",
		Status:       "syncing",
		LastSync:     &lastSync,
		SyncInterval: 600,
		ErrorMessage: "",
		ResourceURI:  "/api/admin/status/v1/conference_sync/1/",
	}

	client.On("GetJSON", t.Context(), "status/v1/conference_sync/1/", mock.AnythingOfType("*status.ConferenceSync")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*ConferenceSync)
		*result = *expectedSync
	})

	service := New(client)
	result, err := service.GetConferenceSync(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedSync, result)
	client.AssertExpectations(t)
}
