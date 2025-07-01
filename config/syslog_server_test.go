package config

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListSyslogServers(t *testing.T) {
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
				expectedResponse := &SyslogServerListResponse{
					Objects: []SyslogServer{
						{ID: 1, Address: "192.168.1.100", Port: 514, Transport: "udp", Description: "Main syslog server"},
						{ID: 2, Address: "syslog.example.com", Port: 6514, Transport: "tls", Description: "Remote syslog server"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/syslog_server/", mock.AnythingOfType("*config.SyslogServerListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*SyslogServerListResponse)
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
				Search: "main",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				expectedResponse := &SyslogServerListResponse{
					Objects: []SyslogServer{
						{ID: 1, Address: "192.168.1.100", Port: 514, Transport: "udp", Description: "Main syslog server"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/syslog_server/?limit=5&name__icontains=main", mock.AnythingOfType("*config.SyslogServerListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*SyslogServerListResponse)
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
			result, err := service.ListSyslogServers(t.Context(), tt.opts)

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

func TestService_GetSyslogServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedServer := &SyslogServer{
		ID:          1,
		Address:     "192.168.1.100",
		Description: "Main syslog server",
		Port:        514,
		Transport:   "udp",
		AuditLog:    true,
		SupportLog:  true,
		WebLog:      false,
	}

	client.On("GetJSON", t.Context(), "configuration/v1/syslog_server/1/", mock.AnythingOfType("*config.SyslogServer")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*SyslogServer)
		*result = *expectedServer
	})

	service := New(client)
	result, err := service.GetSyslogServer(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedServer, result)
	client.AssertExpectations(t)
}

func TestService_CreateSyslogServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &SyslogServerCreateRequest{
		Address:     "syslog.example.com",
		Description: "Remote syslog server",
		Port:        6514,
		Transport:   "tls",
		AuditLog:    true,
		SupportLog:  true,
		WebLog:      false,
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/syslog_server/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/syslog_server/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateSyslogServer(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateSyslogServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	updateRequest := &SyslogServerUpdateRequest{
		Description: "Updated syslog server",
		Port:        1514,
	}

	expectedServer := &SyslogServer{
		ID:          1,
		Address:     "192.168.1.100",
		Description: "Updated syslog server",
		Port:        1514,
		Transport:   "udp",
		AuditLog:    true,
		SupportLog:  true,
		WebLog:      false,
	}

	client.On("PutJSON", t.Context(), "configuration/v1/syslog_server/1/", updateRequest, mock.AnythingOfType("*config.SyslogServer")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*SyslogServer)
		*result = *expectedServer
	})

	service := New(client)
	result, err := service.UpdateSyslogServer(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedServer, result)
	client.AssertExpectations(t)
}

func TestService_DeleteSyslogServer(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/syslog_server/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteSyslogServer(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
