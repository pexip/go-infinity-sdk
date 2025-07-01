package config

import (
	"context"
	"fmt"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListLicenceRequests retrieves a list of licence requests
func (s *Service) ListLicenceRequests(ctx context.Context, opts *ListOptions) (*LicenceRequestListResponse, error) {
	endpoint := "configuration/v1/licence_request/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result LicenceRequestListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetLicenceRequest retrieves a specific licence request by sequence number
func (s *Service) GetLicenceRequest(ctx context.Context, sequenceNumber string) (*LicenceRequest, error) {
	endpoint := fmt.Sprintf("configuration/v1/licence_request/%s/", sequenceNumber)

	var result LicenceRequest
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateLicenceRequest creates a new licence request
func (s *Service) CreateLicenceRequest(ctx context.Context, req *LicenceRequestCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/licence_request/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}
