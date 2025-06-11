package status

import (
	"testing"
	"time"

	mockClient "github.com/pexip/go-infinity-sdk/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListSchedulingOperations(t *testing.T) {
	client := &mockClient.Client{}

	createdTime := time.Now().Add(-1 * time.Hour)
	completedTime := time.Now().Add(-30 * time.Minute)

	expectedResponse := &SchedulingOperationListResponse{
		Objects: []SchedulingOperation{
			{
				ID:             1,
				OperationType:  "create",
				Status:         "completed",
				CreatedTime:    &createdTime,
				CompletedTime:  &completedTime,
				ErrorMessage:   "",
				ConferenceName: "Test Meeting",
				ResourceURI:    "/api/admin/status/v1/scheduling_operation/1/",
			},
		},
	}

	client.On("GetJSON", t.Context(), "status/v1/scheduling_operation/", mock.AnythingOfType("*status.SchedulingOperationListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*SchedulingOperationListResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ListSchedulingOperations(t.Context(), nil)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "create", result.Objects[0].OperationType)
	assert.Equal(t, "completed", result.Objects[0].Status)
	client.AssertExpectations(t)
}

func TestService_ListSchedulingOperations_WithOptions(t *testing.T) {
	client := &mockClient.Client{}
	service := New(client)

	opts := &ListOptions{
		Limit:  12,
		Offset: 3,
	}

	createdTime := time.Now().Add(-1 * time.Hour)
	expectedResponse := &SchedulingOperationListResponse{
		Objects: []SchedulingOperation{
			{
				ID:             2,
				OperationType:  "delete",
				Status:         "completed",
				CreatedTime:    &createdTime,
				ErrorMessage:   "",
				ConferenceName: "Test Meeting With Options",
			},
		},
	}

	client.On("GetJSON", t.Context(), mock.MatchedBy(func(endpoint string) bool {
		return endpoint != "status/v1/scheduling_operation/"
	}), mock.AnythingOfType("*status.SchedulingOperationListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*SchedulingOperationListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListSchedulingOperations(t.Context(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "delete", result.Objects[0].OperationType)

	client.AssertExpectations(t)
}

func TestService_GetSchedulingOperation(t *testing.T) {
	client := &mockClient.Client{}

	createdTime := time.Now().Add(-2 * time.Hour)
	expectedOperation := &SchedulingOperation{
		ID:             1,
		OperationType:  "update",
		Status:         "running",
		CreatedTime:    &createdTime,
		ErrorMessage:   "",
		ConferenceName: "Updated Meeting",
		ResourceURI:    "/api/admin/status/v1/scheduling_operation/1/",
	}

	client.On("GetJSON", t.Context(), "status/v1/scheduling_operation/1/", mock.AnythingOfType("*status.SchedulingOperation")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*SchedulingOperation)
		*result = *expectedOperation
	})

	service := New(client)
	result, err := service.GetSchedulingOperation(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedOperation, result)
	client.AssertExpectations(t)
}