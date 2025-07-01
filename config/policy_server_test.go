package config

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListPolicyServers(t *testing.T) {
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
				expectedResponse := &PolicyServerListResponse{
					Objects: []PolicyServer{
						{ID: 1, Name: "policy-server1", Description: "Primary policy server", URL: "https://policy.example.com", EnableServiceLookup: true, EnableParticipantLookup: true},
						{ID: 2, Name: "policy-server2", Description: "Secondary policy server", URL: "https://policy2.example.com", EnableDirectoryLookup: true, EnableAvatarLookup: true},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/policy_server/", mock.AnythingOfType("*config.PolicyServerListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*PolicyServerListResponse)
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
				Search: "policy-server1",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				expectedResponse := &PolicyServerListResponse{
					Objects: []PolicyServer{
						{ID: 1, Name: "policy-server1", Description: "Primary policy server", URL: "https://policy.example.com", EnableServiceLookup: true, EnableParticipantLookup: true},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/policy_server/?limit=5&name__icontains=policy-server1", mock.AnythingOfType("*config.PolicyServerListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*PolicyServerListResponse)
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
			result, err := service.ListPolicyServers(t.Context(), tt.opts)

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

func TestService_GetPolicyServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedPolicyServer := &PolicyServer{
		ID:                                  1,
		Name:                                "test-policy-server",
		Description:                         "Test Policy Server",
		URL:                                 "https://policy.example.com",
		Username:                            "testuser",
		Password:                            "testpass",
		EnableServiceLookup:                 true,
		EnableParticipantLookup:             true,
		EnableRegistrationLookup:            true,
		EnableDirectoryLookup:               false,
		EnableAvatarLookup:                  true,
		EnableMediaLocationLookup:           false,
		EnableInternalServicePolicy:         true,
		EnableInternalParticipantPolicy:     false,
		EnableInternalMediaLocationPolicy:   false,
		InternalServicePolicyTemplate:       "service_template",
		InternalParticipantPolicyTemplate:   "participant_template",
		InternalMediaLocationPolicyTemplate: "media_template",
		PreferLocalAvatarConfiguration:      true,
	}

	client.On("GetJSON", t.Context(), "configuration/v1/policy_server/1/", mock.AnythingOfType("*config.PolicyServer")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*PolicyServer)
		*result = *expectedPolicyServer
	})

	service := New(client)
	result, err := service.GetPolicyServer(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedPolicyServer, result)
	client.AssertExpectations(t)
}

func TestService_CreatePolicyServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &PolicyServerCreateRequest{
		Name:                                "new-policy-server",
		Description:                         "New Policy Server",
		URL:                                 "https://new-policy.example.com",
		Username:                            "newuser",
		Password:                            "newpass",
		EnableServiceLookup:                 true,
		EnableParticipantLookup:             true,
		EnableRegistrationLookup:            false,
		EnableDirectoryLookup:               true,
		EnableAvatarLookup:                  false,
		EnableMediaLocationLookup:           true,
		EnableInternalServicePolicy:         false,
		EnableInternalParticipantPolicy:     true,
		EnableInternalMediaLocationPolicy:   false,
		InternalServicePolicyTemplate:       "new_service_template",
		InternalParticipantPolicyTemplate:   "new_participant_template",
		InternalMediaLocationPolicyTemplate: "new_media_template",
		PreferLocalAvatarConfiguration:      false,
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/policy_server/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/policy_server/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreatePolicyServer(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdatePolicyServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	enableServiceLookup := false
	enableAvatarLookup := true

	updateRequest := &PolicyServerUpdateRequest{
		Description:         "Updated Policy Server",
		URL:                 "https://updated-policy.example.com",
		EnableServiceLookup: &enableServiceLookup,
		EnableAvatarLookup:  &enableAvatarLookup,
	}

	expectedPolicyServer := &PolicyServer{
		ID:                                  1,
		Name:                                "test-policy-server",
		Description:                         "Updated Policy Server",
		URL:                                 "https://updated-policy.example.com",
		Username:                            "testuser",
		Password:                            "testpass",
		EnableServiceLookup:                 false,
		EnableParticipantLookup:             true,
		EnableRegistrationLookup:            true,
		EnableDirectoryLookup:               false,
		EnableAvatarLookup:                  true,
		EnableMediaLocationLookup:           false,
		EnableInternalServicePolicy:         true,
		EnableInternalParticipantPolicy:     false,
		EnableInternalMediaLocationPolicy:   false,
		InternalServicePolicyTemplate:       "service_template",
		InternalParticipantPolicyTemplate:   "participant_template",
		InternalMediaLocationPolicyTemplate: "media_template",
		PreferLocalAvatarConfiguration:      true,
	}

	client.On("PutJSON", t.Context(), "configuration/v1/policy_server/1/", updateRequest, mock.AnythingOfType("*config.PolicyServer")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*PolicyServer)
		*result = *expectedPolicyServer
	})

	service := New(client)
	result, err := service.UpdatePolicyServer(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedPolicyServer, result)
	client.AssertExpectations(t)
}

func TestService_DeletePolicyServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/policy_server/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeletePolicyServer(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
