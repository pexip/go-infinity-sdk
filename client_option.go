/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package infinity

import (
	"fmt"
	"github.com/pexip/go-infinity-sdk/v38/auth"
	"net/http"
	"net/url"
)

// ClientOption is a function that configures a Client
type ClientOption func(*Client) error

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

// WithTransport sets a custom HTTP transport for the client.
func WithTransport(transport http.RoundTripper) ClientOption {
	return func(c *Client) error {
		if transport == nil {
			return fmt.Errorf("transport cannot be nil")
		}
		// If no custom HTTP client was set, create one with the transport
		if c.httpClient.Transport == http.DefaultTransport {
			c.httpClient = &http.Client{
				Timeout:   c.httpClient.Timeout,
				Transport: transport,
			}
		} else {
			// Modify existing client's transport
			c.httpClient.Transport = transport
		}
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

// WithRetryConfig sets the retry configuration for the client
func WithRetryConfig(config *RetryConfig) ClientOption {
	return func(c *Client) error {
		if config == nil {
			return fmt.Errorf("retry config cannot be nil")
		}
		c.retryConfig = config
		return nil
	}
}

// WithMaxRetries sets the maximum number of retries (convenience function)
func WithMaxRetries(maxRetries int) ClientOption {
	return func(c *Client) error {
		if maxRetries < 0 {
			return fmt.Errorf("max retries cannot be negative")
		}
		c.retryConfig.MaxRetries = maxRetries
		return nil
	}
}

// WithNoRetries disables retries completely
func WithNoRetries() ClientOption {
	return WithMaxRetries(0)
}

// WithUserAgent sets the User-Agent header for HTTP requests
func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) error {
		if userAgent == "" {
			return fmt.Errorf("user agent cannot be empty")
		}
		c.userAgent = userAgent
		return nil
	}
}
