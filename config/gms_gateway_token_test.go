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

func TestService_GetGMSGatewayToken(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	intermediateCert := "-----BEGIN CERTIFICATE-----\nintermediate cert\n-----END CERTIFICATE-----"
	leafCert := "-----BEGIN CERTIFICATE-----\nleaf cert\n-----END CERTIFICATE-----"
	privateKey := "-----BEGIN PRIVATE KEY-----\nprivate key\n-----END PRIVATE KEY-----"
	supportsDirectGuestJoin := true

	expectedToken := &GMSGatewayToken{
		ID:                      1,
		Certificate:             "-----BEGIN CERTIFICATE-----\nmain cert\n-----END CERTIFICATE-----",
		IntermediateCertificate: &intermediateCert,
		LeafCertificate:         &leafCert,
		PrivateKey:              &privateKey,
		SupportsDirectGuestJoin: &supportsDirectGuestJoin,
		ResourceURI:             "/api/admin/configuration/v1/gms_gateway_token/1/",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/gms_gateway_token/1/", mock.AnythingOfType("*config.GMSGatewayToken")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*GMSGatewayToken)
		*result = *expectedToken
	})

	service := New(client)
	result, err := service.GetGMSGatewayToken(t.Context())

	assert.NoError(t, err)
	assert.Equal(t, expectedToken, result)
	client.AssertExpectations(t)
}

func TestService_UpdateGMSGatewayToken(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	newIntermediateCert := "-----BEGIN CERTIFICATE-----\nnew intermediate cert\n-----END CERTIFICATE-----"
	supportsDirectGuestJoin := false

	updateRequest := &GMSGatewayTokenUpdateRequest{
		Certificate:             "-----BEGIN CERTIFICATE-----\nupdated cert\n-----END CERTIFICATE-----",
		IntermediateCertificate: &newIntermediateCert,
		SupportsDirectGuestJoin: &supportsDirectGuestJoin,
	}

	expectedToken := &GMSGatewayToken{
		ID:                      1,
		Certificate:             "-----BEGIN CERTIFICATE-----\nupdated cert\n-----END CERTIFICATE-----",
		IntermediateCertificate: &newIntermediateCert,
		SupportsDirectGuestJoin: &supportsDirectGuestJoin,
		ResourceURI:             "/api/admin/configuration/v1/gms_gateway_token/1/",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/gms_gateway_token/1/", updateRequest, mock.AnythingOfType("*config.GMSGatewayToken")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*GMSGatewayToken)
		*result = *expectedToken
	})

	service := New(client)
	result, err := service.UpdateGMSGatewayToken(t.Context(), updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedToken, result)
	client.AssertExpectations(t)
}
