/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package status

import (
	"testing"
	"time"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListExchangeSchedulers(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	lastModified := time.Now().Add(-1 * time.Hour)

	expectedResponse := &ExchangeSchedulerListResponse{
		Meta: Meta{
			Limit:      100,
			Next:       "",
			Offset:     0,
			Previous:   "",
			TotalCount: 1,
		},
		Objects: []ExchangeScheduler{
			{
				ExchangeConnectorID: 10,
				ID:                  1,
				LastModifiedTime:    &util.InfinityTime{Time: lastModified},
				ResourceURI:         "/api/admin/status/v1/exchange_scheduler/1/",
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
	assert.Equal(t, 1, result.Objects[0].ID)
	assert.Equal(t, 10, result.Objects[0].ExchangeConnectorID)
	assert.Equal(t, "/api/admin/status/v1/exchange_scheduler/1/", result.Objects[0].ResourceURI)
	assert.Equal(t, lastModified.Unix(), result.Objects[0].LastModifiedTime.Time.Unix())
	client.AssertExpectations(t)
}

func TestService_ListExchangeSchedulers_WithOptions(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	opts := &ListOptions{
		Limit:  10,
		Offset: 0,
	}

	expectedResponse := &ExchangeSchedulerListResponse{
		Meta: Meta{
			Limit:      10,
			Next:       "",
			Offset:     0,
			Previous:   "",
			TotalCount: 1,
		},
		Objects: []ExchangeScheduler{
			{
				ExchangeConnectorID: 20,
				ID:                  2,
				LastModifiedTime:    nil,
				ResourceURI:         "/api/admin/status/v1/exchange_scheduler/2/",
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
	assert.Equal(t, 2, result.Objects[0].ID)
	assert.Equal(t, 20, result.Objects[0].ExchangeConnectorID)
	assert.Equal(t, "/api/admin/status/v1/exchange_scheduler/2/", result.Objects[0].ResourceURI)
	assert.Nil(t, result.Objects[0].LastModifiedTime)

	client.AssertExpectations(t)
}

func TestService_GetExchangeScheduler(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	lastModified := time.Now().Add(-30 * time.Minute)
	expectedScheduler := &ExchangeScheduler{
		ExchangeConnectorID: 10,
		ID:                  1,
		LastModifiedTime:    &util.InfinityTime{Time: lastModified},
		ResourceURI:         "/api/admin/status/v1/exchange_scheduler/1/",
	}

	client.On("GetJSON", t.Context(), "status/v1/exchange_scheduler/1/", mock.AnythingOfType("*status.ExchangeScheduler")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*ExchangeScheduler)
		*result = *expectedScheduler
	})

	service := New(client)
	result, err := service.GetExchangeScheduler(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedScheduler, result)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, 10, result.ExchangeConnectorID)
	assert.Equal(t, "/api/admin/status/v1/exchange_scheduler/1/", result.ResourceURI)
	assert.Equal(t, lastModified.Unix(), result.LastModifiedTime.Time.Unix())
	client.AssertExpectations(t)
}
