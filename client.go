package infinity

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/pexip/go-infinity-sdk/auth"
	"github.com/pexip/go-infinity-sdk/command"
	"github.com/pexip/go-infinity-sdk/config"
	"github.com/pexip/go-infinity-sdk/history"
	"github.com/pexip/go-infinity-sdk/status"
)

const (
	DefaultBaseURL = "https://admin.example.com"
	DefaultTimeout = 30 * time.Second
	APIPrefix      = "/api/admin/"
)

// Client represents a Pexip Infinity Management API client
type Client struct {
	baseURL    *url.URL
	httpClient *http.Client
	auth       auth.Authenticator

	// API services
	Config  *config.Service
	Status  *status.Service
	History *history.Service
	Command *command.Service
}

// ClientOption is a function that configures a Client
type ClientOption func(*Client) error

// New creates a new Infinity API client with the given options
func New(options ...ClientOption) (*Client, error) {
	baseURL, err := url.Parse(DefaultBaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse default base URL: %w", err)
	}

	c := &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
	}

	for _, option := range options {
		if err := option(c); err != nil {
			return nil, err
		}
	}

	// Initialize API services
	c.Config = config.New(c)
	c.Status = status.New(c)
	c.History = history.New(c)
	c.Command = command.New(c)

	return c, nil
}

// WithBaseURL sets the base URL for the client
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		u, err := url.Parse(baseURL)
		if err != nil {
			return fmt.Errorf("failed to parse base URL: %w", err)
		}
		c.baseURL = u
		return nil
	}
}

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) error {
		c.httpClient = httpClient
		return nil
	}
}

// WithAuth sets the authentication method
func WithAuth(authenticator auth.Authenticator) ClientOption {
	return func(c *Client) error {
		c.auth = authenticator
		return nil
	}
}

// WithBasicAuth sets basic authentication credentials
func WithBasicAuth(username, password string) ClientOption {
	return func(c *Client) error {
		c.auth = auth.NewBasicAuth(username, password)
		return nil
	}
}

// WithTokenAuth sets token-based authentication
func WithTokenAuth(token string) ClientOption {
	return func(c *Client) error {
		c.auth = auth.NewTokenAuth(token)
		return nil
	}
}

// Request represents an API request
type Request struct {
	Method      string
	Endpoint    string
	Body        interface{}
	QueryParams url.Values
	Headers     map[string]string
}

// Response represents an API response
type Response struct {
	StatusCode int
	Body       []byte
	Headers    http.Header
}

// DoRequest performs an HTTP request to the Infinity API
func (c *Client) DoRequest(ctx context.Context, req *Request) (*Response, error) {
	// Build the full URL
	u := *c.baseURL
	u.Path = path.Join(u.Path, APIPrefix, strings.TrimPrefix(req.Endpoint, "/"))
	if req.QueryParams != nil {
		u.RawQuery = req.QueryParams.Encode()
	}

	// Prepare request body
	var body io.Reader
	if req.Body != nil {
		jsonBody, err := json.Marshal(req.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		body = bytes.NewBuffer(jsonBody)
	}

	// Create HTTP request
	httpReq, err := http.NewRequestWithContext(ctx, req.Method, u.String(), body)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Set default headers
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "application/json")

	// Set custom headers
	for key, value := range req.Headers {
		httpReq.Header.Set(key, value)
	}

	// Apply authentication
	if c.auth != nil {
		if err := c.auth.Authenticate(httpReq); err != nil {
			return nil, fmt.Errorf("failed to authenticate request: %w", err)
		}
	}

	// Perform the request
	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to perform HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	response := &Response{
		StatusCode: resp.StatusCode,
		Body:       respBody,
		Headers:    resp.Header,
	}

	// Check for API errors
	if resp.StatusCode >= 400 {
		return response, c.handleAPIError(response)
	}

	return response, nil
}

// APIError represents an error returned by the Infinity API
type APIError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Details    string `json:"details,omitempty"`
}

func (e *APIError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("API error %d: %s (%s)", e.StatusCode, e.Message, e.Details)
	}
	return fmt.Sprintf("API error %d: %s", e.StatusCode, e.Message)
}

// handleAPIError processes API error responses
func (c *Client) handleAPIError(resp *Response) error {
	apiErr := &APIError{
		StatusCode: resp.StatusCode,
		Message:    http.StatusText(resp.StatusCode),
	}

	// Try to parse error details from response body
	if len(resp.Body) > 0 {
		var errorResp map[string]interface{}
		if err := json.Unmarshal(resp.Body, &errorResp); err == nil {
			if msg, ok := errorResp["error"].(string); ok {
				apiErr.Message = msg
			}
			if details, ok := errorResp["details"].(string); ok {
				apiErr.Details = details
			}
		}
	}

	return apiErr
}

// GetJSON performs a GET request and unmarshals the JSON response
func (c *Client) GetJSON(ctx context.Context, endpoint string, result interface{}) error {
	req := &Request{
		Method:   http.MethodGet,
		Endpoint: endpoint,
	}

	resp, err := c.DoRequest(ctx, req)
	if err != nil {
		return err
	}

	if result != nil && len(resp.Body) > 0 {
		if err := json.Unmarshal(resp.Body, result); err != nil {
			return fmt.Errorf("failed to unmarshal JSON response: %w", err)
		}
	}

	return nil
}

// PostJSON performs a POST request with JSON body and unmarshals the JSON response
func (c *Client) PostJSON(ctx context.Context, endpoint string, body interface{}, result interface{}) error {
	req := &Request{
		Method:   http.MethodPost,
		Endpoint: endpoint,
		Body:     body,
	}

	resp, err := c.DoRequest(ctx, req)
	if err != nil {
		return err
	}

	if result != nil && len(resp.Body) > 0 {
		if err := json.Unmarshal(resp.Body, result); err != nil {
			return fmt.Errorf("failed to unmarshal JSON response: %w", err)
		}
	}

	return nil
}

// PutJSON performs a PUT request with JSON body and unmarshals the JSON response
func (c *Client) PutJSON(ctx context.Context, endpoint string, body interface{}, result interface{}) error {
	req := &Request{
		Method:   http.MethodPut,
		Endpoint: endpoint,
		Body:     body,
	}

	resp, err := c.DoRequest(ctx, req)
	if err != nil {
		return err
	}

	if result != nil && len(resp.Body) > 0 {
		if err := json.Unmarshal(resp.Body, result); err != nil {
			return fmt.Errorf("failed to unmarshal JSON response: %w", err)
		}
	}

	return nil
}

// DeleteJSON performs a DELETE request and unmarshals the JSON response
func (c *Client) DeleteJSON(ctx context.Context, endpoint string, result interface{}) error {
	req := &Request{
		Method:   http.MethodDelete,
		Endpoint: endpoint,
	}

	resp, err := c.DoRequest(ctx, req)
	if err != nil {
		return err
	}

	if result != nil && len(resp.Body) > 0 {
		if err := json.Unmarshal(resp.Body, result); err != nil {
			return fmt.Errorf("failed to unmarshal JSON response: %w", err)
		}
	}

	return nil
}