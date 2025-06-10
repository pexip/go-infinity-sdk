package status

import (
	"testing"
	"time"

	mockClient "github.com/pexip/go-infinity-sdk/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestService_GetSystemStatus(t *testing.T) {
	mockClient := &mockClient.Client{}

	expectedStatus := &SystemStatus{
		Status:      "healthy",
		Version:     "29.0.0",
		Uptime:      3600,
		Timestamp:   time.Now(),
		HostName:    "pexip-mgmt",
		TotalMemory: 8589934592,
		UsedMemory:  4294967296,
		CPULoad:     25.5,
	}

	mockClient.On("GetJSON", t.Context(), "status/v1/system_status/", mock.AnythingOfType("*status.SystemStatus")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*SystemStatus)
		*result = *expectedStatus
	})

	service := New(mockClient)
	result, err := service.GetSystemStatus(t.Context())

	assert.NoError(t, err)
	assert.Equal(t, expectedStatus, result)
	mockClient.AssertExpectations(t)
}

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
						{ID: 1, Name: "Test Conference 1", Started: true},
						{ID: 2, Name: "Test Conference 2", Started: false},
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
						{ID: 3, Name: "Test Conference 3", Started: true},
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
	expectedConference := &ConferenceStatus{
		ID:               1,
		Name:             "Test Conference",
		ServiceType:      "conference",
		Started:          true,
		ParticipantCount: 5,
	}

	mockClient.On("GetJSON", t.Context(), "status/v1/conference/1/", mock.AnythingOfType("*status.ConferenceStatus")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*ConferenceStatus)
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
				ID:          "participant-1",
				UUID:        "uuid-1",
				DisplayName: "John Doe",
				Role:        "chair",
				IsMuted:     false,
			},
			{
				ID:          "participant-2",
				UUID:        "uuid-2",
				DisplayName: "Jane Smith",
				Role:        "guest",
				IsMuted:     true,
			},
		},
	}

	mockClient.On("GetJSON", t.Context(), "status/v1/participant/", mock.AnythingOfType("*status.ParticipantListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*ParticipantListResponse)
		*result = *expectedResponse
	})

	service := New(mockClient)
	result, err := service.ListParticipants(t.Context(), nil)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result.Objects))
	assert.Equal(t, "John Doe", result.Objects[0].DisplayName)
	assert.Equal(t, "chair", result.Objects[0].Role)
	mockClient.AssertExpectations(t)
}

func TestService_GetParticipant(t *testing.T) {
	mockClient := &mockClient.Client{}
	expectedParticipant := &Participant{
		ID:             "participant-1",
		UUID:           "test-uuid",
		DisplayName:    "John Doe",
		Role:           "chair",
		IsMuted:        false,
		IsPresenting:   true,
		ConferenceID:   1,
		ConferenceName: "Test Conference",
	}

	mockClient.On("GetJSON", t.Context(), "status/v1/participant/test-uuid/", mock.AnythingOfType("*status.Participant")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*Participant)
		*result = *expectedParticipant
	})

	service := New(mockClient)
	result, err := service.GetParticipant(t.Context(), "test-uuid")

	assert.NoError(t, err)
	assert.Equal(t, expectedParticipant, result)
	mockClient.AssertExpectations(t)
}

func TestService_ListWorkers(t *testing.T) {
	mockClient := &mockClient.Client{}

	expectedResponse := &WorkerListResponse{
		Objects: []Worker{
			{
				NodeID:       "worker-1",
				HostName:     "pexip-worker-1",
				Status:       "online",
				CPU:          25.5,
				Memory:       60.0,
				Conferences:  5,
				Participants: 25,
			},
			{
				NodeID:       "worker-2",
				HostName:     "pexip-worker-2",
				Status:       "online",
				CPU:          30.2,
				Memory:       45.8,
				Conferences:  3,
				Participants: 15,
			},
		},
	}

	mockClient.On("GetJSON", t.Context(), "status/v1/worker/", mock.AnythingOfType("*status.WorkerListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*WorkerListResponse)
		*result = *expectedResponse
	})

	service := New(mockClient)
	result, err := service.ListWorkers(t.Context(), nil)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result.Objects))
	assert.Equal(t, "pexip-worker-1", result.Objects[0].HostName)
	mockClient.AssertExpectations(t)
}

func TestService_GetWorker(t *testing.T) {
	mockClient := &mockClient.Client{}
	expectedWorker := &Worker{
		NodeID:       "worker-1",
		HostName:     "pexip-worker-1",
		Status:       "online",
		CPU:          25.5,
		Memory:       60.0,
		Conferences:  5,
		Participants: 25,
		LastSeen:     time.Now(),
	}

	mockClient.On("GetJSON", t.Context(), "status/v1/worker/worker-1/", mock.AnythingOfType("*status.Worker")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*Worker)
		*result = *expectedWorker
	})

	service := New(mockClient)
	result, err := service.GetWorker(t.Context(), "worker-1")

	assert.NoError(t, err)
	assert.Equal(t, expectedWorker, result)
	mockClient.AssertExpectations(t)
}

func TestService_ListAlarms(t *testing.T) {
	mockClient := &mockClient.Client{}

	expectedResponse := &AlarmListResponse{
		Objects: []Alarm{
			{
				ID:       1,
				Level:    "warning",
				Name:     "High CPU Usage",
				Details:  "CPU usage is above 80%",
				Instance: "worker-1",
				NodeID:   "worker-1",
				Cleared:  false,
			},
			{
				ID:       2,
				Level:    "critical",
				Name:     "Service Down",
				Details:  "Media service is not responding",
				Instance: "worker-2",
				NodeID:   "worker-2",
				Cleared:  true,
			},
		},
	}

	mockClient.On("GetJSON", t.Context(), "status/v1/alarm/", mock.AnythingOfType("*status.AlarmListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*AlarmListResponse)
		*result = *expectedResponse
	})

	service := New(mockClient)
	result, err := service.ListAlarms(t.Context(), nil)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result.Objects))
	assert.Equal(t, "High CPU Usage", result.Objects[0].Name)
	assert.Equal(t, "warning", result.Objects[0].Level)
	mockClient.AssertExpectations(t)
}

func TestService_GetAlarm(t *testing.T) {
	mockClient := &mockClient.Client{}
	expectedAlarm := &Alarm{
		ID:       1,
		Level:    "warning",
		Name:     "High CPU Usage",
		Details:  "CPU usage is above 80%",
		Instance: "worker-1",
		NodeID:   "worker-1",
		Cleared:  false,
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	mockClient.On("GetJSON", t.Context(), "status/v1/alarm/1/", mock.AnythingOfType("*status.Alarm")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*Alarm)
		*result = *expectedAlarm
	})

	service := New(mockClient)
	result, err := service.GetAlarm(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedAlarm, result)
	mockClient.AssertExpectations(t)
}

func TestNew(t *testing.T) {
	mockClient := &mockClient.Client{}
	service := New(mockClient)

	require.NotNil(t, service)
	assert.Equal(t, mockClient, service.client)
}
