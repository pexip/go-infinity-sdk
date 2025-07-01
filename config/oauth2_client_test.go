package config

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListOAuth2Clients(t *testing.T) {
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
				privateKeyJWT1 := "private-key-jwt-1"
				privateKeyJWT2 := "private-key-jwt-2"
				expectedResponse := &OAuth2ClientListResponse{
					Objects: []OAuth2Client{
						{ClientID: "client-1", ClientName: "Primary OAuth2 Client", Role: "admin", PrivateKeyJWT: &privateKeyJWT1},
						{ClientID: "client-2", ClientName: "Secondary OAuth2 Client", Role: "user", PrivateKeyJWT: &privateKeyJWT2},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/oauth2_client/", mock.AnythingOfType("*config.OAuth2ClientListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*OAuth2ClientListResponse)
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
				Search: "Primary",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				privateKeyJWT := "private-key-jwt-1"
				expectedResponse := &OAuth2ClientListResponse{
					Objects: []OAuth2Client{
						{ClientID: "client-1", ClientName: "Primary OAuth2 Client", Role: "admin", PrivateKeyJWT: &privateKeyJWT},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/oauth2_client/?limit=5&name__icontains=Primary", mock.AnythingOfType("*config.OAuth2ClientListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*OAuth2ClientListResponse)
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
			result, err := service.ListOAuth2Clients(t.Context(), tt.opts)

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

func TestService_GetOAuth2Client(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	privateKeyJWT := "test-private-key-jwt"
	expectedOAuth2Client := &OAuth2Client{
		ClientID:      "test-client-id",
		ClientName:    "Test OAuth2 Client",
		Role:          "admin",
		PrivateKeyJWT: &privateKeyJWT,
	}

	client.On("GetJSON", t.Context(), "configuration/v1/oauth2_client/test-client-id/", mock.AnythingOfType("*config.OAuth2Client")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*OAuth2Client)
		*result = *expectedOAuth2Client
	})

	service := New(client)
	result, err := service.GetOAuth2Client(t.Context(), "test-client-id")

	assert.NoError(t, err)
	assert.Equal(t, expectedOAuth2Client, result)
	client.AssertExpectations(t)
}

func TestService_CreateOAuth2Client(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &OAuth2ClientCreateRequest{
		ClientName: "New OAuth2 Client",
		Role:       "user",
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/oauth2_client/new-client-id/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/oauth2_client/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateOAuth2Client(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateOAuth2Client(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	updateRequest := &OAuth2ClientUpdateRequest{
		ClientName: "Updated OAuth2 Client",
		Role:       "admin",
	}

	privateKeyJWT := "test-private-key-jwt"
	expectedOAuth2Client := &OAuth2Client{
		ClientID:      "test-client-id",
		ClientName:    "Updated OAuth2 Client",
		Role:          "admin",
		PrivateKeyJWT: &privateKeyJWT,
	}

	client.On("PutJSON", t.Context(), "configuration/v1/oauth2_client/test-client-id/", updateRequest, mock.AnythingOfType("*config.OAuth2Client")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*OAuth2Client)
		*result = *expectedOAuth2Client
	})

	service := New(client)
	result, err := service.UpdateOAuth2Client(t.Context(), "test-client-id", updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedOAuth2Client, result)
	client.AssertExpectations(t)
}

func TestService_DeleteOAuth2Client(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/oauth2_client/test-client-id/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteOAuth2Client(t.Context(), "test-client-id")

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
