// Package config provides access to the Pexip Infinity Configuration API.
// It allows management of conferences, locations, and other configuration resources
// with full CRUD operations and search capabilities.
package config

import (
	"context"
)

// ClientInterface defines the interface for the HTTP client
type ClientInterface interface {
	GetJSON(ctx context.Context, endpoint string, result interface{}) error
	PostJSON(ctx context.Context, endpoint string, body interface{}, result interface{}) error
	PutJSON(ctx context.Context, endpoint string, body interface{}, result interface{}) error
	DeleteJSON(ctx context.Context, endpoint string, result interface{}) error
}

// Service handles Configuration API endpoints
type Service struct {
	client ClientInterface
}

// New creates a new Configuration API service
func New(client ClientInterface) *Service {
	return &Service{
		client: client,
	}
}