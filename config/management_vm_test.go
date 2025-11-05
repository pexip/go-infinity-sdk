/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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
	result, err := service.GetManagementVM(t.Context())

	assert.NoError(t, err)
	assert.Equal(t, expectedManagementVM, result)
	client.AssertExpectations(t)
}
