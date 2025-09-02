/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_GetRegistration(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedRegistration := &Registration{
		ID:                         1,
		Enable:                     true,
		RefreshStrategy:            "adaptive",
		AdaptiveMinRefresh:         60,
		AdaptiveMaxRefresh:         3600,
		MaximumMinRefresh:          30,
		MaximumMaxRefresh:          7200,
		NattedMinRefresh:           120,
		NattedMaxRefresh:           600,
		RouteViaRegistrar:          true,
		EnablePushNotifications:    true,
		EnableGoogleCloudMessaging: false,
		PushToken:                  "push-token-123",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/registration/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.Registration")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*Registration)
		*result = *expectedRegistration
	})

	service := New(client)
	result, err := service.GetRegistration(t.Context())

	assert.NoError(t, err)
	assert.Equal(t, expectedRegistration, result)
	client.AssertExpectations(t)
}

func TestService_UpdateRegistration(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	enable := false
	adaptiveMinRefresh := 90
	enablePushNotifications := false

	updateRequest := &RegistrationUpdateRequest{
		Enable:                  &enable,
		RefreshStrategy:         "maximum",
		AdaptiveMinRefresh:      &adaptiveMinRefresh,
		EnablePushNotifications: &enablePushNotifications,
		PushToken:               "updated-push-token",
	}

	expectedRegistration := &Registration{
		ID:                         1,
		Enable:                     false,
		RefreshStrategy:            "maximum",
		AdaptiveMinRefresh:         90,
		AdaptiveMaxRefresh:         3600,
		MaximumMinRefresh:          30,
		MaximumMaxRefresh:          7200,
		NattedMinRefresh:           120,
		NattedMaxRefresh:           600,
		RouteViaRegistrar:          true,
		EnablePushNotifications:    false,
		EnableGoogleCloudMessaging: false,
		PushToken:                  "updated-push-token",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/registration/1/", updateRequest, mock.AnythingOfType("*config.Registration")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*Registration)
		*result = *expectedRegistration
	})

	service := New(client)
	result, err := service.UpdateRegistration(t.Context(), updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedRegistration, result)
	client.AssertExpectations(t)
}
