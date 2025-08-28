/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/stretchr/testify/assert"
)

func TestService_CreateUpgrade(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	packageName := "upgrade-package-v28.0.0.tar.gz"
	createRequest := &UpgradeCreateRequest{
		Package: &packageName,
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/upgrade/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/upgrade/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateUpgrade(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}
