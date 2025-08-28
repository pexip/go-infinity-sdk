/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// MjxExchangeDeployment represents a MJX Exchange deployment configuration
type MjxExchangeDeployment struct {
	ID                             int      `json:"id,omitempty"`
	Name                           string   `json:"name"`
	Description                    string   `json:"description,omitempty"`
	ServiceAccountUsername         string   `json:"service_account_username"`
	ServiceAccountPassword         string   `json:"service_account_password,omitempty"`
	AuthenticationMethod           string   `json:"authentication_method"`
	EWSURL                         string   `json:"ews_url,omitempty"`
	DisableProxy                   bool     `json:"disable_proxy"`
	FindItemsRequestQuota          int      `json:"find_items_request_quota"`
	KerberosRealm                  string   `json:"kerberos_realm,omitempty"`
	KerberosKDC                    string   `json:"kerberos_kdc,omitempty"`
	KerberosExchangeSPN            string   `json:"kerberos_exchange_spn,omitempty"`
	KerberosAuthEveryRequest       bool     `json:"kerberos_auth_every_request"`
	KerberosEnableTLS              bool     `json:"kerberos_enable_tls"`
	KerberosKDCHTTPSProxy          string   `json:"kerberos_kdc_https_proxy,omitempty"`
	KerberosVerifyTLSUsingCustomCA bool     `json:"kerberos_verify_tls_using_custom_ca"`
	OAuthClientID                  *string  `json:"oauth_client_id,omitempty"`
	OAuthAuthEndpoint              string   `json:"oauth_auth_endpoint,omitempty"`
	OAuthTokenEndpoint             string   `json:"oauth_token_endpoint,omitempty"`
	OAuthRedirectURI               string   `json:"oauth_redirect_uri,omitempty"`
	OAuthRefreshToken              string   `json:"oauth_refresh_token,omitempty"`
	OAuthState                     *string  `json:"oauth_state,omitempty"`
	AutodiscoverURLs               []string `json:"autodiscover_urls,omitempty"`
	MjxIntegrations                []string `json:"mjx_integrations,omitempty"`
	ResourceURI                    string   `json:"resource_uri,omitempty"`
}

// MjxExchangeDeploymentCreateRequest represents a request to create a MJX Exchange deployment
type MjxExchangeDeploymentCreateRequest struct {
	Name                           string   `json:"name"`
	Description                    string   `json:"description,omitempty"`
	ServiceAccountUsername         string   `json:"service_account_username"`
	ServiceAccountPassword         string   `json:"service_account_password,omitempty"`
	AuthenticationMethod           string   `json:"authentication_method"`
	EWSURL                         string   `json:"ews_url,omitempty"`
	DisableProxy                   bool     `json:"disable_proxy"`
	FindItemsRequestQuota          int      `json:"find_items_request_quota"`
	KerberosRealm                  string   `json:"kerberos_realm,omitempty"`
	KerberosKDC                    string   `json:"kerberos_kdc,omitempty"`
	KerberosExchangeSPN            string   `json:"kerberos_exchange_spn,omitempty"`
	KerberosAuthEveryRequest       bool     `json:"kerberos_auth_every_request"`
	KerberosEnableTLS              bool     `json:"kerberos_enable_tls"`
	KerberosKDCHTTPSProxy          string   `json:"kerberos_kdc_https_proxy,omitempty"`
	KerberosVerifyTLSUsingCustomCA bool     `json:"kerberos_verify_tls_using_custom_ca"`
	OAuthClientID                  *string  `json:"oauth_client_id,omitempty"`
	OAuthAuthEndpoint              string   `json:"oauth_auth_endpoint,omitempty"`
	OAuthTokenEndpoint             string   `json:"oauth_token_endpoint,omitempty"`
	OAuthRedirectURI               string   `json:"oauth_redirect_uri,omitempty"`
	OAuthRefreshToken              string   `json:"oauth_refresh_token,omitempty"`
	AutodiscoverURLs               []string `json:"autodiscover_urls,omitempty"`
}

// MjxExchangeDeploymentUpdateRequest represents a request to update a MJX Exchange deployment
type MjxExchangeDeploymentUpdateRequest struct {
	Name                           string   `json:"name,omitempty"`
	Description                    string   `json:"description,omitempty"`
	ServiceAccountUsername         string   `json:"service_account_username,omitempty"`
	ServiceAccountPassword         string   `json:"service_account_password,omitempty"`
	AuthenticationMethod           string   `json:"authentication_method,omitempty"`
	EWSURL                         string   `json:"ews_url,omitempty"`
	DisableProxy                   *bool    `json:"disable_proxy,omitempty"`
	FindItemsRequestQuota          *int     `json:"find_items_request_quota,omitempty"`
	KerberosRealm                  string   `json:"kerberos_realm,omitempty"`
	KerberosKDC                    string   `json:"kerberos_kdc,omitempty"`
	KerberosExchangeSPN            string   `json:"kerberos_exchange_spn,omitempty"`
	KerberosAuthEveryRequest       *bool    `json:"kerberos_auth_every_request,omitempty"`
	KerberosEnableTLS              *bool    `json:"kerberos_enable_tls,omitempty"`
	KerberosKDCHTTPSProxy          string   `json:"kerberos_kdc_https_proxy,omitempty"`
	KerberosVerifyTLSUsingCustomCA *bool    `json:"kerberos_verify_tls_using_custom_ca,omitempty"`
	OAuthClientID                  *string  `json:"oauth_client_id,omitempty"`
	OAuthAuthEndpoint              string   `json:"oauth_auth_endpoint,omitempty"`
	OAuthTokenEndpoint             string   `json:"oauth_token_endpoint,omitempty"`
	OAuthRedirectURI               string   `json:"oauth_redirect_uri,omitempty"`
	OAuthRefreshToken              string   `json:"oauth_refresh_token,omitempty"`
	AutodiscoverURLs               []string `json:"autodiscover_urls,omitempty"`
}

// MjxExchangeDeploymentListResponse represents the response from listing MJX Exchange deployments
type MjxExchangeDeploymentListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []MjxExchangeDeployment `json:"objects"`
}
