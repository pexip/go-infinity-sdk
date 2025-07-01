package config

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListMjxEndpoints(t *testing.T) {
	tests := []struct {
		name    string
		opts    *ListOptions
		setup   func(m *interfaces.HTTPClientMock)
		wantErr bool
	}{
		{
			name: "successful list without options",
			opts: nil,
			setup: func(m *interfaces.HTTPClientMock) {
				group1 := "/api/admin/configuration/v1/mjx_endpoint_group/1/"
				apiAddress1 := "192.168.1.100"
				apiPort1 := 443
				apiUsername1 := "admin"
				apiPassword1 := "password123"
				polyUsername1 := "polyuser"
				polyPassword1 := "polypass"
				webexDeviceID1 := "webex-device-123"

				group2 := "/api/admin/configuration/v1/mjx_endpoint_group/2/"
				apiAddress2 := "192.168.1.101"

				expectedResponse := &MjxEndpointListResponse{
					Objects: []MjxEndpoint{
						{ID: 1, Name: "conf-room-01", Description: "Conference Room 01", EndpointType: "cisco", RoomResourceEmail: "room01@example.com", MjxEndpointGroup: &group1, APIAddress: &apiAddress1, APIPort: &apiPort1, APIUsername: &apiUsername1, APIPassword: &apiPassword1, UseHTTPS: "yes", VerifyCert: "yes", PolyUsername: &polyUsername1, PolyPassword: &polyPassword1, PolyRaiseAlarmsForThisEndpoint: true, WebexDeviceID: &webexDeviceID1},
						{ID: 2, Name: "conf-room-02", Description: "Conference Room 02", EndpointType: "poly", RoomResourceEmail: "room02@example.com", MjxEndpointGroup: &group2, APIAddress: &apiAddress2, UseHTTPS: "no", VerifyCert: "no", PolyRaiseAlarmsForThisEndpoint: false},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/mjx_endpoint/", mock.AnythingOfType("*config.MjxEndpointListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*MjxEndpointListResponse)
					*result = *expectedResponse
				})
			},
			wantErr: false,
		},
		{
			name: "successful list with options",
			opts: &ListOptions{
				BaseListOptions: options.BaseListOptions{
					Limit: 5,
				},
				Search: "conf-room-01",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				group := "/api/admin/configuration/v1/mjx_endpoint_group/1/"
				apiAddress := "192.168.1.100"
				apiPort := 443
				apiUsername := "admin"
				apiPassword := "password123"
				webexDeviceID := "webex-device-123"

				expectedResponse := &MjxEndpointListResponse{
					Objects: []MjxEndpoint{
						{ID: 1, Name: "conf-room-01", Description: "Conference Room 01", EndpointType: "cisco", RoomResourceEmail: "room01@example.com", MjxEndpointGroup: &group, APIAddress: &apiAddress, APIPort: &apiPort, APIUsername: &apiUsername, APIPassword: &apiPassword, UseHTTPS: "yes", VerifyCert: "yes", PolyRaiseAlarmsForThisEndpoint: true, WebexDeviceID: &webexDeviceID},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/mjx_endpoint/?limit=5&name__icontains=conf-room-01", mock.AnythingOfType("*config.MjxEndpointListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*MjxEndpointListResponse)
					*result = *expectedResponse
				})
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := interfaces.NewHTTPClientMock()
			tt.setup(client)

			service := New(client)
			result, err := service.ListMjxEndpoints(t.Context(), tt.opts)

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

func TestService_GetMjxEndpoint(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	group := "/api/admin/configuration/v1/mjx_endpoint_group/1/"
	apiAddress := "192.168.1.100"
	apiPort := 443
	apiUsername := "admin"
	apiPassword := "secret123"
	polyUsername := "polyuser"
	polyPassword := "polypass"
	webexDeviceID := "webex-device-abc123"

	expectedMjxEndpoint := &MjxEndpoint{
		ID:                             1,
		Name:                           "test-endpoint",
		Description:                    "Test MJX endpoint",
		EndpointType:                   "cisco",
		RoomResourceEmail:              "testroom@example.com",
		MjxEndpointGroup:               &group,
		APIAddress:                     &apiAddress,
		APIPort:                        &apiPort,
		APIUsername:                    &apiUsername,
		APIPassword:                    &apiPassword,
		UseHTTPS:                       "yes",
		VerifyCert:                     "yes",
		PolyUsername:                   &polyUsername,
		PolyPassword:                   &polyPassword,
		PolyRaiseAlarmsForThisEndpoint: true,
		WebexDeviceID:                  &webexDeviceID,
	}

	client.On("GetJSON", t.Context(), "configuration/v1/mjx_endpoint/1/", mock.AnythingOfType("*config.MjxEndpoint")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*MjxEndpoint)
		*result = *expectedMjxEndpoint
	})

	service := New(client)
	result, err := service.GetMjxEndpoint(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedMjxEndpoint, result)
	client.AssertExpectations(t)
}

func TestService_CreateMjxEndpoint(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	group := "/api/admin/configuration/v1/mjx_endpoint_group/1/"
	apiAddress := "192.168.1.102"
	apiPort := 443
	apiUsername := "newadmin"
	apiPassword := "newpassword"
	webexDeviceID := "webex-new-device"

	createRequest := &MjxEndpointCreateRequest{
		Name:                           "new-endpoint",
		Description:                    "New MJX endpoint",
		EndpointType:                   "webex",
		RoomResourceEmail:              "newroom@example.com",
		MjxEndpointGroup:               &group,
		APIAddress:                     &apiAddress,
		APIPort:                        &apiPort,
		APIUsername:                    &apiUsername,
		APIPassword:                    &apiPassword,
		UseHTTPS:                       "yes",
		VerifyCert:                     "no",
		PolyRaiseAlarmsForThisEndpoint: false,
		WebexDeviceID:                  &webexDeviceID,
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/mjx_endpoint/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/mjx_endpoint/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateMjxEndpoint(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateMjxEndpoint(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	raiseAlarms := false
	updateRequest := &MjxEndpointUpdateRequest{
		Description:                    "Updated MJX endpoint",
		UseHTTPS:                       "no",
		PolyRaiseAlarmsForThisEndpoint: &raiseAlarms,
	}

	group := "/api/admin/configuration/v1/mjx_endpoint_group/1/"
	apiAddress := "192.168.1.100"
	apiPort := 443
	apiUsername := "admin"
	apiPassword := "secret123"
	webexDeviceID := "webex-device-abc123"

	expectedMjxEndpoint := &MjxEndpoint{
		ID:                             1,
		Name:                           "test-endpoint",
		Description:                    "Updated MJX endpoint",
		EndpointType:                   "cisco",
		RoomResourceEmail:              "testroom@example.com",
		MjxEndpointGroup:               &group,
		APIAddress:                     &apiAddress,
		APIPort:                        &apiPort,
		APIUsername:                    &apiUsername,
		APIPassword:                    &apiPassword,
		UseHTTPS:                       "no",
		VerifyCert:                     "yes",
		PolyRaiseAlarmsForThisEndpoint: false,
		WebexDeviceID:                  &webexDeviceID,
	}

	client.On("PutJSON", t.Context(), "configuration/v1/mjx_endpoint/1/", updateRequest, mock.AnythingOfType("*config.MjxEndpoint")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*MjxEndpoint)
		*result = *expectedMjxEndpoint
	})

	service := New(client)
	result, err := service.UpdateMjxEndpoint(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedMjxEndpoint, result)
	client.AssertExpectations(t)
}

func TestService_DeleteMjxEndpoint(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/mjx_endpoint/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteMjxEndpoint(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
