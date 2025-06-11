package command

import (
	"testing"

	mockClient "github.com/pexip/go-infinity-sdk/v38/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_SendConferenceEmail(t *testing.T) {
	tests := []struct {
		name                     string
		conferenceID             int
		conferenceSyncTemplateID *int
		wantErr                  bool
	}{
		{
			name:                     "send email without template",
			conferenceID:             123,
			conferenceSyncTemplateID: nil,
			wantErr:                  false,
		},
		{
			name:                     "send email with template",
			conferenceID:             456,
			conferenceSyncTemplateID: intPtr(789),
			wantErr:                  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &mockClient.Client{}

			expectedRequest := &ConferenceSendEmailRequest{
				ConferenceID:             tt.conferenceID,
				ConferenceSyncTemplateID: tt.conferenceSyncTemplateID,
			}

			expectedResponse := &CommandResponse{
				Status:  "success",
				Message: "Conference email sent",
			}

			client.On("PostJSON", t.Context(), "command/v1/conference/send_email/", expectedRequest, mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
				result := args.Get(3).(*CommandResponse)
				*result = *expectedResponse
			})

			service := New(client)
			result, err := service.SendConferenceEmail(t.Context(), tt.conferenceID, tt.conferenceSyncTemplateID)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, expectedResponse, result)
			}

			client.AssertExpectations(t)
		})
	}
}

func TestService_SendDeviceEmail(t *testing.T) {
	tests := []struct {
		name                     string
		deviceID                 int
		conferenceSyncTemplateID *int
		wantErr                  bool
	}{
		{
			name:                     "send device email without template",
			deviceID:                 321,
			conferenceSyncTemplateID: nil,
			wantErr:                  false,
		},
		{
			name:                     "send device email with template",
			deviceID:                 654,
			conferenceSyncTemplateID: intPtr(987),
			wantErr:                  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &mockClient.Client{}

			expectedRequest := &DeviceSendEmailRequest{
				DeviceID:                 tt.deviceID,
				ConferenceSyncTemplateID: tt.conferenceSyncTemplateID,
			}

			expectedResponse := &CommandResponse{
				Status:  "success",
				Message: "Device email sent",
			}

			client.On("PostJSON", t.Context(), "command/v1/device/send_email/", expectedRequest, mock.AnythingOfType("*command.CommandResponse")).Return(nil).Run(func(args mock.Arguments) {
				result := args.Get(3).(*CommandResponse)
				*result = *expectedResponse
			})

			service := New(client)
			result, err := service.SendDeviceEmail(t.Context(), tt.deviceID, tt.conferenceSyncTemplateID)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, expectedResponse, result)
			}

			client.AssertExpectations(t)
		})
	}
}

// intPtr is a helper function to get a pointer to an int value
func intPtr(i int) *int {
	return &i
}
