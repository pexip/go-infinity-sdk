package status

import (
	"testing"
	"time"

	mockClient "github.com/pexip/go-infinity-sdk/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListSnapshotRequests(t *testing.T) {
	client := &mockClient.Client{}

	created := time.Now().Add(-3 * time.Hour)
	started := time.Now().Add(-2 * time.Hour)
	completed := time.Now().Add(-1 * time.Hour)

	expectedResponse := &SnapshotRequestListResponse{
		Objects: []SnapshotRequest{
			{
				ID:          1,
				Status:      "completed",
				Created:     &created,
				Started:     &started,
				Completed:   &completed,
				Size:        512000000,
				Description: "Daily snapshot",
				ResourceURI: "/api/admin/status/v1/snapshot_request/1/",
			},
		},
	}

	client.On("GetJSON", t.Context(), "status/v1/snapshot_request/", mock.AnythingOfType("*status.SnapshotRequestListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*SnapshotRequestListResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ListSnapshotRequests(t.Context(), nil)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "completed", result.Objects[0].Status)
	assert.Equal(t, "Daily snapshot", result.Objects[0].Description)
	assert.Equal(t, int64(512000000), result.Objects[0].Size)
	client.AssertExpectations(t)
}

func TestService_ListSnapshotRequests_WithOptions(t *testing.T) {
	client := &mockClient.Client{}
	service := New(client)

	opts := &ListOptions{
		Limit:  7,
		Offset: 1,
	}

	created := time.Now().Add(-2 * time.Hour)
	completed := time.Now().Add(-30 * time.Minute)
	expectedResponse := &SnapshotRequestListResponse{
		Objects: []SnapshotRequest{
			{
				ID:          2,
				Status:      "completed",
				Created:     &created,
				Completed:   &completed,
				Size:        1024000000,
				Description: "Options Test Snapshot",
			},
		},
	}

	client.On("GetJSON", t.Context(), mock.MatchedBy(func(endpoint string) bool {
		return endpoint != "status/v1/snapshot_request/"
	}), mock.AnythingOfType("*status.SnapshotRequestListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*SnapshotRequestListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListSnapshotRequests(t.Context(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "Options Test Snapshot", result.Objects[0].Description)

	client.AssertExpectations(t)
}

func TestService_GetSnapshotRequest(t *testing.T) {
	client := &mockClient.Client{}

	created := time.Now().Add(-1 * time.Hour)
	started := time.Now().Add(-30 * time.Minute)
	expectedRequest := &SnapshotRequest{
		ID:          1,
		Status:      "running",
		Created:     &created,
		Started:     &started,
		Size:        0,
		Description: "Manual snapshot",
		ResourceURI: "/api/admin/status/v1/snapshot_request/1/",
	}

	client.On("GetJSON", t.Context(), "status/v1/snapshot_request/1/", mock.AnythingOfType("*status.SnapshotRequest")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*SnapshotRequest)
		*result = *expectedRequest
	})

	service := New(client)
	result, err := service.GetSnapshotRequest(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedRequest, result)
	client.AssertExpectations(t)
}