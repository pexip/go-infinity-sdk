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

func TestService_ListBreakInAllowListAddresses(t *testing.T) {
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
				expectedResponse := &BreakInAllowListAddressListResponse{
					Objects: []BreakInAllowListAddress{
						{ID: 1, Name: "corporate-network", Description: "Corporate network range", Address: "192.168.1.0", Prefix: 24, AllowlistEntryType: "all", IgnoreIncorrectAliases: true, IgnoreIncorrectPins: false},
						{ID: 2, Name: "vpn-range", Description: "VPN IP range", Address: "10.0.0.0", Prefix: 16, AllowlistEntryType: "alias_only", IgnoreIncorrectAliases: false, IgnoreIncorrectPins: true},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/break_in_allow_list_address/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.BreakInAllowListAddressListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*BreakInAllowListAddressListResponse)
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
				Search: "corporate",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				expectedResponse := &BreakInAllowListAddressListResponse{
					Objects: []BreakInAllowListAddress{
						{ID: 1, Name: "corporate-network", Description: "Corporate network range", Address: "192.168.1.0", Prefix: 24, AllowlistEntryType: "all", IgnoreIncorrectAliases: true, IgnoreIncorrectPins: false},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/break_in_allow_list_address/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.BreakInAllowListAddressListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*BreakInAllowListAddressListResponse)
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
			result, err := service.ListBreakInAllowListAddresses(t.Context(), tt.opts)

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

func TestService_GetBreakInAllowListAddress(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedBreakInAllowListAddress := &BreakInAllowListAddress{
		ID:                     1,
		Name:                   "test-allowlist",
		Description:            "Test allowlist entry",
		Address:                "192.168.1.0",
		Prefix:                 24,
		AllowlistEntryType:     "all",
		IgnoreIncorrectAliases: true,
		IgnoreIncorrectPins:    false,
	}

	client.On("GetJSON", t.Context(), "configuration/v1/break_in_allow_list_address/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.BreakInAllowListAddress")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*BreakInAllowListAddress)
		*result = *expectedBreakInAllowListAddress
	})

	service := New(client)
	result, err := service.GetBreakInAllowListAddress(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedBreakInAllowListAddress, result)
	client.AssertExpectations(t)
}

func TestService_CreateBreakInAllowListAddress(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &BreakInAllowListAddressCreateRequest{
		Name:                   "new-allowlist",
		Description:            "New allowlist entry",
		Address:                "10.0.0.0",
		Prefix:                 16,
		AllowlistEntryType:     "alias_only",
		IgnoreIncorrectAliases: false,
		IgnoreIncorrectPins:    true,
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/break_in_allow_list_address/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/break_in_allow_list_address/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateBreakInAllowListAddress(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateBreakInAllowListAddress(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	prefix := 20
	ignoreAliases := false
	updateRequest := &BreakInAllowListAddressUpdateRequest{
		Description:            "Updated allowlist entry",
		Prefix:                 &prefix,
		IgnoreIncorrectAliases: &ignoreAliases,
	}

	expectedBreakInAllowListAddress := &BreakInAllowListAddress{
		ID:                     1,
		Name:                   "test-allowlist",
		Description:            "Updated allowlist entry",
		Address:                "192.168.1.0",
		Prefix:                 20,
		AllowlistEntryType:     "all",
		IgnoreIncorrectAliases: false,
		IgnoreIncorrectPins:    false,
	}

	client.On("PutJSON", t.Context(), "configuration/v1/break_in_allow_list_address/1/", updateRequest, mock.AnythingOfType("*config.BreakInAllowListAddress")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*BreakInAllowListAddress)
		*result = *expectedBreakInAllowListAddress
	})

	service := New(client)
	result, err := service.UpdateBreakInAllowListAddress(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedBreakInAllowListAddress, result)
	client.AssertExpectations(t)
}

func TestService_DeleteBreakInAllowListAddress(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/break_in_allow_list_address/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteBreakInAllowListAddress(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
