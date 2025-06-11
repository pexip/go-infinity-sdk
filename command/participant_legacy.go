package command

import (
	"context"
)

// DisconnectParticipant disconnects a participant from a conference
func (s *Service) DisconnectParticipant(ctx context.Context, participantUUID string) (*CommandResponse, error) {
	endpoint := "command/v1/participant/disconnect/"

	req := &ParticipantDisconnectRequestLegacy{
		ParticipantUUID: participantUUID,
	}

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// MuteParticipant mutes a participant
func (s *Service) MuteParticipant(ctx context.Context, participantUUID string) (*CommandResponse, error) {
	endpoint := "command/v1/participant/mute/"

	req := &ParticipantMuteRequestLegacy{
		ParticipantUUID: participantUUID,
		Setting:         "mute",
	}

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// UnmuteParticipant unmutes a participant
func (s *Service) UnmuteParticipant(ctx context.Context, participantUUID string) (*CommandResponse, error) {
	endpoint := "command/v1/participant/mute/"

	req := &ParticipantMuteRequestLegacy{
		ParticipantUUID: participantUUID,
		Setting:         "unmute",
	}

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// ToggleMuteParticipant toggles the mute status of a participant
func (s *Service) ToggleMuteParticipant(ctx context.Context, participantUUID string) (*CommandResponse, error) {
	endpoint := "command/v1/participant/mute/"

	req := &ParticipantMuteRequestLegacy{
		ParticipantUUID: participantUUID,
		Setting:         "toggle",
	}

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// SpotlightParticipant enables spotlight for a participant
func (s *Service) SpotlightParticipant(ctx context.Context, participantUUID string) (*CommandResponse, error) {
	endpoint := "command/v1/participant/spotlight/"

	req := &ParticipantSpotlightRequest{
		ParticipantUUID: participantUUID,
		Setting:         "on",
	}

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// UnspotlightParticipant disables spotlight for a participant
func (s *Service) UnspotlightParticipant(ctx context.Context, participantUUID string) (*CommandResponse, error) {
	endpoint := "command/v1/participant/spotlight/"

	req := &ParticipantSpotlightRequest{
		ParticipantUUID: participantUUID,
		Setting:         "off",
	}

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// ToggleSpotlightParticipant toggles the spotlight status of a participant
func (s *Service) ToggleSpotlightParticipant(ctx context.Context, participantUUID string) (*CommandResponse, error) {
	endpoint := "command/v1/participant/spotlight/"

	req := &ParticipantSpotlightRequest{
		ParticipantUUID: participantUUID,
		Setting:         "toggle",
	}

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// SendMessageToParticipant sends a message to a specific participant
func (s *Service) SendMessageToParticipant(ctx context.Context, participantUUID, message string) (*CommandResponse, error) {
	endpoint := "command/v1/participant/message/"

	req := &ParticipantMessageRequest{
		ParticipantUUID: participantUUID,
		Message:         message,
	}

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// TransferParticipant transfers a participant to another conference
func (s *Service) TransferParticipant(ctx context.Context, participantUUID, conferenceAlias string, opts *TransferOptions) (*CommandResponse, error) {
	endpoint := "command/v1/participant/transfer/"

	req := &ParticipantTransferRequestLegacy{
		ParticipantUUID: participantUUID,
		ConferenceAlias: conferenceAlias,
	}

	if opts != nil {
		req.Role = opts.Role
		req.PIN = opts.PIN
	}

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// ChangeParticipantRole changes a participant's role
func (s *Service) ChangeParticipantRole(ctx context.Context, participantUUID, role string) (*CommandResponse, error) {
	endpoint := "command/v1/participant/role/"

	req := &ParticipantRoleRequestLegacy{
		ParticipantUUID: participantUUID,
		Role:            role,
	}

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// PromoteParticipant promotes a participant to chair role
func (s *Service) PromoteParticipant(ctx context.Context, participantUUID string) (*CommandResponse, error) {
	return s.ChangeParticipantRole(ctx, participantUUID, "chair")
}

// DemoteParticipant demotes a participant to guest role
func (s *Service) DemoteParticipant(ctx context.Context, participantUUID string) (*CommandResponse, error) {
	return s.ChangeParticipantRole(ctx, participantUUID, "guest")
}
