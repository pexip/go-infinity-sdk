// Package history provides access to the Pexip Infinity history API.
// It allows retrieval of historical data for conferences, participants, and media streams
// with support for time-based filtering and search capabilities.
package history

import (
	"github.com/pexip/go-infinity-sdk/v38/interfaces"
)

// Service handles history API endpoints
type Service struct {
	client interfaces.HTTPClient
}

// New creates a new history API service
func New(client interfaces.HTTPClient) *Service {
	return &Service{
		client: client,
	}
}
