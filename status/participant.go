package status

import (
	"context"
	"fmt"
)

// ListParticipants retrieves a list of participants
func (s *Service) ListParticipants(ctx context.Context, opts *ListOptions) (*ParticipantListResponse, error) {
	endpoint := "status/v1/participant/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result ParticipantListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetParticipant retrieves a specific participant by UUID
func (s *Service) GetParticipant(ctx context.Context, uuid string) (*Participant, error) {
	endpoint := fmt.Sprintf("status/v1/participant/%s/", uuid)

	var result Participant
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}
