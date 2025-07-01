package config

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListMsExchangeConnectors(t *testing.T) {
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
				roomEmailAddress := "conference.room@example.com"
				scheduledAliasPrefix := "meet"
				oauthClientID := "oauth-client-123"
				oauthState := "oauth-state"
				addinApplicationID := "addin-app-123"
				domains := "/api/admin/configuration/v1/domains/1/"
				ivrTheme := "/api/admin/configuration/v1/ivr_theme/1/"

				expectedResponse := &MsExchangeConnectorListResponse{
					Objects: []MsExchangeConnector{
						{
							ID:                           1,
							Name:                         "primary-exchange-connector",
							Description:                  "Primary Microsoft Exchange connector",
							RoomMailboxEmailAddress:      &roomEmailAddress,
							RoomMailboxName:              "Conference Room",
							URL:                          "https://exchange.example.com/ews/exchange.asmx",
							Username:                     "svc-exchange@example.com",
							AuthenticationMethod:         "oauth",
							AuthProvider:                 "microsoft",
							UUID:                         "12345678-1234-1234-1234-123456789012",
							ScheduledAliasPrefix:         &scheduledAliasPrefix,
							ScheduledAliasDomain:         "meet.example.com",
							ScheduledAliasSuffixLength:   8,
							MeetingBufferBefore:          5,
							MeetingBufferAfter:           5,
							EnableDynamicVmrs:            true,
							EnablePersonalVmrs:           true,
							AllowNewUsers:                true,
							DisableProxy:                 false,
							UseCustomAddInSources:        false,
							EnableAddinDebugLogs:         false,
							OauthClientID:                &oauthClientID,
							OauthAuthEndpoint:            "https://login.microsoftonline.com/common/oauth2/v2.0/authorize",
							OauthTokenEndpoint:           "https://login.microsoftonline.com/common/oauth2/v2.0/token",
							OauthRedirectURI:             "https://pexip.example.com/oauth/callback",
							OauthState:                   &oauthState,
							AddinServerDomain:            "pexip.example.com",
							AddinDisplayName:             "Pexip Scheduler",
							AddinDescription:             "Schedule video meetings",
							AddinProviderName:            "Pexip",
							AddinButtonLabel:             "Add Video Meeting",
							AddinGroupLabel:              "Pexip",
							AddinSupertipTitle:           "Add Pexip Video Meeting",
							AddinSupertipDescription:     "Add a Pexip video meeting to this appointment",
							AddinApplicationID:           &addinApplicationID,
							AddinAuthenticationMethod:    "oauth",
							OfficeJsURL:                  "https://appsforoffice.microsoft.com/lib/1/hosted/office.js",
							MicrosoftFabricURL:           "https://static2.sharepointonline.com/files/fabric/office-ui-fabric-core/11.0.0/css/fabric.min.css",
							MicrosoftFabricComponentsURL: "https://static2.sharepointonline.com/files/fabric/office-ui-fabric-js/1.4.0/js/fabric.min.js",
							Domains:                      &domains,
							IvrTheme:                     &ivrTheme,
							NonIdpParticipants:           "allowed",
						},
						{
							ID:                           2,
							Name:                         "backup-exchange-connector",
							Description:                  "Backup Microsoft Exchange connector",
							AuthenticationMethod:         "kerberos",
							AuthProvider:                 "microsoft",
							UUID:                         "87654321-4321-4321-4321-210987654321",
							ScheduledAliasDomain:         "backup.example.com",
							ScheduledAliasSuffixLength:   6,
							MeetingBufferBefore:          3,
							MeetingBufferAfter:           3,
							EnableDynamicVmrs:            false,
							EnablePersonalVmrs:           false,
							AllowNewUsers:                false,
							DisableProxy:                 true,
							KerberosRealm:                "EXAMPLE.COM",
							KerberosKdc:                  "kdc.example.com",
							KerberosExchangeSpn:          "HTTP/exchange.example.com",
							KerberosEnableTls:            true,
							KerberosAuthEveryRequest:     false,
							AddinServerDomain:            "backup.example.com",
							AddinDisplayName:             "Backup Pexip Scheduler",
							OfficeJsURL:                  "https://appsforoffice.microsoft.com/lib/1/hosted/office.js",
							MicrosoftFabricURL:           "https://static2.sharepointonline.com/files/fabric/office-ui-fabric-core/11.0.0/css/fabric.min.css",
							MicrosoftFabricComponentsURL: "https://static2.sharepointonline.com/files/fabric/office-ui-fabric-js/1.4.0/js/fabric.min.js",
							NonIdpParticipants:           "denied",
						},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/ms_exchange_connector/", mock.AnythingOfType("*config.MsExchangeConnectorListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*MsExchangeConnectorListResponse)
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
				roomEmailAddress := "conference.room@example.com"
				oauthClientID := "oauth-client-123"

				expectedResponse := &MsExchangeConnectorListResponse{
					Objects: []MsExchangeConnector{
						{
							ID:                      1,
							Name:                    "primary-exchange-connector",
							Description:             "Primary Microsoft Exchange connector",
							RoomMailboxEmailAddress: &roomEmailAddress,
							AuthenticationMethod:    "oauth",
							AuthProvider:            "microsoft",
							UUID:                    "12345678-1234-1234-1234-123456789012",
							OauthClientID:           &oauthClientID,
							EnableDynamicVmrs:       true,
							EnablePersonalVmrs:      true,
						},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/ms_exchange_connector/?limit=5&name__icontains=primary", mock.AnythingOfType("*config.MsExchangeConnectorListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*MsExchangeConnectorListResponse)
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
			result, err := service.ListMsExchangeConnectors(t.Context(), tt.opts)

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

func TestService_GetMsExchangeConnector(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	roomEmailAddress := "test.room@example.com"
	scheduledAliasPrefix := "test"
	oauthClientID := "test-oauth-client"
	oauthState := "test-oauth-state"
	addinApplicationID := "test-addin-app"
	addinNaaWebApiAppID := "test-naa-app"
	personalVmrOauthClientID := "test-personal-vmr-client"
	domains := "/api/admin/configuration/v1/domains/1/"
	hostIdpGroup := "/api/admin/configuration/v1/identity_provider_group/1/"
	ivrTheme := "/api/admin/configuration/v1/ivr_theme/1/"

	expectedConnector := &MsExchangeConnector{
		ID:                         1,
		Name:                       "test-exchange-connector",
		Description:                "Test Microsoft Exchange connector",
		RoomMailboxEmailAddress:    &roomEmailAddress,
		RoomMailboxName:            "Test Conference Room",
		URL:                        "https://test-exchange.example.com/ews/exchange.asmx",
		Username:                   "svc-test@example.com",
		Password:                   "test-password",
		AuthenticationMethod:       "oauth",
		AuthProvider:               "microsoft",
		UUID:                       "test-uuid-1234-5678-9012",
		ScheduledAliasPrefix:       &scheduledAliasPrefix,
		ScheduledAliasDomain:       "test.example.com",
		ScheduledAliasSuffixLength: 10,
		MeetingBufferBefore:        10,
		MeetingBufferAfter:         10,
		EnableDynamicVmrs:          true,
		EnablePersonalVmrs:         true,
		AllowNewUsers:              true,
		DisableProxy:               false,
		UseCustomAddInSources:      true,
		EnableAddinDebugLogs:       true,
		// OAuth fields
		OauthClientID:      &oauthClientID,
		OauthClientSecret:  "test-oauth-secret",
		OauthAuthEndpoint:  "https://login.microsoftonline.com/test/oauth2/v2.0/authorize",
		OauthTokenEndpoint: "https://login.microsoftonline.com/test/oauth2/v2.0/token",
		OauthRedirectURI:   "https://test.example.com/oauth/callback",
		OauthRefreshToken:  "test-refresh-token",
		OauthState:         &oauthState,
		// Kerberos fields
		KerberosRealm:                  "TEST.EXAMPLE.COM",
		KerberosKdc:                    "kdc.test.example.com",
		KerberosKdcHttpsProxy:          "proxy.test.example.com:8080",
		KerberosExchangeSpn:            "HTTP/test-exchange.example.com",
		KerberosEnableTls:              true,
		KerberosAuthEveryRequest:       false,
		KerberosVerifyTlsUsingCustomCa: false,
		// Add-in fields
		AddinServerDomain:           "test.example.com",
		AddinDisplayName:            "Test Pexip Scheduler",
		AddinDescription:            "Test video meeting scheduler",
		AddinProviderName:           "Test Pexip",
		AddinButtonLabel:            "Add Test Video Meeting",
		AddinGroupLabel:             "Test Pexip",
		AddinSupertipTitle:          "Add Test Video Meeting",
		AddinSupertipDescription:    "Add a test video meeting to this appointment",
		AddinApplicationID:          &addinApplicationID,
		AddinAuthorityURL:           "https://login.microsoftonline.com/test",
		AddinOidcMetadataURL:        "https://login.microsoftonline.com/test/v2.0/.well-known/openid_configuration",
		AddinAuthenticationMethod:   "oauth",
		AddinNaaWebApiApplicationID: &addinNaaWebApiAppID,
		// Add-in pane fields
		AddinPaneTitle:                                   "Test Video Meeting",
		AddinPaneDescription:                             "Add a test video meeting",
		AddinPaneButtonTitle:                             "Add Video Meeting",
		AddinPaneSuccessHeading:                          "Success",
		AddinPaneSuccessMessage:                          "Video meeting added successfully",
		AddinPaneAlreadyVideoMeetingHeading:              "Already has video meeting",
		AddinPaneAlreadyVideoMeetingMessage:              "This meeting already has video",
		AddinPaneGeneralErrorHeading:                     "Error",
		AddinPaneGeneralErrorMessage:                     "An error occurred",
		AddinPaneManagementNodeDownHeading:               "Service Unavailable",
		AddinPaneManagementNodeDownMessage:               "Management node is down",
		AddinPanePersonalVmrAddButton:                    "Add Personal VMR",
		AddinPanePersonalVmrSignInButton:                 "Sign In",
		AddinPanePersonalVmrSelectMessage:                "Select your personal VMR",
		AddinPanePersonalVmrNoneMessage:                  "No personal VMRs available",
		AddinPanePersonalVmrErrorGettingMessage:          "Error getting personal VMRs",
		AddinPanePersonalVmrErrorSigningInMessage:        "Error signing in",
		AddinPanePersonalVmrErrorInsertingMeetingMessage: "Error inserting meeting",
		// Personal VMR OAuth fields
		PersonalVmrOauthClientID:                   &personalVmrOauthClientID,
		PersonalVmrOauthClientSecret:               "personal-vmr-secret",
		PersonalVmrOauthAuthEndpoint:               "https://login.microsoftonline.com/test/oauth2/v2.0/authorize",
		PersonalVmrOauthTokenEndpoint:              "https://login.microsoftonline.com/test/oauth2/v2.0/token",
		PersonalVmrAdfsRelyingPartyTrustIdentifier: "test-adfs-identifier",
		// Template fields
		MeetingInstructionsTemplate:         "Join meeting: {{meeting_url}}",
		PersonalVmrInstructionsTemplate:     "Personal VMR: {{vmr_url}}",
		PersonalVmrLocationTemplate:         "{{vmr_url}}",
		PersonalVmrNameTemplate:             "{{user_name}}'s Personal VMR",
		PersonalVmrDescriptionTemplate:      "Personal VMR for {{user_name}}",
		PlaceholderInstructionsTemplate:     "Video meeting details will be added",
		ConferenceNameTemplate:              "{{subject}}",
		ConferenceDescriptionTemplate:       "Conference for {{subject}}",
		ConferenceSubjectTemplate:           "Video Conference: {{subject}}",
		ScheduledAliasDescriptionTemplate:   "Scheduled meeting: {{subject}}",
		AcceptNewSingleMeetingTemplate:      "Accepted: {{subject}}",
		AcceptNewRecurringSeriesTemplate:    "Accepted recurring: {{subject}}",
		AcceptEditedSingleMeetingTemplate:   "Accepted edited: {{subject}}",
		AcceptEditedRecurringSeriesTemplate: "Accepted edited series: {{subject}}",
		AcceptEditedOccurrenceTemplate:      "Accepted edited occurrence: {{subject}}",
		RejectGeneralErrorTemplate:          "Error: Cannot process meeting",
		RejectAliasConflictTemplate:         "Error: Alias conflict",
		RejectAliasDeletedTemplate:          "Error: Alias deleted",
		RejectInvalidAliasIDTemplate:        "Error: Invalid alias ID",
		RejectSingleMeetingPast:             "Error: Meeting is in the past",
		RejectRecurringSeriesPastTemplate:   "Error: Series is in the past",
		// JavaScript and CSS URLs
		OfficeJsURL:                  "https://appsforoffice.microsoft.com/lib/1/hosted/office.js",
		MicrosoftFabricURL:           "https://static2.sharepointonline.com/files/fabric/office-ui-fabric-core/11.0.0/css/fabric.min.css",
		MicrosoftFabricComponentsURL: "https://static2.sharepointonline.com/files/fabric/office-ui-fabric-js/1.4.0/js/fabric.min.js",
		AdditionalAddInScriptSources: "https://test.example.com/custom.js",
		// Related resources
		Domains:                   &domains,
		HostIdentityProviderGroup: &hostIdpGroup,
		IvrTheme:                  &ivrTheme,
		NonIdpParticipants:        "allowed",
		PublicKey:                 "test-public-key",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/ms_exchange_connector/1/", mock.AnythingOfType("*config.MsExchangeConnector")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*MsExchangeConnector)
		*result = *expectedConnector
	})

	service := New(client)
	result, err := service.GetMsExchangeConnector(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedConnector, result)
	client.AssertExpectations(t)
}

func TestService_CreateMsExchangeConnector(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	roomEmailAddress := "new.room@example.com"
	scheduledAliasPrefix := "new"
	oauthClientID := "new-oauth-client"
	addinApplicationID := "new-addin-app"
	domains := "/api/admin/configuration/v1/domains/1/"
	ivrTheme := "/api/admin/configuration/v1/ivr_theme/1/"

	createRequest := &MsExchangeConnectorCreateRequest{
		Name:                         "new-exchange-connector",
		Description:                  "New Microsoft Exchange connector",
		RoomMailboxEmailAddress:      &roomEmailAddress,
		RoomMailboxName:              "New Conference Room",
		URL:                          "https://new-exchange.example.com/ews/exchange.asmx",
		Username:                     "svc-new@example.com",
		Password:                     "new-password",
		AuthenticationMethod:         "oauth",
		AuthProvider:                 "microsoft",
		UUID:                         "new-uuid-1234-5678-9012",
		ScheduledAliasPrefix:         &scheduledAliasPrefix,
		ScheduledAliasDomain:         "new.example.com",
		ScheduledAliasSuffixLength:   8,
		MeetingBufferBefore:          5,
		MeetingBufferAfter:           5,
		EnableDynamicVmrs:            true,
		EnablePersonalVmrs:           true,
		AllowNewUsers:                true,
		DisableProxy:                 false,
		UseCustomAddInSources:        false,
		EnableAddinDebugLogs:         false,
		OauthClientID:                &oauthClientID,
		OauthClientSecret:            "new-oauth-secret",
		OauthAuthEndpoint:            "https://login.microsoftonline.com/new/oauth2/v2.0/authorize",
		OauthTokenEndpoint:           "https://login.microsoftonline.com/new/oauth2/v2.0/token",
		OauthRedirectURI:             "https://new.example.com/oauth/callback",
		OauthRefreshToken:            "new-refresh-token",
		AddinServerDomain:            "new.example.com",
		AddinDisplayName:             "New Pexip Scheduler",
		AddinDescription:             "New video meeting scheduler",
		AddinProviderName:            "New Pexip",
		AddinButtonLabel:             "Add New Video Meeting",
		AddinGroupLabel:              "New Pexip",
		AddinSupertipTitle:           "Add New Video Meeting",
		AddinSupertipDescription:     "Add a new video meeting to this appointment",
		AddinApplicationID:           &addinApplicationID,
		AddinAuthenticationMethod:    "oauth",
		AddinPaneTitle:               "New Video Meeting",
		AddinPaneDescription:         "Add a new video meeting",
		AddinPaneButtonTitle:         "Add Video Meeting",
		AddinPaneSuccessHeading:      "Success",
		AddinPaneSuccessMessage:      "Video meeting added successfully",
		OfficeJsURL:                  "https://appsforoffice.microsoft.com/lib/1/hosted/office.js",
		MicrosoftFabricURL:           "https://static2.sharepointonline.com/files/fabric/office-ui-fabric-core/11.0.0/css/fabric.min.css",
		MicrosoftFabricComponentsURL: "https://static2.sharepointonline.com/files/fabric/office-ui-fabric-js/1.4.0/js/fabric.min.js",
		Domains:                      &domains,
		IvrTheme:                     &ivrTheme,
		NonIdpParticipants:           "allowed",
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/ms_exchange_connector/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/ms_exchange_connector/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateMsExchangeConnector(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateMsExchangeConnector(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	meetingBufferBefore := 15
	meetingBufferAfter := 15
	enableDynamicVmrs := false
	disableProxy := true

	updateRequest := &MsExchangeConnectorUpdateRequest{
		Description:         "Updated Microsoft Exchange connector",
		MeetingBufferBefore: &meetingBufferBefore,
		MeetingBufferAfter:  &meetingBufferAfter,
		EnableDynamicVmrs:   &enableDynamicVmrs,
		DisableProxy:        &disableProxy,
		AddinDisplayName:    "Updated Pexip Scheduler",
	}

	roomEmailAddress := "test.room@example.com"
	oauthClientID := "test-oauth-client"
	expectedConnector := &MsExchangeConnector{
		ID:                           1,
		Name:                         "test-exchange-connector",
		Description:                  "Updated Microsoft Exchange connector",
		RoomMailboxEmailAddress:      &roomEmailAddress,
		AuthenticationMethod:         "oauth",
		AuthProvider:                 "microsoft",
		UUID:                         "test-uuid-1234-5678-9012",
		MeetingBufferBefore:          15,
		MeetingBufferAfter:           15,
		EnableDynamicVmrs:            false,
		EnablePersonalVmrs:           true,
		AllowNewUsers:                true,
		DisableProxy:                 true,
		OauthClientID:                &oauthClientID,
		AddinDisplayName:             "Updated Pexip Scheduler",
		OfficeJsURL:                  "https://appsforoffice.microsoft.com/lib/1/hosted/office.js",
		MicrosoftFabricURL:           "https://static2.sharepointonline.com/files/fabric/office-ui-fabric-core/11.0.0/css/fabric.min.css",
		MicrosoftFabricComponentsURL: "https://static2.sharepointonline.com/files/fabric/office-ui-fabric-js/1.4.0/js/fabric.min.js",
		NonIdpParticipants:           "allowed",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/ms_exchange_connector/1/", updateRequest, mock.AnythingOfType("*config.MsExchangeConnector")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*MsExchangeConnector)
		*result = *expectedConnector
	})

	service := New(client)
	result, err := service.UpdateMsExchangeConnector(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedConnector, result)
	client.AssertExpectations(t)
}

func TestService_DeleteMsExchangeConnector(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/ms_exchange_connector/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteMsExchangeConnector(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
