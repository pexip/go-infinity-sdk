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

func TestService_ListSMTPServers(t *testing.T) {
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
				expectedResponse := &SMTPServerListResponse{
					Objects: []SMTPServer{
						{ID: 1, Name: "primary-smtp", Description: "Primary SMTP server", Address: "smtp.example.com", Port: 587, ConnectionSecurity: "STARTTLS", FromEmailAddress: "noreply@example.com"},
						{ID: 2, Name: "backup-smtp", Description: "Backup SMTP server", Address: "backup-smtp.example.com", Port: 25, ConnectionSecurity: "NONE", FromEmailAddress: "noreply@example.com"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/smtp_server/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.SMTPServerListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*SMTPServerListResponse)
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
				expectedResponse := &SMTPServerListResponse{
					Objects: []SMTPServer{
						{ID: 1, Name: "primary-smtp", Description: "Primary SMTP server", Address: "smtp.example.com", Port: 587, ConnectionSecurity: "STARTTLS", FromEmailAddress: "noreply@example.com"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/smtp_server/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.SMTPServerListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*SMTPServerListResponse)
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
			result, err := service.ListSMTPServers(t.Context(), tt.opts)

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

func TestService_GetSMTPServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedSMTPServer := &SMTPServer{
		ID:                 1,
		Name:               "test-smtp",
		Description:        "Test SMTP server",
		Address:            "smtp.example.com",
		Port:               587,
		Username:           "testuser",
		Password:           "testpass",
		FromEmailAddress:   "noreply@example.com",
		ConnectionSecurity: "STARTTLS",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/smtp_server/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.SMTPServer")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*SMTPServer)
		*result = *expectedSMTPServer
	})

	service := New(client)
	result, err := service.GetSMTPServer(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedSMTPServer, result)
	client.AssertExpectations(t)
}

func TestService_CreateSMTPServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &SMTPServerCreateRequest{
		Name:               "new-smtp",
		Description:        "New SMTP server",
		Address:            "new-smtp.example.com",
		Port:               587,
		Username:           "newuser",
		Password:           "newpass",
		FromEmailAddress:   "noreply@example.com",
		ConnectionSecurity: "STARTTLS",
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/smtp_server/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/smtp_server/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateSMTPServer(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateSMTPServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	newPort := 465
	updateRequest := &SMTPServerUpdateRequest{
		Description:        "Updated SMTP server",
		Port:               &newPort,
		ConnectionSecurity: "STARTTLS",
	}

	expectedSMTPServer := &SMTPServer{
		ID:                 1,
		Name:               "test-smtp",
		Description:        "Updated SMTP server",
		Address:            "smtp.example.com",
		Port:               465,
		Username:           "testuser",
		Password:           "testpass",
		FromEmailAddress:   "noreply@example.com",
		ConnectionSecurity: "STARTTLS",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/smtp_server/1/", updateRequest, mock.AnythingOfType("*config.SMTPServer")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*SMTPServer)
		*result = *expectedSMTPServer
	})

	service := New(client)
	result, err := service.UpdateSMTPServer(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedSMTPServer, result)
	client.AssertExpectations(t)
}

func TestService_DeleteSMTPServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/smtp_server/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteSMTPServer(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
