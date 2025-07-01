// Package config provides access to the Pexip Infinity Configuration API.
// It allows management of conferences, locations, and other configuration resources
// with full CRUD operations and search capabilities.
package config

import (
	"github.com/pexip/go-infinity-sdk/v38/interfaces"
)

// Service handles Configuration API endpoints
type Service struct {
	client interfaces.HTTPClient
}

// New creates a new Configuration API service
func New(client interfaces.HTTPClient) *Service {
	return &Service{
		client: client,
	}
}
