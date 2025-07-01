package config

import (
	"context"
	"fmt"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListDiagnosticGraphs retrieves a list of diagnostic graphs
func (s *Service) ListDiagnosticGraphs(ctx context.Context, opts *ListOptions) (*DiagnosticGraphListResponse, error) {
	endpoint := "configuration/v1/diagnostic_graphs/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result DiagnosticGraphListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetDiagnosticGraph retrieves a specific diagnostic graph by ID
func (s *Service) GetDiagnosticGraph(ctx context.Context, id int) (*DiagnosticGraph, error) {
	endpoint := fmt.Sprintf("configuration/v1/diagnostic_graphs/%d/", id)

	var result DiagnosticGraph
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateDiagnosticGraph creates a new diagnostic graph
func (s *Service) CreateDiagnosticGraph(ctx context.Context, req *DiagnosticGraphCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/diagnostic_graphs/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateDiagnosticGraph updates an existing diagnostic graph
func (s *Service) UpdateDiagnosticGraph(ctx context.Context, id int, req *DiagnosticGraphUpdateRequest) (*DiagnosticGraph, error) {
	endpoint := fmt.Sprintf("configuration/v1/diagnostic_graphs/%d/", id)

	var result DiagnosticGraph
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteDiagnosticGraph deletes a diagnostic graph
func (s *Service) DeleteDiagnosticGraph(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/diagnostic_graphs/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
