package status

import (
	"context"
	"fmt"
)

// GetLicensing retrieves the current licensing status
func (s *Service) GetLicensing(ctx context.Context) (*Licensing, error) {
	endpoint := "status/v1/licensing/"

	var result LicensingResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	if err != nil {
		return nil, err
	}
	if len(result.Objects) == 0 {
		return nil, fmt.Errorf("no licensing data returned")
	}
	// Assuming we always expect only one object in the response
	return &result.Objects[0], nil
}
