package status

import (
	"testing"
	"time"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListBackupRequests(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	created := time.Now().Add(-2 * time.Hour)
	updated := time.Now().Add(-30 * time.Minute)

	expectedResponse := &BackupRequestListResponse{
		Objects: []BackupRequest{
			{
				CreatedAt:   &util.InfinityTime{Time: created},
				UpdatedAt:   &util.InfinityTime{Time: updated},
				DownloadURI: "/downloads/backup/1",
				Message:     "Backup completed successfully",
				ResourceURI: "/api/admin/status/v1/backup_request/1/",
				State:       "completed",
			},
			{
				CreatedAt:   &util.InfinityTime{Time: created},
				UpdatedAt:   nil,
				DownloadURI: "",
				Message:     "Backup in progress",
				ResourceURI: "/api/admin/status/v1/backup_request/2/",
				State:       "running",
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
	assert.Equal(t, "completed", result.Objects[0].State)
	assert.Equal(t, "Backup completed successfully", result.Objects[0].Message)
	assert.Equal(t, "/downloads/backup/1", result.Objects[0].DownloadURI)
	assert.Equal(t, "/api/admin/status/v1/backup_request/1/", result.Objects[0].ResourceURI)
	assert.Equal(t, created.Unix(), result.Objects[0].CreatedAt.Time.Unix())
	assert.Equal(t, updated.Unix(), result.Objects[0].UpdatedAt.Time.Unix())
	assert.Equal(t, "running", result.Objects[1].State)
	assert.Equal(t, "Backup in progress", result.Objects[1].Message)
	assert.Equal(t, "", result.Objects[1].DownloadURI)
	assert.Equal(t, "/api/admin/status/v1/backup_request/2/", result.Objects[1].ResourceURI)
	assert.Equal(t, created.Unix(), result.Objects[1].CreatedAt.Time.Unix())
	assert.Nil(t, result.Objects[1].UpdatedAt)
	client.AssertExpectations(t)
}

func TestService_GetBackupRequest(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	created := time.Now().Add(-3 * time.Hour)
	completed := time.Now().Add(-1 * time.Hour)

	expectedRequest := &BackupRequest{
		CreatedAt:   &util.InfinityTime{Time: created},
		UpdatedAt:   &util.InfinityTime{Time: completed},
		DownloadURI: "/downloads/backup/1",
		Message:     "Full system backup completed",
		ResourceURI: "/api/admin/status/v1/backup_request/1/",
		State:       "completed",
	}

	client.On("GetJSON", t.Context(), "status/v1/backup_request/1/", mock.AnythingOfType("*status.BackupRequest")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*BackupRequest)
		*result = *expectedRequest
	})

	service := New(client)
	result, err := service.GetBackupRequest(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedRequest, result)
	assert.Equal(t, "completed", result.State)
	assert.Equal(t, "Full system backup completed", result.Message)
	assert.Equal(t, "/downloads/backup/1", result.DownloadURI)
	assert.Equal(t, "/api/admin/status/v1/backup_request/1/", result.ResourceURI)
	assert.Equal(t, created.Unix(), result.CreatedAt.Time.Unix())
	assert.Equal(t, completed.Unix(), result.UpdatedAt.Time.Unix())
	client.AssertExpectations(t)
}

func TestService_ListBackupRequests_WithOptions(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	opts := &ListOptions{
		Limit:  10,
		Offset: 0,
	}

	expectedResponse := &BackupRequestListResponse{
		Meta: Meta{
			Limit:      10,
			Next:       "",
			Offset:     0,
			Previous:   "",
			TotalCount: 1,
		},
		Objects: []BackupRequest{
			{
				CreatedAt:   nil,
				UpdatedAt:   nil,
				DownloadURI: "",
				Message:     "Automated backup pending",
				ResourceURI: "/api/admin/status/v1/backup_request/3/",
				State:       "pending",
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
	assert.Equal(t, "pending", result.Objects[0].State)
	assert.Equal(t, "Automated backup pending", result.Objects[0].Message)
	assert.Equal(t, "/api/admin/status/v1/backup_request/3/", result.Objects[0].ResourceURI)
	assert.Equal(t, 1, result.Meta.TotalCount)

	client.AssertExpectations(t)
}
