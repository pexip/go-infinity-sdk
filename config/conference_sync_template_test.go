/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListConferenceSyncTemplates(t *testing.T) {
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
				expectedResponse := &ConferenceSyncTemplateListResponse{
					Objects: []ConferenceSyncTemplate{
						{ID: 1, Name: "primary-template", Description: "Primary sync template", LdapUserFilter: "(objectClass=user)", EnableAutomaticSync: true, SyncConferences: true, SyncDevices: false, SyncEndUsers: true, ServiceType: "conference", CallType: "video", AllowGuests: true},
						{ID: 2, Name: "secondary-template", Description: "Secondary sync template", LdapUserFilter: "(objectClass=person)", EnableAutomaticSync: false, SyncConferences: false, SyncDevices: true, SyncEndUsers: false, ServiceType: "virtual_meeting_room", CallType: "audio", AllowGuests: false},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/conference_sync_template/", mock.AnythingOfType("*config.ConferenceSyncTemplateListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*ConferenceSyncTemplateListResponse)
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
				expectedResponse := &ConferenceSyncTemplateListResponse{
					Objects: []ConferenceSyncTemplate{
						{ID: 1, Name: "primary-template", Description: "Primary sync template", LdapUserFilter: "(objectClass=user)", EnableAutomaticSync: true, SyncConferences: true, SyncDevices: false, SyncEndUsers: true, ServiceType: "conference", CallType: "video", AllowGuests: true},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/conference_sync_template/?limit=5&name__icontains=primary", mock.AnythingOfType("*config.ConferenceSyncTemplateListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*ConferenceSyncTemplateListResponse)
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
			result, err := service.ListConferenceSyncTemplates(t.Context(), tt.opts)

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

func TestService_GetConferenceSyncTemplate(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	maxCallrateIn := 2048
	maxCallrateOut := 1024
	participantLimit := 50
	expectedConferenceSyncTemplate := &ConferenceSyncTemplate{
		ID:                              1,
		Name:                            "test-template",
		Description:                     "Test sync template",
		LdapSyncSource:                  "/api/admin/configuration/v1/ldap_sync_source/1/",
		LdapUserFilter:                  "(objectClass=user)",
		LdapUserSearchDN:                "ou=users,dc=example,dc=com",
		EnableAutomaticSync:             true,
		SyncConferences:                 true,
		SyncDevices:                     false,
		SyncEndUsers:                    true,
		EnableServiceEmails:             false,
		ServiceType:                     "conference",
		Tag:                             "test-tag",
		TagOverridable:                  true,
		PIN:                             "1234",
		GuestPIN:                        "5678",
		PINSettingsOverridable:          false,
		AllowGuests:                     true,
		CallType:                        "video",
		CallTypeOverridable:             true,
		CryptoMode:                      "auto",
		CryptoModeOverridable:           false,
		EnableChat:                      "yes",
		EnableChatOverridable:           true,
		EnableOverlayText:               false,
		EnableActiveSpeakerIndication:   true,
		HostView:                        "one_main_seven_pips",
		HostViewOverridable:             false,
		GuestsCanPresent:                true,
		GuestsCanPresentOverridable:     false,
		DirectMedia:                     "yes",
		DirectMediaOverridable:          true,
		DirectMediaNotificationDuration: 10,
		DirectMediaNotificationDurationOverridable: false,
		MaxCallrateIn:                       &maxCallrateIn,
		MaxCallrateOut:                      &maxCallrateOut,
		CallratesOverridable:                true,
		MaxPixelsPerSecond:                  "1966080",
		MaxPixelsPerSecondOverridable:       false,
		ParticipantLimit:                    &participantLimit,
		ParticipantLimitOverridable:         true,
		PrimaryOwnerEmailAddress:            "admin@example.com",
		PrimaryOwnerEmailAddressOverridable: false,
		NonIdpParticipants:                  "allow",
		IdpSettingsOverridable:              true,
		Alias1:                              "{{conference_name}}",
		Alias1Description:                   "Primary alias",
		AliasesOverridable:                  false,
		DeviceUsername:                      "{{device_username}}",
		DeviceUsernameOverridable:           true,
		DevicePassword:                      "{{device_password}}",
		DevicePasswordOverridable:           false,
		DeviceAlias:                         "{{device_alias}}",
		DeviceDescription:                   "{{device_description}}",
		DeviceDescriptionOverridable:        true,
		DeviceEnableSIP:                     true,
		DeviceEnableH323:                    false,
		DeviceEnableInfinityConnectNonSSO:   true,
		DeviceEnableInfinityConnectSSO:      false,
		DeviceEnableStandardSSO:             false,
		DeviceRegistrationTypesOverridable:  true,
		DeviceSyncIfAccountDisabled:         false,
		EndUserUUID:                         "{{end_user_uuid}}",
		EndUserFirstName:                    "{{end_user_first_name}}",
		EndUserLastName:                     "{{end_user_last_name}}",
		EndUserDisplayName:                  "{{end_user_display_name}}",
		EndUserNamesOverridable:             true,
		EndUserDescription:                  "{{end_user_description}}",
		EndUserDescriptionOverridable:       false,
		EndUserTelephoneNumber:              "{{end_user_telephone_number}}",
		EndUserMobileNumber:                 "{{end_user_mobile_number}}",
		EndUserContactsOverridable:          true,
		EndUserAdvancedOverridable:          false,
	}

	client.On("GetJSON", t.Context(), "configuration/v1/conference_sync_template/1/", mock.AnythingOfType("*config.ConferenceSyncTemplate")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*ConferenceSyncTemplate)
		*result = *expectedConferenceSyncTemplate
	})

	service := New(client)
	result, err := service.GetConferenceSyncTemplate(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedConferenceSyncTemplate, result)
	client.AssertExpectations(t)
}

func TestService_CreateConferenceSyncTemplate(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	maxCallrateIn := 1024
	participantLimit := 25
	createRequest := &ConferenceSyncTemplateCreateRequest{
		Name:                              "new-template",
		Description:                       "New sync template",
		LdapUserFilter:                    "(objectClass=person)",
		EnableAutomaticSync:               true,
		SyncConferences:                   true,
		SyncDevices:                       true,
		SyncEndUsers:                      true,
		EnableServiceEmails:               false,
		ServiceType:                       "virtual_meeting_room",
		Tag:                               "new-tag",
		TagOverridable:                    false,
		PIN:                               "9999",
		AllowGuests:                       false,
		CallType:                          "audio",
		CallTypeOverridable:               false,
		EnableChat:                        "no",
		EnableOverlayText:                 true,
		EnableActiveSpeakerIndication:     false,
		HostView:                          "one_main_zero_pips",
		GuestsCanPresent:                  false,
		DirectMedia:                       "no",
		DirectMediaNotificationDuration:   5,
		MaxCallrateIn:                     &maxCallrateIn,
		ParticipantLimit:                  &participantLimit,
		CallratesOverridable:              false,
		PrimaryOwnerEmailAddress:          "newadmin@example.com",
		NonIdpParticipants:                "deny",
		DeviceEnableSIP:                   true,
		DeviceEnableH323:                  true,
		DeviceEnableInfinityConnectNonSSO: false,
		EndUserNamesOverridable:           true,
		EndUserContactsOverridable:        false,
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/conference_sync_template/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/conference_sync_template/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateConferenceSyncTemplate(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateConferenceSyncTemplate(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	enableAutoSync := false
	allowGuests := false
	maxCallrateOut := 512
	updateRequest := &ConferenceSyncTemplateUpdateRequest{
		Description:         "Updated sync template",
		EnableAutomaticSync: &enableAutoSync,
		AllowGuests:         &allowGuests,
		MaxCallrateOut:      &maxCallrateOut,
		CallType:            "audio",
	}

	maxCallrateIn := 2048
	participantLimit := 50
	expectedConferenceSyncTemplate := &ConferenceSyncTemplate{
		ID:                  1,
		Name:                "test-template",
		Description:         "Updated sync template",
		LdapUserFilter:      "(objectClass=user)",
		EnableAutomaticSync: false,
		SyncConferences:     true,
		SyncDevices:         false,
		SyncEndUsers:        true,
		ServiceType:         "conference",
		AllowGuests:         false,
		CallType:            "audio",
		MaxCallrateIn:       &maxCallrateIn,
		MaxCallrateOut:      &maxCallrateOut,
		ParticipantLimit:    &participantLimit,
		DeviceEnableSIP:     true,
		DeviceEnableH323:    false,
	}

	client.On("PutJSON", t.Context(), "configuration/v1/conference_sync_template/1/", updateRequest, mock.AnythingOfType("*config.ConferenceSyncTemplate")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*ConferenceSyncTemplate)
		*result = *expectedConferenceSyncTemplate
	})

	service := New(client)
	result, err := service.UpdateConferenceSyncTemplate(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedConferenceSyncTemplate, result)
	client.AssertExpectations(t)
}

func TestService_DeleteConferenceSyncTemplate(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/conference_sync_template/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteConferenceSyncTemplate(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
