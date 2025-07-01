package config

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListSystemTuneables(t *testing.T) {
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
				expectedResponse := &SystemTuneableListResponse{
					Objects: []SystemTuneable{
						{ID: 1, Name: "max_connections", Setting: "1000"},
						{ID: 2, Name: "timeout_interval", Setting: "30"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/system_tuneable/", mock.AnythingOfType("*config.SystemTuneableListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*SystemTuneableListResponse)
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
				Search: "max",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				expectedResponse := &SystemTuneableListResponse{
					Objects: []SystemTuneable{
						{ID: 1, Name: "max_connections", Setting: "1000"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/system_tuneable/?limit=5&name__icontains=max", mock.AnythingOfType("*config.SystemTuneableListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*SystemTuneableListResponse)
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
			result, err := service.ListSystemTuneables(t.Context(), tt.opts)

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

func TestService_GetSystemTuneable(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedSystemTuneable := &SystemTuneable{
		ID:          1,
		Name:        "max_connections",
		Setting:     "1000",
		ResourceURI: "/api/admin/configuration/v1/system_tuneable/1/",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/system_tuneable/1/", mock.AnythingOfType("*config.SystemTuneable")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*SystemTuneable)
		*result = *expectedSystemTuneable
	})

	service := New(client)
	result, err := service.GetSystemTuneable(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedSystemTuneable, result)
	client.AssertExpectations(t)
}

func TestService_CreateSystemTuneable(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &SystemTuneableCreateRequest{
		Name:    "new_setting",
		Setting: "500",
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/system_tuneable/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/system_tuneable/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateSystemTuneable(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateSystemTuneable(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	updateRequest := &SystemTuneableUpdateRequest{
		Setting: "2000",
	}

	expectedSystemTuneable := &SystemTuneable{
		ID:          1,
		Name:        "max_connections",
		Setting:     "2000",
		ResourceURI: "/api/admin/configuration/v1/system_tuneable/1/",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/system_tuneable/1/", updateRequest, mock.AnythingOfType("*config.SystemTuneable")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*SystemTuneable)
		*result = *expectedSystemTuneable
	})

	service := New(client)
	result, err := service.UpdateSystemTuneable(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedSystemTuneable, result)
	client.AssertExpectations(t)
}

func TestService_DeleteSystemTuneable(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/system_tuneable/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteSystemTuneable(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
