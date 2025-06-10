package history

import (
	"testing"
	"time"

	mockClient "github.com/pexip/go-infinity-sdk/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
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
					Objects: []Conference{
						{ID: 1, Name: "Test Conference 1", DurationSeconds: 3600},
						{ID: 2, Name: "Test Conference 2", DurationSeconds: 1800},
					},
				}
				m.On("GetJSON", t.Context(), "history/v1/conference/", mock.AnythingOfType("*history.ConferenceListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*ConferenceListResponse)
					*result = *expectedResponse
				})
			},
			wantErr: false,
		},
		{
			name: "successful list with time range and search",
			opts: &ListOptions{
				Limit:     10,
				Offset:    5,
				StartTime: &time.Time{},
				EndTime:   &time.Time{},
				Search:    "test",
			},
			setup: func(m *mockClient.Client) {
				expectedResponse := &ConferenceListResponse{
					Objects: []Conference{
						{ID: 1, Name: "Test Conference"},
					},
				}
				// Note: The exact query string will vary based on time formatting
				m.On("GetJSON", t.Context(), mock.MatchedBy(func(endpoint string) bool {
					return endpoint != "" && endpoint != "history/v1/conference/"
				}), mock.AnythingOfType("*history.ConferenceListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*ConferenceListResponse)
					*result = *expectedResponse
				})
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := &mockClient.Client{}
			tt.setup(mockClient)

			service := New(mockClient)
			result, err := service.ListConferences(t.Context(), tt.opts)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
			}

			mockClient.AssertExpectations(t)
		})
	}
}

func TestService_GetConference(t *testing.T) {
	mockClient := &mockClient.Client{}
	expectedConference := &Conference{
		ID:                1,
		Name:              "Test Conference",
		ServiceType:       "conference",
		DurationSeconds:   3600,
		TotalParticipants: 5,
	}

	mockClient.On("GetJSON", t.Context(), "history/v1/conference/1/", mock.AnythingOfType("*history.Conference")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*Conference)
		*result = *expectedConference
	})

	service := New(mockClient)
	result, err := service.GetConference(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedConference, result)
	mockClient.AssertExpectations(t)
}

func TestService_ListParticipants(t *testing.T) {
	mockClient := &mockClient.Client{}

	expectedResponse := &ParticipantListResponse{
		Objects: []Participant{
			{
				ID:               1,
				ConferenceID:     1,
				ConferenceName:   "Test Conference",
				DisplayName:      "John Doe",
				Role:             "chair",
				DurationSeconds:  1800,
				DisconnectReason: "normal",
			},
		},
	}

	mockClient.On("GetJSON", t.Context(), "history/v1/participant/", mock.AnythingOfType("*history.ParticipantListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*ParticipantListResponse)
		*result = *expectedResponse
	})

	service := New(mockClient)
	result, err := service.ListParticipants(t.Context(), nil)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "John Doe", result.Objects[0].DisplayName)
	mockClient.AssertExpectations(t)
}

func TestService_GetParticipant(t *testing.T) {
	mockClient := &mockClient.Client{}
	expectedParticipant := &Participant{
		ID:               1,
		ConferenceID:     1,
		ConferenceName:   "Test Conference",
		DisplayName:      "John Doe",
		Role:             "chair",
		DurationSeconds:  1800,
		DisconnectReason: "normal",
	}

	mockClient.On("GetJSON", t.Context(), "history/v1/participant/1/", mock.AnythingOfType("*history.Participant")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*Participant)
		*result = *expectedParticipant
	})

	service := New(mockClient)
	result, err := service.GetParticipant(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedParticipant, result)
	mockClient.AssertExpectations(t)
}

func TestService_ListParticipantsByConference(t *testing.T) {
	mockClient := &mockClient.Client{}

	expectedResponse := &ParticipantListResponse{
		Objects: []Participant{
			{
				ID:             1,
				ConferenceID:   123,
				ConferenceName: "Test Conference",
				DisplayName:    "Alice",
				Role:           "chair",
			},
			{
				ID:             2,
				ConferenceID:   123,
				ConferenceName: "Test Conference",
				DisplayName:    "Bob",
				Role:           "guest",
			},
		},
	}

	mockClient.On("GetJSON", t.Context(), mock.MatchedBy(func(endpoint string) bool {
		return endpoint != "" && endpoint != "history/v1/participant/"
	}), mock.AnythingOfType("*history.ParticipantListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*ParticipantListResponse)
		*result = *expectedResponse
	})

	service := New(mockClient)
	result, err := service.ListParticipantsByConference(t.Context(), 123, nil)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result.Objects))
	assert.Equal(t, "Alice", result.Objects[0].DisplayName)
	assert.Equal(t, "Bob", result.Objects[1].DisplayName)
	mockClient.AssertExpectations(t)
}

func TestService_ListMediaStreams(t *testing.T) {
	mockClient := &mockClient.Client{}

	expectedResponse := &MediaStreamListResponse{
		Objects: []MediaStream{
			{
				ID:              1,
				ParticipantID:   1,
				StreamType:      "video",
				Direction:       "sendrecv",
				Codec:           "H.264",
				Resolution:      "1920x1080",
				Bitrate:         2000000,
				DurationSeconds: 1800,
			},
		},
	}

	mockClient.On("GetJSON", t.Context(), "history/v1/media_stream/", mock.AnythingOfType("*history.MediaStreamListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*MediaStreamListResponse)
		*result = *expectedResponse
	})

	service := New(mockClient)
	result, err := service.ListMediaStreams(t.Context(), nil)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "video", result.Objects[0].StreamType)
	assert.Equal(t, "H.264", result.Objects[0].Codec)
	mockClient.AssertExpectations(t)
}

func TestService_GetMediaStream(t *testing.T) {
	mockClient := &mockClient.Client{}
	expectedStream := &MediaStream{
		ID:              1,
		ParticipantID:   1,
		StreamType:      "audio",
		Direction:       "sendrecv",
		Codec:           "Opus",
		Bitrate:         64000,
		DurationSeconds: 1800,
	}

	mockClient.On("GetJSON", t.Context(), "history/v1/media_stream/1/", mock.AnythingOfType("*history.MediaStream")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*MediaStream)
		*result = *expectedStream
	})

	service := New(mockClient)
	result, err := service.GetMediaStream(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedStream, result)
	mockClient.AssertExpectations(t)
}

func TestService_ListMediaStreamsByParticipant(t *testing.T) {
	mockClient := &mockClient.Client{}

	expectedResponse := &MediaStreamListResponse{
		Objects: []MediaStream{
			{
				ID:            1,
				ParticipantID: 456,
				StreamType:    "video",
				Direction:     "sendrecv",
				Codec:         "H.264",
			},
			{
				ID:            2,
				ParticipantID: 456,
				StreamType:    "audio",
				Direction:     "sendrecv",
				Codec:         "Opus",
			},
		},
	}

	mockClient.On("GetJSON", t.Context(), mock.MatchedBy(func(endpoint string) bool {
		return endpoint != "" && endpoint != "history/v1/media_stream/"
	}), mock.AnythingOfType("*history.MediaStreamListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*MediaStreamListResponse)
		*result = *expectedResponse
	})

	service := New(mockClient)
	result, err := service.ListMediaStreamsByParticipant(t.Context(), 456, nil)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result.Objects))
	assert.Equal(t, "video", result.Objects[0].StreamType)
	assert.Equal(t, "audio", result.Objects[1].StreamType)
	mockClient.AssertExpectations(t)
}

func TestNew(t *testing.T) {
	mockClient := &mockClient.Client{}
	service := New(mockClient)

	require.NotNil(t, service)
	assert.Equal(t, mockClient, service.client)
}
