package infinity

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/pexip/go-infinity-sdk/auth"
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
			assert.NotNil(t, client.Config)
			assert.NotNil(t, client.Status)
			assert.NotNil(t, client.History)
			assert.NotNil(t, client.Command)
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
	err = client.GetJSON(t.Context(), "test", &result)

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
			expectedDet:  "",
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

			apiErr, ok := err.(*APIError)
			require.True(t, ok, "Error should be of type *APIError")

			assert.Equal(t, tt.statusCode, apiErr.StatusCode)
			assert.Equal(t, tt.expectedMsg, apiErr.Message)
			assert.Equal(t, tt.expectedDet, apiErr.Details)
		})
	}
}
