package config

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListIdentityProviderAttributes(t *testing.T) {
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
				expectedResponse := &IdentityProviderAttributeListResponse{
					Objects: []IdentityProviderAttribute{
						{ID: 1, Name: "displayName", Description: "User display name attribute"},
						{ID: 2, Name: "email", Description: "User email address attribute"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/identity_provider_attribute/", mock.AnythingOfType("*config.IdentityProviderAttributeListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*IdentityProviderAttributeListResponse)
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
				Search: "email",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				expectedResponse := &IdentityProviderAttributeListResponse{
					Objects: []IdentityProviderAttribute{
						{ID: 2, Name: "email", Description: "User email address attribute"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/identity_provider_attribute/?limit=5&name__icontains=email", mock.AnythingOfType("*config.IdentityProviderAttributeListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*IdentityProviderAttributeListResponse)
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
			result, err := service.ListIdentityProviderAttributes(t.Context(), tt.opts)

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

func TestService_GetIdentityProviderAttribute(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedAttribute := &IdentityProviderAttribute{
		ID:          1,
		Name:        "displayName",
		Description: "User display name attribute mapping",
		ResourceURI: "/api/admin/configuration/v1/identity_provider_attribute/1/",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/identity_provider_attribute/1/", mock.AnythingOfType("*config.IdentityProviderAttribute")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*IdentityProviderAttribute)
		*result = *expectedAttribute
	})

	service := New(client)
	result, err := service.GetIdentityProviderAttribute(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedAttribute, result)
	client.AssertExpectations(t)
}

func TestService_CreateIdentityProviderAttribute(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &IdentityProviderAttributeCreateRequest{
		Name:        "department",
		Description: "User department attribute",
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/identity_provider_attribute/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/identity_provider_attribute/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateIdentityProviderAttribute(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateIdentityProviderAttribute(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	updateRequest := &IdentityProviderAttributeUpdateRequest{
		Name:        "updatedDisplayName",
		Description: "Updated user display name attribute",
	}

	expectedAttribute := &IdentityProviderAttribute{
		ID:          1,
		Name:        "updatedDisplayName",
		Description: "Updated user display name attribute",
		ResourceURI: "/api/admin/configuration/v1/identity_provider_attribute/1/",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/identity_provider_attribute/1/", updateRequest, mock.AnythingOfType("*config.IdentityProviderAttribute")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*IdentityProviderAttribute)
		*result = *expectedAttribute
	})

	service := New(client)
	result, err := service.UpdateIdentityProviderAttribute(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedAttribute, result)
	client.AssertExpectations(t)
}

func TestService_DeleteIdentityProviderAttribute(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/identity_provider_attribute/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteIdentityProviderAttribute(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
