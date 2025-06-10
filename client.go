// Package infinity provides a comprehensive Go client library for the Pexip Infinity Management API.
// It offers complete support for all four API categories: Configuration, Status, History, and Command APIs
// with features including type-safe operations, automatic retry with exponential backoff, and flexible authentication.
package infinity

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pexip/go-infinity-sdk/auth"
	"github.com/pexip/go-infinity-sdk/command"
	"github.com/pexip/go-infinity-sdk/config"
	"github.com/pexip/go-infinity-sdk/history"
	"github.com/pexip/go-infinity-sdk/status"
)

const (
	DefaultBaseURL           = "https://admin.example.com"
	DefaultTimeout           = 30 * time.Second
	APIPrefix                = "/api/admin/"
	DefaultMaxRetries        = 3
	DefaultBackoffMin        = 1 * time.Second
	DefaultBackoffMax        = 30 * time.Second
	DefaultBackoffMultiplier = 2.0
	DefaultJitterFactor      = 0.1
)

// Client represents a Pexip Infinity Management API client
type Client struct {
	baseURL     *url.URL
	httpClient  *http.Client
	auth        auth.Authenticator
	retryConfig *RetryConfig

	// API services
	Config  *config.Service
	Status  *status.Service
	History *history.Service
	Command *command.Service
}

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
		retryConfig: DefaultRetryConfig(),
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

// DoRequest performs an HTTP request to the Infinity API with retry logic
func (c *Client) DoRequest(ctx context.Context, req *Request) (*Response, error) {
	fullURL := c.baseURL.JoinPath(APIPrefix, strings.TrimPrefix(req.Endpoint, "/"))
	if req.QueryParams != nil {
		fullURL.RawQuery = req.QueryParams.Encode()
	}

	// Pre-marshal request body once for all retries
	var jsonBody []byte
	var err error
	if req.Body != nil {
		jsonBody, err = json.Marshal(req.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
	}

	var lastResponse *Response
	var lastError error

	// Perform the request with retry logic
	for attempt := 0; attempt <= c.retryConfig.MaxRetries; attempt++ {
		// Check context cancellation before each attempt
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		var bodyReader io.Reader
		if jsonBody != nil {
			bodyReader = bytes.NewReader(jsonBody)
		}

		httpReq, err := http.NewRequestWithContext(ctx, req.Method, fullURL.String(), bodyReader)
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

		// Perform the HTTP request
		resp, err := c.httpClient.Do(httpReq)
		if err != nil {
			lastError = err

			// Check if this error should be retried
			if attempt < c.retryConfig.MaxRetries && c.retryConfig.IsRetriable(0, err) {
				if !c.sleepWithContext(ctx, c.retryConfig.CalculateBackoff(attempt+1)) {
					return nil, ctx.Err()
				}
				continue
			}
			// No more retries or error is not retriable
			return nil, fmt.Errorf("failed to perform HTTP request after %d attempts: %w", attempt+1, err)
		}

		// Read response body
		respBody, err := io.ReadAll(resp.Body)
		resp.Body.Close() // Always close the body
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %w", err)
		}

		response := &Response{
			StatusCode: resp.StatusCode,
			Body:       respBody,
			Headers:    resp.Header,
		}

		// Check if we should retry based on status code
		if resp.StatusCode >= 400 {
			lastResponse = response

			// Check if this status code should be retried
			if attempt < c.retryConfig.MaxRetries && c.retryConfig.IsRetriable(resp.StatusCode, nil) {
				if !c.sleepWithContext(ctx, c.retryConfig.CalculateBackoff(attempt+1)) {
					return nil, ctx.Err()
				}
				continue
			}

			// No more retries or status code is not retriable - return error
			return response, c.handleAPIError(response)
		}

		// Success - return the response
		return response, nil
	}

	// This should not be reached, but handle it just in case
	if lastResponse != nil {
		return lastResponse, c.handleAPIError(lastResponse)
	}
	return nil, fmt.Errorf("request failed after %d attempts: %w", c.retryConfig.MaxRetries+1, lastError)
}

// handleAPIError processes API error responses
func (c *Client) handleAPIError(resp *Response) error {
	apiErr := &APIError{
		StatusCode: resp.StatusCode,
		Message:    http.StatusText(resp.StatusCode),
	}

	// Try to parse error details from response body using the APIError's UnmarshalJSON method
	if len(resp.Body) > 0 {
		_ = json.Unmarshal(resp.Body, apiErr)
	}
	return apiErr
}

// GetJSON performs a GET request and unmarshal the JSON response
func (c *Client) GetJSON(ctx context.Context, endpoint string, result interface{}) error {
	return c.performJSONRequest(ctx, http.MethodGet, endpoint, nil, result)
}

// PostJSON performs a POST request with JSON body and unmarshal the JSON response
func (c *Client) PostJSON(ctx context.Context, endpoint string, body interface{}, result interface{}) error {
	return c.performJSONRequest(ctx, http.MethodPost, endpoint, body, result)
}

// PutJSON performs a PUT request with JSON body and unmarshal the JSON response
func (c *Client) PutJSON(ctx context.Context, endpoint string, body interface{}, result interface{}) error {
	return c.performJSONRequest(ctx, http.MethodPut, endpoint, body, result)
}

// DeleteJSON performs a DELETE request and unmarshal the JSON response
func (c *Client) DeleteJSON(ctx context.Context, endpoint string, result interface{}) error {
	return c.performJSONRequest(ctx, http.MethodDelete, endpoint, nil, result)
}

func (c *Client) performJSONRequest(ctx context.Context, method string, endpoint string, requestBody interface{}, result interface{}) error {
	req := &Request{
		Method:   method,
		Endpoint: endpoint,
		Body:     requestBody,
	}

	resp, err := c.DoRequest(ctx, req)
	if err != nil {
		return err
	}

	return unmarshalResponseBody(resp.Body, result)
}

func unmarshalResponseBody(body []byte, result interface{}) error {
	if result != nil && len(body) > 0 {
		if err := json.Unmarshal(body, result); err != nil {
			return fmt.Errorf("failed to unmarshal JSON response: %w", err)
		}
	}
	return nil
}

func (c *Client) sleepWithContext(ctx context.Context, duration time.Duration) bool {
	if duration <= 0 {
		return true
	}
	timer := time.NewTimer(duration)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return false
	case <-timer.C:
		return true
	}
}
