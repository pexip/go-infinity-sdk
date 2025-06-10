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
	auth := NewCustomAuth("X-API-Key", "my-api-key")
	req, err := http.NewRequest("GET", "https://example.com", nil)
	require.NoError(t, err)

	err = auth.Authenticate(req)
	assert.NoError(t, err)
	assert.Equal(t, "my-api-key", req.Header.Get("X-API-Key"))
}

func TestAuthenticators_HandleSpecialCharacters(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() (Authenticator, *http.Request)
		validate func(t *testing.T, req *http.Request)
	}{
		{
			name: "basic auth with special characters",
			setup: func() (Authenticator, *http.Request) {
				auth := NewBasicAuth("user@domain.com", "p@ssw0rd!#$%")
				req, _ := http.NewRequest("GET", "https://example.com", nil)
				return auth, req
			},
			validate: func(t *testing.T, req *http.Request) {
				expected := "Basic " + base64.StdEncoding.EncodeToString([]byte("user@domain.com:p@ssw0rd!#$%"))
				assert.Equal(t, expected, req.Header.Get("Authorization"))
			},
		},
		{
			name: "token auth with unicode characters",
			setup: func() (Authenticator, *http.Request) {
				auth := NewTokenAuth("tôken-with-ünïcödé")
				req, _ := http.NewRequest("GET", "https://example.com", nil)
				return auth, req
			},
			validate: func(t *testing.T, req *http.Request) {
				assert.Equal(t, "Token tôken-with-ünïcödé", req.Header.Get("Authorization"))
			},
		},
		{
			name: "custom auth with multiple header overwrites",
			setup: func() (Authenticator, *http.Request) {
				auth := NewCustomAuth("X-API-Key", "new-value")
				req, _ := http.NewRequest("GET", "https://example.com", nil)
				req.Header.Set("X-API-Key", "old-value")
				return auth, req
			},
			validate: func(t *testing.T, req *http.Request) {
				assert.Equal(t, "new-value", req.Header.Get("X-API-Key"))
				// Should only have one value, not both
				assert.Len(t, req.Header.Values("X-API-Key"), 1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auth, req := tt.setup()
			err := auth.Authenticate(req)
			assert.NoError(t, err)
			tt.validate(t, req)
		})
	}
}

func TestBasicAuth_EmptyCredentials(t *testing.T) {
	tests := []struct {
		name     string
		username string
		password string
		expected string
	}{
		{
			name:     "empty username",
			username: "",
			password: "password",
			expected: "Basic " + base64.StdEncoding.EncodeToString([]byte(":password")),
		},
		{
			name:     "empty password",
			username: "username",
			password: "",
			expected: "Basic " + base64.StdEncoding.EncodeToString([]byte("username:")),
		},
		{
			name:     "both empty",
			username: "",
			password: "",
			expected: "Basic " + base64.StdEncoding.EncodeToString([]byte(":")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auth := NewBasicAuth(tt.username, tt.password)
			req, err := http.NewRequest("GET", "https://example.com", nil)
			require.NoError(t, err)

			err = auth.Authenticate(req)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, req.Header.Get("Authorization"))
		})
	}
}

func TestTokenAuth_EmptyToken(t *testing.T) {
	auth := NewTokenAuth("")
	req, err := http.NewRequest("GET", "https://example.com", nil)
	require.NoError(t, err)

	err = auth.Authenticate(req)
	assert.NoError(t, err)
	assert.Equal(t, "Token ", req.Header.Get("Authorization"))
}

func TestBearerAuth_EmptyToken(t *testing.T) {
	auth := NewBearerAuth("")
	req, err := http.NewRequest("GET", "https://example.com", nil)
	require.NoError(t, err)

	err = auth.Authenticate(req)
	assert.NoError(t, err)
	assert.Equal(t, "Bearer ", req.Header.Get("Authorization"))
}

func TestCustomAuth_EdgeCases(t *testing.T) {
	tests := []struct {
		name        string
		headerName  string
		headerValue string
		wantErr     bool
		expectedErr string
	}{
		{
			name:        "empty header name",
			headerName:  "",
			headerValue: "value",
			wantErr:     true,
			expectedErr: "header name cannot be empty",
		},
		{
			name:        "whitespace-only header name",
			headerName:  "   ",
			headerValue: "value",
			wantErr:     false, // HTTP allows this, though it's unusual
		},
		{
			name:        "empty header value",
			headerName:  "X-API-Key",
			headerValue: "",
			wantErr:     false,
		},
		{
			name:        "header name with special characters",
			headerName:  "X-My-Custom-Header",
			headerValue: "value",
			wantErr:     false,
		},
		{
			name:        "header value with newlines",
			headerName:  "X-API-Key",
			headerValue: "value\nwith\nnewlines",
			wantErr:     false, // HTTP will handle this
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
				if tt.headerName != "" {
					assert.Equal(t, tt.headerValue, req.Header.Get(tt.headerName))
				}
			}
		})
	}
}
