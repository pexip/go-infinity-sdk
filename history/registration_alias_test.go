package history

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/pexip/go-infinity-sdk/v38/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListRegistrationAliases(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	expectedResponse := &RegistrationAliasListResponse{
		Meta: struct {
			Limit      int    `json:"limit"`
			Next       string `json:"next"`
			Offset     int    `json:"offset"`
			Previous   string `json:"previous"`
			TotalCount int    `json:"total_count"`
		}{
			Limit:      20,
			TotalCount: 1,
		},
		Objects: []RegistrationAlias{
			{
				ID:             1,
				Alias:          "john.doe@example.com",
				RegistrationID: "reg-123",
				Username:       "john.doe",
				Node:           "node1.example.com",
				Protocol:       "SIP",
				RemoteAddress:  "192.168.1.100",
				IsNatted:       false,
				StartTime:      &util.InfinityTime{},
				EndTime:        &util.InfinityTime{},
				ResourceURI:    "/api/admin/history/v1/registration_alias/1/",
			},
		},
	}

	client.On("GetJSON", context.Background(), "history/v1/registration_alias/", mock.AnythingOfType("*history.RegistrationAliasListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*RegistrationAliasListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListRegistrationAliases(context.Background(), nil)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "john.doe@example.com", result.Objects[0].Alias)
	assert.Equal(t, "SIP", result.Objects[0].Protocol)
	assert.Equal(t, false, result.Objects[0].IsNatted)

	client.AssertExpectations(t)
}

func TestService_GetRegistrationAlias(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	expectedAlias := &RegistrationAlias{
		ID:             1,
		Alias:          "jane.smith@example.com",
		RegistrationID: "reg-456",
		Username:       "jane.smith",
		Node:           "node2.example.com",
		Protocol:       "H323",
		RemoteAddress:  "10.0.0.50",
		IsNatted:       true,
		StartTime:      &util.InfinityTime{},
		EndTime:        &util.InfinityTime{},
		ResourceURI:    "/api/admin/history/v1/registration_alias/1/",
	}

	client.On("GetJSON", context.Background(), "history/v1/registration_alias/1/", mock.AnythingOfType("*history.RegistrationAlias")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*RegistrationAlias)
		*result = *expectedAlias
	})

	result, err := service.GetRegistrationAlias(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, expectedAlias, result)

	client.AssertExpectations(t)
}

func TestService_ListRegistrationAliases_WithOptions(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	startTime := time.Now().Add(-6 * time.Hour)
	endTime := time.Now()
	opts := &ListOptions{
		SearchableListOptions: options.SearchableListOptions{
			BaseListOptions: options.BaseListOptions{
				Limit:  50,
				Offset: 20,
			},
			Search: "example.com",
		},
		StartTime: &startTime,
		EndTime:   &endTime,
	}

	expectedResponse := &RegistrationAliasListResponse{
		Meta: struct {
			Limit      int    `json:"limit"`
			Next       string `json:"next"`
			Offset     int    `json:"offset"`
			Previous   string `json:"previous"`
			TotalCount int    `json:"total_count"`
		}{
			Limit:      50,
			Offset:     20,
			TotalCount: 5,
		},
		Objects: []RegistrationAlias{
			{
				ID:       1,
				Alias:    "test@example.com",
				Protocol: "SIP",
			},
		},
	}

	client.On("GetJSON", context.Background(), mock.MatchedBy(func(endpoint string) bool {
		return endpoint != "history/v1/registration_alias/"
	}), mock.AnythingOfType("*history.RegistrationAliasListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*RegistrationAliasListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListRegistrationAliases(context.Background(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 50, result.Meta.Limit)
	assert.Equal(t, 20, result.Meta.Offset)
	assert.Equal(t, 1, len(result.Objects))

	client.AssertExpectations(t)
}

func TestService_ListRegistrationAliases_Error(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	client.On("GetJSON", context.Background(), "history/v1/registration_alias/", mock.AnythingOfType("*history.RegistrationAliasListResponse")).Return(errors.New("server error"))

	_, err := service.ListRegistrationAliases(context.Background(), nil)
	assert.Error(t, err)
	assert.Equal(t, "server error", err.Error())

	client.AssertExpectations(t)
}

func TestService_GetRegistrationAlias_NotFound(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	client.On("GetJSON", context.Background(), "history/v1/registration_alias/999/", mock.AnythingOfType("*history.RegistrationAlias")).Return(errors.New("registration alias not found"))

	_, err := service.GetRegistrationAlias(context.Background(), 999)
	assert.Error(t, err)

	client.AssertExpectations(t)
}
