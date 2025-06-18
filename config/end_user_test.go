package config

import (
	"testing"

	mockClient "github.com/pexip/go-infinity-sdk/v38/internal/mock"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListEndUsers(t *testing.T) {
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
				expectedResponse := &EndUserListResponse{
					Objects: []EndUser{
						{ID: 1, PrimaryEmailAddress: "user1@example.com", FirstName: "John", LastName: "Doe"},
						{ID: 2, PrimaryEmailAddress: "user2@example.com", FirstName: "Jane", LastName: "Smith"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/end_user/", mock.AnythingOfType("*config.EndUserListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*EndUserListResponse)
					*result = *expectedResponse
				})
			},
			wantErr: false,
		},
		{
			name: "successful list with options",
			opts: &ListOptions{
				BaseListOptions: options.BaseListOptions{
					Limit:  5,
					Offset: 10,
				},
				Search: "john",
			},
			setup: func(m *mockClient.Client) {
				expectedResponse := &EndUserListResponse{
					Objects: []EndUser{
						{ID: 1, PrimaryEmailAddress: "john@example.com", FirstName: "John", LastName: "Doe"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/end_user/?limit=5&name__icontains=john&offset=10", mock.AnythingOfType("*config.EndUserListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*EndUserListResponse)
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
			result, err := service.ListEndUsers(t.Context(), tt.opts)

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

func TestService_GetEndUser(t *testing.T) {
	client := &mockClient.Client{}
	expectedUser := &EndUser{
		ID:                  1,
		PrimaryEmailAddress: "john@example.com",
		FirstName:           "John",
		LastName:            "Doe",
		DisplayName:         "John Doe",
		TelephoneNumber:     "+1234567890",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/end_user/1/", mock.AnythingOfType("*config.EndUser")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*EndUser)
		*result = *expectedUser
	})

	service := New(client)
	result, err := service.GetEndUser(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, result)
	client.AssertExpectations(t)
}

func TestService_CreateEndUser(t *testing.T) {
	client := &mockClient.Client{}

	createRequest := &EndUserCreateRequest{
		PrimaryEmailAddress: "newuser@example.com",
		FirstName:           "New",
		LastName:            "User",
		DisplayName:         "New User",
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/end_user/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/end_user/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateEndUser(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateEndUser(t *testing.T) {
	client := &mockClient.Client{}

	updateRequest := &EndUserUpdateRequest{
		DisplayName: "Updated User",
		Title:       "Senior Developer",
	}

	expectedUser := &EndUser{
		ID:          1,
		DisplayName: "Updated User",
		Title:       "Senior Developer",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/end_user/1/", updateRequest, mock.AnythingOfType("*config.EndUser")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*EndUser)
		*result = *expectedUser
	})

	service := New(client)
	result, err := service.UpdateEndUser(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, result)
	client.AssertExpectations(t)
}

func TestService_DeleteEndUser(t *testing.T) {
	client := &mockClient.Client{}

	client.On("DeleteJSON", t.Context(), "configuration/v1/end_user/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteEndUser(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
