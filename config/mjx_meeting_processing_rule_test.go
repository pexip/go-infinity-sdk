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

func TestService_ListMjxMeetingProcessingRules(t *testing.T) {
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
				expectedResponse := &MjxMeetingProcessingRuleListResponse{
					Objects: []MjxMeetingProcessingRule{
						{
							ID:                       1,
							Name:                     "primary-processing-rule",
							Description:              "Primary meeting processing rule",
							Priority:                 100,
							Enabled:                  true,
							MeetingType:              "exchange",
							MjxIntegration:           "/api/admin/configuration/v1/mjx_integration/1/",
							MatchString:              "meeting@example.com",
							ReplaceString:            "conference@example.com",
							TransformRule:            "replace",
							CustomTemplate:           "{{subject}} - {{location}}",
							Domain:                   "example.com",
							CompanyID:                "company123",
							IncludePin:               true,
							DefaultProcessingEnabled: true,
						},
						{
							ID:                       2,
							Name:                     "backup-processing-rule",
							Description:              "Backup meeting processing rule",
							Priority:                 50,
							Enabled:                  false,
							MeetingType:              "google",
							MjxIntegration:           "/api/admin/configuration/v1/mjx_integration/2/",
							MatchString:              "backup@example.com",
							ReplaceString:            "backup-conf@example.com",
							TransformRule:            "template",
							CustomTemplate:           "Backup: {{subject}}",
							Domain:                   "backup.example.com",
							IncludePin:               false,
							DefaultProcessingEnabled: false,
						},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/mjx_meeting_processing_rule/", mock.AnythingOfType("*config.MjxMeetingProcessingRuleListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*MjxMeetingProcessingRuleListResponse)
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
				expectedResponse := &MjxMeetingProcessingRuleListResponse{
					Objects: []MjxMeetingProcessingRule{
						{
							ID:                       1,
							Name:                     "primary-processing-rule",
							Description:              "Primary meeting processing rule",
							Priority:                 100,
							Enabled:                  true,
							MeetingType:              "exchange",
							MjxIntegration:           "/api/admin/configuration/v1/mjx_integration/1/",
							DefaultProcessingEnabled: true,
						},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/mjx_meeting_processing_rule/?limit=5&name__icontains=primary", mock.AnythingOfType("*config.MjxMeetingProcessingRuleListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*MjxMeetingProcessingRuleListResponse)
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
			result, err := service.ListMjxMeetingProcessingRules(t.Context(), tt.opts)

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

func TestService_GetMjxMeetingProcessingRule(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	expectedRule := &MjxMeetingProcessingRule{
		ID:                       1,
		Name:                     "test-processing-rule",
		Description:              "Test meeting processing rule",
		Priority:                 75,
		Enabled:                  true,
		MeetingType:              "exchange",
		MjxIntegration:           "/api/admin/configuration/v1/mjx_integration/1/",
		MatchString:              "test@example.com",
		ReplaceString:            "test-conf@example.com",
		TransformRule:            "replace",
		CustomTemplate:           "{{subject}} - Conference Room {{location}}",
		Domain:                   "test.example.com",
		CompanyID:                "testcompany456",
		IncludePin:               true,
		DefaultProcessingEnabled: true,
	}

	client.On("GetJSON", t.Context(), "configuration/v1/mjx_meeting_processing_rule/1/", mock.AnythingOfType("*config.MjxMeetingProcessingRule")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*MjxMeetingProcessingRule)
		*result = *expectedRule
	})

	service := New(client)
	result, err := service.GetMjxMeetingProcessingRule(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedRule, result)
	client.AssertExpectations(t)
}

func TestService_CreateMjxMeetingProcessingRule(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &MjxMeetingProcessingRuleCreateRequest{
		Name:                     "new-processing-rule",
		Description:              "New meeting processing rule",
		Priority:                 90,
		Enabled:                  true,
		MeetingType:              "graph",
		MjxIntegration:           "/api/admin/configuration/v1/mjx_integration/3/",
		MatchString:              "new@example.com",
		ReplaceString:            "new-conf@example.com",
		TransformRule:            "template",
		CustomTemplate:           "New Meeting: {{subject}} in {{location}}",
		Domain:                   "new.example.com",
		CompanyID:                "newcompany789",
		IncludePin:               false,
		DefaultProcessingEnabled: true,
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/mjx_meeting_processing_rule/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/mjx_meeting_processing_rule/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateMjxMeetingProcessingRule(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateMjxMeetingProcessingRule(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	priority := 80
	enabled := false
	includePin := false
	defaultProcessingEnabled := false

	updateRequest := &MjxMeetingProcessingRuleUpdateRequest{
		Description:              "Updated meeting processing rule",
		Priority:                 &priority,
		Enabled:                  &enabled,
		CustomTemplate:           "Updated: {{subject}} at {{location}}",
		IncludePin:               &includePin,
		DefaultProcessingEnabled: &defaultProcessingEnabled,
	}

	expectedRule := &MjxMeetingProcessingRule{
		ID:                       1,
		Name:                     "test-processing-rule",
		Description:              "Updated meeting processing rule",
		Priority:                 80,
		Enabled:                  false,
		MeetingType:              "exchange",
		MjxIntegration:           "/api/admin/configuration/v1/mjx_integration/1/",
		MatchString:              "test@example.com",
		ReplaceString:            "test-conf@example.com",
		TransformRule:            "replace",
		CustomTemplate:           "Updated: {{subject}} at {{location}}",
		Domain:                   "test.example.com",
		CompanyID:                "testcompany456",
		IncludePin:               false,
		DefaultProcessingEnabled: false,
	}

	client.On("PutJSON", t.Context(), "configuration/v1/mjx_meeting_processing_rule/1/", updateRequest, mock.AnythingOfType("*config.MjxMeetingProcessingRule")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*MjxMeetingProcessingRule)
		*result = *expectedRule
	})

	service := New(client)
	result, err := service.UpdateMjxMeetingProcessingRule(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedRule, result)
	client.AssertExpectations(t)
}

func TestService_DeleteMjxMeetingProcessingRule(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/mjx_meeting_processing_rule/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteMjxMeetingProcessingRule(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
