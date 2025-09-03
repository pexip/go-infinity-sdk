/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package infinity

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/pexip/go-infinity-sdk/v38/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		options []ClientOption
		wantErr bool
	}{
		{
			name:    "default client",
			options: nil,
			wantErr: false,
		},
		{
			name: "with custom base URL",
			options: []ClientOption{
				WithBaseURL("https://api.example.com"),
			},
			wantErr: false,
		},
		{
			name: "with invalid base URL",
			options: []ClientOption{
				WithBaseURL(":/invalid-url"),
			},
			wantErr: true,
		},
		{
			name: "with basic auth",
			options: []ClientOption{
				WithBasicAuth("admin", "password"),
			},
			wantErr: false,
		},
		{
			name: "with token auth",
			options: []ClientOption{
				WithTokenAuth("token123"),
			},
			wantErr: false,
		},
		{
			name: "with user agent",
			options: []ClientOption{
				WithUserAgent("TestApp/1.0"),
			},
			wantErr: false,
		},
		{
			name: "with empty user agent",
			options: []ClientOption{
				WithUserAgent(""),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := New(tt.options...)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, client)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, client)
			assert.NotNil(t, client.config)
			assert.NotNil(t, client.status)
			assert.NotNil(t, client.history)
			assert.NotNil(t, client.command)
		})
	}
}

func TestWithBaseURL(t *testing.T) {
	tests := []struct {
		name    string
		baseURL string
		wantErr bool
	}{
		{
			name:    "valid URL",
			baseURL: "https://api.example.com",
			wantErr: false,
		},
		{
			name:    "invalid URL",
			baseURL: ":/invalid-url",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{}
			option := WithBaseURL(tt.baseURL)
			err := option(client)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			expectedURL, _ := url.Parse(tt.baseURL)
			assert.Equal(t, expectedURL, client.baseURL)
		})
	}
}

func TestWithHTTPClient(t *testing.T) {
	customClient := &http.Client{
		Timeout: 60 * time.Second,
	}

	client := &Client{}
	option := WithHTTPClient(customClient)
	err := option(client)

	assert.NoError(t, err)
	assert.Equal(t, customClient, client.httpClient)
}

func TestWithAuth(t *testing.T) {
	authenticator := auth.NewBasicAuth("user", "pass")

	client := &Client{}
	option := WithAuth(authenticator)
	err := option(client)

	assert.NoError(t, err)
	assert.Equal(t, authenticator, client.auth)
}

func TestWithUserAgent(t *testing.T) {
	tests := []struct {
		name      string
		userAgent string
		wantErr   bool
		errorMsg  string
	}{
		{
			name:      "valid user agent",
			userAgent: "MyApp/1.0 (https://example.com)",
			wantErr:   false,
		},
		{
			name:      "another valid user agent",
			userAgent: "go-infinity-sdk/v38",
			wantErr:   false,
		},
		{
			name:      "empty user agent",
			userAgent: "",
			wantErr:   true,
			errorMsg:  "user agent cannot be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{}
			option := WithUserAgent(tt.userAgent)
			err := option(client)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.userAgent, client.userAgent)
		})
	}
}

func TestWithTransport(t *testing.T) {
	customTransport := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	tests := []struct {
		name        string
		setupClient func() *Client
		transport   http.RoundTripper
		wantErr     bool
		errorMsg    string
		validate    func(t *testing.T, client *Client, transport http.RoundTripper)
	}{
		{
			name: "set transport on default client",
			setupClient: func() *Client {
				client, _ := New()
				return client
			},
			transport: customTransport,
			wantErr:   false,
			validate: func(t *testing.T, client *Client, transport http.RoundTripper) {
				assert.Equal(t, transport, client.httpClient.Transport)
				assert.Equal(t, DefaultTimeout, client.httpClient.Timeout)
			},
		},
		{
			name: "set transport on client with custom timeout",
			setupClient: func() *Client {
				customClient := &http.Client{Timeout: 60 * time.Second}
				client, _ := New(WithHTTPClient(customClient))
				return client
			},
			transport: customTransport,
			wantErr:   false,
			validate: func(t *testing.T, client *Client, transport http.RoundTripper) {
				assert.Equal(t, transport, client.httpClient.Transport)
				assert.Equal(t, 60*time.Second, client.httpClient.Timeout)
			},
		},
		{
			name: "nil transport returns error",
			setupClient: func() *Client {
				client, _ := New()
				return client
			},
			transport: nil,
			wantErr:   true,
			errorMsg:  "transport cannot be nil",
			validate: func(t *testing.T, client *Client, transport http.RoundTripper) {
				// Should not be called for error case
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := tt.setupClient()
			option := WithTransport(tt.transport)
			err := option(client)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)
			} else {
				assert.NoError(t, err)
				tt.validate(t, client, tt.transport)
			}
		})
	}
}

func TestWithTransport_Integration(t *testing.T) {
	// Test that WithTransport works when creating a new client
	customTransport := &http.Transport{
		MaxIdleConns:       20,
		IdleConnTimeout:    60 * time.Second,
		DisableCompression: true,
	}

	client, err := New(
		WithBaseURL("https://api.example.com"),
		WithTransport(customTransport),
		WithBasicAuth("user", "pass"),
	)

	require.NoError(t, err)
	assert.Equal(t, customTransport, client.httpClient.Transport)
	assert.Equal(t, DefaultTimeout, client.httpClient.Timeout)
}

func TestWithTransport_WithCustomHTTPClient(t *testing.T) {
	// Test precedence when both WithHTTPClient and WithTransport are used
	customTransport := &http.Transport{
		MaxIdleConns: 50,
	}

	customClient := &http.Client{
		Timeout: 45 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns: 100, // This should be overridden
		},
	}

	tests := []struct {
		name    string
		options []ClientOption
		check   func(t *testing.T, client *Client)
	}{
		{
			name: "WithTransport after WithHTTPClient",
			options: []ClientOption{
				WithHTTPClient(customClient),
				WithTransport(customTransport),
			},
			check: func(t *testing.T, client *Client) {
				assert.Equal(t, customTransport, client.httpClient.Transport)
				assert.Equal(t, 45*time.Second, client.httpClient.Timeout)
			},
		},
		{
			name: "WithHTTPClient after WithTransport",
			options: []ClientOption{
				WithTransport(customTransport),
				WithHTTPClient(customClient),
			},
			check: func(t *testing.T, client *Client) {
				// WithHTTPClient should completely replace the client
				assert.Equal(t, customClient.Transport, client.httpClient.Transport)
				assert.Equal(t, 45*time.Second, client.httpClient.Timeout)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := New(tt.options...)
			require.NoError(t, err)
			tt.check(t, client)
		})
	}
}

func TestDoRequest(t *testing.T) {
	tests := []struct {
		name           string
		setupServer    func() *httptest.Server
		request        *Request
		expectedStatus int
		expectError    bool
		errorType      string
	}{
		{
			name: "successful GET request",
			setupServer: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					assert.Equal(t, http.MethodGet, r.Method)
					assert.Equal(t, "/api/admin/test", r.URL.Path)
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
				}))
			},
			request: &Request{
				Method:   http.MethodGet,
				Endpoint: "test",
			},
			expectedStatus: http.StatusOK,
			expectError:    false,
		},
		{
			name: "successful POST request with body",
			setupServer: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					assert.Equal(t, http.MethodPost, r.Method)
					assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

					var body map[string]string
					_ = json.NewDecoder(r.Body).Decode(&body)
					assert.Equal(t, "test", body["name"])

					w.WriteHeader(http.StatusCreated)
					_ = json.NewEncoder(w).Encode(map[string]interface{}{"id": 1, "name": "test"})
				}))
			},
			request: &Request{
				Method:   http.MethodPost,
				Endpoint: "test",
				Body:     map[string]string{"name": "test"},
			},
			expectedStatus: http.StatusCreated,
			expectError:    false,
		},
		{
			name: "API error response",
			setupServer: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusBadRequest)
					_ = json.NewEncoder(w).Encode(map[string]interface{}{
						"error":   "Bad Request",
						"details": "Invalid parameter",
					})
				}))
			},
			request: &Request{
				Method:   http.MethodGet,
				Endpoint: "test",
			},
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
			errorType:      "*infinity.APIError",
		},
		{
			name: "request with query parameters",
			setupServer: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					assert.Equal(t, "10", r.URL.Query().Get("limit"))
					assert.Equal(t, "5", r.URL.Query().Get("offset"))
					w.WriteHeader(http.StatusOK)
					_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
				}))
			},
			request: &Request{
				Method:   http.MethodGet,
				Endpoint: "test",
				QueryParams: url.Values{
					"limit":  []string{"10"},
					"offset": []string{"5"},
				},
			},
			expectedStatus: http.StatusOK,
			expectError:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := tt.setupServer()
			defer server.Close()

			client, err := New(WithBaseURL(server.URL))
			require.NoError(t, err)

			resp, err := client.DoRequest(t.Context(), tt.request)

			if tt.expectError {
				assert.Error(t, err)
				if tt.errorType != "" {
					assert.Contains(t, fmt.Sprintf("%T", err), tt.errorType)
				}
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)
		})
	}
}

func TestClientWithAuthentication(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeaderString := r.Header.Get("Authorization")
		assert.Contains(t, authHeaderString, "Basic")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "authenticated"})
	}))
	defer server.Close()

	client, err := New(
		WithBaseURL(server.URL),
		WithBasicAuth("admin", "password"),
	)
	require.NoError(t, err)

	req := &Request{
		Method:   http.MethodGet,
		Endpoint: "test",
	}

	resp, err := client.DoRequest(t.Context(), req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestClientWithUserAgent(t *testing.T) {
	expectedUserAgent := "TestApp/1.0 (test environment)"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userAgent := r.Header.Get("User-Agent")
		assert.Equal(t, expectedUserAgent, userAgent)
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	}))
	defer server.Close()

	client, err := New(
		WithBaseURL(server.URL),
		WithUserAgent(expectedUserAgent),
	)
	require.NoError(t, err)

	req := &Request{
		Method:   http.MethodGet,
		Endpoint: "test",
	}

	resp, err := client.DoRequest(t.Context(), req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestClientWithoutUserAgent(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userAgent := r.Header.Get("User-Agent")
		// Should use the default Go HTTP client User-Agent when no custom one is set
		assert.Equal(t, "Go-http-client/1.1", userAgent)
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	}))
	defer server.Close()

	client, err := New(
		WithBaseURL(server.URL),
	)
	require.NoError(t, err)

	req := &Request{
		Method:   http.MethodGet,
		Endpoint: "test",
	}

	resp, err := client.DoRequest(t.Context(), req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"id":   1,
			"name": "test",
		})
	}))
	defer server.Close()

	client, err := New(WithBaseURL(server.URL))
	require.NoError(t, err)

	var result map[string]interface{}
	err = client.GetJSON(t.Context(), "test", nil, &result)

	assert.NoError(t, err)
	assert.Equal(t, float64(1), result["id"])
	assert.Equal(t, "test", result["name"])
}

func TestPostJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)

		var body map[string]string
		_ = json.NewDecoder(r.Body).Decode(&body)
		assert.Equal(t, "test", body["name"])

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"id":   1,
			"name": body["name"],
		})
	}))
	defer server.Close()

	client, err := New(WithBaseURL(server.URL))
	require.NoError(t, err)

	requestBody := map[string]string{"name": "test"}
	var result map[string]interface{}
	err = client.PostJSON(t.Context(), "test", requestBody, &result)

	assert.NoError(t, err)
	assert.Equal(t, float64(1), result["id"])
	assert.Equal(t, "test", result["name"])
}

func TestAPIError(t *testing.T) {
	err := &APIError{
		StatusCode: 400,
		Message:    "Bad Request",
		Details:    "Invalid parameter",
	}

	expected := "API error 400: Bad Request (Invalid parameter)"
	assert.Equal(t, expected, err.Error())

	errWithoutDetails := &APIError{
		StatusCode: 404,
		Message:    "Not Found",
	}

	expected = "API error 404: Not Found"
	assert.Equal(t, expected, errWithoutDetails.Error())
}

func TestAPIError_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name          string
		jsonData      string
		expectedError APIError
	}{
		{
			name:     "error field format",
			jsonData: `{"error": "Authentication failed", "details": "Invalid credentials"}`,
			expectedError: APIError{
				Message: "Authentication failed",
				Details: "Invalid credentials",
			},
		},
		{
			name:     "message field format",
			jsonData: `{"message": "Resource not found", "detail": "Conference with ID 123 does not exist"}`,
			expectedError: APIError{
				Message: "Resource not found",
				Details: "Conference with ID 123 does not exist",
			},
		},
		{
			name:     "error field takes priority over message",
			jsonData: `{"error": "Primary error", "message": "Secondary message", "details": "Error details"}`,
			expectedError: APIError{
				Message: "Primary error",
				Details: "Error details",
			},
		},
		{
			name:     "details field takes priority over detail",
			jsonData: `{"error": "Error message", "details": "Primary details", "detail": "Secondary detail"}`,
			expectedError: APIError{
				Message: "Error message",
				Details: "Primary details",
			},
		},
		{
			name:     "empty json object",
			jsonData: `{}`,
			expectedError: APIError{
				Message: "",
				Details: "",
			},
		},
		{
			name:     "only error field",
			jsonData: `{"error": "Simple error"}`,
			expectedError: APIError{
				Message: "Simple error",
				Details: "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var apiErr APIError
			err := json.Unmarshal([]byte(tt.jsonData), &apiErr)

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedError.Message, apiErr.Message)
			assert.Equal(t, tt.expectedError.Details, apiErr.Details)
		})
	}
}

func TestAPIError_UnmarshalJSON_InvalidJSON(t *testing.T) {
	var apiErr APIError
	err := json.Unmarshal([]byte(`{invalid json`), &apiErr)

	assert.Error(t, err)
}

func TestClient_handleAPIError(t *testing.T) {
	client, err := New()
	require.NoError(t, err)

	tests := []struct {
		name         string
		statusCode   int
		responseBody string
		expectedMsg  string
		expectedDet  string
	}{
		{
			name:         "error with structured JSON response",
			statusCode:   400,
			responseBody: `{"error": "Invalid request", "details": "Missing required field"}`,
			expectedMsg:  "Invalid request",
			expectedDet:  "Missing required field",
		},
		{
			name:         "error with message JSON response",
			statusCode:   404,
			responseBody: `{"message": "Not found", "detail": "Resource does not exist"}`,
			expectedMsg:  "Not found",
			expectedDet:  "Resource does not exist",
		},
		{
			name:         "error with invalid JSON response",
			statusCode:   500,
			responseBody: `{invalid json}`,
			expectedMsg:  "Internal Server Error", // Falls back to HTTP status text
			expectedDet:  "{invalid json}",
		},
		{
			name:         "error with empty response body",
			statusCode:   403,
			responseBody: "",
			expectedMsg:  "Forbidden", // Falls back to HTTP status text
			expectedDet:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &Response{
				StatusCode: tt.statusCode,
				Body:       []byte(tt.responseBody),
			}

			err := client.handleAPIError(resp)
			require.Error(t, err)

			var apiErr *APIError
			ok := errors.As(err, &apiErr)
			require.True(t, ok, "Error should be of type *APIError")

			assert.Equal(t, tt.statusCode, apiErr.StatusCode)
			assert.Equal(t, tt.expectedMsg, apiErr.Message)
			assert.Equal(t, tt.expectedDet, apiErr.Details)
		})
	}
}

func TestClient_DoRequest_EdgeCases(t *testing.T) {
	tests := []struct {
		name          string
		request       *Request
		setupClient   func() *Client
		expectError   bool
		errorContains string
	}{
		{
			name: "malformed request body JSON",
			request: &Request{
				Method:   "POST",
				Endpoint: "test",
				Body:     make(chan int), // This cannot be marshaled to JSON
			},
			setupClient: func() *Client {
				client, _ := New(WithBaseURL("https://example.com"))
				return client
			},
			expectError:   true,
			errorContains: "failed to marshal request body",
		},
		{
			name: "authentication error",
			request: &Request{
				Method:   "GET",
				Endpoint: "test",
			},
			setupClient: func() *Client {
				client, _ := New(
					WithBaseURL("https://example.com"),
					WithAuth(&FailingAuth{}),
				)
				return client
			},
			expectError:   true,
			errorContains: "failed to authenticate request",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := tt.setupClient()
			_, err := client.DoRequest(context.Background(), tt.request)

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorContains)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestClient_sleepWithContext(t *testing.T) {
	client, err := New()
	require.NoError(t, err)

	t.Run("zero duration returns immediately", func(t *testing.T) {
		start := time.Now()
		result := client.sleepWithContext(context.Background(), 0)
		duration := time.Since(start)

		assert.True(t, result)
		assert.Less(t, duration, 10*time.Millisecond)
	})

	t.Run("negative duration returns immediately", func(t *testing.T) {
		start := time.Now()
		result := client.sleepWithContext(context.Background(), -100*time.Millisecond)
		duration := time.Since(start)

		assert.True(t, result)
		assert.Less(t, duration, 10*time.Millisecond)
	})

	t.Run("context cancellation interrupts sleep", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())

		// Cancel context after 10ms
		go func() {
			time.Sleep(10 * time.Millisecond)
			cancel()
		}()

		start := time.Now()
		result := client.sleepWithContext(ctx, 100*time.Millisecond)
		duration := time.Since(start)

		assert.False(t, result)
		assert.Less(t, duration, 50*time.Millisecond) // Should be much less than 100ms
	})

	t.Run("normal sleep completes", func(t *testing.T) {
		start := time.Now()
		result := client.sleepWithContext(context.Background(), 20*time.Millisecond)
		duration := time.Since(start)

		assert.True(t, result)
		assert.GreaterOrEqual(t, duration, 15*time.Millisecond)
		assert.Less(t, duration, 50*time.Millisecond)
	})
}

func TestUnmarshalResponseBody_EdgeCases(t *testing.T) {
	t.Run("nil result pointer", func(t *testing.T) {
		body := []byte(`{"test": "value"}`)
		err := unmarshalResponseBody(body, nil)
		assert.NoError(t, err)
	})

	t.Run("empty body with non-nil result", func(t *testing.T) {
		var result map[string]string
		err := unmarshalResponseBody([]byte{}, &result)
		assert.NoError(t, err)
		assert.Nil(t, result)
	})

	t.Run("malformed JSON", func(t *testing.T) {
		var result map[string]string
		body := []byte(`{"test": "value",}`)
		err := unmarshalResponseBody(body, &result)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to unmarshal JSON response")
	})

	t.Run("type mismatch", func(t *testing.T) {
		var result int
		body := []byte(`{"test": "value"}`)
		err := unmarshalResponseBody(body, &result)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to unmarshal JSON response")
	})
}

// FailingAuth is a test authenticator that always fails
type FailingAuth struct{}

func (f *FailingAuth) Authenticate(req *http.Request) error {
	return errors.New("authentication failed")
}

func TestClient_HttpClient(t *testing.T) {
	customClient := &http.Client{
		Timeout: 45 * time.Second,
	}

	client, err := New(WithHTTPClient(customClient))
	require.NoError(t, err)

	// Test that HttpClient returns the same client we set
	retrievedClient := client.HttpClient()
	assert.Equal(t, customClient, retrievedClient)
	assert.Equal(t, 45*time.Second, retrievedClient.Timeout)

	// Test with default client
	defaultClient, err := New()
	require.NoError(t, err)

	retrievedDefault := defaultClient.HttpClient()
	assert.NotNil(t, retrievedDefault)
	assert.Equal(t, DefaultTimeout, retrievedDefault.Timeout)
}
