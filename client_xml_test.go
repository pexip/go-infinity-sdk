package infinity

import (
	"context"
	"encoding/xml"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestResponse represents a test response structure for XML unmarshaling
type TestResponse struct {
	XMLName xml.Name `xml:"response"`
	Message string   `xml:"message"`
	ID      int      `xml:"id"`
}

func TestPostWithResponse_XML(t *testing.T) {
	// Test XML response with location header
	xmlResponse := `<?xml version="1.0" encoding="UTF-8"?>
<response>
    <message>Resource created successfully</message>
    <id>12345</id>
</response>`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		w.Header().Set("Location", "https://example.com/api/admin/resource/12345/")
		w.Header().Set("Content-Type", "application/xml")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(xmlResponse))
	}))
	defer server.Close()

	client, err := New(WithBaseURL(server.URL))
	require.NoError(t, err)

	var result TestResponse
	postResp, err := client.PostWithResponse(context.Background(), "test/endpoint", map[string]string{"name": "test"}, &result)

	require.NoError(t, err)
	assert.NotNil(t, postResp)
	assert.Equal(t, "https://example.com/api/admin/resource/12345/", postResp.ResourceURI)
	assert.Equal(t, "Resource created successfully", result.Message)
	assert.Equal(t, 12345, result.ID)
}

func TestPostWithResponse_JSON(t *testing.T) {
	// Test JSON response with location header
	jsonResponse := `{"message": "Resource created successfully", "id": 12345}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)

		w.Header().Set("Location", "https://example.com/api/admin/resource/12345/")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(jsonResponse))
	}))
	defer server.Close()

	client, err := New(WithBaseURL(server.URL))
	require.NoError(t, err)

	var result map[string]interface{}
	postResp, err := client.PostWithResponse(context.Background(), "test/endpoint", map[string]string{"name": "test"}, &result)

	require.NoError(t, err)
	assert.NotNil(t, postResp)
	assert.Equal(t, "https://example.com/api/admin/resource/12345/", postResp.ResourceURI)
	assert.Equal(t, "Resource created successfully", result["message"])
	assert.Equal(t, float64(12345), result["id"]) // JSON numbers become float64
}

func TestPostWithResponse_NoLocationHeader(t *testing.T) {
	// Test response without location header
	jsonResponse := `{"message": "Success"}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonResponse))
	}))
	defer server.Close()

	client, err := New(WithBaseURL(server.URL))
	require.NoError(t, err)

	var result map[string]interface{}
	postResp, err := client.PostWithResponse(context.Background(), "test/endpoint", nil, &result)

	require.NoError(t, err)
	assert.NotNil(t, postResp)
	assert.Empty(t, postResp.ResourceURI) // Should be empty when no location header
	assert.Equal(t, "Success", result["message"])
}

func TestUnmarshalResponseBodyAuto(t *testing.T) {
	tests := []struct {
		name        string
		body        []byte
		expected    interface{}
		expectError bool
	}{
		{
			name: "valid JSON",
			body: []byte(`{"message": "test", "value": 123}`),
			expected: map[string]interface{}{
				"message": "test",
				"value":   float64(123),
			},
			expectError: false,
		},
		{
			name: "valid XML",
			body: []byte(`<?xml version="1.0"?><response><message>test</message><id>123</id></response>`),
			expected: TestResponse{
				XMLName: xml.Name{Local: "response"},
				Message: "test",
				ID:      123,
			},
			expectError: false,
		},
		{
			name:        "invalid both JSON and XML",
			body:        []byte(`invalid content`),
			expected:    nil,
			expectError: true,
		},
		{
			name:        "empty body",
			body:        []byte{},
			expected:    nil,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result interface{}

			if tt.name == "valid XML" {
				result = &TestResponse{}
			} else {
				result = &map[string]interface{}{}
			}

			err := unmarshalResponseBodyAuto(tt.body, result)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				if tt.expected != nil {
					if tt.name == "valid XML" {
						assert.Equal(t, tt.expected, *result.(*TestResponse))
					} else if tt.name == "valid JSON" {
						assert.Equal(t, tt.expected, *result.(*map[string]interface{}))
					}
				}
			}
		})
	}
}

func TestTypes_PostResponse(t *testing.T) {
	// Test that the PostResponse type is properly accessible
	postResp := &types.PostResponse{
		Body:        []byte("test body"),
		ResourceURI: "https://example.com/resource/123",
	}

	assert.Equal(t, []byte("test body"), postResp.Body)
	assert.Equal(t, "https://example.com/resource/123", postResp.ResourceURI)
}
