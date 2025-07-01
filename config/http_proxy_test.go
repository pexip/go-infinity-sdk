package config

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListHTTPProxies(t *testing.T) {
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
				port1 := 8080
				port2 := 3128
				expectedResponse := &HTTPProxyListResponse{
					Objects: []HTTPProxy{
						{ID: 1, Name: "primary-proxy", Address: "proxy.example.com", Port: &port1, Username: "user1", Password: "pass1", Protocol: "http"},
						{ID: 2, Name: "backup-proxy", Address: "backup.example.com", Port: &port2, Username: "user2", Password: "pass2", Protocol: "https"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/http_proxy/", mock.AnythingOfType("*config.HTTPProxyListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*HTTPProxyListResponse)
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
				port := 8080
				expectedResponse := &HTTPProxyListResponse{
					Objects: []HTTPProxy{
						{ID: 1, Name: "primary-proxy", Address: "proxy.example.com", Port: &port, Username: "user1", Password: "pass1", Protocol: "http"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/http_proxy/?limit=5&name__icontains=primary", mock.AnythingOfType("*config.HTTPProxyListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*HTTPProxyListResponse)
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
			result, err := service.ListHTTPProxies(t.Context(), tt.opts)

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

func TestService_GetHTTPProxy(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	port := 8080
	expectedProxy := &HTTPProxy{
		ID:          1,
		Name:        "test-proxy",
		Address:     "proxy.example.com",
		Port:        &port,
		Username:    "testuser",
		Password:    "testpass",
		Protocol:    "http",
		ResourceURI: "/api/admin/configuration/v1/http_proxy/1/",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/http_proxy/1/", mock.AnythingOfType("*config.HTTPProxy")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*HTTPProxy)
		*result = *expectedProxy
	})

	service := New(client)
	result, err := service.GetHTTPProxy(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedProxy, result)
	client.AssertExpectations(t)
}

func TestService_CreateHTTPProxy(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	port := 8080
	createRequest := &HTTPProxyCreateRequest{
		Name:     "new-proxy",
		Address:  "new-proxy.example.com",
		Port:     &port,
		Username: "newuser",
		Password: "newpass",
		Protocol: "http",
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/http_proxy/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/http_proxy/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateHTTPProxy(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateHTTPProxy(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	newPort := 3128
	updateRequest := &HTTPProxyUpdateRequest{
		Name:     "updated-proxy",
		Port:     &newPort,
		Protocol: "https",
	}

	expectedProxy := &HTTPProxy{
		ID:          1,
		Name:        "updated-proxy",
		Address:     "proxy.example.com",
		Port:        &newPort,
		Username:    "testuser",
		Password:    "testpass",
		Protocol:    "https",
		ResourceURI: "/api/admin/configuration/v1/http_proxy/1/",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/http_proxy/1/", updateRequest, mock.AnythingOfType("*config.HTTPProxy")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*HTTPProxy)
		*result = *expectedProxy
	})

	service := New(client)
	result, err := service.UpdateHTTPProxy(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedProxy, result)
	client.AssertExpectations(t)
}

func TestService_DeleteHTTPProxy(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/http_proxy/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteHTTPProxy(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
