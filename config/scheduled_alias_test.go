/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"testing"
	"time"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/pexip/go-infinity-sdk/v38/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListScheduledAliases(t *testing.T) {
	tests := []struct {
		name    string
		opts    *ListOptions
		setup   func(m *interfaces.HTTPClientMock)
		wantErr bool
	}{
		{
			name: "successful list without options",
			opts: nil,
			setup: func(m *interfaces.HTTPClientMock) {
				ewsItemUID1 := "ews-uid-1"
				ewsItemUID2 := "ews-uid-2"
				creationTime := util.InfinityTime{Time: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC)}
				deletionTime := util.InfinityTime{Time: time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC)}

				expectedResponse := &ScheduledAliasListResponse{
					Objects: []ScheduledAlias{
						{ID: 1, Alias: "meeting.room1", AliasNumber: 1001, NumericAlias: "1001", UUID: "uuid-1", ExchangeConnector: "exchange1", IsUsed: true, EWSItemUID: &ewsItemUID1, CreationTime: creationTime},
						{ID: 2, Alias: "meeting.room2", AliasNumber: 1002, NumericAlias: "1002", UUID: "uuid-2", ExchangeConnector: "exchange2", IsUsed: false, EWSItemUID: &ewsItemUID2, ConferenceDeletionTime: &deletionTime},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/scheduled_alias/", mock.AnythingOfType("*config.ScheduledAliasListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*ScheduledAliasListResponse)
					*result = *expectedResponse
				})
			},
			wantErr: false,
		},
		{
			name: "successful list with options",
			opts: &ListOptions{
				BaseListOptions: options.BaseListOptions{
					Limit: 5,
				},
				Search: "room1",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				ewsItemUID := "ews-uid-1"
				creationTime := util.InfinityTime{Time: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC)}

				expectedResponse := &ScheduledAliasListResponse{
					Objects: []ScheduledAlias{
						{ID: 1, Alias: "meeting.room1", AliasNumber: 1001, NumericAlias: "1001", UUID: "uuid-1", ExchangeConnector: "exchange1", IsUsed: true, EWSItemUID: &ewsItemUID, CreationTime: creationTime},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/scheduled_alias/?limit=5&name__icontains=room1", mock.AnythingOfType("*config.ScheduledAliasListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*ScheduledAliasListResponse)
					*result = *expectedResponse
				})
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := interfaces.NewHTTPClientMock()
			tt.setup(client)

			service := New(client)
			result, err := service.ListScheduledAliases(t.Context(), tt.opts)

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

func TestService_GetScheduledAlias(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	ewsItemUID := "test-ews-uid"
	creationTime := util.InfinityTime{Time: time.Date(2023, 6, 15, 14, 30, 0, 0, time.UTC)}
	deletionTime := util.InfinityTime{Time: time.Date(2023, 12, 25, 18, 0, 0, 0, time.UTC)}

	expectedScheduledAlias := &ScheduledAlias{
		ID:                     1,
		Alias:                  "test.meeting",
		AliasNumber:            1234,
		NumericAlias:           "1234",
		UUID:                   "test-uuid-123",
		ExchangeConnector:      "test-exchange",
		IsUsed:                 true,
		EWSItemUID:             &ewsItemUID,
		CreationTime:           creationTime,
		ConferenceDeletionTime: &deletionTime,
	}

	client.On("GetJSON", t.Context(), "configuration/v1/scheduled_alias/1/", mock.AnythingOfType("*config.ScheduledAlias")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*ScheduledAlias)
		*result = *expectedScheduledAlias
	})

	service := New(client)
	result, err := service.GetScheduledAlias(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedScheduledAlias, result)
	client.AssertExpectations(t)
}

func TestService_CreateScheduledAlias(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	ewsItemUID := "new-ews-uid"
	createRequest := &ScheduledAliasCreateRequest{
		Alias:             "new.meeting",
		AliasNumber:       5678,
		NumericAlias:      "5678",
		UUID:              "new-uuid-456",
		ExchangeConnector: "new-exchange",
		IsUsed:            false,
		EWSItemUID:        &ewsItemUID,
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/scheduled_alias/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/scheduled_alias/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateScheduledAlias(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateScheduledAlias(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	aliasNumber := 9999
	isUsed := false
	ewsItemUID := "updated-ews-uid"

	updateRequest := &ScheduledAliasUpdateRequest{
		Alias:        "updated.meeting",
		AliasNumber:  &aliasNumber,
		NumericAlias: "9999",
		IsUsed:       &isUsed,
		EWSItemUID:   &ewsItemUID,
	}

	creationTime := util.InfinityTime{Time: time.Date(2023, 6, 15, 14, 30, 0, 0, time.UTC)}
	expectedScheduledAlias := &ScheduledAlias{
		ID:                1,
		Alias:             "updated.meeting",
		AliasNumber:       9999,
		NumericAlias:      "9999",
		UUID:              "test-uuid-123",
		ExchangeConnector: "test-exchange",
		IsUsed:            false,
		EWSItemUID:        &ewsItemUID,
		CreationTime:      creationTime,
	}

	client.On("PutJSON", t.Context(), "configuration/v1/scheduled_alias/1/", updateRequest, mock.AnythingOfType("*config.ScheduledAlias")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*ScheduledAlias)
		*result = *expectedScheduledAlias
	})

	service := New(client)
	result, err := service.UpdateScheduledAlias(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedScheduledAlias, result)
	client.AssertExpectations(t)
}

func TestService_DeleteScheduledAlias(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/scheduled_alias/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteScheduledAlias(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
