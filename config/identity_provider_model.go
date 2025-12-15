/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// IdentityProvider represents an identity provider configuration
type IdentityProvider struct {
	ID                                  int    `json:"id,omitempty"`
	Name                                string `json:"name"`
	Description                         string `json:"description,omitempty"`
	IdpType                             string `json:"idp_type"`
	UUID                                string `json:"uuid"`
	SSOUrl                              string `json:"sso_url,omitempty"`
	IdpEntityID                         string `json:"idp_entity_id,omitempty"`
	IdpPublicKey                        string `json:"idp_public_key,omitempty"`
	ServiceEntityID                     string `json:"service_entity_id,omitempty"`
	ServicePublicKey                    string `json:"service_public_key,omitempty"`
	ServicePrivateKey                   string `json:"service_private_key,omitempty"`
	SignatureAlgorithm                  string `json:"signature_algorithm"`
	DigestAlgorithm                     string `json:"digest_algorithm"`
	DisplayNameAttributeName            string `json:"display_name_attribute_name,omitempty"`
	RegistrationAliasAttributeName      string `json:"registration_alias_attribute_name,omitempty"`
	AssertionConsumerServiceURL         string `json:"assertion_consumer_service_url"`
	AssertionConsumerServiceURL2        string `json:"assertion_consumer_service_url2,omitempty"`
	AssertionConsumerServiceURL3        string `json:"assertion_consumer_service_url3,omitempty"`
	AssertionConsumerServiceURL4        string `json:"assertion_consumer_service_url4,omitempty"`
	AssertionConsumerServiceURL5        string `json:"assertion_consumer_service_url5,omitempty"`
	AssertionConsumerServiceURL6        string `json:"assertion_consumer_service_url6,omitempty"`
	AssertionConsumerServiceURL7        string `json:"assertion_consumer_service_url7,omitempty"`
	AssertionConsumerServiceURL8        string `json:"assertion_consumer_service_url8,omitempty"`
	AssertionConsumerServiceURL9        string `json:"assertion_consumer_service_url9,omitempty"`
	AssertionConsumerServiceURL10       string `json:"assertion_consumer_service_url10,omitempty"`
	WorkerFQDNACSURLs                   bool   `json:"worker_fqdn_acs_urls"`
	DisablePopupFlow                    bool   `json:"disable_popup_flow"`
	OidcFlow                            string `json:"oidc_flow"`
	OidcClientID                        string `json:"oidc_client_id,omitempty"`
	OidcClientSecret                    string `json:"oidc_client_secret,omitempty"`
	OidcTokenURL                        string `json:"oidc_token_url,omitempty"`
	OidcUserInfoURL                     string `json:"oidc_user_info_url,omitempty"`
	OidcJWKSURL                         string `json:"oidc_jwks_url,omitempty"`
	OidcTokenEndpointAuthScheme         string `json:"oidc_token_endpoint_auth_scheme"`
	OidcTokenSignatureScheme            string `json:"oidc_token_signature_scheme"`
	OidcDisplayNameClaimName            string `json:"oidc_display_name_claim_name,omitempty"`
	OidcRegistrationAliasClaimName      string `json:"oidc_registration_alias_claim_name,omitempty"`
	OidcAdditionalScopes                string `json:"oidc_additional_scopes,omitempty"`
	OidcFranceConnectRequiredEidasLevel string `json:"oidc_france_connect_required_eidas_level"`
	Attributes                          string `json:"attributes,omitempty"`
	ResourceURI                         string `json:"resource_uri,omitempty"`
}

// IdentityProviderCreateRequest represents a request to create an identity provider
type IdentityProviderCreateRequest struct {
	Name                                string `json:"name"`
	Description                         string `json:"description,omitempty"`
	IdpType                             string `json:"idp_type"`
	UUID                                string `json:"uuid,omitempty"`
	SSOUrl                              string `json:"sso_url,omitempty"`
	IdpEntityID                         string `json:"idp_entity_id,omitempty"`
	IdpPublicKey                        string `json:"idp_public_key,omitempty"`
	ServiceEntityID                     string `json:"service_entity_id,omitempty"`
	ServicePublicKey                    string `json:"service_public_key,omitempty"`
	ServicePrivateKey                   string `json:"service_private_key,omitempty"`
	SignatureAlgorithm                  string `json:"signature_algorithm"`
	DigestAlgorithm                     string `json:"digest_algorithm"`
	DisplayNameAttributeName            string `json:"display_name_attribute_name,omitempty"`
	RegistrationAliasAttributeName      string `json:"registration_alias_attribute_name,omitempty"`
	AssertionConsumerServiceURL         string `json:"assertion_consumer_service_url,omitempty"`
	AssertionConsumerServiceURL2        string `json:"assertion_consumer_service_url2,omitempty"`
	AssertionConsumerServiceURL3        string `json:"assertion_consumer_service_url3,omitempty"`
	AssertionConsumerServiceURL4        string `json:"assertion_consumer_service_url4,omitempty"`
	AssertionConsumerServiceURL5        string `json:"assertion_consumer_service_url5,omitempty"`
	AssertionConsumerServiceURL6        string `json:"assertion_consumer_service_url6,omitempty"`
	AssertionConsumerServiceURL7        string `json:"assertion_consumer_service_url7,omitempty"`
	AssertionConsumerServiceURL8        string `json:"assertion_consumer_service_url8,omitempty"`
	AssertionConsumerServiceURL9        string `json:"assertion_consumer_service_url9,omitempty"`
	AssertionConsumerServiceURL10       string `json:"assertion_consumer_service_url10,omitempty"`
	WorkerFQDNACSURLs                   bool   `json:"worker_fqdn_acs_urls"`
	DisablePopupFlow                    bool   `json:"disable_popup_flow"`
	OidcFlow                            string `json:"oidc_flow"`
	OidcClientID                        string `json:"oidc_client_id,omitempty"`
	OidcClientSecret                    string `json:"oidc_client_secret,omitempty"`
	OidcTokenURL                        string `json:"oidc_token_url,omitempty"`
	OidcUserInfoURL                     string `json:"oidc_user_info_url,omitempty"`
	OidcJWKSURL                         string `json:"oidc_jwks_url,omitempty"`
	OidcTokenEndpointAuthScheme         string `json:"oidc_token_endpoint_auth_scheme"`
	OidcTokenSignatureScheme            string `json:"oidc_token_signature_scheme"`
	OidcDisplayNameClaimName            string `json:"oidc_display_name_claim_name,omitempty"`
	OidcRegistrationAliasClaimName      string `json:"oidc_registration_alias_claim_name,omitempty"`
	OidcAdditionalScopes                string `json:"oidc_additional_scopes,omitempty"`
	OidcFranceConnectRequiredEidasLevel string `json:"oidc_france_connect_required_eidas_level"`
	Attributes                          string `json:"attributes,omitempty"`
}

// IdentityProviderUpdateRequest represents a request to update an identity provider
type IdentityProviderUpdateRequest struct {
	Name                                string `json:"name,omitempty"`
	Description                         string `json:"description,omitempty"`
	IdpType                             string `json:"idp_type,omitempty"`
	SSOUrl                              string `json:"sso_url,omitempty"`
	IdpEntityID                         string `json:"idp_entity_id,omitempty"`
	IdpPublicKey                        string `json:"idp_public_key,omitempty"`
	ServiceEntityID                     string `json:"service_entity_id,omitempty"`
	ServicePublicKey                    string `json:"service_public_key,omitempty"`
	ServicePrivateKey                   string `json:"service_private_key,omitempty"`
	SignatureAlgorithm                  string `json:"signature_algorithm,omitempty"`
	DigestAlgorithm                     string `json:"digest_algorithm,omitempty"`
	DisplayNameAttributeName            string `json:"display_name_attribute_name,omitempty"`
	RegistrationAliasAttributeName      string `json:"registration_alias_attribute_name,omitempty"`
	AssertionConsumerServiceURL         string `json:"assertion_consumer_service_url,omitempty"`
	AssertionConsumerServiceURL2        string `json:"assertion_consumer_service_url2,omitempty"`
	AssertionConsumerServiceURL3        string `json:"assertion_consumer_service_url3,omitempty"`
	AssertionConsumerServiceURL4        string `json:"assertion_consumer_service_url4,omitempty"`
	AssertionConsumerServiceURL5        string `json:"assertion_consumer_service_url5,omitempty"`
	AssertionConsumerServiceURL6        string `json:"assertion_consumer_service_url6,omitempty"`
	AssertionConsumerServiceURL7        string `json:"assertion_consumer_service_url7,omitempty"`
	AssertionConsumerServiceURL8        string `json:"assertion_consumer_service_url8,omitempty"`
	AssertionConsumerServiceURL9        string `json:"assertion_consumer_service_url9,omitempty"`
	AssertionConsumerServiceURL10       string `json:"assertion_consumer_service_url10,omitempty"`
	WorkerFQDNACSURLs                   *bool  `json:"worker_fqdn_acs_urls,omitempty"`
	DisablePopupFlow                    *bool  `json:"disable_popup_flow,omitempty"`
	OidcFlow                            string `json:"oidc_flow,omitempty"`
	OidcClientID                        string `json:"oidc_client_id,omitempty"`
	OidcClientSecret                    string `json:"oidc_client_secret,omitempty"`
	OidcTokenURL                        string `json:"oidc_token_url,omitempty"`
	OidcUserInfoURL                     string `json:"oidc_user_info_url,omitempty"`
	OidcJWKSURL                         string `json:"oidc_jwks_url,omitempty"`
	OidcTokenEndpointAuthScheme         string `json:"oidc_token_endpoint_auth_scheme,omitempty"`
	OidcTokenSignatureScheme            string `json:"oidc_token_signature_scheme,omitempty"`
	OidcDisplayNameClaimName            string `json:"oidc_display_name_claim_name,omitempty"`
	OidcRegistrationAliasClaimName      string `json:"oidc_registration_alias_claim_name,omitempty"`
	OidcAdditionalScopes                string `json:"oidc_additional_scopes,omitempty"`
	OidcFranceConnectRequiredEidasLevel string `json:"oidc_france_connect_required_eidas_level,omitempty"`
	Attributes                          string `json:"attributes,omitempty"`
}

// IdentityProviderListResponse represents the response from listing identity providers
type IdentityProviderListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []IdentityProvider `json:"objects"`
}
