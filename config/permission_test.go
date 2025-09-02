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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListPermissions(t *testing.T) {
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
				expectedResponse := &PermissionListResponse{
					Objects: []Permission{
						{
							ID:       1,
							Name:     "Can add conference",
							Codename: "add_conference",
						},
						{
							ID:       2,
							Name:     "Can change conference",
							Codename: "change_conference",
						},
						{
							ID:       3,
							Name:     "Can delete conference",
							Codename: "delete_conference",
						},
						{
							ID:       4,
							Name:     "Can view conference",
							Codename: "view_conference",
						},
						{
							ID:       5,
							Name:     "Can add user",
							Codename: "add_user",
						},
						{
							ID:       6,
							Name:     "Can change user",
							Codename: "change_user",
						},
						{
							ID:       7,
							Name:     "Can delete user",
							Codename: "delete_user",
						},
						{
							ID:       8,
							Name:     "Can view user",
							Codename: "view_user",
						},
						{
							ID:       9,
							Name:     "Can add device",
							Codename: "add_device",
						},
						{
							ID:       10,
							Name:     "Can change device",
							Codename: "change_device",
						},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/permission/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.PermissionListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*PermissionListResponse)
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
				Search: "conference",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				expectedResponse := &PermissionListResponse{
					Objects: []Permission{
						{
							ID:       1,
							Name:     "Can add conference",
							Codename: "add_conference",
						},
						{
							ID:       2,
							Name:     "Can change conference",
							Codename: "change_conference",
						},
						{
							ID:       3,
							Name:     "Can delete conference",
							Codename: "delete_conference",
						},
						{
							ID:       4,
							Name:     "Can view conference",
							Codename: "view_conference",
						},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/permission/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.PermissionListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*PermissionListResponse)
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
			result, err := service.ListPermissions(t.Context(), tt.opts)

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

func TestService_GetPermission(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	expectedPermission := &Permission{
		ID:       1,
		Name:     "Can add conference",
		Codename: "add_conference",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/permission/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.Permission")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*Permission)
		*result = *expectedPermission
	})

	service := New(client)
	result, err := service.GetPermission(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedPermission, result)
	client.AssertExpectations(t)
}
