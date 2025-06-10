package config

import (
	"testing"

	mockClient "github.com/pexip/go-infinity-sdk/internal/mock"
	"github.com/pexip/go-infinity-sdk/options"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

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
					Objects: []Conference{
						{ID: 1, Name: "Test Conference 1"},
						{ID: 2, Name: "Test Conference 2"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/conference/", mock.AnythingOfType("*config.ConferenceListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*ConferenceListResponse)
					*result = *expectedResponse
				})
			},
			wantErr: false,
		},
		{
			name: "successful list with options",
			opts: &ListOptions{
				BaseListOptions: options.BaseListOptions{
					Limit:  10,
					Offset: 5,
				},
				Search: "test",
			},
			setup: func(m *mockClient.Client) {
				expectedResponse := &ConferenceListResponse{
					Objects: []Conference{
						{ID: 1, Name: "Test Conference"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/conference/?limit=10&name__icontains=test&offset=5", mock.AnythingOfType("*config.ConferenceListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*ConferenceListResponse)
					*result = *expectedResponse
				})
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &mockClient.Client{}
			tt.setup(client)

			service := New(client)
			result, err := service.ListConferences(t.Context(), tt.opts)

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

func TestService_GetConference(t *testing.T) {
	client := &mockClient.Client{}
	expectedConference := &Conference{
		ID:   1,
		Name: "Test Conference",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/conference/1/", mock.AnythingOfType("*config.Conference")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*Conference)
		*result = *expectedConference
	})

	service := New(client)
	result, err := service.GetConference(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedConference, result)
	client.AssertExpectations(t)
}

func TestService_CreateConference(t *testing.T) {
	client := &mockClient.Client{}

	createRequest := &ConferenceCreateRequest{
		Name:        "New Conference",
		ServiceType: "conference",
		AllowGuests: true,
	}

	expectedConference := &Conference{
		ID:          1,
		Name:        "New Conference",
		ServiceType: "conference",
		AllowGuests: true,
	}

	client.On("PostJSON", t.Context(), "configuration/v1/conference/", createRequest, mock.AnythingOfType("*config.Conference")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*Conference)
		*result = *expectedConference
	})

	service := New(client)
	result, err := service.CreateConference(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedConference, result)
	client.AssertExpectations(t)
}

func TestService_UpdateConference(t *testing.T) {
	client := &mockClient.Client{}

	updateRequest := &ConferenceUpdateRequest{
		Name: "Updated Conference",
	}

	expectedConference := &Conference{
		ID:   1,
		Name: "Updated Conference",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/conference/1/", updateRequest, mock.AnythingOfType("*config.Conference")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*Conference)
		*result = *expectedConference
	})

	service := New(client)
	result, err := service.UpdateConference(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedConference, result)
	client.AssertExpectations(t)
}

func TestService_DeleteConference(t *testing.T) {
	client := &mockClient.Client{}

	client.On("DeleteJSON", t.Context(), "configuration/v1/conference/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteConference(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}

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

func TestNew(t *testing.T) {
	client := &mockClient.Client{}
	service := New(client)

	require.NotNil(t, service)
	assert.Equal(t, client, service.client)
}
