package command

import (
	"context"
)

// UnlockParticipant unlocks a participant (removes their waiting state)
func (s *Service) UnlockParticipant(ctx context.Context, participantID string) (*CommandResponse, error) {
	endpoint := "command/v1/participant/unlock/"

	req := &ParticipantUnlockRequest{
		ParticipantID: participantID,
	}

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}
