package config

import (
	"testing"

	mockClient "github.com/pexip/go-infinity-sdk/internal/mock"
	"github.com/pexip/go-infinity-sdk/options"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListStaticRoutes(t *testing.T) {
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
				expectedResponse := &StaticRouteListResponse{
					Objects: []StaticRoute{
						{ID: 1, Name: "route-1", Address: "10.0.0.0", Prefix: 24, Gateway: "192.168.1.1"},
						{ID: 2, Name: "route-2", Address: "172.16.0.0", Prefix: 16, Gateway: "192.168.1.1"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/static_route/", mock.AnythingOfType("*config.StaticRouteListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*StaticRouteListResponse)
					*result = *expectedResponse
				})
			},
			wantErr: false,
		},
		{
			name: "successful list with options",
			opts: &ListOptions{
				BaseListOptions: options.BaseListOptions{
					Limit: 2,
				},
				Search: "route",
			},
			setup: func(m *mockClient.Client) {
				expectedResponse := &StaticRouteListResponse{
					Objects: []StaticRoute{
						{ID: 1, Name: "test-route", Address: "10.0.0.0", Prefix: 24, Gateway: "192.168.1.1"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/static_route/?limit=2&name__icontains=route", mock.AnythingOfType("*config.StaticRouteListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*StaticRouteListResponse)
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
			result, err := service.ListStaticRoutes(t.Context(), tt.opts)

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

func TestService_GetStaticRoute(t *testing.T) {
	client := &mockClient.Client{}
	expectedRoute := &StaticRoute{
		ID:      1,
		Name:    "test-route",
		Address: "10.0.0.0",
		Prefix:  24,
		Gateway: "192.168.1.1",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/static_route/1/", mock.AnythingOfType("*config.StaticRoute")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*StaticRoute)
		*result = *expectedRoute
	})

	service := New(client)
	result, err := service.GetStaticRoute(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedRoute, result)
	client.AssertExpectations(t)
}

func TestService_CreateStaticRoute(t *testing.T) {
	client := &mockClient.Client{}

	createRequest := &StaticRouteCreateRequest{
		Name:    "new-route",
		Address: "172.16.0.0",
		Prefix:  16,
		Gateway: "192.168.1.1",
	}

	expectedRoute := &StaticRoute{
		ID:      1,
		Name:    "new-route",
		Address: "172.16.0.0",
		Prefix:  16,
		Gateway: "192.168.1.1",
	}

	client.On("PostJSON", t.Context(), "configuration/v1/static_route/", createRequest, mock.AnythingOfType("*config.StaticRoute")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*StaticRoute)
		*result = *expectedRoute
	})

	service := New(client)
	result, err := service.CreateStaticRoute(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedRoute, result)
	client.AssertExpectations(t)
}

func TestService_UpdateStaticRoute(t *testing.T) {
	client := &mockClient.Client{}

	updateRequest := &StaticRouteUpdateRequest{
		Name:    "updated-route",
		Gateway: "192.168.1.254",
	}

	expectedRoute := &StaticRoute{
		ID:      1,
		Name:    "updated-route",
		Address: "10.0.0.0",
		Prefix:  24,
		Gateway: "192.168.1.254",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/static_route/1/", updateRequest, mock.AnythingOfType("*config.StaticRoute")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*StaticRoute)
		*result = *expectedRoute
	})

	service := New(client)
	result, err := service.UpdateStaticRoute(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedRoute, result)
	client.AssertExpectations(t)
}

func TestService_DeleteStaticRoute(t *testing.T) {
	client := &mockClient.Client{}

	client.On("DeleteJSON", t.Context(), "configuration/v1/static_route/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteStaticRoute(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}