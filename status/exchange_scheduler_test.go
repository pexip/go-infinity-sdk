package status

import (
	"testing"
	"time"

	mockClient "github.com/pexip/go-infinity-sdk/v38/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListExchangeSchedulers(t *testing.T) {
	client := &mockClient.Client{}

	lastSync := time.Now().Add(-1 * time.Hour)
	nextSync := time.Now().Add(1 * time.Hour)

	expectedResponse := &ExchangeSchedulerListResponse{
		Objects: []ExchangeScheduler{
			{
				ID:                1,
				Name:              "Primary Exchange",
				Status:            "active",
				LastSync:          &lastSync,
				NextSync:          &nextSync,
				ProcessedMeetings: 45,
				ErrorCount:        0,
				ResourceURI:       "/api/admin/status/v1/exchange_scheduler/1/",
			},
		},
	}

	client.On("GetJSON", t.Context(), "status/v1/exchange_scheduler/", mock.AnythingOfType("*status.ExchangeSchedulerListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*ExchangeSchedulerListResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ListExchangeSchedulers(t.Context(), nil)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "Primary Exchange", result.Objects[0].Name)
	assert.Equal(t, "active", result.Objects[0].Status)
	assert.Equal(t, 45, result.Objects[0].ProcessedMeetings)
	client.AssertExpectations(t)
}

func TestService_ListExchangeSchedulers_WithOptions(t *testing.T) {
	client := &mockClient.Client{}
	service := New(client)

	opts := &ListOptions{
		Limit:  10,
		Offset: 0,
	}

	expectedResponse := &ExchangeSchedulerListResponse{
		Objects: []ExchangeScheduler{
			{
				ID:                2,
				Name:              "Test Exchange With Options",
				Status:            "active",
				ProcessedMeetings: 25,
				ErrorCount:        0,
			},
		},
	}

	client.On("GetJSON", t.Context(), mock.MatchedBy(func(endpoint string) bool {
		return endpoint != "status/v1/exchange_scheduler/"
	}), mock.AnythingOfType("*status.ExchangeSchedulerListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*ExchangeSchedulerListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListExchangeSchedulers(t.Context(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "Test Exchange With Options", result.Objects[0].Name)

	client.AssertExpectations(t)
}

func TestService_GetExchangeScheduler(t *testing.T) {
	client := &mockClient.Client{}

	lastSync := time.Now().Add(-30 * time.Minute)
	nextSync := time.Now().Add(30 * time.Minute)
	expectedScheduler := &ExchangeScheduler{
		ID:                1,
		Name:              "Test Exchange",
		Status:            "syncing",
		LastSync:          &lastSync,
		NextSync:          &nextSync,
		ProcessedMeetings: 12,
		ErrorCount:        1,
		ResourceURI:       "/api/admin/status/v1/exchange_scheduler/1/",
	}

	client.On("GetJSON", t.Context(), "status/v1/exchange_scheduler/1/", mock.AnythingOfType("*status.ExchangeScheduler")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*ExchangeScheduler)
		*result = *expectedScheduler
	})

	service := New(client)
	result, err := service.GetExchangeScheduler(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedScheduler, result)
	client.AssertExpectations(t)
}
