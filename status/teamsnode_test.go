package status

import (
	"testing"
	"time"

	mockClient "github.com/pexip/go-infinity-sdk/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListTeamsNodes(t *testing.T) {
	client := &mockClient.Client{}

	lastContact := time.Now().Add(-10 * time.Minute)

	expectedResponse := &TeamsNodeListResponse{
		Objects: []TeamsNode{
			{
				ID:          1,
				Name:        "teams-node-1",
				Status:      "active",
				Version:     "2.0.1",
				LastContact: &lastContact,
				ActiveCalls: 5,
				ResourceURI: "/api/admin/status/v1/teamsnode/1/",
			},
			{
				ID:          2,
				Name:        "teams-node-2",
				Status:      "inactive",
				Version:     "2.0.0",
				LastContact: &lastContact,
				ActiveCalls: 0,
				ResourceURI: "/api/admin/status/v1/teamsnode/2/",
			},
		},
	}

	client.On("GetJSON", t.Context(), "status/v1/teamsnode/", mock.AnythingOfType("*status.TeamsNodeListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*TeamsNodeListResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ListTeamsNodes(t.Context(), nil)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result.Objects))
	assert.Equal(t, "teams-node-1", result.Objects[0].Name)
	assert.Equal(t, "active", result.Objects[0].Status)
	assert.Equal(t, "2.0.1", result.Objects[0].Version)
	assert.Equal(t, 5, result.Objects[0].ActiveCalls)
	assert.Equal(t, "teams-node-2", result.Objects[1].Name)
	assert.Equal(t, "inactive", result.Objects[1].Status)
	assert.Equal(t, 0, result.Objects[1].ActiveCalls)
	client.AssertExpectations(t)
}

func TestService_GetTeamsNode(t *testing.T) {
	client := &mockClient.Client{}
	lastContact := time.Now().Add(-5 * time.Minute)

	expectedNode := &TeamsNode{
		ID:          1,
		Name:        "teams-node-primary",
		Status:      "active",
		Version:     "2.1.0",
		LastContact: &lastContact,
		ActiveCalls: 12,
		ResourceURI: "/api/admin/status/v1/teamsnode/1/",
	}

	client.On("GetJSON", t.Context(), "status/v1/teamsnode/1/", mock.AnythingOfType("*status.TeamsNode")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*TeamsNode)
		*result = *expectedNode
	})

	service := New(client)
	result, err := service.GetTeamsNode(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedNode, result)
	assert.Equal(t, "teams-node-primary", result.Name)
	assert.Equal(t, "active", result.Status)
	assert.Equal(t, "2.1.0", result.Version)
	assert.Equal(t, 12, result.ActiveCalls)
	client.AssertExpectations(t)
}

func TestService_ListTeamsNodes_WithOptions(t *testing.T) {
	client := &mockClient.Client{}
	service := New(client)

	opts := &ListOptions{
		Limit:  3,
		Offset: 0,
	}

	expectedResponse := &TeamsNodeListResponse{
		Objects: []TeamsNode{
			{
				ID:          1,
				Name:        "teams-node-test",
				Status:      "active",
				Version:     "2.0.2",
				ActiveCalls: 8,
			},
		},
	}

	client.On("GetJSON", t.Context(), mock.MatchedBy(func(endpoint string) bool {
		return endpoint != "status/v1/teamsnode/"
	}), mock.AnythingOfType("*status.TeamsNodeListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*TeamsNodeListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListTeamsNodes(t.Context(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "teams-node-test", result.Objects[0].Name)
	assert.Equal(t, "active", result.Objects[0].Status)
	assert.Equal(t, 8, result.Objects[0].ActiveCalls)

	client.AssertExpectations(t)
}