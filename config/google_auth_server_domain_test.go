package config

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListGoogleAuthServerDomains(t *testing.T) {
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
				expectedResponse := &GoogleAuthServerDomainListResponse{
					Objects: []GoogleAuthServerDomain{
						{ID: 1, Domain: "example.com", Description: "Primary domain", GoogleAuthServer: "/api/admin/configuration/v1/google_auth_server/1/"},
						{ID: 2, Domain: "test.com", Description: "Test domain", GoogleAuthServer: "/api/admin/configuration/v1/google_auth_server/2/"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/google_auth_server_domain/", mock.AnythingOfType("*config.GoogleAuthServerDomainListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*GoogleAuthServerDomainListResponse)
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
				Search: "example",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				expectedResponse := &GoogleAuthServerDomainListResponse{
					Objects: []GoogleAuthServerDomain{
						{ID: 1, Domain: "example.com", Description: "Primary domain", GoogleAuthServer: "/api/admin/configuration/v1/google_auth_server/1/"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/google_auth_server_domain/?limit=5&name__icontains=example", mock.AnythingOfType("*config.GoogleAuthServerDomainListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*GoogleAuthServerDomainListResponse)
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
			result, err := service.ListGoogleAuthServerDomains(t.Context(), tt.opts)

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

func TestService_GetGoogleAuthServerDomain(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedDomain := &GoogleAuthServerDomain{
		ID:               1,
		Domain:           "example.com",
		Description:      "Test domain for Google OAuth",
		GoogleAuthServer: "/api/admin/configuration/v1/google_auth_server/1/",
		ResourceURI:      "/api/admin/configuration/v1/google_auth_server_domain/1/",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/google_auth_server_domain/1/", mock.AnythingOfType("*config.GoogleAuthServerDomain")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*GoogleAuthServerDomain)
		*result = *expectedDomain
	})

	service := New(client)
	result, err := service.GetGoogleAuthServerDomain(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedDomain, result)
	client.AssertExpectations(t)
}

func TestService_CreateGoogleAuthServerDomain(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &GoogleAuthServerDomainCreateRequest{
		Domain:           "newdomain.com",
		Description:      "New domain for Google OAuth",
		GoogleAuthServer: "/api/admin/configuration/v1/google_auth_server/1/",
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/google_auth_server_domain/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/google_auth_server_domain/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateGoogleAuthServerDomain(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateGoogleAuthServerDomain(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	updateRequest := &GoogleAuthServerDomainUpdateRequest{
		Domain:      "updated-domain.com",
		Description: "Updated domain description",
	}

	expectedDomain := &GoogleAuthServerDomain{
		ID:               1,
		Domain:           "updated-domain.com",
		Description:      "Updated domain description",
		GoogleAuthServer: "/api/admin/configuration/v1/google_auth_server/1/",
		ResourceURI:      "/api/admin/configuration/v1/google_auth_server_domain/1/",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/google_auth_server_domain/1/", updateRequest, mock.AnythingOfType("*config.GoogleAuthServerDomain")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*GoogleAuthServerDomain)
		*result = *expectedDomain
	})

	service := New(client)
	result, err := service.UpdateGoogleAuthServerDomain(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedDomain, result)
	client.AssertExpectations(t)
}

func TestService_DeleteGoogleAuthServerDomain(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/google_auth_server_domain/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteGoogleAuthServerDomain(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
