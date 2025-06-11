package command

import (
	"context"
)

// MuteGuests mutes all guest participants in a conference
func (s *Service) MuteGuests(ctx context.Context, conferenceID string) (*CommandResponse, error) {
	endpoint := "command/v1/conference/mute_guests/"

	req := &ConferenceMuteGuestsRequest{
		ConferenceID: conferenceID,
	}

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// UnmuteGuests unmutes all guest participants in a conference
func (s *Service) UnmuteGuests(ctx context.Context, conferenceID string) (*CommandResponse, error) {
	endpoint := "command/v1/conference/unmute_guests/"

	req := &ConferenceUnmuteGuestsRequest{
		ConferenceID: conferenceID,
	}

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}
