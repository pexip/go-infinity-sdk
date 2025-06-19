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
				ID:             1,
				Alias:          "conference.room@domain.com",
				IsNatted:       false,
				Node:           "node-01",
				Protocol:       "SIP",
				PushToken:      "",
				RegistrationID: "",
				RemoteAddress:  "",
				ResourceURI:    "/api/admin/status/v1/registration_alias/1/",
				StartTime:      nil,
				Username:       "",
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
	assert.Equal(t, 1, result.Objects[0].ID)
	assert.Equal(t, "conference.room@domain.com", result.Objects[0].Alias)
	assert.Equal(t, false, result.Objects[0].IsNatted)
	assert.Equal(t, "node-01", result.Objects[0].Node)
	assert.Equal(t, "SIP", result.Objects[0].Protocol)
	assert.Equal(t, "/api/admin/status/v1/registration_alias/1/", result.Objects[0].ResourceURI)
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
				ID:             2,
				Alias:          "options.test@example.com",
				IsNatted:       true,
				Node:           "node-03",
				Protocol:       "SIP",
				PushToken:      "",
				RegistrationID: "",
				RemoteAddress:  "",
				ResourceURI:    "",
				StartTime:      nil,
				Username:       "",
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
	assert.Equal(t, 2, result.Objects[0].ID)
	assert.Equal(t, "options.test@example.com", result.Objects[0].Alias)
	assert.Equal(t, true, result.Objects[0].IsNatted)
	assert.Equal(t, "node-03", result.Objects[0].Node)
	assert.Equal(t, "SIP", result.Objects[0].Protocol)

	client.AssertExpectations(t)
}

func TestService_GetRegistrationAlias(t *testing.T) {
	client := &mockClient.Client{}

	expectedAlias := &RegistrationAlias{
		ID:             1,
		Alias:          "test.room@example.com",
		IsNatted:       false,
		Node:           "node-02",
		Protocol:       "H.323",
		PushToken:      "",
		RegistrationID: "",
		RemoteAddress:  "",
		ResourceURI:    "/api/admin/status/v1/registration_alias/1/",
		StartTime:      nil,
		Username:       "",
	}

	client.On("GetJSON", t.Context(), "status/v1/registration_alias/1/", mock.AnythingOfType("*status.RegistrationAlias")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*RegistrationAlias)
		*result = *expectedAlias
	})

	service := New(client)
	result, err := service.GetRegistrationAlias(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedAlias, result)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "test.room@example.com", result.Alias)
	assert.Equal(t, false, result.IsNatted)
	assert.Equal(t, "node-02", result.Node)
	assert.Equal(t, "H.323", result.Protocol)
	assert.Equal(t, "/api/admin/status/v1/registration_alias/1/", result.ResourceURI)
	client.AssertExpectations(t)
}
