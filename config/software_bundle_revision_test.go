/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListSoftwareBundleRevisions(t *testing.T) {
	tests := []struct {
		name    string
		opts    *ListOptions
		setup   func(m *interfaces.HTTPClientMock)
		wantErr bool
	}{
		{
			name: "successful list without options",
			opts: nil,
			setup: func(m *interfaces.HTTPClientMock) {
				expectedResponse := &SoftwareBundleRevisionListResponse{
					Objects: []SoftwareBundleRevision{
						{ID: 1, BundleType: "core", Core: true, Revision: "v27.4.1", Version: "27.4.1"},
						{ID: 2, BundleType: "conferencing", Core: false, Revision: "v28.0.0", Version: "28.0.0"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/software_bundle_revision/", mock.AnythingOfType("*config.SoftwareBundleRevisionListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*SoftwareBundleRevisionListResponse)
					*result = *expectedResponse
				})
			},
			wantErr: false,
		},
		{
			name: "successful list with options",
			opts: &ListOptions{
				BaseListOptions: options.BaseListOptions{
					Limit: 5,
				},
				Search: "v27",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				expectedResponse := &SoftwareBundleRevisionListResponse{
					Objects: []SoftwareBundleRevision{
						{ID: 1, BundleType: "core", Core: true, Revision: "v27.4.1", Version: "27.4.1"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/software_bundle_revision/?limit=5&name__icontains=v27", mock.AnythingOfType("*config.SoftwareBundleRevisionListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*SoftwareBundleRevisionListResponse)
					*result = *expectedResponse
				})
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := interfaces.NewHTTPClientMock()
			tt.setup(client)

			service := New(client)
			result, err := service.ListSoftwareBundleRevisions(t.Context(), tt.opts)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
			}

			client.AssertExpectations(t)
		})
	}
}

func TestService_GetSoftwareBundleRevision(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedSoftwareBundleRevision := &SoftwareBundleRevision{
		ID:          1,
		BundleType:  "core",
		Core:        true,
		Revision:    "v27.4.1",
		Version:     "27.4.1",
		ResourceURI: "/api/admin/configuration/v1/software_bundle_revision/1/",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/software_bundle_revision/1/", mock.AnythingOfType("*config.SoftwareBundleRevision")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*SoftwareBundleRevision)
		*result = *expectedSoftwareBundleRevision
	})

	service := New(client)
	result, err := service.GetSoftwareBundleRevision(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedSoftwareBundleRevision, result)
	client.AssertExpectations(t)
}
