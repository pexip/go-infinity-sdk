package status

import (
	"testing"
	"time"

	mockClient "github.com/pexip/go-infinity-sdk/v38/internal/mock"
	"github.com/pexip/go-infinity-sdk/v38/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListSnapshotRequests(t *testing.T) {
	client := &mockClient.Client{}

	created := time.Now().Add(-3 * time.Hour)
	updated := time.Now().Add(-2 * time.Hour)

	expectedResponse := &SnapshotRequestListResponse{
		Objects: []SnapshotRequest{
			{
				CreatedAt:   &util.InfinityTime{Time: created},
				DownloadURI: "/downloads/snapshot-1",
				Message:     "Snapshot completed",
				ResourceURI: "/api/admin/status/v1/snapshot_request/1/",
				State:       "completed",
				UpdatedAt:   &util.InfinityTime{Time: updated},
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
	assert.Equal(t, "completed", result.Objects[0].State)
	assert.Equal(t, "Snapshot completed", result.Objects[0].Message)
	assert.Equal(t, "/downloads/snapshot-1", result.Objects[0].DownloadURI)
	assert.Equal(t, "/api/admin/status/v1/snapshot_request/1/", result.Objects[0].ResourceURI)
	assert.Equal(t, created.Unix(), result.Objects[0].CreatedAt.Time.Unix())
	assert.Equal(t, updated.Unix(), result.Objects[0].UpdatedAt.Time.Unix())
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
	updated := time.Now().Add(-30 * time.Minute)
	expectedResponse := &SnapshotRequestListResponse{
		Objects: []SnapshotRequest{
			{
				CreatedAt:   &util.InfinityTime{Time: created},
				DownloadURI: "/downloads/snapshot-2",
				Message:     "Options Test Snapshot",
				ResourceURI: "",
				State:       "completed",
				UpdatedAt:   &util.InfinityTime{Time: updated},
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
	assert.Equal(t, "Options Test Snapshot", result.Objects[0].Message)
	assert.Equal(t, "completed", result.Objects[0].State)
	assert.Equal(t, "/downloads/snapshot-2", result.Objects[0].DownloadURI)
	assert.Equal(t, created.Unix(), result.Objects[0].CreatedAt.Time.Unix())
	assert.Equal(t, updated.Unix(), result.Objects[0].UpdatedAt.Time.Unix())

	client.AssertExpectations(t)
}

func TestService_GetSnapshotRequest(t *testing.T) {
	client := &mockClient.Client{}

	created := time.Now().Add(-1 * time.Hour)
	updated := time.Now().Add(-30 * time.Minute)
	expectedRequest := &SnapshotRequest{
		CreatedAt:   &util.InfinityTime{Time: created},
		DownloadURI: "/downloads/snapshot-3",
		Message:     "Manual snapshot running",
		ResourceURI: "/api/admin/status/v1/snapshot_request/1/",
		State:       "running",
		UpdatedAt:   &util.InfinityTime{Time: updated},
	}

	client.On("GetJSON", t.Context(), "status/v1/snapshot_request/1/", mock.AnythingOfType("*status.SnapshotRequest")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*SnapshotRequest)
		*result = *expectedRequest
	})

	service := New(client)
	result, err := service.GetSnapshotRequest(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedRequest, result)
	assert.Equal(t, "running", result.State)
	assert.Equal(t, "Manual snapshot running", result.Message)
	assert.Equal(t, "/downloads/snapshot-3", result.DownloadURI)
	assert.Equal(t, "/api/admin/status/v1/snapshot_request/1/", result.ResourceURI)
	assert.Equal(t, created.Unix(), result.CreatedAt.Time.Unix())
	assert.Equal(t, updated.Unix(), result.UpdatedAt.Time.Unix())
	client.AssertExpectations(t)
}
