/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// Authentication represents authentication configuration (singleton)
type Authentication struct {
	ID                        int    `json:"id,omitempty"`
	Source                    string `json:"source"`
	ClientCertificate         string `json:"client_certificate"`
	ApiOauth2DisableBasic     bool   `json:"api_oauth2_disable_basic"`
	ApiOauth2AllowAllPerms    bool   `json:"api_oauth2_allow_all_perms"`
	ApiOauth2Expiration       int    `json:"api_oauth2_expiration"`
	LdapServer                string `json:"ldap_server,omitempty"`
	LdapBaseDN                string `json:"ldap_base_dn,omitempty"`
	LdapBindUsername          string `json:"ldap_bind_username,omitempty"`
	LdapBindPassword          string `json:"ldap_bind_password,omitempty"`
	LdapUserSearchDN          string `json:"ldap_user_search_dn,omitempty"`
	LdapUserFilter            string `json:"ldap_user_filter"`
	LdapUserSearchFilter      string `json:"ldap_user_search_filter"`
	LdapUserGroupAttributes   string `json:"ldap_user_group_attributes,omitempty"`
	LdapGroupSearchDN         string `json:"ldap_group_search_dn,omitempty"`
	LdapGroupFilter           string `json:"ldap_group_filter"`
	LdapGroupMembershipFilter string `json:"ldap_group_membership_filter"`
	LdapUseGlobalCatalog      bool   `json:"ldap_use_global_catalog"`
	LdapPermitNoTLS           bool   `json:"ldap_permit_no_tls"`
	OidcMetadataURL           string `json:"oidc_metadata_url,omitempty"`
	OidcMetadata              string `json:"oidc_metadata,omitempty"`
	OidcClientID              string `json:"oidc_client_id,omitempty"`
	OidcClientSecret          string `json:"oidc_client_secret,omitempty"`
	OidcPrivateKey            string `json:"oidc_private_key,omitempty"`
	OidcAuthMethod            string `json:"oidc_auth_method"`
	OidcScope                 string `json:"oidc_scope,omitempty"`
	OidcAuthorizeURL          string `json:"oidc_authorize_url,omitempty"`
	OidcTokenEndpointURL      string `json:"oidc_token_endpoint_url,omitempty"`
	OidcUsernameField         string `json:"oidc_username_field"`
	OidcGroupsField           string `json:"oidc_groups_field,omitempty"`
	OidcRequiredKey           string `json:"oidc_required_key,omitempty"`
	OidcRequiredValue         string `json:"oidc_required_value,omitempty"`
	OidcDomainHint            string `json:"oidc_domain_hint,omitempty"`
	OidcLoginButton           string `json:"oidc_login_button,omitempty"`
	ResourceURI               string `json:"resource_uri,omitempty"`
}

// AuthenticationUpdateRequest represents a request to update authentication configuration
type AuthenticationUpdateRequest struct {
	Source                    string `json:"source,omitempty"`
	ClientCertificate         string `json:"client_certificate,omitempty"`
	ApiOauth2DisableBasic     *bool  `json:"api_oauth2_disable_basic,omitempty"`
	ApiOauth2AllowAllPerms    *bool  `json:"api_oauth2_allow_all_perms,omitempty"`
	ApiOauth2Expiration       *int   `json:"api_oauth2_expiration,omitempty"`
	LdapServer                string `json:"ldap_server,omitempty"`
	LdapBaseDN                string `json:"ldap_base_dn,omitempty"`
	LdapBindUsername          string `json:"ldap_bind_username,omitempty"`
	LdapBindPassword          string `json:"ldap_bind_password,omitempty"`
	LdapUserSearchDN          string `json:"ldap_user_search_dn,omitempty"`
	LdapUserFilter            string `json:"ldap_user_filter,omitempty"`
	LdapUserSearchFilter      string `json:"ldap_user_search_filter,omitempty"`
	LdapUserGroupAttributes   string `json:"ldap_user_group_attributes,omitempty"`
	LdapGroupSearchDN         string `json:"ldap_group_search_dn,omitempty"`
	LdapGroupFilter           string `json:"ldap_group_filter,omitempty"`
	LdapGroupMembershipFilter string `json:"ldap_group_membership_filter,omitempty"`
	LdapUseGlobalCatalog      *bool  `json:"ldap_use_global_catalog,omitempty"`
	LdapPermitNoTLS           *bool  `json:"ldap_permit_no_tls,omitempty"`
	OidcMetadataURL           string `json:"oidc_metadata_url,omitempty"`
	OidcMetadata              string `json:"oidc_metadata,omitempty"`
	OidcClientID              string `json:"oidc_client_id,omitempty"`
	OidcClientSecret          string `json:"oidc_client_secret,omitempty"`
	OidcPrivateKey            string `json:"oidc_private_key,omitempty"`
	OidcAuthMethod            string `json:"oidc_auth_method,omitempty"`
	OidcScope                 string `json:"oidc_scope,omitempty"`
	OidcAuthorizeURL          string `json:"oidc_authorize_url,omitempty"`
	OidcTokenEndpointURL      string `json:"oidc_token_endpoint_url,omitempty"`
	OidcUsernameField         string `json:"oidc_username_field,omitempty"`
	OidcGroupsField           string `json:"oidc_groups_field,omitempty"`
	OidcRequiredKey           string `json:"oidc_required_key,omitempty"`
	OidcRequiredValue         string `json:"oidc_required_value,omitempty"`
	OidcDomainHint            string `json:"oidc_domain_hint,omitempty"`
	OidcLoginButton           string `json:"oidc_login_button,omitempty"`
}
