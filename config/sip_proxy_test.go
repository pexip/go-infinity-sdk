package config

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListSIPProxies(t *testing.T) {
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
				port1 := 5060
				port2 := 5061
				expectedResponse := &SIPProxyListResponse{
					Objects: []SIPProxy{
						{ID: 1, Name: "primary-proxy", Description: "Primary SIP proxy", Address: "sip1.example.com", Port: &port1, Transport: "UDP"},
						{ID: 2, Name: "secondary-proxy", Description: "Secondary SIP proxy", Address: "sip2.example.com", Port: &port2, Transport: "TCP"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/sip_proxy/", mock.AnythingOfType("*config.SIPProxyListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*SIPProxyListResponse)
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
				port := 5060
				expectedResponse := &SIPProxyListResponse{
					Objects: []SIPProxy{
						{ID: 1, Name: "primary-proxy", Description: "Primary SIP proxy", Address: "sip1.example.com", Port: &port, Transport: "UDP"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/sip_proxy/?limit=5&name__icontains=primary", mock.AnythingOfType("*config.SIPProxyListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*SIPProxyListResponse)
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
			result, err := service.ListSIPProxies(t.Context(), tt.opts)

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

func TestService_GetSIPProxy(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	port := 5080
	expectedSIPProxy := &SIPProxy{
		ID:          1,
		Name:        "test-sip-proxy",
		Description: "Test SIP Proxy",
		Address:     "test-sip.example.com",
		Port:        &port,
		Transport:   "TLS",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/sip_proxy/1/", mock.AnythingOfType("*config.SIPProxy")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*SIPProxy)
		*result = *expectedSIPProxy
	})

	service := New(client)
	result, err := service.GetSIPProxy(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedSIPProxy, result)
	client.AssertExpectations(t)
}

func TestService_CreateSIPProxy(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	port := 5070
	createRequest := &SIPProxyCreateRequest{
		Name:        "new-sip-proxy",
		Description: "New SIP Proxy",
		Address:     "new-sip.example.com",
		Port:        &port,
		Transport:   "UDP",
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/sip_proxy/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/sip_proxy/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateSIPProxy(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateSIPProxy(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	port := 5090
	updateRequest := &SIPProxyUpdateRequest{
		Name:        "updated-sip-proxy",
		Description: "Updated SIP Proxy",
		Port:        &port,
		Transport:   "TCP",
	}

	expectedSIPProxy := &SIPProxy{
		ID:          1,
		Name:        "updated-sip-proxy",
		Description: "Updated SIP Proxy",
		Address:     "test-sip.example.com",
		Port:        &port,
		Transport:   "TCP",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/sip_proxy/1/", updateRequest, mock.AnythingOfType("*config.SIPProxy")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*SIPProxy)
		*result = *expectedSIPProxy
	})

	service := New(client)
	result, err := service.UpdateSIPProxy(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedSIPProxy, result)
	client.AssertExpectations(t)
}

func TestService_DeleteSIPProxy(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/sip_proxy/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteSIPProxy(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
