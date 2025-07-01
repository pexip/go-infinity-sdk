package config

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListTelehealthProfiles(t *testing.T) {
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
				expectedResponse := &TelehealthProfileListResponse{
					Objects: []TelehealthProfile{
						{ID: 1, Name: "epic-profile-1", Description: "Epic Telehealth Profile 1", UUID: "123e4567-e89b-12d3-a456-426614174000", TelehealthCallDomain: "telehealth.example.com", EpicEncryptionAlgorithm: "AES256", EpicEncryptionKeyType: "direct"},
						{ID: 2, Name: "epic-profile-2", Description: "Epic Telehealth Profile 2", UUID: "123e4567-e89b-12d3-a456-426614174001", TelehealthCallDomain: "telehealth2.example.com", EpicEncryptionAlgorithm: "AES128", EpicEncryptionKeyType: "jwks"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/telehealth_profile/", mock.AnythingOfType("*config.TelehealthProfileListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*TelehealthProfileListResponse)
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
				Search: "epic-profile-1",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				expectedResponse := &TelehealthProfileListResponse{
					Objects: []TelehealthProfile{
						{ID: 1, Name: "epic-profile-1", Description: "Epic Telehealth Profile 1", UUID: "123e4567-e89b-12d3-a456-426614174000", TelehealthCallDomain: "telehealth.example.com", EpicEncryptionAlgorithm: "AES256", EpicEncryptionKeyType: "direct"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/telehealth_profile/?limit=5&name__icontains=epic-profile-1", mock.AnythingOfType("*config.TelehealthProfileListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*TelehealthProfileListResponse)
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
			result, err := service.ListTelehealthProfiles(t.Context(), tt.opts)

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

func TestService_GetTelehealthProfile(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedTelehealthProfile := &TelehealthProfile{
		ID:                                    1,
		Name:                                  "test-epic-profile",
		Description:                           "Test Epic Telehealth Profile",
		UUID:                                  "123e4567-e89b-12d3-a456-426614174000",
		TelehealthCallDomain:                  "telehealth.example.com",
		TelehealthIntegrationBaseURL:          "https://api.example.com",
		TelehealthIntegrationOauth2BaseAPIURL: "https://oauth.example.com",
		InfinityWebappServerBaseURL:           "https://webapp.example.com",
		EpicPatientAppClientID:                "patient-client-id",
		EpicProviderAppClientID:               "provider-client-id",
		EpicEncryptionAlgorithm:               "AES256",
		EpicEncryptionKeyType:                 "direct",
		PatientOauth2RedirectURL:              "https://example.com/patient/redirect",
		ProviderOauth2RedirectURL:             "https://example.com/provider/redirect",
		ServiceNameTemplate:                   "Telehealth Meeting {appointment_id}",
		PatientAliasTemplate:                  "patient-{patient_id}",
		ProviderAliasTemplate:                 "provider-{provider_id}",
		PatientDisplayNameTemplate:            "{patient_name}",
		ProviderDisplayNameTemplate:           "Dr. {provider_name}",
		PatientWebJoinLinkTemplate:            "https://meet.example.com/patient/{meeting_id}",
		ProviderWebJoinLinkTemplate:           "https://meet.example.com/provider/{meeting_id}",
		ResourceURI:                           "/api/admin/configuration/v1/telehealth_profile/1/",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/telehealth_profile/1/", mock.AnythingOfType("*config.TelehealthProfile")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*TelehealthProfile)
		*result = *expectedTelehealthProfile
	})

	service := New(client)
	result, err := service.GetTelehealthProfile(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedTelehealthProfile, result)
	client.AssertExpectations(t)
}

func TestService_CreateTelehealthProfile(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &TelehealthProfileCreateRequest{
		Name:                                  "new-epic-profile",
		Description:                           "New Epic Telehealth Profile",
		UUID:                                  "123e4567-e89b-12d3-a456-426614174002",
		TelehealthCallDomain:                  "new-telehealth.example.com",
		TelehealthIntegrationBaseURL:          "https://new-api.example.com",
		TelehealthIntegrationOauth2BaseAPIURL: "https://new-oauth.example.com",
		InfinityWebappServerBaseURL:           "https://new-webapp.example.com",
		EpicPatientAppClientID:                "new-patient-client-id",
		EpicProviderAppClientID:               "new-provider-client-id",
		EpicEncryptionAlgorithm:               "AES256",
		EpicEncryptionKeyType:                 "direct",
		PatientOauth2RedirectURL:              "https://new.example.com/patient/redirect",
		ProviderOauth2RedirectURL:             "https://new.example.com/provider/redirect",
		ServiceNameTemplate:                   "New Telehealth Meeting {appointment_id}",
		PatientAliasTemplate:                  "new-patient-{patient_id}",
		ProviderAliasTemplate:                 "new-provider-{provider_id}",
		PatientDisplayNameTemplate:            "{patient_name}",
		ProviderDisplayNameTemplate:           "Dr. {provider_name}",
		PatientWebJoinLinkTemplate:            "https://new-meet.example.com/patient/{meeting_id}",
		ProviderWebJoinLinkTemplate:           "https://new-meet.example.com/provider/{meeting_id}",
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/telehealth_profile/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/telehealth_profile/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateTelehealthProfile(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateTelehealthProfile(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	updateRequest := &TelehealthProfileUpdateRequest{
		Description:          "Updated Epic Telehealth Profile",
		TelehealthCallDomain: "updated-telehealth.example.com",
		ServiceNameTemplate:  "Updated Meeting {appointment_id}",
	}

	expectedTelehealthProfile := &TelehealthProfile{
		ID:                                    1,
		Name:                                  "test-epic-profile",
		Description:                           "Updated Epic Telehealth Profile",
		UUID:                                  "123e4567-e89b-12d3-a456-426614174000",
		TelehealthCallDomain:                  "updated-telehealth.example.com",
		TelehealthIntegrationBaseURL:          "https://api.example.com",
		TelehealthIntegrationOauth2BaseAPIURL: "https://oauth.example.com",
		InfinityWebappServerBaseURL:           "https://webapp.example.com",
		EpicPatientAppClientID:                "patient-client-id",
		EpicProviderAppClientID:               "provider-client-id",
		EpicEncryptionAlgorithm:               "AES256",
		EpicEncryptionKeyType:                 "direct",
		PatientOauth2RedirectURL:              "https://example.com/patient/redirect",
		ProviderOauth2RedirectURL:             "https://example.com/provider/redirect",
		ServiceNameTemplate:                   "Updated Meeting {appointment_id}",
		PatientAliasTemplate:                  "patient-{patient_id}",
		ProviderAliasTemplate:                 "provider-{provider_id}",
		PatientDisplayNameTemplate:            "{patient_name}",
		ProviderDisplayNameTemplate:           "Dr. {provider_name}",
		PatientWebJoinLinkTemplate:            "https://meet.example.com/patient/{meeting_id}",
		ProviderWebJoinLinkTemplate:           "https://meet.example.com/provider/{meeting_id}",
		ResourceURI:                           "/api/admin/configuration/v1/telehealth_profile/1/",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/telehealth_profile/1/", updateRequest, mock.AnythingOfType("*config.TelehealthProfile")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*TelehealthProfile)
		*result = *expectedTelehealthProfile
	})

	service := New(client)
	result, err := service.UpdateTelehealthProfile(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedTelehealthProfile, result)
	client.AssertExpectations(t)
}

func TestService_DeleteTelehealthProfile(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/telehealth_profile/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteTelehealthProfile(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
