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

func TestService_ListGatewayRoutingRules(t *testing.T) {
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
				sipProxy := "/api/admin/configuration/v1/sip_proxy/1/"
				location := "/api/admin/configuration/v1/system_location/1/"
				expectedResponse := &GatewayRoutingRuleListResponse{
					Objects: []GatewayRoutingRule{
						{ID: 1, Name: "primary-rule", Description: "Primary routing rule", Priority: 100, Enable: true, MatchString: ".*", MatchStringFull: false, CalledDeviceType: "external", OutgoingProtocol: "sip", CallType: "video", MatchIncomingCalls: true, MatchOutgoingCalls: false, MatchIncomingSIP: true, SIPProxy: &sipProxy, OutgoingLocation: &location, TreatAsTrusted: false},
						{ID: 2, Name: "secondary-rule", Description: "Secondary routing rule", Priority: 200, Enable: false, MatchString: "test@.*", MatchStringFull: true, CalledDeviceType: "internal", OutgoingProtocol: "h323", CallType: "audio", MatchIncomingCalls: false, MatchOutgoingCalls: true, MatchIncomingH323: true, TreatAsTrusted: true},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/gateway_routing_rule/", mock.AnythingOfType("*config.GatewayRoutingRuleListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*GatewayRoutingRuleListResponse)
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
				sipProxy := "/api/admin/configuration/v1/sip_proxy/1/"
				location := "/api/admin/configuration/v1/system_location/1/"
				expectedResponse := &GatewayRoutingRuleListResponse{
					Objects: []GatewayRoutingRule{
						{ID: 1, Name: "primary-rule", Description: "Primary routing rule", Priority: 100, Enable: true, MatchString: ".*", MatchStringFull: false, CalledDeviceType: "external", OutgoingProtocol: "sip", CallType: "video", MatchIncomingCalls: true, MatchOutgoingCalls: false, MatchIncomingSIP: true, SIPProxy: &sipProxy, OutgoingLocation: &location, TreatAsTrusted: false},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/gateway_routing_rule/?limit=5&name__icontains=primary", mock.AnythingOfType("*config.GatewayRoutingRuleListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*GatewayRoutingRuleListResponse)
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
			result, err := service.ListGatewayRoutingRules(t.Context(), tt.opts)

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

func TestService_GetGatewayRoutingRule(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	sipProxy := "/api/admin/configuration/v1/sip_proxy/1/"
	h323Gatekeeper := "/api/admin/configuration/v1/h323_gatekeeper/1/"
	location := "/api/admin/configuration/v1/system_location/1/"
	maxCallrateIn := 2048
	maxCallrateOut := 1024
	cryptoMode := "auto"
	maxPixelsPerSecond := "1966080"

	expectedGatewayRoutingRule := &GatewayRoutingRule{
		ID:                            1,
		Name:                          "test-rule",
		Description:                   "Test gateway routing rule",
		Priority:                      150,
		Enable:                        true,
		MatchString:                   "example@.*",
		MatchStringFull:               true,
		ReplaceString:                 "test-$1",
		CalledDeviceType:              "external",
		OutgoingProtocol:              "sip",
		CallType:                      "video",
		MatchIncomingCalls:            true,
		MatchOutgoingCalls:            false,
		MatchIncomingSIP:              true,
		MatchIncomingH323:             false,
		MatchIncomingMSSIP:            false,
		MatchIncomingWebRTC:           true,
		MatchIncomingTeams:            false,
		MatchIncomingOnlyIfRegistered: false,
		MatchSourceLocation:           &location,
		OutgoingLocation:              &location,
		SIPProxy:                      &sipProxy,
		H323Gatekeeper:                &h323Gatekeeper,
		MaxPixelsPerSecond:            &maxPixelsPerSecond,
		MaxCallrateIn:                 &maxCallrateIn,
		MaxCallrateOut:                &maxCallrateOut,
		CryptoMode:                    &cryptoMode,
		DenoiseAudio:                  false,
		LiveCaptionsEnabled:           "disabled",
		TreatAsTrusted:                true,
		Tag:                           "test-tag",
		DisabledCodecs:                []string{"h264", "vp8"},
	}

	client.On("GetJSON", t.Context(), "configuration/v1/gateway_routing_rule/1/", mock.AnythingOfType("*config.GatewayRoutingRule")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*GatewayRoutingRule)
		*result = *expectedGatewayRoutingRule
	})

	service := New(client)
	result, err := service.GetGatewayRoutingRule(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedGatewayRoutingRule, result)
	client.AssertExpectations(t)
}

func TestService_CreateGatewayRoutingRule(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	sipProxy := "/api/admin/configuration/v1/sip_proxy/1/"
	location := "/api/admin/configuration/v1/system_location/1/"
	maxCallrateIn := 1024
	cryptoMode := "besteffort"

	createRequest := &GatewayRoutingRuleCreateRequest{
		Name:                          "new-rule",
		Description:                   "New gateway routing rule",
		Priority:                      300,
		Enable:                        true,
		MatchString:                   "new@.*",
		MatchStringFull:               false,
		ReplaceString:                 "updated-$1",
		CalledDeviceType:              "internal",
		OutgoingProtocol:              "sip",
		CallType:                      "audio",
		MatchIncomingCalls:            true,
		MatchOutgoingCalls:            true,
		MatchIncomingSIP:              true,
		MatchIncomingH323:             false,
		MatchIncomingMSSIP:            false,
		MatchIncomingWebRTC:           false,
		MatchIncomingTeams:            false,
		MatchIncomingOnlyIfRegistered: true,
		OutgoingLocation:              &location,
		SIPProxy:                      &sipProxy,
		MaxCallrateIn:                 &maxCallrateIn,
		CryptoMode:                    &cryptoMode,
		DenoiseAudio:                  true,
		LiveCaptionsEnabled:           "auto",
		TreatAsTrusted:                false,
		Tag:                           "new-tag",
		DisabledCodecs:                []string{"vp9"},
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/gateway_routing_rule/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/gateway_routing_rule/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateGatewayRoutingRule(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateGatewayRoutingRule(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	priority := 250
	enable := false
	matchIncomingWebRTC := false
	updateRequest := &GatewayRoutingRuleUpdateRequest{
		Description:         "Updated gateway routing rule",
		Priority:            &priority,
		Enable:              &enable,
		MatchIncomingWebRTC: &matchIncomingWebRTC,
		CallType:            "audio",
	}

	sipProxy := "/api/admin/configuration/v1/sip_proxy/1/"
	location := "/api/admin/configuration/v1/system_location/1/"
	maxCallrateIn := 2048
	maxCallrateOut := 1024
	cryptoMode := "auto"

	expectedGatewayRoutingRule := &GatewayRoutingRule{
		ID:                  1,
		Name:                "test-rule",
		Description:         "Updated gateway routing rule",
		Priority:            250,
		Enable:              false,
		MatchString:         "example@.*",
		MatchStringFull:     true,
		CalledDeviceType:    "external",
		OutgoingProtocol:    "sip",
		CallType:            "audio",
		MatchIncomingCalls:  true,
		MatchOutgoingCalls:  false,
		MatchIncomingSIP:    true,
		MatchIncomingWebRTC: false,
		OutgoingLocation:    &location,
		SIPProxy:            &sipProxy,
		MaxCallrateIn:       &maxCallrateIn,
		MaxCallrateOut:      &maxCallrateOut,
		CryptoMode:          &cryptoMode,
		TreatAsTrusted:      true,
	}

	client.On("PutJSON", t.Context(), "configuration/v1/gateway_routing_rule/1/", updateRequest, mock.AnythingOfType("*config.GatewayRoutingRule")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*GatewayRoutingRule)
		*result = *expectedGatewayRoutingRule
	})

	service := New(client)
	result, err := service.UpdateGatewayRoutingRule(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedGatewayRoutingRule, result)
	client.AssertExpectations(t)
}

func TestService_DeleteGatewayRoutingRule(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/gateway_routing_rule/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteGatewayRoutingRule(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
