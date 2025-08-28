/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package command

import (
	"context"
)

// TransformLayout transforms the layout of a conference with various options
func (s *Service) TransformLayout(ctx context.Context, req *ConferenceTransformLayoutRequest) (*CommandResponse, error) {
	endpoint := "command/v1/conference/transform_layout/"

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// TransformLayoutSimple transforms the layout of a conference with basic parameters
func (s *Service) TransformLayoutSimple(ctx context.Context, conferenceID, layout string) (*CommandResponse, error) {
	req := &ConferenceTransformLayoutRequest{
		ConferenceID: conferenceID,
		Layout:       layout,
	}

	return s.TransformLayout(ctx, req)
}

// TransformLayoutWithOptions transforms the layout of a conference with options
func (s *Service) TransformLayoutWithOptions(ctx context.Context, conferenceID string, opts *TransformLayoutOptions) (*CommandResponse, error) {
	req := &ConferenceTransformLayoutRequest{
		ConferenceID: conferenceID,
	}

	if opts != nil {
		req.Layout = opts.Layout
		req.HostLayout = opts.HostLayout
		req.GuestLayout = opts.GuestLayout
		req.EnableOverlayText = opts.EnableOverlayText
		req.FreeFormOverlayText = opts.FreeFormOverlayText
		req.RecordingIndicator = opts.RecordingIndicator
		req.StreamingIndicator = opts.StreamingIndicator
		req.LiveCaptionsIndicator = opts.LiveCaptionsIndicator
		req.TranscribingIndicator = opts.TranscribingIndicator
		req.AIEnabledIndicator = opts.AIEnabledIndicator
		req.PlusNPipEnabled = opts.PlusNPipEnabled
	}

	return s.TransformLayout(ctx, req)
}

// TransformLayoutOptions contains options for transforming conference layout
type TransformLayoutOptions struct {
	Layout                string
	HostLayout            string
	GuestLayout           string
	EnableOverlayText     *bool
	FreeFormOverlayText   string
	RecordingIndicator    *bool
	StreamingIndicator    *bool
	LiveCaptionsIndicator *bool
	TranscribingIndicator *bool
	AIEnabledIndicator    *bool
	PlusNPipEnabled       *bool
}
