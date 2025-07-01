package config

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListSnmpNetworkManagementSystems(t *testing.T) {
	tests := []struct {
		name    string
		opts    *ListOptions
		setup   func(m *interfaces.HTTPClientMock)
		wantErr bool
	}{
		{
			name: "successful list without options",
			opts: nil,
			setup: func(m *interfaces.HTTPClientMock) {
				expectedResponse := &SnmpNetworkManagementSystemListResponse{
					Objects: []SnmpNetworkManagementSystem{
						{ID: 1, Name: "primary-nms", Description: "Primary network management system", Address: "nms1.example.com", Port: 162, SnmpTrapCommunity: "public"},
						{ID: 2, Name: "backup-nms", Description: "Backup network management system", Address: "nms2.example.com", Port: 162, SnmpTrapCommunity: "private"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/snmp_network_management_system/", mock.AnythingOfType("*config.SnmpNetworkManagementSystemListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*SnmpNetworkManagementSystemListResponse)
					*result = *expectedResponse
				})
			},
			wantErr: false,
		},
		{
			name: "successful list with options",
			opts: &ListOptions{
				BaseListOptions: options.BaseListOptions{
					Limit: 5,
				},
				Search: "primary",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				expectedResponse := &SnmpNetworkManagementSystemListResponse{
					Objects: []SnmpNetworkManagementSystem{
						{ID: 1, Name: "primary-nms", Description: "Primary network management system", Address: "nms1.example.com", Port: 162, SnmpTrapCommunity: "public"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/snmp_network_management_system/?limit=5&name__icontains=primary", mock.AnythingOfType("*config.SnmpNetworkManagementSystemListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*SnmpNetworkManagementSystemListResponse)
					*result = *expectedResponse
				})
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := interfaces.NewHTTPClientMock()
			tt.setup(client)

			service := New(client)
			result, err := service.ListSnmpNetworkManagementSystems(t.Context(), tt.opts)

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

func TestService_GetSnmpNetworkManagementSystem(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedSnmpNMS := &SnmpNetworkManagementSystem{
		ID:                1,
		Name:              "test-nms",
		Description:       "Test SNMP Network Management System",
		Address:           "test-nms.example.com",
		Port:              162,
		SnmpTrapCommunity: "test-community",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/snmp_network_management_system/1/", mock.AnythingOfType("*config.SnmpNetworkManagementSystem")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*SnmpNetworkManagementSystem)
		*result = *expectedSnmpNMS
	})

	service := New(client)
	result, err := service.GetSnmpNetworkManagementSystem(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedSnmpNMS, result)
	client.AssertExpectations(t)
}

func TestService_CreateSnmpNetworkManagementSystem(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &SnmpNetworkManagementSystemCreateRequest{
		Name:              "new-nms",
		Description:       "New SNMP Network Management System",
		Address:           "new-nms.example.com",
		Port:              1162,
		SnmpTrapCommunity: "new-community",
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/snmp_network_management_system/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/snmp_network_management_system/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateSnmpNetworkManagementSystem(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateSnmpNetworkManagementSystem(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	port := 2162
	updateRequest := &SnmpNetworkManagementSystemUpdateRequest{
		Name:              "updated-nms",
		Description:       "Updated SNMP Network Management System",
		Port:              &port,
		SnmpTrapCommunity: "updated-community",
	}

	expectedSnmpNMS := &SnmpNetworkManagementSystem{
		ID:                1,
		Name:              "updated-nms",
		Description:       "Updated SNMP Network Management System",
		Address:           "test-nms.example.com",
		Port:              2162,
		SnmpTrapCommunity: "updated-community",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/snmp_network_management_system/1/", updateRequest, mock.AnythingOfType("*config.SnmpNetworkManagementSystem")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*SnmpNetworkManagementSystem)
		*result = *expectedSnmpNMS
	})

	service := New(client)
	result, err := service.UpdateSnmpNetworkManagementSystem(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedSnmpNMS, result)
	client.AssertExpectations(t)
}

func TestService_DeleteSnmpNetworkManagementSystem(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/snmp_network_management_system/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteSnmpNetworkManagementSystem(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
