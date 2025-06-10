package mock

import (
	"context"

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
