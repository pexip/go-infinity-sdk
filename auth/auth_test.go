package auth

import (
	"encoding/base64"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBasicAuth_Authenticate(t *testing.T) {
	auth := NewBasicAuth("admin", "password")
	req, err := http.NewRequest("GET", "https://example.com", nil)
	require.NoError(t, err)

	err = auth.Authenticate(req)
	assert.NoError(t, err)

	expectedAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte("admin:password"))
	assert.Equal(t, expectedAuth, req.Header.Get("Authorization"))
}

func TestTokenAuth_Authenticate(t *testing.T) {
	auth := NewTokenAuth("my-secret-token")
	req, err := http.NewRequest("GET", "https://example.com", nil)
	require.NoError(t, err)

	err = auth.Authenticate(req)
	assert.NoError(t, err)

	assert.Equal(t, "Token my-secret-token", req.Header.Get("Authorization"))
}

func TestBearerAuth_Authenticate(t *testing.T) {
	auth := NewBearerAuth("my-bearer-token")
	req, err := http.NewRequest("GET", "https://example.com", nil)
	require.NoError(t, err)

	err = auth.Authenticate(req)
	assert.NoError(t, err)

	assert.Equal(t, "Bearer my-bearer-token", req.Header.Get("Authorization"))
}

func TestCustomAuth_Authenticate(t *testing.T) {
	tests := []struct {
		name        string
		headerName  string
		headerValue string
		wantErr     bool
		expectedErr string
	}{
		{
			name:        "valid custom auth",
			headerName:  "X-API-Key",
			headerValue: "my-api-key",
			wantErr:     false,
		},
		{
			name:        "empty header name",
			headerName:  "",
			headerValue: "value",
			wantErr:     true,
			expectedErr: "header name cannot be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auth := NewCustomAuth(tt.headerName, tt.headerValue)
			req, err := http.NewRequest("GET", "https://example.com", nil)
			require.NoError(t, err)

			err = auth.Authenticate(req)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.headerValue, req.Header.Get(tt.headerName))
			}
		})
	}
}

func TestNewBasicAuth(t *testing.T) {
	auth := NewBasicAuth("user", "pass")
	assert.Equal(t, "user", auth.username)
	assert.Equal(t, "pass", auth.password)
}

func TestNewTokenAuth(t *testing.T) {
	auth := NewTokenAuth("token123")
	assert.Equal(t, "token123", auth.token)
}

func TestNewBearerAuth(t *testing.T) {
	auth := NewBearerAuth("bearer123")
	assert.Equal(t, "bearer123", auth.token)
}

func TestNewCustomAuth(t *testing.T) {
	auth := NewCustomAuth("X-Header", "value")
	assert.Equal(t, "X-Header", auth.headerName)
	assert.Equal(t, "value", auth.headerValue)
}
