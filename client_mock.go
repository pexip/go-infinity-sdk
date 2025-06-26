package infinity

import (
	"context"
	"github.com/pexip/go-infinity-sdk/v38/command"
	"github.com/pexip/go-infinity-sdk/v38/config"
	"github.com/pexip/go-infinity-sdk/v38/history"
	"github.com/pexip/go-infinity-sdk/v38/status"

	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/stretchr/testify/mock"
)

// ClientMock is a mock implementation of the ClientInterface
type ClientMock struct {
	mock.Mock

	// API services
	config  *config.Service
	status  *status.Service
	history *history.Service
	command *command.Service
}

func NewClientMock() *ClientMock {
	client := &ClientMock{}

	// Initialize API services
	client.config = config.New(client)
	client.status = status.New(client)
	client.history = history.New(client)
	client.command = command.New(client)

	return client
}

func (client *ClientMock) Config() *config.Service {
	return client.config
}

func (client *ClientMock) Status() *status.Service {
	return client.status
}

func (client *ClientMock) History() *history.Service {
	return client.history
}

func (client *ClientMock) Command() *command.Service {
	return client.command
}

// GetJSON mocks the GetJSON method
func (m *ClientMock) GetJSON(ctx context.Context, endpoint string, result interface{}) error {
	args := m.Called(ctx, endpoint, result)
	return args.Error(0)
}

// PostJSON mocks the PostJSON method
func (m *ClientMock) PostJSON(ctx context.Context, endpoint string, body interface{}, result interface{}) error {
	args := m.Called(ctx, endpoint, body, result)
	return args.Error(0)
}

// PutJSON mocks the PutJSON method
func (m *ClientMock) PutJSON(ctx context.Context, endpoint string, body interface{}, result interface{}) error {
	args := m.Called(ctx, endpoint, body, result)
	return args.Error(0)
}

// DeleteJSON mocks the DeleteJSON method
func (m *ClientMock) DeleteJSON(ctx context.Context, endpoint string, result interface{}) error {
	args := m.Called(ctx, endpoint, result)
	return args.Error(0)
}

// PostWithResponse mocks the PostWithResponse method
func (m *ClientMock) PostWithResponse(ctx context.Context, endpoint string, body interface{}, result interface{}) (*types.PostResponse, error) {
	args := m.Called(ctx, endpoint, body, result)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*types.PostResponse), args.Error(1)
}
