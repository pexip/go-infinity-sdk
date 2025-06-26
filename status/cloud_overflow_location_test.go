package status

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListCloudOverflowLocations(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	expectedResponse := &CloudOverflowLocationListResponse{
		Meta: Meta{
			Limit:      100,
			Next:       "",
			Offset:     0,
			Previous:   "",
			TotalCount: 1,
		},
		Objects: []CloudOverflowLocation{
			{
				ID:               1,
				Name:             "Overflow US East",
				FreeHDCalls:      5,
				MaxHDCalls:       10,
				MediaLoad:        20,
				ResourceURI:      "/api/admin/status/v1/cloud_overflow_location/1/",
				SystemLocationID: 101,
			},
		},
	}

	client.On("GetJSON", t.Context(), "status/v1/cloud_overflow_location/", mock.AnythingOfType("*status.CloudOverflowLocationListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*CloudOverflowLocationListResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ListCloudOverflowLocations(t.Context(), nil)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, 1, result.Objects[0].ID)
	assert.Equal(t, "Overflow US East", result.Objects[0].Name)
	assert.Equal(t, 5, result.Objects[0].FreeHDCalls)
	assert.Equal(t, 10, result.Objects[0].MaxHDCalls)
	assert.Equal(t, 20, result.Objects[0].MediaLoad)
	assert.Equal(t, 101, result.Objects[0].SystemLocationID)
	assert.Equal(t, "/api/admin/status/v1/cloud_overflow_location/1/", result.Objects[0].ResourceURI)
	client.AssertExpectations(t)
}

func TestService_ListCloudOverflowLocations_WithOptions(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	opts := &ListOptions{
		Limit:  3,
		Offset: 1,
	}

	expectedResponse := &CloudOverflowLocationListResponse{
		Meta: Meta{
			Limit:      3,
			Next:       "",
			Offset:     1,
			Previous:   "",
			TotalCount: 1,
		},
		Objects: []CloudOverflowLocation{
			{
				ID:               2,
				Name:             "Test Overflow",
				FreeHDCalls:      3,
				MaxHDCalls:       7,
				MediaLoad:        15,
				ResourceURI:      "/api/admin/status/v1/cloud_overflow_location/2/",
				SystemLocationID: 102,
			},
		},
	}

	client.On("GetJSON", t.Context(), mock.MatchedBy(func(endpoint string) bool {
		return endpoint != "status/v1/cloud_overflow_location/"
	}), mock.AnythingOfType("*status.CloudOverflowLocationListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*CloudOverflowLocationListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListCloudOverflowLocations(t.Context(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, 2, result.Objects[0].ID)
	assert.Equal(t, "Test Overflow", result.Objects[0].Name)
	assert.Equal(t, 3, result.Objects[0].FreeHDCalls)
	assert.Equal(t, 7, result.Objects[0].MaxHDCalls)
	assert.Equal(t, 15, result.Objects[0].MediaLoad)
	assert.Equal(t, 102, result.Objects[0].SystemLocationID)
	assert.Equal(t, "/api/admin/status/v1/cloud_overflow_location/2/", result.Objects[0].ResourceURI)

	client.AssertExpectations(t)
}

func TestService_GetCloudOverflowLocation(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	expectedLocation := &CloudOverflowLocation{
		ID:               1,
		Name:             "Primary Overflow",
		FreeHDCalls:      6,
		MaxHDCalls:       12,
		MediaLoad:        25,
		ResourceURI:      "/api/admin/status/v1/cloud_overflow_location/1/",
		SystemLocationID: 103,
	}

	client.On("GetJSON", t.Context(), "status/v1/cloud_overflow_location/1/", mock.AnythingOfType("*status.CloudOverflowLocation")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*CloudOverflowLocation)
		*result = *expectedLocation
	})

	service := New(client)
	result, err := service.GetCloudOverflowLocation(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedLocation, result)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Primary Overflow", result.Name)
	assert.Equal(t, 6, result.FreeHDCalls)
	assert.Equal(t, 12, result.MaxHDCalls)
	assert.Equal(t, 25, result.MediaLoad)
	assert.Equal(t, 103, result.SystemLocationID)
	assert.Equal(t, "/api/admin/status/v1/cloud_overflow_location/1/", result.ResourceURI)
	client.AssertExpectations(t)
}
