package status

import (
	"testing"
	"time"

	mockClient "github.com/pexip/go-infinity-sdk/v38/internal/mock"
	"github.com/pexip/go-infinity-sdk/v38/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListManagementVMs(t *testing.T) {
	client := &mockClient.Client{}

	expectedResponse := &ManagementVMListResponse{
		Objects: []ManagementVM{
			{
				ID:              1,
				ConfigurationID: 101,
				Name:            "mgmt-vm-1",
				Primary:         true,
				SyncStatus:      "SYNCED",
				UpgradeStatus:   "IDLE",
				Version:         "30.0.0",
				ResourceURI:     "/api/admin/status/v1/management_vm/1/",
			},
			{
				ID:              2,
				ConfigurationID: 102,
				Name:            "mgmt-vm-2",
				Primary:         false,
				SyncStatus:      "SYNCED",
				UpgradeStatus:   "IDLE",
				Version:         "30.0.0",
				ResourceURI:     "/api/admin/status/v1/management_vm/2/",
			},
		},
	}

	client.On("GetJSON", t.Context(), "status/v1/management_vm/", mock.AnythingOfType("*status.ManagementVMListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*ManagementVMListResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ListManagementVMs(t.Context(), nil)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result.Objects))
	assert.Equal(t, "mgmt-vm-1", result.Objects[0].Name)
	assert.True(t, result.Objects[0].Primary)
	assert.Equal(t, "mgmt-vm-2", result.Objects[1].Name)
	assert.False(t, result.Objects[1].Primary)
	client.AssertExpectations(t)
}

func TestService_GetManagementVM(t *testing.T) {
	client := &mockClient.Client{}
	lastUpdated := time.Now()
	lastAttemptedContact := time.Now().Add(-1 * time.Hour)

	expectedVM := &ManagementVM{
		ID:                   1,
		ConfigurationID:      101,
		Name:                 "mgmt-vm-primary",
		Primary:              true,
		SyncStatus:           "SYNCED",
		UpgradeStatus:        "COMPLETE",
		Version:              "30.1.0",
		LastAttemptedContact: &util.InfinityTime{Time: lastAttemptedContact},
		LastUpdated:          &util.InfinityTime{Time: lastUpdated},
		ResourceURI:          "/api/admin/status/v1/management_vm/1/",
	}

	client.On("GetJSON", t.Context(), "status/v1/management_vm/1/", mock.AnythingOfType("*status.ManagementVM")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*ManagementVM)
		*result = *expectedVM
	})

	service := New(client)
	result, err := service.GetManagementVM(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedVM, result)
	assert.Equal(t, "mgmt-vm-primary", result.Name)
	assert.True(t, result.Primary)
	assert.Equal(t, "COMPLETE", result.UpgradeStatus)
	client.AssertExpectations(t)
}

func TestService_ListManagementVMs_WithOptions(t *testing.T) {
	client := &mockClient.Client{}
	service := New(client)

	opts := &ListOptions{
		Limit:  5,
		Offset: 0,
	}

	expectedResponse := &ManagementVMListResponse{
		Objects: []ManagementVM{
			{
				ID:            1,
				Name:          "mgmt-primary",
				Primary:       true,
				SyncStatus:    "SYNCED",
				UpgradeStatus: "IDLE",
				Version:       "30.0.0",
			},
		},
	}

	client.On("GetJSON", t.Context(), mock.MatchedBy(func(endpoint string) bool {
		return endpoint != "status/v1/management_vm/"
	}), mock.AnythingOfType("*status.ManagementVMListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*ManagementVMListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListManagementVMs(t.Context(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "mgmt-primary", result.Objects[0].Name)
	assert.True(t, result.Objects[0].Primary)

	client.AssertExpectations(t)
}
