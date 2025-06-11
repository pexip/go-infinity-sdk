package status

import (
	"testing"

	mockClient "github.com/pexip/go-infinity-sdk/v38/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListConferences(t *testing.T) {
	tests := []struct {
		name    string
		opts    *ListOptions
		setup   func(m *mockClient.Client)
		wantErr bool
	}{
		{
			name: "successful list without options",
			opts: nil,
			setup: func(m *mockClient.Client) {
				expectedResponse := &ConferenceListResponse{
					Objects: []ConferenceStatus{
						{ID: "1", Name: "Test Conference 1", IsStarted: true, ServiceType: "conference"},
						{ID: "2", Name: "Test Conference 2", IsStarted: false, ServiceType: "conference"},
					},
				}
				m.On("GetJSON", t.Context(), "status/v1/conference/", mock.AnythingOfType("*status.ConferenceListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*ConferenceListResponse)
					*result = *expectedResponse
				})
			},
			wantErr: false,
		},
		{
			name: "successful list with options",
			opts: &ListOptions{
				Limit:  5,
				Offset: 10,
			},
			setup: func(m *mockClient.Client) {
				expectedResponse := &ConferenceListResponse{
					Objects: []ConferenceStatus{
						{ID: "3", Name: "Test Conference 3", IsStarted: true, ServiceType: "conference"},
					},
				}
				m.On("GetJSON", t.Context(), "status/v1/conference/?limit=5&offset=10", mock.AnythingOfType("*status.ConferenceListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*ConferenceListResponse)
					*result = *expectedResponse
				})
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &mockClient.Client{}
			tt.setup(client)

			service := New(client)
			result, err := service.ListConferences(t.Context(), tt.opts)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
			}

			client.AssertExpectations(t)
		})
	}
}

func TestService_GetConference(t *testing.T) {
	client := &mockClient.Client{}
	expectedConference := &ConferenceStatus{
		ID:          "1",
		Name:        "Test Conference",
		ServiceType: "conference",
		IsStarted:   true,
	}

	client.On("GetJSON", t.Context(), "status/v1/conference/1/", mock.AnythingOfType("*status.ConferenceStatus")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*ConferenceStatus)
		*result = *expectedConference
	})

	service := New(client)
	result, err := service.GetConference(t.Context(), "1")

	assert.NoError(t, err)
	assert.Equal(t, expectedConference, result)
	client.AssertExpectations(t)
}
