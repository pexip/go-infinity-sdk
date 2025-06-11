package status

import (
	"testing"

	mockClient "github.com/pexip/go-infinity-sdk/v38/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListWorkerVMs(t *testing.T) {
	client := &mockClient.Client{}

	expectedResponse := &WorkerVMListResponse{
		Objects: []WorkerVM{
			{
				ID:             1,
				Name:           "pexip-worker-1",
				SyncStatus:     "SYNCED",
				NodeType:       "CONFERENCING",
				MediaLoad:      25,
				SignalingCount: 5,
			},
			{
				ID:             2,
				Name:           "pexip-worker-2",
				SyncStatus:     "SYNCED",
				NodeType:       "CONFERENCING",
				MediaLoad:      30,
				SignalingCount: 3,
			},
		},
	}

	client.On("GetJSON", t.Context(), "status/v1/worker_vm/", mock.AnythingOfType("*status.WorkerVMListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*WorkerVMListResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ListWorkerVMs(t.Context(), nil)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result.Objects))
	assert.Equal(t, "pexip-worker-1", result.Objects[0].Name)
	client.AssertExpectations(t)
}

func TestService_ListWorkerVMs_WithOptions(t *testing.T) {
	client := &mockClient.Client{}
	service := New(client)

	opts := &ListOptions{
		Limit:  5,
		Offset: 2,
	}

	expectedResponse := &WorkerVMListResponse{
		Objects: []WorkerVM{
			{
				ID:             3,
				Name:           "pexip-worker-3",
				SyncStatus:     "SYNCED",
				NodeType:       "CONFERENCING",
				MediaLoad:      15,
				SignalingCount: 2,
			},
		},
	}

	client.On("GetJSON", t.Context(), mock.MatchedBy(func(endpoint string) bool {
		return endpoint != "status/v1/worker_vm/"
	}), mock.AnythingOfType("*status.WorkerVMListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*WorkerVMListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListWorkerVMs(t.Context(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "pexip-worker-3", result.Objects[0].Name)

	client.AssertExpectations(t)
}

func TestService_GetWorkerVM(t *testing.T) {
	client := &mockClient.Client{}
	expectedWorkerVM := &WorkerVM{
		ID:             1,
		Name:           "pexip-worker-1",
		SyncStatus:     "SYNCED",
		NodeType:       "CONFERENCING",
		MediaLoad:      25,
		SignalingCount: 5,
		UpgradeStatus:  "IDLE",
	}

	client.On("GetJSON", t.Context(), "status/v1/worker_vm/1/", mock.AnythingOfType("*status.WorkerVM")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*WorkerVM)
		*result = *expectedWorkerVM
	})

	service := New(client)
	result, err := service.GetWorkerVM(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedWorkerVM, result)
	client.AssertExpectations(t)
}
