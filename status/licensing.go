/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package status

import (
	"context"
	"errors"
)

var ErrNoLicensingData = errors.New("no licensing data returned")

// GetLicensing retrieves the current licensing status, taken from the first element of a LicensingResponse.
// It returns a Licensing object or an error if the request fails or no data is returned.
func (s *Service) GetLicensing(ctx context.Context) (*Licensing, error) {
	endpoint := "status/v1/licensing/"

	var result LicensingResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	if err != nil {
		return nil, err
	}
	if len(result.Objects) == 0 {
		return nil, ErrNoLicensingData
	}
	// Assuming we always expect only one object in the response
	return &result.Objects[0], nil
}
