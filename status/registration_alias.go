package status

import (
	"context"
	"fmt"
)

// ListRegistrationAliases retrieves a list of registration alias statuses
func (s *Service) ListRegistrationAliases(ctx context.Context, opts *ListOptions) (*RegistrationAliasListResponse, error) {
	endpoint := "status/v1/registration_alias/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result RegistrationAliasListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetRegistrationAlias retrieves a specific registration alias status by ID
func (s *Service) GetRegistrationAlias(ctx context.Context, id int) (*RegistrationAlias, error) {
	endpoint := fmt.Sprintf("status/v1/registration_alias/%d/", id)

	var result RegistrationAlias
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}