package config

import (
	"context"
	"fmt"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListHTTPProxies retrieves a list of HTTP proxies
func (s *Service) ListHTTPProxies(ctx context.Context, opts *ListOptions) (*HTTPProxyListResponse, error) {
	endpoint := "configuration/v1/http_proxy/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result HTTPProxyListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetHTTPProxy retrieves a specific HTTP proxy by ID
func (s *Service) GetHTTPProxy(ctx context.Context, id int) (*HTTPProxy, error) {
	endpoint := fmt.Sprintf("configuration/v1/http_proxy/%d/", id)

	var result HTTPProxy
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateHTTPProxy creates a new HTTP proxy
func (s *Service) CreateHTTPProxy(ctx context.Context, req *HTTPProxyCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/http_proxy/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateHTTPProxy updates an existing HTTP proxy
func (s *Service) UpdateHTTPProxy(ctx context.Context, id int, req *HTTPProxyUpdateRequest) (*HTTPProxy, error) {
	endpoint := fmt.Sprintf("configuration/v1/http_proxy/%d/", id)

	var result HTTPProxy
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteHTTPProxy deletes an HTTP proxy
func (s *Service) DeleteHTTPProxy(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/http_proxy/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
