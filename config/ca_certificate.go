package config

import (
	"context"
	"fmt"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListCACertificates retrieves a list of CA certificates
func (s *Service) ListCACertificates(ctx context.Context, opts *ListOptions) (*CACertificateListResponse, error) {
	endpoint := "configuration/v1/ca_certificate/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result CACertificateListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetCACertificate retrieves a specific CA certificate by ID
func (s *Service) GetCACertificate(ctx context.Context, id int) (*CACertificate, error) {
	endpoint := fmt.Sprintf("configuration/v1/ca_certificate/%d/", id)

	var result CACertificate
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateCACertificate creates a new CA certificate
func (s *Service) CreateCACertificate(ctx context.Context, req *CACertificateCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/ca_certificate/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateCACertificate updates an existing CA certificate (partial update)
func (s *Service) UpdateCACertificate(ctx context.Context, id int, req *CACertificateUpdateRequest) (*CACertificate, error) {
	endpoint := fmt.Sprintf("configuration/v1/ca_certificate/%d/", id)

	var result CACertificate
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteCACertificate deletes a CA certificate
func (s *Service) DeleteCACertificate(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/ca_certificate/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
