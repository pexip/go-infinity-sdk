/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/pexip/go-infinity-sdk/v38/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListTLSCertificates(t *testing.T) {
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
				startDate1 := util.InfinityTime{}
				endDate1 := util.InfinityTime{}
				keyID1 := "key123"

				expectedResponse := &TLSCertificateListResponse{
					Objects: []TLSCertificate{
						{ID: 1, Certificate: "-----BEGIN CERTIFICATE-----\nMIIB...", PrivateKey: "-----BEGIN PRIVATE KEY-----\nMIIE...", Nodes: []string{"management"}, StartDate: startDate1, EndDate: endDate1, SubjectName: "CN=example.com", KeyID: &keyID1},
						{ID: 2, Certificate: "-----BEGIN CERTIFICATE-----\nMIIC...", PrivateKey: "-----BEGIN PRIVATE KEY-----\nMIIF...", Nodes: []string{"conferencing"}, StartDate: startDate1, EndDate: endDate1, SubjectName: "CN=meeting.example.com"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/tls_certificate/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.TLSCertificateListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*TLSCertificateListResponse)
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
				Search: "example.com",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				startDate := util.InfinityTime{}
				endDate := util.InfinityTime{}
				keyID := "key123"

				expectedResponse := &TLSCertificateListResponse{
					Objects: []TLSCertificate{
						{ID: 1, Certificate: "-----BEGIN CERTIFICATE-----\nMIIB...", PrivateKey: "-----BEGIN PRIVATE KEY-----\nMIIE...", Nodes: []string{"management"}, StartDate: startDate, EndDate: endDate, SubjectName: "CN=example.com", KeyID: &keyID},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/tls_certificate/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.TLSCertificateListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*TLSCertificateListResponse)
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
			result, err := service.ListTLSCertificates(t.Context(), tt.opts)

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

func TestService_GetTLSCertificate(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	startDate := util.InfinityTime{}
	endDate := util.InfinityTime{}
	keyID := "key123"
	issuerKeyID := "issuerkey456"

	expectedTLSCertificate := &TLSCertificate{
		ID:                   1,
		Certificate:          "-----BEGIN CERTIFICATE-----\nMIIB...",
		PrivateKey:           "-----BEGIN PRIVATE KEY-----\nMIIE...",
		PrivateKeyPassphrase: "secret",
		Parameters:           "RSA-2048",
		Nodes:                []string{"management", "conferencing"},
		StartDate:            startDate,
		EndDate:              endDate,
		SubjectName:          "CN=example.com",
		SubjectHash:          "abc123",
		SubjectAltNames:      "DNS:example.com,DNS:*.example.com",
		IssuerName:           "CN=CA Authority",
		IssuerHash:           "def456",
		SerialNo:             "123456789",
		KeyID:                &keyID,
		IssuerKeyID:          &issuerKeyID,
		ResourceURI:          "/api/admin/configuration/v1/tls_certificate/1/",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/tls_certificate/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.TLSCertificate")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*TLSCertificate)
		*result = *expectedTLSCertificate
	})

	service := New(client)
	result, err := service.GetTLSCertificate(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedTLSCertificate, result)
	client.AssertExpectations(t)
}

func TestService_CreateTLSCertificate(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &TLSCertificateCreateRequest{
		Certificate:          "-----BEGIN CERTIFICATE-----\nMIIB...",
		PrivateKey:           "-----BEGIN PRIVATE KEY-----\nMIIE...",
		PrivateKeyPassphrase: "secret",
		Parameters:           "RSA-2048",
		Nodes:                []string{"management"},
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/tls_certificate/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/tls_certificate/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateTLSCertificate(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateTLSCertificate(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	updateRequest := &TLSCertificateUpdateRequest{
		PrivateKeyPassphrase: "newsecret",
		Nodes:                []string{"management", "conferencing"},
	}

	startDate := util.InfinityTime{}
	endDate := util.InfinityTime{}
	keyID := "key123"

	expectedTLSCertificate := &TLSCertificate{
		ID:                   1,
		Certificate:          "-----BEGIN CERTIFICATE-----\nMIIB...",
		PrivateKey:           "-----BEGIN PRIVATE KEY-----\nMIIE...",
		PrivateKeyPassphrase: "newsecret",
		Parameters:           "RSA-2048",
		Nodes:                []string{"management", "conferencing"},
		StartDate:            startDate,
		EndDate:              endDate,
		SubjectName:          "CN=example.com",
		KeyID:                &keyID,
		ResourceURI:          "/api/admin/configuration/v1/tls_certificate/1/",
	}

	client.On("PatchJSON", t.Context(), "configuration/v1/tls_certificate/1/", updateRequest, mock.AnythingOfType("*config.TLSCertificate")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*TLSCertificate)
		*result = *expectedTLSCertificate
	})

	service := New(client)
	result, err := service.UpdateTLSCertificate(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedTLSCertificate, result)
	client.AssertExpectations(t)
}

func TestService_DeleteTLSCertificate(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/tls_certificate/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteTLSCertificate(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
