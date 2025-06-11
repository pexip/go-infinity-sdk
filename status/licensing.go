package status

import (
	"context"
)

// GetLicensing retrieves the current licensing status
func (s *Service) GetLicensing(ctx context.Context) (*Licensing, error) {
	endpoint := "status/v1/licensing/"

	var result Licensing
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}