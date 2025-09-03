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

func TestService_ListSchedulingOperations(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	creationTime := time.Now().Add(-1 * time.Hour)

	expectedResponse := &SchedulingOperationListResponse{
		Objects: []SchedulingOperation{
			{
				CreationTime:     &util.InfinityTime{Time: creationTime},
				ErrorCode:        "",
				ErrorDescription: "",
				ID:               1,
				ResourceID:       nil,
				ResourceURI:      "/api/admin/status/v1/scheduling_operation/1/",
				Success:          true,
				TransactionUUID:  "uuid-1",
			},
		},
	}

	client.On("GetJSON", t.Context(), "status/v1/scheduling_operation/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*status.SchedulingOperationListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*SchedulingOperationListResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ListSchedulingOperations(t.Context(), nil)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, 1, result.Objects[0].ID)
	assert.Equal(t, "/api/admin/status/v1/scheduling_operation/1/", result.Objects[0].ResourceURI)
	assert.Equal(t, true, result.Objects[0].Success)
	assert.Equal(t, "uuid-1", result.Objects[0].TransactionUUID)
	assert.Equal(t, creationTime.Unix(), result.Objects[0].CreationTime.Time.Unix())
	client.AssertExpectations(t)
}

func TestService_ListSchedulingOperations_WithOptions(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	opts := &ListOptions{
		Limit:  12,
		Offset: 3,
	}

	creationTime := time.Now().Add(-1 * time.Hour)
	expectedResponse := &SchedulingOperationListResponse{
		Objects: []SchedulingOperation{
			{
				CreationTime:     &util.InfinityTime{Time: creationTime},
				ErrorCode:        "ERR-1",
				ErrorDescription: "Some error",
				ID:               2,
				ResourceID:       nil,
				ResourceURI:      "",
				Success:          false,
				TransactionUUID:  "uuid-2",
			},
		},
	}

	client.On("GetJSON", t.Context(), "status/v1/scheduling_operation/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*status.SchedulingOperationListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*SchedulingOperationListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListSchedulingOperations(t.Context(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, 2, result.Objects[0].ID)
	assert.Equal(t, "ERR-1", result.Objects[0].ErrorCode)
	assert.Equal(t, "Some error", result.Objects[0].ErrorDescription)
	assert.Equal(t, false, result.Objects[0].Success)
	assert.Equal(t, "uuid-2", result.Objects[0].TransactionUUID)
	assert.Equal(t, creationTime.Unix(), result.Objects[0].CreationTime.Time.Unix())

	client.AssertExpectations(t)
}

func TestService_GetSchedulingOperation(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	creationTime := time.Now().Add(-2 * time.Hour)
	resourceID := 42
	expectedOperation := &SchedulingOperation{
		CreationTime:     &util.InfinityTime{Time: creationTime},
		ErrorCode:        "",
		ErrorDescription: "",
		ID:               1,
		ResourceID:       &resourceID,
		ResourceURI:      "/api/admin/status/v1/scheduling_operation/1/",
		Success:          true,
		TransactionUUID:  "uuid-3",
	}

	client.On("GetJSON", t.Context(), "status/v1/scheduling_operation/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*status.SchedulingOperation")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*SchedulingOperation)
		*result = *expectedOperation
	})

	service := New(client)
	result, err := service.GetSchedulingOperation(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedOperation, result)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "/api/admin/status/v1/scheduling_operation/1/", result.ResourceURI)
	assert.Equal(t, true, result.Success)
	assert.Equal(t, "uuid-3", result.TransactionUUID)
	assert.Equal(t, creationTime.Unix(), result.CreationTime.Time.Unix())
	assert.NotNil(t, result.ResourceID)
	assert.Equal(t, 42, *result.ResourceID)
	client.AssertExpectations(t)
}
