package command

import (
	"context"
)

// LockConferenceByID locks a conference using string conference ID (schema-compliant)
func (s *Service) LockConferenceByID(ctx context.Context, conferenceID string) (*CommandResponse, error) {
	endpoint := "command/v1/conference/lock/"

	req := &ConferenceLockRequest{
		ConferenceID: conferenceID,
	}

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// UnlockConferenceByID unlocks a conference using string conference ID (schema-compliant)
func (s *Service) UnlockConferenceByID(ctx context.Context, conferenceID string) (*CommandResponse, error) {
	endpoint := "command/v1/conference/unlock/"

	req := &ConferenceUnlockRequest{
		ConferenceID: conferenceID,
	}

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}
