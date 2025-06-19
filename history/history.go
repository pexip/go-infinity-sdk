// Package history provides access to the Pexip Infinity History API.
// It allows retrieval of historical data for conferences, participants, and media streams
// with support for time-based filtering and search capabilities.
package history

import (
	"context"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ClientInterface defines the interface for the HTTP client
type ClientInterface interface {
	GetJSON(ctx context.Context, endpoint string, result interface{}) error
	PostWithResponse(ctx context.Context, endpoint string, body interface{}, result interface{}) (*types.PostResponse, error)
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
