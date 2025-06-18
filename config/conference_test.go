package config

import (
	"errors"
	"testing"

	mockClient "github.com/pexip/go-infinity-sdk/v38/internal/mock"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/conference/123/",
	}
	client.On("PostWithResponse", t.Context(), "configuration/v1/conference/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateConference(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
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

func TestService_ListConferences_QueryParameterValidation(t *testing.T) {
	tests := []struct {
		name          string
		opts          *ListOptions
		expectedQuery string
	}{
		{
			name:          "nil options",
			opts:          nil,
			expectedQuery: "configuration/v1/conference/",
		},
		{
			name: "empty search string",
			opts: &ListOptions{
				BaseListOptions: options.BaseListOptions{Limit: 10},
				Search:          "",
			},
			expectedQuery: "configuration/v1/conference/?limit=10",
		},
		{
			name: "search with special characters",
			opts: &ListOptions{
				BaseListOptions: options.BaseListOptions{Limit: 5},
				Search:          "test@domain.com",
			},
			expectedQuery: "configuration/v1/conference/?limit=5&name__icontains=test%40domain.com",
		},
		{
			name: "unicode search",
			opts: &ListOptions{
				BaseListOptions: options.BaseListOptions{Limit: 5},
				Search:          "cönférence",
			},
			expectedQuery: "configuration/v1/conference/?limit=5&name__icontains=c%C3%B6nf%C3%A9rence",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &mockClient.Client{}
			expectedResponse := &ConferenceListResponse{Objects: []Conference{}}

			client.On("GetJSON", t.Context(), tt.expectedQuery, mock.AnythingOfType("*config.ConferenceListResponse")).Return(nil).Run(func(args mock.Arguments) {
				result := args.Get(2).(*ConferenceListResponse)
				*result = *expectedResponse
			})

			service := New(client)
			_, err := service.ListConferences(t.Context(), tt.opts)

			assert.NoError(t, err)
			client.AssertExpectations(t)
		})
	}
}

func TestService_ConferenceErrorHandling(t *testing.T) {
	tests := []struct {
		name      string
		setup     func(m *mockClient.Client)
		operation func(service *Service) error
	}{
		{
			name: "ListConferences client error",
			setup: func(m *mockClient.Client) {
				m.On("GetJSON", t.Context(), "configuration/v1/conference/", mock.AnythingOfType("*config.ConferenceListResponse")).Return(errors.New("network error"))
			},
			operation: func(service *Service) error {
				_, err := service.ListConferences(t.Context(), nil)
				return err
			},
		},
		{
			name: "GetConference client error",
			setup: func(m *mockClient.Client) {
				m.On("GetJSON", t.Context(), "configuration/v1/conference/1/", mock.AnythingOfType("*config.Conference")).Return(errors.New("not found"))
			},
			operation: func(service *Service) error {
				_, err := service.GetConference(t.Context(), 1)
				return err
			},
		},
		{
			name: "CreateConference client error",
			setup: func(m *mockClient.Client) {
				m.On("PostWithResponse", t.Context(), "configuration/v1/conference/", mock.Anything, nil).Return(nil, errors.New("validation error"))
			},
			operation: func(service *Service) error {
				_, err := service.CreateConference(t.Context(), &ConferenceCreateRequest{Name: "Test"})
				return err
			},
		},
		{
			name: "UpdateConference client error",
			setup: func(m *mockClient.Client) {
				m.On("PutJSON", t.Context(), "configuration/v1/conference/1/", mock.Anything, mock.AnythingOfType("*config.Conference")).Return(errors.New("update failed"))
			},
			operation: func(service *Service) error {
				_, err := service.UpdateConference(t.Context(), 1, &ConferenceUpdateRequest{Name: "Updated"})
				return err
			},
		},
		{
			name: "DeleteConference client error",
			setup: func(m *mockClient.Client) {
				m.On("DeleteJSON", t.Context(), "configuration/v1/conference/1/", mock.Anything).Return(errors.New("delete failed"))
			},
			operation: func(service *Service) error {
				return service.DeleteConference(t.Context(), 1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &mockClient.Client{}
			tt.setup(client)

			service := New(client)
			err := tt.operation(service)

			assert.Error(t, err)
			client.AssertExpectations(t)
		})
	}
}
