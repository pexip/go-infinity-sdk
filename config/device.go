/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"context"
	"fmt"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListDevices retrieves a list of devices
func (s *Service) ListDevices(ctx context.Context, opts *ListOptions) (*DeviceListResponse, error) {
	endpoint := "configuration/v1/device/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result DeviceListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetDevice retrieves a specific device by ID
func (s *Service) GetDevice(ctx context.Context, id int) (*Device, error) {
	endpoint := fmt.Sprintf("configuration/v1/device/%d/", id)

	var result Device
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateDevice creates a new device
func (s *Service) CreateDevice(ctx context.Context, req *DeviceCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/device/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateDevice updates an existing device
func (s *Service) UpdateDevice(ctx context.Context, id int, req *DeviceUpdateRequest) (*Device, error) {
	endpoint := fmt.Sprintf("configuration/v1/device/%d/", id)

	var result Device
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteDevice deletes a device
func (s *Service) DeleteDevice(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/device/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
