/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package command

import (
	"context"
)

// DialParticipant dials out to a participant to join a conference
func (s *Service) DialParticipant(ctx context.Context, req *ParticipantDialRequest) (*CommandResponse, error) {
	endpoint := "command/v1/participant/dial/"

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DialParticipantWithOptions dials out to a participant with simplified parameters
func (s *Service) DialParticipantWithOptions(ctx context.Context, conferenceAlias, destination string, opts *DialOptions) (*CommandResponse, error) {
	req := &ParticipantDialRequest{
		ConferenceAlias: conferenceAlias,
		Destination:     destination,
		Streaming:       false, // Required field, defaults to false
	}

	if opts != nil {
		req.CallType = opts.CallType
		req.Protocol = opts.Protocol
		req.Role = opts.Role
		req.LocalDisplayName = opts.LocalDisplayName
		req.RemoteDisplayName = opts.RemoteDisplayName
		req.DTMFSequence = opts.DTMFSequence
		req.SystemLocation = opts.SystemLocation
		req.Node = opts.Node
		req.Streaming = opts.Streaming
		req.CustomSIPHeaders = opts.CustomSIPHeaders
		req.PresentationURL = opts.PresentationURL
		req.KeepConferenceAlive = opts.KeepConferenceAlive
		req.Routing = opts.Routing
	}

	return s.DialParticipant(ctx, req)
}

// DialOptions contains optional parameters for dialing participants
type DialOptions struct {
	CallType            string // "audio", "video", "video-only"
	Protocol            string // "gms", "h323", "sip", "mssip", "rtmp", "teams"
	Role                string // "guest", "chair"
	LocalDisplayName    string
	RemoteDisplayName   string
	DTMFSequence        string
	SystemLocation      string
	Node                string
	Streaming           bool
	CustomSIPHeaders    string
	PresentationURL     string
	KeepConferenceAlive string // "keep_conference_alive", "keep_conference_alive_if_multiple", "keep_conference_alive_never"
	Routing             string // "manual", "routing_rule"
}
