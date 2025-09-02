/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package status

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListParticipants(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	expectedResponse := &ParticipantListResponse{
		Objects: []Participant{
			{
				ID:                      "participant-1",
				Bandwidth:               intPtr(1600),
				CallDirection:           "in",
				CallQuality:             "0_unknown",
				CallTag:                 "",
				CallUUID:                "uuid-1",
				Conference:              "conf-1",
				ConnectTime:             &util.InfinityTime{},
				ConversationID:          "uuid-1",
				DestinationAlias:        "alice@example.com",
				DisplayName:             "John Doe",
				Encryption:              "On",
				HasMedia:                true,
				IdpUUID:                 "",
				IsClientMuted:           false,
				IsDirect:                false,
				IsDisconnectSupported:   true,
				IsIdpAuthenticated:      false,
				IsMuteSupported:         true,
				IsMuted:                 false,
				IsOnHold:                false,
				IsPresentationSupported: true,
				IsPresenting:            false,
				IsRecording:             false,
				IsStreaming:             false,
				IsTranscribing:          false,
				IsTransferSupported:     true,
				LicenseCount:            1,
				LicenseType:             "port",
				MediaNode:               "172.27.1.31",
				ParentID:                "",
				ParticipantAlias:        "John Doe",
				Protocol:                "WebRTC",
				ProxyNode:               "",
				RemoteAddress:           "172.24.8.2",
				RemotePort:              62410,
				ResourceURI:             "/api/admin/status/v1/participant/participant-1/",
				Role:                    "chair",
				RxBandwidth:             intPtr(3381),
				ServiceTag:              "cvp-dev",
				ServiceType:             "conference",
				SignallingNode:          "172.27.1.21",
				SourceAlias:             "John Doe",
				SystemLocation:          "LON1",
				TranscodingEnabled:      true,
				TxBandwidth:             intPtr(1600),
				Vendor:                  "Mozilla/5.0",
			},
			{
				ID:                      "participant-2",
				Bandwidth:               intPtr(0),
				CallDirection:           "in",
				CallQuality:             "0_unknown",
				CallTag:                 "",
				CallUUID:                "uuid-2",
				Conference:              "conf-2",
				ConnectTime:             &util.InfinityTime{},
				ConversationID:          "uuid-2",
				DestinationAlias:        "bob@example.com",
				DisplayName:             "Jane Smith",
				Encryption:              "On",
				HasMedia:                false,
				IdpUUID:                 "",
				IsClientMuted:           false,
				IsDirect:                false,
				IsDisconnectSupported:   true,
				IsIdpAuthenticated:      false,
				IsMuteSupported:         true,
				IsMuted:                 true,
				IsOnHold:                false,
				IsPresentationSupported: false,
				IsPresenting:            false,
				IsRecording:             false,
				IsStreaming:             false,
				IsTranscribing:          false,
				IsTransferSupported:     true,
				LicenseCount:            0,
				LicenseType:             "nolicense",
				MediaNode:               "172.27.1.31",
				ParentID:                "",
				ParticipantAlias:        "Jane Smith",
				Protocol:                "WebRTC",
				ProxyNode:               "",
				RemoteAddress:           "172.24.8.2",
				RemotePort:              62370,
				ResourceURI:             "/api/admin/status/v1/participant/participant-2/",
				Role:                    "guest",
				RxBandwidth:             intPtr(0),
				ServiceTag:              "cvp-dev",
				ServiceType:             "conference",
				SignallingNode:          "172.27.1.21",
				SourceAlias:             "Jane Smith",
				SystemLocation:          "LON1",
				TranscodingEnabled:      true,
				TxBandwidth:             intPtr(0),
				Vendor:                  "Viju Pexip agent",
			},
		},
	}

	client.On("GetJSON", t.Context(), "status/v1/participant/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*status.ParticipantListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*ParticipantListResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ListParticipants(t.Context(), nil)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result.Objects))
	assert.Equal(t, "John Doe", result.Objects[0].DisplayName)
	assert.Equal(t, "chair", result.Objects[0].Role)
	assert.Equal(t, "Jane Smith", result.Objects[1].DisplayName)
	assert.Equal(t, "guest", result.Objects[1].Role)
	client.AssertExpectations(t)
}

func TestService_ListParticipants_WithOptions(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	opts := &ListOptions{
		Limit:  20,
		Offset: 10,
	}

	expectedResponse := &ParticipantListResponse{
		Objects: []Participant{
			{
				ID:                      "participant-3",
				Bandwidth:               intPtr(800),
				CallDirection:           "out",
				CallQuality:             "good",
				CallTag:                 "",
				CallUUID:                "uuid-3",
				Conference:              "conf-3",
				ConnectTime:             &util.InfinityTime{},
				ConversationID:          "uuid-3",
				DestinationAlias:        "carol@example.com",
				DisplayName:             "Alice Smith",
				Encryption:              "On",
				HasMedia:                true,
				IdpUUID:                 "",
				IsClientMuted:           false,
				IsDirect:                false,
				IsDisconnectSupported:   true,
				IsIdpAuthenticated:      false,
				IsMuteSupported:         true,
				IsMuted:                 false,
				IsOnHold:                false,
				IsPresentationSupported: false,
				IsPresenting:            false,
				IsRecording:             false,
				IsStreaming:             false,
				IsTranscribing:          false,
				IsTransferSupported:     true,
				LicenseCount:            1,
				LicenseType:             "port",
				MediaNode:               "172.27.1.32",
				ParentID:                "",
				ParticipantAlias:        "Alice Smith",
				Protocol:                "WebRTC",
				ProxyNode:               "",
				RemoteAddress:           "172.24.8.3",
				RemotePort:              62500,
				ResourceURI:             "/api/admin/status/v1/participant/participant-3/",
				Role:                    "guest",
				RxBandwidth:             intPtr(1200),
				ServiceTag:              "cvp-dev",
				ServiceType:             "conference",
				SignallingNode:          "172.27.1.22",
				SourceAlias:             "Alice Smith",
				SystemLocation:          "LON2",
				TranscodingEnabled:      true,
				TxBandwidth:             intPtr(800),
				Vendor:                  "Mozilla/5.0",
			},
		},
	}

	client.On("GetJSON", t.Context(), "status/v1/participant/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*status.ParticipantListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*ParticipantListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListParticipants(t.Context(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "Alice Smith", result.Objects[0].DisplayName)
	assert.Equal(t, "guest", result.Objects[0].Role)

	client.AssertExpectations(t)
}

func TestService_GetParticipant(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedParticipant := &Participant{
		ID:                      "participant-1",
		Bandwidth:               intPtr(1600),
		CallDirection:           "in",
		CallQuality:             "0_unknown",
		CallTag:                 "",
		CallUUID:                "uuid-1",
		Conference:              "conf-1",
		ConnectTime:             &util.InfinityTime{},
		ConversationID:          "uuid-1",
		DestinationAlias:        "alice@example.com",
		DisplayName:             "John Doe",
		Encryption:              "On",
		HasMedia:                true,
		IdpUUID:                 "",
		IsClientMuted:           false,
		IsDirect:                false,
		IsDisconnectSupported:   true,
		IsIdpAuthenticated:      false,
		IsMuteSupported:         true,
		IsMuted:                 false,
		IsOnHold:                false,
		IsPresentationSupported: true,
		IsPresenting:            false,
		IsRecording:             false,
		IsStreaming:             false,
		IsTranscribing:          false,
		IsTransferSupported:     true,
		LicenseCount:            1,
		LicenseType:             "port",
		MediaNode:               "172.27.1.31",
		ParentID:                "",
		ParticipantAlias:        "John Doe",
		Protocol:                "WebRTC",
		ProxyNode:               "",
		RemoteAddress:           "172.24.8.2",
		RemotePort:              62410,
		ResourceURI:             "/api/admin/status/v1/participant/participant-1/",
		Role:                    "chair",
		RxBandwidth:             intPtr(3381),
		ServiceTag:              "cvp-dev",
		ServiceType:             "conference",
		SignallingNode:          "172.27.1.21",
		SourceAlias:             "John Doe",
		SystemLocation:          "LON1",
		TranscodingEnabled:      true,
		TxBandwidth:             intPtr(1600),
		Vendor:                  "Mozilla/5.0",
	}

	client.On("GetJSON", t.Context(), "status/v1/participant/participant-1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*status.Participant")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*Participant)
		*result = *expectedParticipant
	})

	service := New(client)
	result, err := service.GetParticipant(t.Context(), "participant-1")

	assert.NoError(t, err)
	assert.Equal(t, expectedParticipant, result)
	client.AssertExpectations(t)
}
