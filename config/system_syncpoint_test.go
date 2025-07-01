package config

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/pexip/go-infinity-sdk/v38/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_CreateSystemSyncpoint(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &SystemSyncpointCreateRequest{}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/system_syncpoint/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/system_syncpoint/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateSystemSyncpoint(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_GetSystemSyncpoint(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	creationTime := util.InfinityTime{}

	expectedSystemSyncpoint := &SystemSyncpoint{
		ID:           1,
		CreationTime: creationTime,
		ResourceURI:  "/api/admin/configuration/v1/system_syncpoint/1/",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/system_syncpoint/1/", mock.AnythingOfType("*config.SystemSyncpoint")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*SystemSyncpoint)
		*result = *expectedSystemSyncpoint
	})

	service := New(client)
	result, err := service.GetSystemSyncpoint(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedSystemSyncpoint, result)
	client.AssertExpectations(t)
}
