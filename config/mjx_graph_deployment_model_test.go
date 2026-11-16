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

// marshalMjxGraphDeploymentToMap marshals v to JSON and decodes it into a map
// so tests can check which keys are present and their values.
func marshalMjxGraphDeploymentToMap(t *testing.T, v any) map[string]any {
	t.Helper()
	data, err := json.Marshal(v)
	require.NoError(t, err)
	var out map[string]any
	require.NoError(t, json.Unmarshal(data, &out))
	return out
}

func TestMjxGraphDeploymentUpdateRequest_OAuthJSON(t *testing.T) {
	t.Run("omitted private key is not sent", func(t *testing.T) {
		body := marshalMjxGraphDeploymentToMap(t, &MjxGraphDeploymentUpdateRequest{})

		// Sending null would wipe the stored key, so a nil key must be omitted.
		assert.NotContains(t, body, "oauth_private_key")
	})

	t.Run("supplied private key is sent", func(t *testing.T) {
		key := "-----BEGIN PRIVATE KEY-----\ntest key\n-----END PRIVATE KEY-----"
		body := marshalMjxGraphDeploymentToMap(t, &MjxGraphDeploymentUpdateRequest{OAuthPrivateKey: &key})

		assert.Equal(t, key, body["oauth_private_key"])
	})

	t.Run("empty certificate is sent so it can be cleared", func(t *testing.T) {
		body := marshalMjxGraphDeploymentToMap(t, &MjxGraphDeploymentUpdateRequest{})

		require.Contains(t, body, "oauth_certificate")
		assert.Equal(t, "", body["oauth_certificate"])
	})

	t.Run("unset authentication method is omitted", func(t *testing.T) {
		body := marshalMjxGraphDeploymentToMap(t, &MjxGraphDeploymentUpdateRequest{})

		assert.NotContains(t, body, "authentication_method")
	})
}
