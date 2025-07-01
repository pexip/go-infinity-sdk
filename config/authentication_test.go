package config

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_GetAuthentication(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedAuthentication := &Authentication{
		ID:                        1,
		Source:                    "local",
		ClientCertificate:         "none",
		ApiOauth2DisableBasic:     false,
		ApiOauth2AllowAllPerms:    false,
		ApiOauth2Expiration:       3600,
		LdapServer:                "",
		LdapBaseDN:                "",
		LdapBindUsername:          "",
		LdapBindPassword:          "",
		LdapUserSearchDN:          "",
		LdapUserFilter:            "objectClass=person",
		LdapUserSearchFilter:      "uid={0}",
		LdapUserGroupAttributes:   "",
		LdapGroupSearchDN:         "",
		LdapGroupFilter:           "objectClass=group",
		LdapGroupMembershipFilter: "member={0}",
		LdapUseGlobalCatalog:      false,
		LdapPermitNoTLS:           false,
		OidcMetadataURL:           "",
		OidcMetadata:              "",
		OidcClientID:              "",
		OidcClientSecret:          "",
		OidcPrivateKey:            "",
		OidcAuthMethod:            "client_secret_post",
		OidcScope:                 "",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/authentication/1/", mock.AnythingOfType("*config.Authentication")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*Authentication)
		*result = *expectedAuthentication
	})

	service := New(client)
	result, err := service.GetAuthentication(t.Context())

	assert.NoError(t, err)
	assert.Equal(t, expectedAuthentication, result)
	client.AssertExpectations(t)
}

func TestService_UpdateAuthentication(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	apiExpiration := 7200
	updateRequest := &AuthenticationUpdateRequest{
		Source:              "ldap",
		ApiOauth2Expiration: &apiExpiration,
		LdapServer:          "ldap.example.com",
		LdapBaseDN:          "dc=example,dc=com",
		LdapBindUsername:    "admin",
		LdapBindPassword:    "password",
	}

	expectedAuthentication := &Authentication{
		ID:                        1,
		Source:                    "ldap",
		ClientCertificate:         "none",
		ApiOauth2DisableBasic:     false,
		ApiOauth2AllowAllPerms:    false,
		ApiOauth2Expiration:       7200,
		LdapServer:                "ldap.example.com",
		LdapBaseDN:                "dc=example,dc=com",
		LdapBindUsername:          "admin",
		LdapBindPassword:          "password",
		LdapUserSearchDN:          "",
		LdapUserFilter:            "objectClass=person",
		LdapUserSearchFilter:      "uid={0}",
		LdapUserGroupAttributes:   "",
		LdapGroupSearchDN:         "",
		LdapGroupFilter:           "objectClass=group",
		LdapGroupMembershipFilter: "member={0}",
		LdapUseGlobalCatalog:      false,
		LdapPermitNoTLS:           false,
		OidcMetadataURL:           "",
		OidcMetadata:              "",
		OidcClientID:              "",
		OidcClientSecret:          "",
		OidcPrivateKey:            "",
		OidcAuthMethod:            "client_secret_post",
		OidcScope:                 "",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/authentication/1/", updateRequest, mock.AnythingOfType("*config.Authentication")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*Authentication)
		*result = *expectedAuthentication
	})

	service := New(client)
	result, err := service.UpdateAuthentication(t.Context(), updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedAuthentication, result)
	client.AssertExpectations(t)
}
