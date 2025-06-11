package status

import (
	"testing"

	mockClient "github.com/pexip/go-infinity-sdk/v38/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListRegistrationAliases(t *testing.T) {
	client := &mockClient.Client{}

	expectedResponse := &RegistrationAliasListResponse{
		Objects: []RegistrationAlias{
			{
				ID:          1,
				Alias:       "conference.room@domain.com",
				Status:      "registered",
				NodeID:      "node-01",
				Protocol:    "SIP",
				ResourceURI: "/api/admin/status/v1/registration_alias/1/",
			},
		},
	}

	client.On("GetJSON", t.Context(), "status/v1/registration_alias/", mock.AnythingOfType("*status.RegistrationAliasListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*RegistrationAliasListResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ListRegistrationAliases(t.Context(), nil)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "conference.room@domain.com", result.Objects[0].Alias)
	assert.Equal(t, "registered", result.Objects[0].Status)
	client.AssertExpectations(t)
}

func TestService_ListRegistrationAliases_WithOptions(t *testing.T) {
	client := &mockClient.Client{}
	service := New(client)

	opts := &ListOptions{
		Limit:  8,
		Offset: 2,
	}

	expectedResponse := &RegistrationAliasListResponse{
		Objects: []RegistrationAlias{
			{
				ID:       2,
				Alias:    "options.test@example.com",
				Status:   "registered",
				NodeID:   "node-03",
				Protocol: "SIP",
			},
		},
	}

	client.On("GetJSON", t.Context(), mock.MatchedBy(func(endpoint string) bool {
		return endpoint != "status/v1/registration_alias/"
	}), mock.AnythingOfType("*status.RegistrationAliasListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*RegistrationAliasListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListRegistrationAliases(t.Context(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "options.test@example.com", result.Objects[0].Alias)

	client.AssertExpectations(t)
}

func TestService_GetRegistrationAlias(t *testing.T) {
	client := &mockClient.Client{}

	expectedAlias := &RegistrationAlias{
		ID:          1,
		Alias:       "test.room@example.com",
		Status:      "unregistered",
		NodeID:      "node-02",
		Protocol:    "H.323",
		ResourceURI: "/api/admin/status/v1/registration_alias/1/",
	}

	client.On("GetJSON", t.Context(), "status/v1/registration_alias/1/", mock.AnythingOfType("*status.RegistrationAlias")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*RegistrationAlias)
		*result = *expectedAlias
	})

	service := New(client)
	result, err := service.GetRegistrationAlias(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedAlias, result)
	client.AssertExpectations(t)
}
