/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package command

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_TransformLayout(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	enableOverlay := true
	expectedRequest := &ConferenceTransformLayoutRequest{
		ConferenceID:      "test-conference-id",
		Layout:            "speaker_view",
		EnableOverlayText: &enableOverlay,
	}

	expectedResponse := &CommandResponse{
		Status:  "success",
		Message: "Layout transformed",
	}

	client.On("PostJSON", t.Context(), "command/v1/conference/transform_layout/", expectedRequest, mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CommandResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.TransformLayout(t.Context(), expectedRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_TransformLayoutSimple(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	expectedResponse := &CommandResponse{
		Status:  "success",
		Message: "Layout transformed",
	}

	client.On("PostJSON", t.Context(), "command/v1/conference/transform_layout/", mock.AnythingOfType("*command.ConferenceTransformLayoutRequest"), mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CommandResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.TransformLayoutSimple(t.Context(), "test-conference-id", "gallery_view")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_TransformLayoutWithOptions(t *testing.T) {
	tests := []struct {
		name         string
		conferenceID string
		opts         *TransformLayoutOptions
		wantErr      bool
	}{
		{
			name:         "transform without options",
			conferenceID: "test-conference-id",
			opts:         nil,
			wantErr:      false,
		},
		{
			name:         "transform with basic options",
			conferenceID: "test-conference-id",
			opts: &TransformLayoutOptions{
				Layout:     "speaker_view",
				HostLayout: "gallery_view",
			},
			wantErr: false,
		},
		{
			name:         "transform with indicators",
			conferenceID: "test-conference-id",
			opts: &TransformLayoutOptions{
				Layout:              "gallery_view",
				EnableOverlayText:   boolPtr(true),
				RecordingIndicator:  boolPtr(true),
				StreamingIndicator:  boolPtr(false),
				AIEnabledIndicator:  boolPtr(true),
				FreeFormOverlayText: "Custom Conference Name",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := interfaces.NewHTTPClientMock()

			expectedResponse := &CommandResponse{
				Status:  "success",
				Message: "Layout transformed",
			}

			client.On("PostJSON", t.Context(), "command/v1/conference/transform_layout/", mock.AnythingOfType("*command.ConferenceTransformLayoutRequest"), mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
				result := args.Get(3).(*CommandResponse)
				*result = *expectedResponse
			})

			service := New(client)
			result, err := service.TransformLayoutWithOptions(t.Context(), tt.conferenceID, tt.opts)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, expectedResponse, result)
			}

			client.AssertExpectations(t)
		})
	}
}

// boolPtr is a helper function to get a pointer to a bool value
func boolPtr(b bool) *bool {
	return &b
}
