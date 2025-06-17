package status

import (
	"testing"
	"time"

	mockClient "github.com/pexip/go-infinity-sdk/v38/internal/mock"
	"github.com/pexip/go-infinity-sdk/v38/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListMJXEndpoints(t *testing.T) {
	client := &mockClient.Client{}

	lastContact := time.Now().Add(-15 * time.Minute)

	expectedResponse := &MJXEndpointListResponse{
		Objects: []MJXEndpoint{
			{
				ID:                1,
				Name:              "mjx-endpoint-1",
				Status:            "connected",
				EndpointType:      "teams",
				LastContact:       &util.InfinityTime{Time: lastContact},
				Version:           "1.5.0",
				ActiveConnections: 3,
				ResourceURI:       "/api/admin/status/v1/mjx_endpoint/1/",
			},
			{
				ID:                2,
				Name:              "mjx-endpoint-2",
				Status:            "disconnected",
				EndpointType:      "google",
				LastContact:       &util.InfinityTime{Time: lastContact},
				Version:           "1.4.2",
				ActiveConnections: 0,
				ResourceURI:       "/api/admin/status/v1/mjx_endpoint/2/",
			},
		},
	}

	client.On("GetJSON", t.Context(), "status/v1/mjx_endpoint/", mock.AnythingOfType("*status.MJXEndpointListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*MJXEndpointListResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ListMJXEndpoints(t.Context(), nil)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result.Objects))
	assert.Equal(t, "mjx-endpoint-1", result.Objects[0].Name)
	assert.Equal(t, "connected", result.Objects[0].Status)
	assert.Equal(t, "teams", result.Objects[0].EndpointType)
	assert.Equal(t, 3, result.Objects[0].ActiveConnections)
	assert.Equal(t, "mjx-endpoint-2", result.Objects[1].Name)
	assert.Equal(t, "disconnected", result.Objects[1].Status)
	assert.Equal(t, "google", result.Objects[1].EndpointType)
	assert.Equal(t, 0, result.Objects[1].ActiveConnections)
	client.AssertExpectations(t)
}

func TestService_GetMJXEndpoint(t *testing.T) {
	client := &mockClient.Client{}
	lastContact := time.Now().Add(-2 * time.Minute)

	expectedEndpoint := &MJXEndpoint{
		ID:                1,
		Name:              "mjx-primary-endpoint",
		Status:            "connected",
		EndpointType:      "teams",
		LastContact:       &util.InfinityTime{Time: lastContact},
		Version:           "1.6.0",
		ActiveConnections: 8,
		ResourceURI:       "/api/admin/status/v1/mjx_endpoint/1/",
	}

	client.On("GetJSON", t.Context(), "status/v1/mjx_endpoint/1/", mock.AnythingOfType("*status.MJXEndpoint")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*MJXEndpoint)
		*result = *expectedEndpoint
	})

	service := New(client)
	result, err := service.GetMJXEndpoint(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedEndpoint, result)
	assert.Equal(t, "mjx-primary-endpoint", result.Name)
	assert.Equal(t, "connected", result.Status)
	assert.Equal(t, "teams", result.EndpointType)
	assert.Equal(t, "1.6.0", result.Version)
	assert.Equal(t, 8, result.ActiveConnections)
	client.AssertExpectations(t)
}

func TestService_ListMJXEndpoints_WithOptions(t *testing.T) {
	client := &mockClient.Client{}
	service := New(client)

	opts := &ListOptions{
		Limit:  10,
		Offset: 0,
	}

	expectedResponse := &MJXEndpointListResponse{
		Objects: []MJXEndpoint{
			{
				ID:                3,
				Name:              "mjx-test-endpoint",
				Status:            "maintenance",
				EndpointType:      "webex",
				Version:           "1.5.1",
				ActiveConnections: 0,
			},
		},
	}

	client.On("GetJSON", t.Context(), mock.MatchedBy(func(endpoint string) bool {
		return endpoint != "status/v1/mjx_endpoint/"
	}), mock.AnythingOfType("*status.MJXEndpointListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*MJXEndpointListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListMJXEndpoints(t.Context(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "mjx-test-endpoint", result.Objects[0].Name)
	assert.Equal(t, "maintenance", result.Objects[0].Status)
	assert.Equal(t, "webex", result.Objects[0].EndpointType)

	client.AssertExpectations(t)
}
