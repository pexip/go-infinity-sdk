// Package mock provides mock implementations for testing the Pexip Infinity SDK.
// It includes mock clients and interfaces for unit testing without making actual API calls.
package mock

import (
	"context"

	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/stretchr/testify/mock"
)

// Client is a mock implementation of the ClientInterface
type Client struct {
	mock.Mock
}

// GetJSON mocks the GetJSON method
func (m *Client) GetJSON(ctx context.Context, endpoint string, result interface{}) error {
	args := m.Called(ctx, endpoint, result)
	return args.Error(0)
}

// PostJSON mocks the PostJSON method
func (m *Client) PostJSON(ctx context.Context, endpoint string, body interface{}, result interface{}) error {
	args := m.Called(ctx, endpoint, body, result)
	return args.Error(0)
}

// PutJSON mocks the PutJSON method
func (m *Client) PutJSON(ctx context.Context, endpoint string, body interface{}, result interface{}) error {
	args := m.Called(ctx, endpoint, body, result)
	return args.Error(0)
}

// DeleteJSON mocks the DeleteJSON method
func (m *Client) DeleteJSON(ctx context.Context, endpoint string, result interface{}) error {
	args := m.Called(ctx, endpoint, result)
	return args.Error(0)
}

// PostWithResponse mocks the PostWithResponse method
func (m *Client) PostWithResponse(ctx context.Context, endpoint string, body interface{}, result interface{}) (*types.PostResponse, error) {
	args := m.Called(ctx, endpoint, body, result)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*types.PostResponse), args.Error(1)
}
