package config

import (
	"context"
	"fmt"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListCertificateSigningRequests retrieves a list of certificate signing requests
func (s *Service) ListCertificateSigningRequests(ctx context.Context, opts *ListOptions) (*CertificateSigningRequestListResponse, error) {
	endpoint := "configuration/v1/certificate_signing_request/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result CertificateSigningRequestListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetCertificateSigningRequest retrieves a specific certificate signing request by ID
func (s *Service) GetCertificateSigningRequest(ctx context.Context, id int) (*CertificateSigningRequest, error) {
	endpoint := fmt.Sprintf("configuration/v1/certificate_signing_request/%d/", id)

	var result CertificateSigningRequest
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateCertificateSigningRequest creates a new certificate signing request
func (s *Service) CreateCertificateSigningRequest(ctx context.Context, req *CertificateSigningRequestCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/certificate_signing_request/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateCertificateSigningRequest updates an existing certificate signing request
func (s *Service) UpdateCertificateSigningRequest(ctx context.Context, id int, req *CertificateSigningRequestUpdateRequest) (*CertificateSigningRequest, error) {
	endpoint := fmt.Sprintf("configuration/v1/certificate_signing_request/%d/", id)

	var result CertificateSigningRequest
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteCertificateSigningRequest deletes a certificate signing request
func (s *Service) DeleteCertificateSigningRequest(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/certificate_signing_request/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
