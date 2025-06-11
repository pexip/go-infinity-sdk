package status

import (
	"testing"

	mockClient "github.com/pexip/go-infinity-sdk/v38/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListSystemLocations(t *testing.T) {
	client := &mockClient.Client{}

	expectedResponse := &SystemLocationListResponse{
		Objects: []SystemLocation{
			{
				ID:          1,
				Name:        "Main Office",
				Description: "Primary datacenter location",
				ActiveNodes: 8,
				TotalNodes:  10,
				ActiveCalls: 25,
				ResourceURI: "/api/admin/status/v1/system_location/1/",
			},
			{
				ID:          2,
				Name:        "Branch Office",
				Description: "Secondary office location",
				ActiveNodes: 3,
				TotalNodes:  5,
				ActiveCalls: 12,
				ResourceURI: "/api/admin/status/v1/system_location/2/",
			},
		},
	}

	client.On("GetJSON", t.Context(), "status/v1/system_location/", mock.AnythingOfType("*status.SystemLocationListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*SystemLocationListResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ListSystemLocations(t.Context(), nil)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result.Objects))
	assert.Equal(t, "Main Office", result.Objects[0].Name)
	assert.Equal(t, "Primary datacenter location", result.Objects[0].Description)
	assert.Equal(t, 8, result.Objects[0].ActiveNodes)
	assert.Equal(t, 10, result.Objects[0].TotalNodes)
	assert.Equal(t, 25, result.Objects[0].ActiveCalls)
	assert.Equal(t, "Branch Office", result.Objects[1].Name)
	assert.Equal(t, 3, result.Objects[1].ActiveNodes)
	assert.Equal(t, 12, result.Objects[1].ActiveCalls)
	client.AssertExpectations(t)
}

func TestService_GetSystemLocation(t *testing.T) {
	client := &mockClient.Client{}

	expectedLocation := &SystemLocation{
		ID:          1,
		Name:        "Headquarters",
		Description: "Main corporate headquarters location",
		ActiveNodes: 15,
		TotalNodes:  20,
		ActiveCalls: 45,
		ResourceURI: "/api/admin/status/v1/system_location/1/",
	}

	client.On("GetJSON", t.Context(), "status/v1/system_location/1/", mock.AnythingOfType("*status.SystemLocation")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*SystemLocation)
		*result = *expectedLocation
	})

	service := New(client)
	result, err := service.GetSystemLocation(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedLocation, result)
	assert.Equal(t, "Headquarters", result.Name)
	assert.Equal(t, "Main corporate headquarters location", result.Description)
	assert.Equal(t, 15, result.ActiveNodes)
	assert.Equal(t, 20, result.TotalNodes)
	assert.Equal(t, 45, result.ActiveCalls)
	client.AssertExpectations(t)
}

func TestService_ListSystemLocations_WithOptions(t *testing.T) {
	client := &mockClient.Client{}
	service := New(client)

	opts := &ListOptions{
		Limit:  5,
		Offset: 0,
	}

	expectedResponse := &SystemLocationListResponse{
		Objects: []SystemLocation{
			{
				ID:          3,
				Name:        "Remote Site",
				Description: "Remote office location",
				ActiveNodes: 2,
				TotalNodes:  3,
				ActiveCalls: 5,
			},
		},
	}

	client.On("GetJSON", t.Context(), mock.MatchedBy(func(endpoint string) bool {
		return endpoint != "status/v1/system_location/"
	}), mock.AnythingOfType("*status.SystemLocationListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*SystemLocationListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListSystemLocations(t.Context(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "Remote Site", result.Objects[0].Name)
	assert.Equal(t, "Remote office location", result.Objects[0].Description)
	assert.Equal(t, 2, result.Objects[0].ActiveNodes)

	client.AssertExpectations(t)
}
