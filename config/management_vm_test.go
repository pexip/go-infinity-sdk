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

func TestService_ListManagementVMs(t *testing.T) {
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
				ipv6Address := "2001:db8::1"
				natAddress := "203.0.113.1"
				expectedResponse := &ManagementVMListResponse{
					Objects: []ManagementVM{
						{ID: 1, Name: "mgmt-vm-01", Description: "Primary management VM", Address: "192.168.1.10", Netmask: "255.255.255.0", Gateway: "192.168.1.1", Hostname: "mgmt01", Domain: "example.com", MTU: 1500, IPV6Address: &ipv6Address, StaticNATAddress: &natAddress, EnableSSH: "enabled", SNMPMode: "v2c", Primary: true, Initializing: false},
						{ID: 2, Name: "mgmt-vm-02", Description: "Secondary management VM", Address: "192.168.1.11", Netmask: "255.255.255.0", Gateway: "192.168.1.1", Hostname: "mgmt02", Domain: "example.com", MTU: 1500, EnableSSH: "disabled", SNMPMode: "disabled", Primary: false, Initializing: false},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/management_vm/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.ManagementVMListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*ManagementVMListResponse)
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
				ipv6Address := "2001:db8::1"
				natAddress := "203.0.113.1"
				expectedResponse := &ManagementVMListResponse{
					Objects: []ManagementVM{
						{ID: 1, Name: "mgmt-vm-01", Description: "Primary management VM", Address: "192.168.1.10", Netmask: "255.255.255.0", Gateway: "192.168.1.1", Hostname: "mgmt01", Domain: "example.com", MTU: 1500, IPV6Address: &ipv6Address, StaticNATAddress: &natAddress, EnableSSH: "enabled", SNMPMode: "v2c", Primary: true, Initializing: false},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/management_vm/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.ManagementVMListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*ManagementVMListResponse)
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
			result, err := service.ListManagementVMs(t.Context(), tt.opts)

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

func TestService_GetManagementVM(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	ipv6Address := "2001:db8::1"
	ipv6Gateway := "2001:db8::1"
	natAddress := "203.0.113.1"
	tlsCert := "/api/admin/configuration/v1/tls_certificate/1/"
	httpProxy := "/api/admin/configuration/v1/http_proxy/1/"

	expectedManagementVM := &ManagementVM{
		ID:                         1,
		Name:                       "test-mgmt-vm",
		Description:                "Test management VM",
		Address:                    "192.168.1.10",
		Netmask:                    "255.255.255.0",
		Gateway:                    "192.168.1.1",
		Hostname:                   "testmgmt",
		Domain:                     "example.com",
		AlternativeFQDN:            "alt.example.com",
		IPV6Address:                &ipv6Address,
		IPV6Gateway:                &ipv6Gateway,
		MTU:                        1500,
		StaticNATAddress:           &natAddress,
		DNSServers:                 []DNSServer{{ResourceURI: "/api/admin/configuration/v1/dns_server/1/"}},
		NTPServers:                 []NTPServer{{ResourceURI: "/api/admin/configuration/v1/ntp_server/1/"}},
		SyslogServers:              []SyslogServer{{ResourceURI: "/api/admin/configuration/v1/syslog_server/1/"}},
		StaticRoutes:               []string{"/api/admin/configuration/v1/static_route/1/"},
		EventSinks:                 []string{"/api/admin/configuration/v1/event_sink/1/"},
		HTTPProxy:                  &httpProxy,
		TLSCertificate:             &tlsCert,
		EnableSSH:                  "enabled",
		SSHAuthorizedKeys:          []string{"/api/admin/configuration/v1/ssh_authorized_key/1/"},
		SSHAuthorizedKeysUseCloud:  false,
		SecondaryConfigPassphrase:  "passphrase123",
		SNMPMode:                   "v3",
		SNMPCommunity:              "public",
		SNMPUsername:               "snmpuser",
		SNMPAuthenticationPassword: "authpass",
		SNMPPrivacyPassword:        "privpass",
		SNMPSystemContact:          "admin@example.com",
		SNMPSystemLocation:         "Data Center 1",
		Primary:                    true,
		Initializing:               false,
	}

	client.On("GetJSON", t.Context(), "configuration/v1/management_vm/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.ManagementVM")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*ManagementVM)
		*result = *expectedManagementVM
	})

	service := New(client)
	result, err := service.GetManagementVM(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedManagementVM, result)
	client.AssertExpectations(t)
}

func TestService_CreateManagementVM(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	ipv6Address := "2001:db8::2"
	natAddress := "203.0.113.2"
	tlsCert := "/api/admin/configuration/v1/tls_certificate/1/"

	createRequest := &ManagementVMCreateRequest{
		Name:                      "new-mgmt-vm",
		Description:               "New management VM",
		Address:                   "192.168.1.20",
		Netmask:                   "255.255.255.0",
		Gateway:                   "192.168.1.1",
		Hostname:                  "newmgmt",
		Domain:                    "example.com",
		AlternativeFQDN:           "newmgmt.example.com",
		IPV6Address:               &ipv6Address,
		MTU:                       1500,
		StaticNATAddress:          &natAddress,
		DNSServers:                []string{"/api/admin/configuration/v1/dns_server/1/"},
		NTPServers:                []string{"/api/admin/configuration/v1/ntp_server/1/"},
		TLSCertificate:            &tlsCert,
		EnableSSH:                 "enabled",
		SSHAuthorizedKeysUseCloud: false,
		SNMPMode:                  "v2c",
		SNMPCommunity:             "public",
		SNMPSystemContact:         "admin@example.com",
		SNMPSystemLocation:        "Data Center 2",
		Initializing:              false,
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/management_vm/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/management_vm/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateManagementVM(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_DeleteManagementVM(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/management_vm/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteManagementVM(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
