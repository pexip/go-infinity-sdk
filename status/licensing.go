package status

import (
	"context"
)

// GetLicensing retrieves the current licensing status
func (s *Service) GetLicensing(ctx context.Context) (*Licensing, error) {
	endpoint := "status/v1/licensing/"

	var result LicensingResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	if len(result.Objects) == 0 {
		return nil, err
	}
	licensing := result.Objects[0]
	return &licensing, err
}
