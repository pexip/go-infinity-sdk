package config

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListUserGroupEntityMappings(t *testing.T) {
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
				expectedResponse := &UserGroupEntityMappingListResponse{
					Objects: []UserGroupEntityMapping{
						{ID: 1, Description: "Conference mapping for admin group", EntityResourceURI: "/api/admin/configuration/v1/conference/1/", UserGroup: "/api/admin/configuration/v1/user_group/1/"},
						{ID: 2, Description: "Location mapping for user group", EntityResourceURI: "/api/admin/configuration/v1/location/1/", UserGroup: "/api/admin/configuration/v1/user_group/2/"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/user_group_entity_mapping/", mock.AnythingOfType("*config.UserGroupEntityMappingListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*UserGroupEntityMappingListResponse)
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
				Search: "conference",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				expectedResponse := &UserGroupEntityMappingListResponse{
					Objects: []UserGroupEntityMapping{
						{ID: 1, Description: "Conference mapping for admin group", EntityResourceURI: "/api/admin/configuration/v1/conference/1/", UserGroup: "/api/admin/configuration/v1/user_group/1/"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/user_group_entity_mapping/?limit=5&name__icontains=conference", mock.AnythingOfType("*config.UserGroupEntityMappingListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*UserGroupEntityMappingListResponse)
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
			result, err := service.ListUserGroupEntityMappings(t.Context(), tt.opts)

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

func TestService_GetUserGroupEntityMapping(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedUserGroupEntityMapping := &UserGroupEntityMapping{
		ID:                1,
		Description:       "Test entity mapping",
		EntityResourceURI: "/api/admin/configuration/v1/conference/1/",
		UserGroup:         "/api/admin/configuration/v1/user_group/1/",
		ResourceURI:       "/api/admin/configuration/v1/user_group_entity_mapping/1/",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/user_group_entity_mapping/1/", mock.AnythingOfType("*config.UserGroupEntityMapping")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*UserGroupEntityMapping)
		*result = *expectedUserGroupEntityMapping
	})

	service := New(client)
	result, err := service.GetUserGroupEntityMapping(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedUserGroupEntityMapping, result)
	client.AssertExpectations(t)
}

func TestService_CreateUserGroupEntityMapping(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &UserGroupEntityMappingCreateRequest{
		Description:       "New entity mapping",
		EntityResourceURI: "/api/admin/configuration/v1/conference/2/",
		UserGroup:         "/api/admin/configuration/v1/user_group/2/",
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/user_group_entity_mapping/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/user_group_entity_mapping/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateUserGroupEntityMapping(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateUserGroupEntityMapping(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	updateRequest := &UserGroupEntityMappingUpdateRequest{
		Description: "Updated entity mapping",
		UserGroup:   "/api/admin/configuration/v1/user_group/3/",
	}

	expectedUserGroupEntityMapping := &UserGroupEntityMapping{
		ID:                1,
		Description:       "Updated entity mapping",
		EntityResourceURI: "/api/admin/configuration/v1/conference/1/",
		UserGroup:         "/api/admin/configuration/v1/user_group/3/",
		ResourceURI:       "/api/admin/configuration/v1/user_group_entity_mapping/1/",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/user_group_entity_mapping/1/", updateRequest, mock.AnythingOfType("*config.UserGroupEntityMapping")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*UserGroupEntityMapping)
		*result = *expectedUserGroupEntityMapping
	})

	service := New(client)
	result, err := service.UpdateUserGroupEntityMapping(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedUserGroupEntityMapping, result)
	client.AssertExpectations(t)
}

func TestService_DeleteUserGroupEntityMapping(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/user_group_entity_mapping/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteUserGroupEntityMapping(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
