package history

import (
	"testing"
	"time"

	mockClient "github.com/pexip/go-infinity-sdk/v38/internal/mock"
	"github.com/pexip/go-infinity-sdk/v38/options"
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
				expectedResponse := &ConferenceRecordListResponse{
					Objects: []ConferenceRecord{
						{ID: 1, Name: "Test ConferenceRecord 1", DurationSeconds: 3600},
						{ID: 2, Name: "Test ConferenceRecord 2", DurationSeconds: 1800},
					},
				}
				m.On("GetJSON", t.Context(), "history/v1/conference/", mock.AnythingOfType("*history.ConferenceRecordListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*ConferenceRecordListResponse)
					*result = *expectedResponse
				})
			},
			wantErr: false,
		},
		{
			name: "successful list with time range and search",
			opts: &ListOptions{
				SearchableListOptions: options.SearchableListOptions{
					BaseListOptions: options.BaseListOptions{
						Limit:  10,
						Offset: 5,
					},
					Search: "test",
				},
				StartTime: &time.Time{},
				EndTime:   &time.Time{},
			},
			setup: func(m *mockClient.Client) {
				expectedResponse := &ConferenceRecordListResponse{
					Objects: []ConferenceRecord{
						{ID: 1, Name: "Test ConferenceRecord"},
					},
				}
				// Note: The exact query string will vary based on time formatting
				m.On("GetJSON", t.Context(), mock.MatchedBy(func(endpoint string) bool {
					return endpoint != "" && endpoint != "history/v1/conference/"
				}), mock.AnythingOfType("*history.ConferenceRecordListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*ConferenceRecordListResponse)
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
			result, err := service.ListConferenceRecords(t.Context(), tt.opts)

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
	expectedConference := &ConferenceRecord{
		ID:                1,
		Name:              "Test ConferenceRecord",
		ServiceType:       "conference",
		DurationSeconds:   3600,
		TotalParticipants: 5,
	}

	client.On("GetJSON", t.Context(), "history/v1/conference/1/", mock.AnythingOfType("*history.ConferenceRecord")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*ConferenceRecord)
		*result = *expectedConference
	})

	service := New(client)
	result, err := service.GetConferenceRecord(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedConference, result)
	client.AssertExpectations(t)
}
