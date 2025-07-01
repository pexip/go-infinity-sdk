package config

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListDNSServers(t *testing.T) {
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
				expectedResponse := &DNSServerListResponse{
					Objects: []DNSServer{
						{ID: 1, Address: "8.8.8.8", Description: "Google DNS"},
						{ID: 2, Address: "1.1.1.1", Description: "Cloudflare DNS"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/dns_server/", mock.AnythingOfType("*config.DNSServerListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*DNSServerListResponse)
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
				Search: "google",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				expectedResponse := &DNSServerListResponse{
					Objects: []DNSServer{
						{ID: 1, Address: "8.8.8.8", Description: "Google DNS"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/dns_server/?limit=5&name__icontains=google", mock.AnythingOfType("*config.DNSServerListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*DNSServerListResponse)
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
			result, err := service.ListDNSServers(t.Context(), tt.opts)

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

func TestService_GetDNSServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedServer := &DNSServer{
		ID:          1,
		Address:     "8.8.8.8",
		Description: "Google DNS Server",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/dns_server/1/", mock.AnythingOfType("*config.DNSServer")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*DNSServer)
		*result = *expectedServer
	})

	service := New(client)
	result, err := service.GetDNSServer(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedServer, result)
	client.AssertExpectations(t)
}

func TestService_CreateDNSServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &DNSServerCreateRequest{
		Address:     "9.9.9.9",
		Description: "Quad9 DNS",
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/dns_server/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/dns_server/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateDNSServer(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateDNSServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	updateRequest := &DNSServerUpdateRequest{
		Description: "Updated DNS Server",
	}

	expectedServer := &DNSServer{
		ID:          1,
		Address:     "8.8.8.8",
		Description: "Updated DNS Server",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/dns_server/1/", updateRequest, mock.AnythingOfType("*config.DNSServer")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*DNSServer)
		*result = *expectedServer
	})

	service := New(client)
	result, err := service.UpdateDNSServer(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedServer, result)
	client.AssertExpectations(t)
}

func TestService_DeleteDNSServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/dns_server/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteDNSServer(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
