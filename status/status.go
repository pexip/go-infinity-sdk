// Package status provides access to the Pexip Infinity status API.
// It allows monitoring of system status, active conferences, participants, worker nodes, and alarms
// with real-time status information and health monitoring capabilities.
package status

import (
	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/options"
)

// Service handles status API endpoints
type Service struct {
	client interfaces.HTTPClient
}

// New creates a new status API service
func New(client interfaces.HTTPClient) *Service {
	return &Service{
		client: client,
	}
}

// ListOptions contains options for listing resources
type ListOptions = options.BaseListOptions
