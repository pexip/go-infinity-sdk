package status

import (
	"testing"
	"time"

	mockClient "github.com/pexip/go-infinity-sdk/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListAlarms(t *testing.T) {
	client := &mockClient.Client{}

	expectedResponse := &AlarmListResponse{
		Objects: []Alarm{
			{
				ID:       1,
				Level:    "warning",
				Name:     "High CPU Usage",
				Details:  "CPU usage is above 80%",
				Instance: "worker-1",
				Node:     "worker-1",
			},
			{
				ID:       2,
				Level:    "critical",
				Name:     "Service Down",
				Details:  "Media service is not responding",
				Instance: "worker-2",
				Node:     "worker-2",
			},
		},
	}

	client.On("GetJSON", t.Context(), "status/v1/alarm/", mock.AnythingOfType("*status.AlarmListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*AlarmListResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ListAlarms(t.Context(), nil)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result.Objects))
	assert.Equal(t, "High CPU Usage", result.Objects[0].Name)
	assert.Equal(t, "warning", result.Objects[0].Level)
	client.AssertExpectations(t)
}

func TestService_ListAlarms_WithOptions(t *testing.T) {
	client := &mockClient.Client{}
	service := New(client)

	opts := &ListOptions{
		Limit:  10,
		Offset: 5,
	}

	expectedResponse := &AlarmListResponse{
		Objects: []Alarm{
			{
				ID:       1,
				Level:    "critical",
				Name:     "Test Alarm",
				Details:  "Test alarm details",
				Instance: "test-instance",
				Node:     "test-node",
			},
		},
	}

	client.On("GetJSON", t.Context(), mock.MatchedBy(func(endpoint string) bool {
		return endpoint != "status/v1/alarm/"
	}), mock.AnythingOfType("*status.AlarmListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*AlarmListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListAlarms(t.Context(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "Test Alarm", result.Objects[0].Name)

	client.AssertExpectations(t)
}

func TestService_GetAlarm(t *testing.T) {
	client := &mockClient.Client{}
	timeRaised := time.Now()
	expectedAlarm := &Alarm{
		ID:         1,
		Level:      "warning",
		Name:       "High CPU Usage",
		Details:    "CPU usage is above 80%",
		Instance:   "worker-1",
		Node:       "worker-1",
		TimeRaised: &timeRaised,
	}

	client.On("GetJSON", t.Context(), "status/v1/alarm/1/", mock.AnythingOfType("*status.Alarm")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*Alarm)
		*result = *expectedAlarm
	})

	service := New(client)
	result, err := service.GetAlarm(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedAlarm, result)
	client.AssertExpectations(t)
}