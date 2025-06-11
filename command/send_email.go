package command

import (
	"context"
)

// SendConferenceEmail sends a reminder email for a conference
func (s *Service) SendConferenceEmail(ctx context.Context, conferenceID int, conferenceSyncTemplateID *int) (*CommandResponse, error) {
	endpoint := "command/v1/conference/send_email/"

	req := &ConferenceSendEmailRequest{
		ConferenceID:             conferenceID,
		ConferenceSyncTemplateID: conferenceSyncTemplateID,
	}

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// SendDeviceEmail sends a reminder email for a device
func (s *Service) SendDeviceEmail(ctx context.Context, deviceID int, conferenceSyncTemplateID *int) (*CommandResponse, error) {
	endpoint := "command/v1/device/send_email/"

	req := &DeviceSendEmailRequest{
		DeviceID:                 deviceID,
		ConferenceSyncTemplateID: conferenceSyncTemplateID,
	}

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}
