// Package history provides access to the Pexip Infinity History API.
// It allows retrieval of historical data for conferences, participants, and media streams
// with support for time-based filtering and search capabilities.
package history

import (
	"context"
)

// ClientInterface defines the interface for the HTTP client
type ClientInterface interface {
	GetJSON(ctx context.Context, endpoint string, result interface{}) error
}

// Service handles History API endpoints
type Service struct {
	client ClientInterface
}

// New creates a new History API service
func New(client ClientInterface) *Service {
	return &Service{
		client: client,
	}
}
