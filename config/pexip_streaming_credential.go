package config

import (
	"context"
	"fmt"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListPexipStreamingCredentials retrieves a list of Pexip Streaming credentials
func (s *Service) ListPexipStreamingCredentials(ctx context.Context, opts *ListOptions) (*PexipStreamingCredentialListResponse, error) {
	endpoint := "configuration/v1/pexip_streaming_credential/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result PexipStreamingCredentialListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetPexipStreamingCredential retrieves a specific Pexip Streaming credential by ID
func (s *Service) GetPexipStreamingCredential(ctx context.Context, id int) (*PexipStreamingCredential, error) {
	endpoint := fmt.Sprintf("configuration/v1/pexip_streaming_credential/%d/", id)

	var result PexipStreamingCredential
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreatePexipStreamingCredential creates a new Pexip Streaming credential
func (s *Service) CreatePexipStreamingCredential(ctx context.Context, req *PexipStreamingCredentialCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/pexip_streaming_credential/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdatePexipStreamingCredential updates an existing Pexip Streaming credential
func (s *Service) UpdatePexipStreamingCredential(ctx context.Context, id int, req *PexipStreamingCredentialUpdateRequest) (*PexipStreamingCredential, error) {
	endpoint := fmt.Sprintf("configuration/v1/pexip_streaming_credential/%d/", id)

	var result PexipStreamingCredential
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeletePexipStreamingCredential deletes a Pexip Streaming credential
func (s *Service) DeletePexipStreamingCredential(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/pexip_streaming_credential/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
