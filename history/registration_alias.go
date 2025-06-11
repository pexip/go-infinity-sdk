package history

import (
	"context"
	"fmt"
)

// ListRegistrationAliases retrieves a list of registration alias history records
func (s *Service) ListRegistrationAliases(ctx context.Context, opts *ListOptions) (*RegistrationAliasListResponse, error) {
	endpoint := "history/v1/registration_alias/"

	if opts != nil {
		params := opts.ToURLValuesWithSearchField("alias__icontains")
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result RegistrationAliasListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetRegistrationAlias retrieves a specific registration alias history record by ID
func (s *Service) GetRegistrationAlias(ctx context.Context, id int) (*RegistrationAlias, error) {
	endpoint := fmt.Sprintf("history/v1/registration_alias/%d/", id)

	var result RegistrationAlias
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}
