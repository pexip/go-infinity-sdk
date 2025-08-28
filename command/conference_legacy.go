/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package command

import (
	"context"
)

// LockConference locks a conference
func (s *Service) LockConference(ctx context.Context, conferenceID int) (*CommandResponse, error) {
	endpoint := "command/v1/conference/lock/"

	req := &ConferenceLockRequestLegacy{
		ConferenceID: conferenceID,
		Setting:      "lock",
	}

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// UnlockConference unlocks a conference
func (s *Service) UnlockConference(ctx context.Context, conferenceID int) (*CommandResponse, error) {
	endpoint := "command/v1/conference/lock/"

	req := &ConferenceLockRequestLegacy{
		ConferenceID: conferenceID,
		Setting:      "unlock",
	}

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// ToggleLockConference toggles the lock status of a conference
func (s *Service) ToggleLockConference(ctx context.Context, conferenceID int) (*CommandResponse, error) {
	endpoint := "command/v1/conference/lock/"

	req := &ConferenceLockRequestLegacy{
		ConferenceID: conferenceID,
		Setting:      "toggle",
	}

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// SendMessageToConference sends a message to all participants in a conference
func (s *Service) SendMessageToConference(ctx context.Context, conferenceID int, message string) (*CommandResponse, error) {
	endpoint := "command/v1/conference/message/"

	req := &ConferenceMessageRequest{
		ConferenceID: conferenceID,
		Message:      message,
	}

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// StartConference starts a conference
func (s *Service) StartConference(ctx context.Context, conferenceAlias string) (*CommandResponse, error) {
	endpoint := "command/v1/conference/start/"

	req := &ConferenceStartRequest{
		ConferenceAlias: conferenceAlias,
	}

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// StopConference stops a conference
func (s *Service) StopConference(ctx context.Context, conferenceID int) (*CommandResponse, error) {
	endpoint := "command/v1/conference/stop/"

	req := &ConferenceStopRequest{
		ConferenceID: conferenceID,
	}

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}
