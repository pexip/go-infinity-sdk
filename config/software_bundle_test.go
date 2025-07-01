package config

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListSoftwareBundles(t *testing.T) {
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
				revision1 := "v27.4.1"
				revision2 := "v28.0.0"
				expectedResponse := &SoftwareBundleListResponse{
					Objects: []SoftwareBundle{
						{ID: 1, BundleType: "core", SelectedRevision: &revision1},
						{ID: 2, BundleType: "conferencing", SelectedRevision: &revision2},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/software_bundle/", mock.AnythingOfType("*config.SoftwareBundleListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*SoftwareBundleListResponse)
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
				Search: "core",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				revision := "v27.4.1"
				expectedResponse := &SoftwareBundleListResponse{
					Objects: []SoftwareBundle{
						{ID: 1, BundleType: "core", SelectedRevision: &revision},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/software_bundle/?limit=5&name__icontains=core", mock.AnythingOfType("*config.SoftwareBundleListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*SoftwareBundleListResponse)
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
			result, err := service.ListSoftwareBundles(t.Context(), tt.opts)

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

func TestService_GetSoftwareBundle(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	revision := "v27.4.1"
	expectedSoftwareBundle := &SoftwareBundle{
		ID:               1,
		BundleType:       "core",
		SelectedRevision: &revision,
		ResourceURI:      "/api/admin/configuration/v1/software_bundle/1/",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/software_bundle/1/", mock.AnythingOfType("*config.SoftwareBundle")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*SoftwareBundle)
		*result = *expectedSoftwareBundle
	})

	service := New(client)
	result, err := service.GetSoftwareBundle(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedSoftwareBundle, result)
	client.AssertExpectations(t)
}

func TestService_UpdateSoftwareBundle(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	newRevision := "v28.0.0"
	updateRequest := &SoftwareBundleUpdateRequest{
		SelectedRevision: &newRevision,
	}

	expectedSoftwareBundle := &SoftwareBundle{
		ID:               1,
		BundleType:       "core",
		SelectedRevision: &newRevision,
		ResourceURI:      "/api/admin/configuration/v1/software_bundle/1/",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/software_bundle/1/", updateRequest, mock.AnythingOfType("*config.SoftwareBundle")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*SoftwareBundle)
		*result = *expectedSoftwareBundle
	})

	service := New(client)
	result, err := service.UpdateSoftwareBundle(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedSoftwareBundle, result)
	client.AssertExpectations(t)
}
