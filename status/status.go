// Package status provides access to the Pexip Infinity Status API.
// It allows monitoring of system status, active conferences, participants, worker nodes, and alarms
// with real-time status information and health monitoring capabilities.
package status

import (
	"context"

	"github.com/pexip/go-infinity-sdk/options"
)

// ClientInterface defines the interface for the HTTP client
type ClientInterface interface {
	GetJSON(ctx context.Context, endpoint string, result interface{}) error
}

// Service handles Status API endpoints
type Service struct {
	client ClientInterface
}

// New creates a new Status API service
func New(client ClientInterface) *Service {
	return &Service{
		client: client,
	}
}

// ListOptions contains options for listing resources
type ListOptions = options.BaseListOptions

