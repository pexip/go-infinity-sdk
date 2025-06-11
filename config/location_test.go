package config

import (
	"testing"

	mockClient "github.com/pexip/go-infinity-sdk/internal/mock"
	"github.com/pexip/go-infinity-sdk/options"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListLocations(t *testing.T) {
	client := &mockClient.Client{}

	expectedResponse := &LocationListResponse{
		Objects: []Location{
			{ID: 1, Name: "Location 1"},
			{ID: 2, Name: "Location 2"},
		},
	}

	client.On("GetJSON", t.Context(), "configuration/v1/location/", mock.AnythingOfType("*config.LocationListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*LocationListResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ListLocations(t.Context(), nil)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result.Objects))
	client.AssertExpectations(t)
}

func TestService_GetLocation(t *testing.T) {
	client := &mockClient.Client{}
	expectedLocation := &Location{
		ID:   1,
		Name: "Test Location",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/location/1/", mock.AnythingOfType("*config.Location")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*Location)
		*result = *expectedLocation
	})

	service := New(client)
	result, err := service.GetLocation(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedLocation, result)
	client.AssertExpectations(t)
}

func TestService_CreateLocation(t *testing.T) {
	client := &mockClient.Client{}

	createRequest := &LocationCreateRequest{
		Name:        "New Location",
		Description: "Test location",
	}

	expectedLocation := &Location{
		ID:          1,
		Name:        "New Location",
		Description: "Test location",
	}

	client.On("PostJSON", t.Context(), "configuration/v1/location/", createRequest, mock.AnythingOfType("*config.Location")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*Location)
		*result = *expectedLocation
	})

	service := New(client)
	result, err := service.CreateLocation(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedLocation, result)
	client.AssertExpectations(t)
}

func TestService_UpdateLocation(t *testing.T) {
	client := &mockClient.Client{}

	updateRequest := &LocationUpdateRequest{
		Name: "Updated Location",
	}

	expectedLocation := &Location{
		ID:   1,
		Name: "Updated Location",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/location/1/", updateRequest, mock.AnythingOfType("*config.Location")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*Location)
		*result = *expectedLocation
	})

	service := New(client)
	result, err := service.UpdateLocation(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedLocation, result)
	client.AssertExpectations(t)
}

func TestService_DeleteLocation(t *testing.T) {
	client := &mockClient.Client{}

	client.On("DeleteJSON", t.Context(), "configuration/v1/location/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteLocation(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}

func TestService_ListLocations_WithOptions(t *testing.T) {
	client := &mockClient.Client{}

	opts := &ListOptions{
		BaseListOptions: options.BaseListOptions{
			Limit:  10,
			Offset: 5,
		},
		Search: "test",
	}

	expectedResponse := &LocationListResponse{
		Objects: []Location{
			{ID: 1, Name: "Test Location"},
		},
	}

	client.On("GetJSON", t.Context(), "configuration/v1/location/?limit=10&name__icontains=test&offset=5", mock.AnythingOfType("*config.LocationListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*LocationListResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ListLocations(t.Context(), opts)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	client.AssertExpectations(t)
}