package config

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListSSHAuthorizedKeys(t *testing.T) {
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
				expectedResponse := &SSHAuthorizedKeyListResponse{
					Objects: []SSHAuthorizedKey{
						{ID: 1, Keytype: "ssh-rsa", Key: "AAAAB3NzaC1yc2EAAAADAQABAAABAQ...", Comment: "admin@example.com", Nodes: []string{"management"}},
						{ID: 2, Keytype: "ssh-ed25519", Key: "AAAAC3NzaC1lZDI1NTE5AAAAID...", Comment: "user@example.com", Nodes: []string{"management", "conferencing"}},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/ssh_authorized_key/", mock.AnythingOfType("*config.SSHAuthorizedKeyListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*SSHAuthorizedKeyListResponse)
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
				Search: "admin",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				expectedResponse := &SSHAuthorizedKeyListResponse{
					Objects: []SSHAuthorizedKey{
						{ID: 1, Keytype: "ssh-rsa", Key: "AAAAB3NzaC1yc2EAAAADAQABAAABAQ...", Comment: "admin@example.com", Nodes: []string{"management"}},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/ssh_authorized_key/?limit=5&name__icontains=admin", mock.AnythingOfType("*config.SSHAuthorizedKeyListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*SSHAuthorizedKeyListResponse)
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
			result, err := service.ListSSHAuthorizedKeys(t.Context(), tt.opts)

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

func TestService_GetSSHAuthorizedKey(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedSSHAuthorizedKey := &SSHAuthorizedKey{
		ID:          1,
		Keytype:     "ssh-rsa",
		Key:         "AAAAB3NzaC1yc2EAAAADAQABAAABAQ...",
		Comment:     "admin@example.com",
		Nodes:       []string{"management"},
		ResourceURI: "/api/admin/configuration/v1/ssh_authorized_key/1/",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/ssh_authorized_key/1/", mock.AnythingOfType("*config.SSHAuthorizedKey")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*SSHAuthorizedKey)
		*result = *expectedSSHAuthorizedKey
	})

	service := New(client)
	result, err := service.GetSSHAuthorizedKey(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedSSHAuthorizedKey, result)
	client.AssertExpectations(t)
}

func TestService_CreateSSHAuthorizedKey(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &SSHAuthorizedKeyCreateRequest{
		Keytype: "ssh-rsa",
		Key:     "AAAAB3NzaC1yc2EAAAADAQABAAABAQ...",
		Comment: "new-admin@example.com",
		Nodes:   []string{"management", "conferencing"},
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/ssh_authorized_key/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/ssh_authorized_key/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateSSHAuthorizedKey(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateSSHAuthorizedKey(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	updateRequest := &SSHAuthorizedKeyUpdateRequest{
		Comment: "updated-admin@example.com",
		Nodes:   []string{"management"},
	}

	expectedSSHAuthorizedKey := &SSHAuthorizedKey{
		ID:          1,
		Keytype:     "ssh-rsa",
		Key:         "AAAAB3NzaC1yc2EAAAADAQABAAABAQ...",
		Comment:     "updated-admin@example.com",
		Nodes:       []string{"management"},
		ResourceURI: "/api/admin/configuration/v1/ssh_authorized_key/1/",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/ssh_authorized_key/1/", updateRequest, mock.AnythingOfType("*config.SSHAuthorizedKey")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*SSHAuthorizedKey)
		*result = *expectedSSHAuthorizedKey
	})

	service := New(client)
	result, err := service.UpdateSSHAuthorizedKey(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedSSHAuthorizedKey, result)
	client.AssertExpectations(t)
}

func TestService_DeleteSSHAuthorizedKey(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/ssh_authorized_key/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteSSHAuthorizedKey(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
