package config

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListMjxExchangeAutodiscoverURLs(t *testing.T) {
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
				exchangeDeployment := "/api/admin/configuration/v1/mjx_exchange_deployment/1/"
				expectedResponse := &MjxExchangeAutodiscoverURLListResponse{
					Objects: []MjxExchangeAutodiscoverURL{
						{ID: 1, Name: "primary-autodiscover", Description: "Primary Exchange autodiscover URL", URL: "https://autodiscover.example.com/autodiscover/autodiscover.xml", ExchangeDeployment: &exchangeDeployment},
						{ID: 2, Name: "backup-autodiscover", Description: "Backup Exchange autodiscover URL", URL: "https://backup.example.com/autodiscover/autodiscover.xml", ExchangeDeployment: &exchangeDeployment},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/mjx_exchange_autodiscover_url/", mock.AnythingOfType("*config.MjxExchangeAutodiscoverURLListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*MjxExchangeAutodiscoverURLListResponse)
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
				exchangeDeployment := "/api/admin/configuration/v1/mjx_exchange_deployment/1/"
				expectedResponse := &MjxExchangeAutodiscoverURLListResponse{
					Objects: []MjxExchangeAutodiscoverURL{
						{ID: 1, Name: "primary-autodiscover", Description: "Primary Exchange autodiscover URL", URL: "https://autodiscover.example.com/autodiscover/autodiscover.xml", ExchangeDeployment: &exchangeDeployment},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/mjx_exchange_autodiscover_url/?limit=5&name__icontains=primary", mock.AnythingOfType("*config.MjxExchangeAutodiscoverURLListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*MjxExchangeAutodiscoverURLListResponse)
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
			result, err := service.ListMjxExchangeAutodiscoverURLs(t.Context(), tt.opts)

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

func TestService_GetMjxExchangeAutodiscoverURL(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	exchangeDeployment := "/api/admin/configuration/v1/mjx_exchange_deployment/1/"

	expectedURL := &MjxExchangeAutodiscoverURL{
		ID:                 1,
		Name:               "test-autodiscover",
		Description:        "Test Exchange autodiscover URL",
		URL:                "https://autodiscover.example.com/autodiscover/autodiscover.xml",
		ExchangeDeployment: &exchangeDeployment,
	}

	client.On("GetJSON", t.Context(), "configuration/v1/mjx_exchange_autodiscover_url/1/", mock.AnythingOfType("*config.MjxExchangeAutodiscoverURL")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*MjxExchangeAutodiscoverURL)
		*result = *expectedURL
	})

	service := New(client)
	result, err := service.GetMjxExchangeAutodiscoverURL(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedURL, result)
	client.AssertExpectations(t)
}

func TestService_CreateMjxExchangeAutodiscoverURL(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	exchangeDeployment := "/api/admin/configuration/v1/mjx_exchange_deployment/1/"
	createRequest := &MjxExchangeAutodiscoverURLCreateRequest{
		Name:               "new-autodiscover",
		Description:        "New Exchange autodiscover URL",
		URL:                "https://new.example.com/autodiscover/autodiscover.xml",
		ExchangeDeployment: &exchangeDeployment,
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/mjx_exchange_autodiscover_url/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/mjx_exchange_autodiscover_url/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateMjxExchangeAutodiscoverURL(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateMjxExchangeAutodiscoverURL(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	updateRequest := &MjxExchangeAutodiscoverURLUpdateRequest{
		Description: "Updated Exchange autodiscover URL",
		URL:         "https://updated.example.com/autodiscover/autodiscover.xml",
	}

	exchangeDeployment := "/api/admin/configuration/v1/mjx_exchange_deployment/1/"
	expectedURL := &MjxExchangeAutodiscoverURL{
		ID:                 1,
		Name:               "test-autodiscover",
		Description:        "Updated Exchange autodiscover URL",
		URL:                "https://updated.example.com/autodiscover/autodiscover.xml",
		ExchangeDeployment: &exchangeDeployment,
	}

	client.On("PutJSON", t.Context(), "configuration/v1/mjx_exchange_autodiscover_url/1/", updateRequest, mock.AnythingOfType("*config.MjxExchangeAutodiscoverURL")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*MjxExchangeAutodiscoverURL)
		*result = *expectedURL
	})

	service := New(client)
	result, err := service.UpdateMjxExchangeAutodiscoverURL(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedURL, result)
	client.AssertExpectations(t)
}

func TestService_DeleteMjxExchangeAutodiscoverURL(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/mjx_exchange_autodiscover_url/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteMjxExchangeAutodiscoverURL(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
