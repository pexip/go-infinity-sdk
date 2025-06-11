// Package command provides access to the Pexip Infinity Command API.
// It allows real-time control of conferences and participants including operations like
// muting, spotlighting, transferring, and sending messages.
package command

import (
	"context"
)

// ClientInterface defines the interface for the HTTP client
type ClientInterface interface {
	PostJSON(ctx context.Context, endpoint string, body interface{}, result interface{}) error
}

// Service handles Command API endpoints
type Service struct {
	client ClientInterface
}

// New creates a new Command API service
func New(client ClientInterface) *Service {
	return &Service{
		client: client,
	}
}
