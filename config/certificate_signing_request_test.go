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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListCertificateSigningRequests(t *testing.T) {
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
				privateKey1 := "-----BEGIN RSA PRIVATE KEY-----\nMIIPrivateKey1\n-----END RSA PRIVATE KEY-----"
				privateKey2 := "-----BEGIN RSA PRIVATE KEY-----\nMIIPrivateKey2\n-----END RSA PRIVATE KEY-----"
				tlsCert1 := "/api/admin/configuration/v1/tls_certificate/1/"
				tlsCert2 := "/api/admin/configuration/v1/tls_certificate/2/"
				expectedResponse := &CertificateSigningRequestListResponse{
					Objects: []CertificateSigningRequest{
						{ID: 1, SubjectName: "CN=example.com", DN: "CN=example.com,O=Test Org,C=US", PrivateKeyType: "rsa_2048", PrivateKey: &privateKey1, AdCompatible: false, TLSCertificate: &tlsCert1, CSR: "-----BEGIN CERTIFICATE REQUEST-----\nMIICSR1\n-----END CERTIFICATE REQUEST-----"},
						{ID: 2, SubjectName: "CN=test.example.com", DN: "CN=test.example.com,O=Test Org,C=US", PrivateKeyType: "rsa_4096", PrivateKey: &privateKey2, AdCompatible: true, TLSCertificate: &tlsCert2, CSR: "-----BEGIN CERTIFICATE REQUEST-----\nMIICSR2\n-----END CERTIFICATE REQUEST-----"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/certificate_signing_request/", mock.AnythingOfType("*config.CertificateSigningRequestListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*CertificateSigningRequestListResponse)
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
				privateKey := "-----BEGIN RSA PRIVATE KEY-----\nMIIPrivateKey1\n-----END RSA PRIVATE KEY-----"
				tlsCert := "/api/admin/configuration/v1/tls_certificate/1/"
				expectedResponse := &CertificateSigningRequestListResponse{
					Objects: []CertificateSigningRequest{
						{ID: 1, SubjectName: "CN=example.com", DN: "CN=example.com,O=Test Org,C=US", PrivateKeyType: "rsa_2048", PrivateKey: &privateKey, AdCompatible: false, TLSCertificate: &tlsCert, CSR: "-----BEGIN CERTIFICATE REQUEST-----\nMIICSR1\n-----END CERTIFICATE REQUEST-----"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/certificate_signing_request/?limit=5&name__icontains=example.com", mock.AnythingOfType("*config.CertificateSigningRequestListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*CertificateSigningRequestListResponse)
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
			result, err := service.ListCertificateSigningRequests(t.Context(), tt.opts)

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

func TestService_GetCertificateSigningRequest(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	privateKey := "-----BEGIN RSA PRIVATE KEY-----\nMIITestPrivateKey\n-----END RSA PRIVATE KEY-----"
	tlsCert := "/api/admin/configuration/v1/tls_certificate/1/"
	expectedCertificateSigningRequest := &CertificateSigningRequest{
		ID:                        1,
		SubjectName:               "CN=test.example.com",
		DN:                        "CN=test.example.com,O=Test Org,C=US",
		AdditionalSubjectAltNames: "DNS:test.example.com,DNS:www.test.example.com",
		PrivateKeyType:            "rsa_2048",
		PrivateKey:                &privateKey,
		PrivateKeyPassphrase:      "test-passphrase",
		AdCompatible:              false,
		TLSCertificate:            &tlsCert,
		CSR:                       "-----BEGIN CERTIFICATE REQUEST-----\nMIITestCSR\n-----END CERTIFICATE REQUEST-----",
		Certificate:               "-----BEGIN CERTIFICATE-----\nMIITestCertificate\n-----END CERTIFICATE-----",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/certificate_signing_request/1/", mock.AnythingOfType("*config.CertificateSigningRequest")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*CertificateSigningRequest)
		*result = *expectedCertificateSigningRequest
	})

	service := New(client)
	result, err := service.GetCertificateSigningRequest(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedCertificateSigningRequest, result)
	client.AssertExpectations(t)
}

func TestService_CreateCertificateSigningRequest(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	privateKey := "-----BEGIN RSA PRIVATE KEY-----\nMIINewPrivateKey\n-----END RSA PRIVATE KEY-----"
	tlsCert := "/api/admin/configuration/v1/tls_certificate/1/"
	createRequest := &CertificateSigningRequestCreateRequest{
		SubjectName:               "CN=new.example.com",
		DN:                        "CN=new.example.com,O=New Org,C=US",
		AdditionalSubjectAltNames: "DNS:new.example.com",
		PrivateKeyType:            "rsa_4096",
		PrivateKey:                &privateKey,
		PrivateKeyPassphrase:      "new-passphrase",
		AdCompatible:              true,
		TLSCertificate:            &tlsCert,
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/certificate_signing_request/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/certificate_signing_request/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateCertificateSigningRequest(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateCertificateSigningRequest(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	adCompatible := true
	updateRequest := &CertificateSigningRequestUpdateRequest{
		SubjectName:    "CN=updated.example.com",
		PrivateKeyType: "rsa_4096",
		AdCompatible:   &adCompatible,
		Certificate:    "-----BEGIN CERTIFICATE-----\nMIIUpdatedCertificate\n-----END CERTIFICATE-----",
	}

	privateKey := "-----BEGIN RSA PRIVATE KEY-----\nMIITestPrivateKey\n-----END RSA PRIVATE KEY-----"
	tlsCert := "/api/admin/configuration/v1/tls_certificate/1/"
	expectedCertificateSigningRequest := &CertificateSigningRequest{
		ID:                        1,
		SubjectName:               "CN=updated.example.com",
		DN:                        "CN=test.example.com,O=Test Org,C=US",
		AdditionalSubjectAltNames: "DNS:test.example.com,DNS:www.test.example.com",
		PrivateKeyType:            "rsa_4096",
		PrivateKey:                &privateKey,
		PrivateKeyPassphrase:      "test-passphrase",
		AdCompatible:              true,
		TLSCertificate:            &tlsCert,
		CSR:                       "-----BEGIN CERTIFICATE REQUEST-----\nMIITestCSR\n-----END CERTIFICATE REQUEST-----",
		Certificate:               "-----BEGIN CERTIFICATE-----\nMIIUpdatedCertificate\n-----END CERTIFICATE-----",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/certificate_signing_request/1/", updateRequest, mock.AnythingOfType("*config.CertificateSigningRequest")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CertificateSigningRequest)
		*result = *expectedCertificateSigningRequest
	})

	service := New(client)
	result, err := service.UpdateCertificateSigningRequest(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedCertificateSigningRequest, result)
	client.AssertExpectations(t)
}

func TestService_DeleteCertificateSigningRequest(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/certificate_signing_request/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteCertificateSigningRequest(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
