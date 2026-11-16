/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// marshalToMap marshals v to JSON and decodes it into a map so tests can
// check which keys are present and their values.
func marshalToMap(t *testing.T, v any) map[string]any {
	t.Helper()
	data, err := json.Marshal(v)
	require.NoError(t, err)
	var out map[string]any
	require.NoError(t, json.Unmarshal(data, &out))
	return out
}

func TestIdentityProviderCreateRequest_DeviceIdpJSON(t *testing.T) {
	t.Run("defaults", func(t *testing.T) {
		body := marshalToMap(t, &IdentityProviderCreateRequest{Name: "idp"})

		// oidc_device_idp has no omitempty, so false is always sent.
		assert.Equal(t, false, body["oidc_device_idp"])
		// Unset base URL and duration are omitted so the API defaults apply.
		assert.NotContains(t, body, "oidc_device_idp_baseurl")
		assert.NotContains(t, body, "oidc_device_idp_session_duration")
	})

	t.Run("values", func(t *testing.T) {
		body := marshalToMap(t, &IdentityProviderCreateRequest{
			Name:                         "idp",
			OidcDeviceIdp:                true,
			OidcDeviceIdpBaseURL:         "https://device-idp.example.com",
			OidcDeviceIdpSessionDuration: 7200,
		})

		assert.Equal(t, true, body["oidc_device_idp"])
		assert.Equal(t, "https://device-idp.example.com", body["oidc_device_idp_baseurl"])
		assert.Equal(t, float64(7200), body["oidc_device_idp_session_duration"])
	})
}

func TestIdentityProviderUpdateRequest_DeviceIdpJSON(t *testing.T) {
	t.Run("nil pointers are omitted", func(t *testing.T) {
		body := marshalToMap(t, &IdentityProviderUpdateRequest{Name: "idp"})

		assert.NotContains(t, body, "oidc_device_idp")
		assert.NotContains(t, body, "oidc_device_idp_session_duration")
	})

	t.Run("explicit false is sent", func(t *testing.T) {
		disabled := false
		body := marshalToMap(t, &IdentityProviderUpdateRequest{Name: "idp", OidcDeviceIdp: &disabled})

		require.Contains(t, body, "oidc_device_idp")
		assert.Equal(t, false, body["oidc_device_idp"])
	})

	t.Run("empty base URL is retained", func(t *testing.T) {
		body := marshalToMap(t, &IdentityProviderUpdateRequest{Name: "idp"})

		require.Contains(t, body, "oidc_device_idp_baseurl")
		assert.Equal(t, "", body["oidc_device_idp_baseurl"])
	})

	t.Run("values", func(t *testing.T) {
		enabled := true
		duration := 3600
		body := marshalToMap(t, &IdentityProviderUpdateRequest{
			Name:                         "idp",
			OidcDeviceIdp:                &enabled,
			OidcDeviceIdpBaseURL:         "https://device-idp.example.com",
			OidcDeviceIdpSessionDuration: &duration,
		})

		assert.Equal(t, true, body["oidc_device_idp"])
		assert.Equal(t, "https://device-idp.example.com", body["oidc_device_idp_baseurl"])
		assert.Equal(t, float64(3600), body["oidc_device_idp_session_duration"])
	})
}

func TestIdentityProvider_DecodeDeviceIdpFields(t *testing.T) {
	fixture := `{
		"id": 1,
		"name": "idp",
		"idp_type": "oidc",
		"oidc_device_idp": true,
		"oidc_device_idp_baseurl": "https://device-idp.example.com",
		"oidc_device_idp_session_duration": 172800,
		"resource_uri": "/api/admin/configuration/v1/identity_provider/1/"
	}`

	var idp IdentityProvider
	require.NoError(t, json.Unmarshal([]byte(fixture), &idp))

	assert.True(t, idp.OidcDeviceIdp)
	assert.Equal(t, "https://device-idp.example.com", idp.OidcDeviceIdpBaseURL)
	assert.Equal(t, 172800, idp.OidcDeviceIdpSessionDuration)
}
