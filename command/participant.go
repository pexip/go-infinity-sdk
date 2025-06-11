package command

import (
	"context"
)

// DisconnectParticipantByID disconnects a participant using participant_id (schema-compliant)
func (s *Service) DisconnectParticipantByID(ctx context.Context, participantID string) (*CommandResponse, error) {
	endpoint := "command/v1/participant/disconnect/"

	req := &ParticipantDisconnectRequest{
		ParticipantID: participantID,
	}

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// MuteParticipantByID mutes a participant using participant_id (schema-compliant)
func (s *Service) MuteParticipantByID(ctx context.Context, participantID string) (*CommandResponse, error) {
	endpoint := "command/v1/participant/mute/"

	req := &ParticipantMuteRequest{
		ParticipantID: participantID,
	}

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// UnmuteParticipantByID unmutes a participant using participant_id (schema-compliant)
func (s *Service) UnmuteParticipantByID(ctx context.Context, participantID string) (*CommandResponse, error) {
	endpoint := "command/v1/participant/unmute/"

	req := &ParticipantUnmuteRequest{
		ParticipantID: participantID,
	}

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// ChangeParticipantRoleByID changes a participant's role using participant_id (schema-compliant)
func (s *Service) ChangeParticipantRoleByID(ctx context.Context, participantID, role string) (*CommandResponse, error) {
	endpoint := "command/v1/participant/role/"

	req := &ParticipantRoleRequest{
		ParticipantID: participantID,
		Role:          role,
	}

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// TransferParticipantByID transfers a participant using participant_id (schema-compliant)
func (s *Service) TransferParticipantByID(ctx context.Context, participantID, conferenceAlias string, role string) (*CommandResponse, error) {
	endpoint := "command/v1/participant/transfer/"

	req := &ParticipantTransferRequest{
		ParticipantID:   participantID,
		ConferenceAlias: conferenceAlias,
		Role:            role,
	}

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// PromoteParticipantByID promotes a participant to chair role using participant_id
func (s *Service) PromoteParticipantByID(ctx context.Context, participantID string) (*CommandResponse, error) {
	return s.ChangeParticipantRoleByID(ctx, participantID, "chair")
}

// DemoteParticipantByID demotes a participant to guest role using participant_id
func (s *Service) DemoteParticipantByID(ctx context.Context, participantID string) (*CommandResponse, error) {
	return s.ChangeParticipantRoleByID(ctx, participantID, "guest")
}
