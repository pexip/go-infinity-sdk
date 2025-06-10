package command

import (
	"context"
)

// ClientInterface defines the interface for the HTTP client
type ClientInterface interface {
	PostJSON(ctx context.Context, endpoint string, body interface{}, result interface{}) error
}

// Service handles Command API endpoints
type Service struct {
	client ClientInterface
}

// New creates a new Command API service
func New(client ClientInterface) *Service {
	return &Service{
		client: client,
	}
}

// ParticipantDisconnectRequest represents a request to disconnect a participant
type ParticipantDisconnectRequest struct {
	ParticipantUUID string `json:"participant_uuid"`
}

// ParticipantMuteRequest represents a request to mute/unmute a participant
type ParticipantMuteRequest struct {
	ParticipantUUID string `json:"participant_uuid"`
	Setting         string `json:"setting"` // "mute", "unmute", or "toggle"
}

// ParticipantSpotlightRequest represents a request to spotlight a participant
type ParticipantSpotlightRequest struct {
	ParticipantUUID string `json:"participant_uuid"`
	Setting         string `json:"setting"` // "on", "off", or "toggle"
}

// ParticipantLockRequest represents a request to lock/unlock a conference
type ConferenceLockRequest struct {
	ConferenceID int    `json:"conference_id"`
	Setting      string `json:"setting"` // "lock", "unlock", or "toggle"
}

// ParticipantMessageRequest represents a request to send a message to a participant
type ParticipantMessageRequest struct {
	ParticipantUUID string `json:"participant_uuid"`
	Message         string `json:"message"`
}

// ConferenceMessageRequest represents a request to send a message to all participants in a conference
type ConferenceMessageRequest struct {
	ConferenceID int    `json:"conference_id"`
	Message      string `json:"message"`
}

// ParticipantTransferRequest represents a request to transfer a participant
type ParticipantTransferRequest struct {
	ParticipantUUID string `json:"participant_uuid"`
	ConferenceAlias string `json:"conference_alias"`
	Role            string `json:"role,omitempty"` // "chair" or "guest"
	PIN             string `json:"pin,omitempty"`
}

// ConferenceStartRequest represents a request to start a conference
type ConferenceStartRequest struct {
	ConferenceAlias string `json:"conference_alias"`
}

// ConferenceStopRequest represents a request to stop a conference
type ConferenceStopRequest struct {
	ConferenceID int `json:"conference_id"`
}

// ParticipantRoleRequest represents a request to change a participant's role
type ParticipantRoleRequest struct {
	ParticipantUUID string `json:"participant_uuid"`
	Role            string `json:"role"` // "chair" or "guest"
}

// CommandResponse represents a generic command response
type CommandResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Result  string `json:"result,omitempty"`
}

// DisconnectParticipant disconnects a participant from a conference
func (s *Service) DisconnectParticipant(ctx context.Context, participantUUID string) (*CommandResponse, error) {
	endpoint := "command/v1/participant/disconnect/"

	req := &ParticipantDisconnectRequest{
		ParticipantUUID: participantUUID,
	}

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// MuteParticipant mutes a participant
func (s *Service) MuteParticipant(ctx context.Context, participantUUID string) (*CommandResponse, error) {
	endpoint := "command/v1/participant/mute/"

	req := &ParticipantMuteRequest{
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

	req := &ParticipantMuteRequest{
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

	req := &ParticipantMuteRequest{
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

// LockConference locks a conference
func (s *Service) LockConference(ctx context.Context, conferenceID int) (*CommandResponse, error) {
	endpoint := "command/v1/conference/lock/"

	req := &ConferenceLockRequest{
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

	req := &ConferenceLockRequest{
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

	req := &ConferenceLockRequest{
		ConferenceID: conferenceID,
		Setting:      "toggle",
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

// TransferParticipant transfers a participant to another conference
func (s *Service) TransferParticipant(ctx context.Context, participantUUID, conferenceAlias string, opts *TransferOptions) (*CommandResponse, error) {
	endpoint := "command/v1/participant/transfer/"

	req := &ParticipantTransferRequest{
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

// TransferOptions contains options for transferring participants
type TransferOptions struct {
	Role string // "chair" or "guest"
	PIN  string
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

// ChangeParticipantRole changes a participant's role
func (s *Service) ChangeParticipantRole(ctx context.Context, participantUUID, role string) (*CommandResponse, error) {
	endpoint := "command/v1/participant/role/"

	req := &ParticipantRoleRequest{
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
