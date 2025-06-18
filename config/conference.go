package config

import (
	"context"
	"fmt"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListConferences retrieves a list of conferences
func (s *Service) ListConferences(ctx context.Context, opts *ListOptions) (*ConferenceListResponse, error) {
	endpoint := "configuration/v1/conference/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result ConferenceListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetConference retrieves a specific conference by ID
func (s *Service) GetConference(ctx context.Context, id int) (*Conference, error) {
	endpoint := fmt.Sprintf("configuration/v1/conference/%d/", id)

	var result Conference
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateConference creates a new conference
func (s *Service) CreateConference(ctx context.Context, req *ConferenceCreateRequest) (*Conference, error) {
	endpoint := "configuration/v1/conference/"

	var result Conference
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// CreateConferenceWithResponse creates a new conference and returns both the response body and location header
func (s *Service) CreateConferenceWithResponse(ctx context.Context, req *ConferenceCreateRequest) (*Conference, *types.PostResponse, error) {
	endpoint := "configuration/v1/conference/"

	var result Conference
	postResp, err := s.client.PostWithResponse(ctx, endpoint, req, &result)
	if err != nil {
		return nil, postResp, err
	}
	return &result, postResp, nil
}

// UpdateConference updates an existing conference
func (s *Service) UpdateConference(ctx context.Context, id int, req *ConferenceUpdateRequest) (*Conference, error) {
	endpoint := fmt.Sprintf("configuration/v1/conference/%d/", id)

	var result Conference
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteConference deletes a conference
func (s *Service) DeleteConference(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/conference/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
