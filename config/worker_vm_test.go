package config

import (
	"testing"

	mockClient "github.com/pexip/go-infinity-sdk/v38/internal/mock"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListWorkerVMs(t *testing.T) {
	tests := []struct {
		name    string
		opts    *ListOptions
		setup   func(m *mockClient.Client)
		wantErr bool
	}{
		{
			name: "successful list without options",
			opts: nil,
			setup: func(m *mockClient.Client) {
				expectedResponse := &WorkerVMListResponse{
					Objects: []WorkerVM{
						{ID: 1, Name: "worker-1", Hostname: "worker1.example.com"},
						{ID: 2, Name: "worker-2", Hostname: "worker2.example.com"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/worker_vm/", mock.AnythingOfType("*config.WorkerVMListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*WorkerVMListResponse)
					*result = *expectedResponse
				})
			},
			wantErr: false,
		},
		{
			name: "successful list with options",
			opts: &ListOptions{
				BaseListOptions: options.BaseListOptions{
					Limit:  2,
					Offset: 4,
				},
				Search: "worker",
			},
			setup: func(m *mockClient.Client) {
				expectedResponse := &WorkerVMListResponse{
					Objects: []WorkerVM{
						{ID: 1, Name: "test-worker", Hostname: "testworker.example.com"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/worker_vm/?limit=2&name__icontains=worker&offset=4", mock.AnythingOfType("*config.WorkerVMListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*WorkerVMListResponse)
					*result = *expectedResponse
				})
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &mockClient.Client{}
			tt.setup(client)

			service := New(client)
			result, err := service.ListWorkerVMs(t.Context(), tt.opts)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
			}

			client.AssertExpectations(t)
		})
	}
}

func TestService_GetWorkerVM(t *testing.T) {
	client := &mockClient.Client{}
	expectedVM := &WorkerVM{
		ID:             1,
		Name:           "worker-1",
		Hostname:       "worker1.example.com",
		Domain:         "example.com",
		Address:        "192.168.1.100",
		Netmask:        "255.255.255.0",
		Gateway:        "192.168.1.1",
		SystemLocation: "/api/admin/configuration/v1/system_location/1/",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/worker_vm/1/", mock.AnythingOfType("*config.WorkerVM")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*WorkerVM)
		*result = *expectedVM
	})

	service := New(client)
	result, err := service.GetWorkerVM(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedVM, result)
	client.AssertExpectations(t)
}

func TestService_CreateWorkerVM(t *testing.T) {
	client := &mockClient.Client{}

	createRequest := &WorkerVMCreateRequest{
		Name:           "new-worker",
		Hostname:       "newworker.example.com",
		Domain:         "example.com",
		Address:        "192.168.1.101",
		Netmask:        "255.255.255.0",
		Gateway:        "192.168.1.1",
		SystemLocation: "/api/admin/configuration/v1/system_location/1/",
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/worker_vm/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/worker_vm/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateWorkerVM(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateWorkerVM(t *testing.T) {
	client := &mockClient.Client{}

	updateRequest := &WorkerVMUpdateRequest{
		Name: "updated-worker",
	}

	expectedVM := &WorkerVM{
		ID:   1,
		Name: "updated-worker",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/worker_vm/1/", updateRequest, mock.AnythingOfType("*config.WorkerVM")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*WorkerVM)
		*result = *expectedVM
	})

	service := New(client)
	result, err := service.UpdateWorkerVM(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedVM, result)
	client.AssertExpectations(t)
}

func TestService_DeleteWorkerVM(t *testing.T) {
	client := &mockClient.Client{}

	client.On("DeleteJSON", t.Context(), "configuration/v1/worker_vm/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteWorkerVM(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
