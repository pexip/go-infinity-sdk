package status

import (
	"testing"
	"time"

	mockClient "github.com/pexip/go-infinity-sdk/v38/internal/mock"
	util "github.com/pexip/go-infinity-sdk/v38/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListConferenceSyncs(t *testing.T) {
	client := &mockClient.Client{}

	lastUpdated := time.Now().Add(-30 * time.Minute)

	expectedResponse := &ConferenceSyncListResponse{
		Meta: Meta{
			Limit:      100,
			Next:       "",
			Offset:     0,
			Previous:   "",
			TotalCount: 1,
		},
		Objects: []ConferenceSync{
			{
				ConfigurationID:          0,
				DevicesCreated:           0,
				DevicesDeleted:           0,
				DevicesUnchanged:         0,
				DevicesUpdated:           0,
				EndUsersCreated:          0,
				EndUsersDeleted:          0,
				EndUsersUnchanged:        0,
				EndUsersUpdated:          0,
				ID:                       1,
				LastUpdated:              &util.InfinityTime{Time: lastUpdated},
				ResourceURI:              "/api/admin/status/v1/conference_sync/1/",
				SyncErrors:               0,
				SyncLastErrorDescription: "",
				SyncProgress:             0,
				SyncStatus:               "active",
				VMRsCreated:              0,
				VMRsDeleted:              0,
				VMRsUnchanged:            0,
				VMRsUpdated:              0,
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
	assert.Equal(t, 1, result.Objects[0].ID)
	assert.Equal(t, "active", result.Objects[0].SyncStatus)
	assert.Equal(t, "/api/admin/status/v1/conference_sync/1/", result.Objects[0].ResourceURI)
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
		Meta: Meta{
			Limit:      5,
			Next:       "",
			Offset:     0,
			Previous:   "",
			TotalCount: 1,
		},
		Objects: []ConferenceSync{
			{
				ConfigurationID:          0,
				DevicesCreated:           0,
				DevicesDeleted:           0,
				DevicesUnchanged:         0,
				DevicesUpdated:           0,
				EndUsersCreated:          0,
				EndUsersDeleted:          0,
				EndUsersUnchanged:        0,
				EndUsersUpdated:          0,
				ID:                       2,
				LastUpdated:              nil,
				ResourceURI:              "/api/admin/status/v1/conference_sync/2/",
				SyncErrors:               0,
				SyncLastErrorDescription: "",
				SyncProgress:             0,
				SyncStatus:               "active",
				VMRsCreated:              0,
				VMRsDeleted:              0,
				VMRsUnchanged:            0,
				VMRsUpdated:              0,
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
	assert.Equal(t, 2, result.Objects[0].ID)
	assert.Equal(t, "active", result.Objects[0].SyncStatus)
	assert.Equal(t, "/api/admin/status/v1/conference_sync/2/", result.Objects[0].ResourceURI)

	client.AssertExpectations(t)
}

func TestService_GetConferenceSync(t *testing.T) {
	client := &mockClient.Client{}

	lastUpdated := time.Now().Add(-15 * time.Minute)
	expectedSync := &ConferenceSync{
		ConfigurationID:          0,
		DevicesCreated:           0,
		DevicesDeleted:           0,
		DevicesUnchanged:         0,
		DevicesUpdated:           0,
		EndUsersCreated:          0,
		EndUsersDeleted:          0,
		EndUsersUnchanged:        0,
		EndUsersUpdated:          0,
		ID:                       1,
		LastUpdated:              &util.InfinityTime{Time: lastUpdated},
		ResourceURI:              "/api/admin/status/v1/conference_sync/1/",
		SyncErrors:               0,
		SyncLastErrorDescription: "",
		SyncProgress:             0,
		SyncStatus:               "syncing",
		VMRsCreated:              0,
		VMRsDeleted:              0,
		VMRsUnchanged:            0,
		VMRsUpdated:              0,
	}

	client.On("GetJSON", t.Context(), "status/v1/conference_sync/1/", mock.AnythingOfType("*status.ConferenceSync")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*ConferenceSync)
		*result = *expectedSync
	})

	service := New(client)
	result, err := service.GetConferenceSync(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedSync, result)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "syncing", result.SyncStatus)
	assert.Equal(t, "/api/admin/status/v1/conference_sync/1/", result.ResourceURI)
	client.AssertExpectations(t)
}
