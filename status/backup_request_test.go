package status

import (
	"testing"
	"time"

	mockClient "github.com/pexip/go-infinity-sdk/v38/internal/mock"
	"github.com/pexip/go-infinity-sdk/v38/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListBackupRequests(t *testing.T) {
	client := &mockClient.Client{}

	created := time.Now().Add(-2 * time.Hour)
	started := time.Now().Add(-1 * time.Hour)
	completed := time.Now()

	expectedResponse := &BackupRequestListResponse{
		Objects: []BackupRequest{
			{
				ID:          1,
				Status:      "completed",
				Created:     &util.InfinityTime{Time: created},
				Started:     &util.InfinityTime{Time: started},
				Completed:   &util.InfinityTime{Time: completed},
				Size:        1024000000,
				Description: "Scheduled backup",
				ResourceURI: "/api/admin/status/v1/backup_request/1/",
			},
			{
				ID:          2,
				Status:      "running",
				Created:     &util.InfinityTime{Time: created},
				Started:     &util.InfinityTime{Time: started},
				Size:        0,
				Description: "Manual backup",
				ResourceURI: "/api/admin/status/v1/backup_request/2/",
			},
		},
	}

	client.On("GetJSON", t.Context(), "status/v1/backup_request/", mock.AnythingOfType("*status.BackupRequestListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*BackupRequestListResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ListBackupRequests(t.Context(), nil)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result.Objects))
	assert.Equal(t, "completed", result.Objects[0].Status)
	assert.Equal(t, "Scheduled backup", result.Objects[0].Description)
	assert.Equal(t, int64(1024000000), result.Objects[0].Size)
	assert.Equal(t, "running", result.Objects[1].Status)
	assert.Equal(t, "Manual backup", result.Objects[1].Description)
	client.AssertExpectations(t)
}

func TestService_GetBackupRequest(t *testing.T) {
	client := &mockClient.Client{}
	created := time.Now().Add(-3 * time.Hour)
	started := time.Now().Add(-2 * time.Hour)
	completed := time.Now().Add(-1 * time.Hour)

	expectedRequest := &BackupRequest{
		ID:          1,
		Status:      "completed",
		Created:     &util.InfinityTime{Time: created},
		Started:     &util.InfinityTime{Time: started},
		Completed:   &util.InfinityTime{Time: completed},
		Size:        2048000000,
		Description: "Full system backup",
		ResourceURI: "/api/admin/status/v1/backup_request/1/",
	}

	client.On("GetJSON", t.Context(), "status/v1/backup_request/1/", mock.AnythingOfType("*status.BackupRequest")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*BackupRequest)
		*result = *expectedRequest
	})

	service := New(client)
	result, err := service.GetBackupRequest(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedRequest, result)
	assert.Equal(t, "completed", result.Status)
	assert.Equal(t, "Full system backup", result.Description)
	assert.Equal(t, int64(2048000000), result.Size)
	client.AssertExpectations(t)
}

func TestService_ListBackupRequests_WithOptions(t *testing.T) {
	client := &mockClient.Client{}
	service := New(client)

	opts := &ListOptions{
		Limit:  10,
		Offset: 0,
	}

	expectedResponse := &BackupRequestListResponse{
		Objects: []BackupRequest{
			{
				ID:          1,
				Status:      "pending",
				Description: "Automated backup",
				Size:        0,
			},
		},
	}

	client.On("GetJSON", t.Context(), mock.MatchedBy(func(endpoint string) bool {
		return endpoint != "status/v1/backup_request/"
	}), mock.AnythingOfType("*status.BackupRequestListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*BackupRequestListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListBackupRequests(t.Context(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "pending", result.Objects[0].Status)
	assert.Equal(t, "Automated backup", result.Objects[0].Description)

	client.AssertExpectations(t)
}
