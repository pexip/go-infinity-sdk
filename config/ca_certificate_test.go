package config

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListCACertificates(t *testing.T) {
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
				expectedResponse := &CACertificateListResponse{
					Objects: []CACertificate{
						{ID: 1, Certificate: "-----BEGIN CERTIFICATE-----\nMIICertificate1\n-----END CERTIFICATE-----", TrustedIntermediate: true, SubjectName: "CN=Root CA 1", IssuerName: "CN=Root CA 1"},
						{ID: 2, Certificate: "-----BEGIN CERTIFICATE-----\nMIICertificate2\n-----END CERTIFICATE-----", TrustedIntermediate: false, SubjectName: "CN=Root CA 2", IssuerName: "CN=Root CA 2"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/ca_certificate/", mock.AnythingOfType("*config.CACertificateListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*CACertificateListResponse)
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
				Search: "RootCA",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				expectedResponse := &CACertificateListResponse{
					Objects: []CACertificate{
						{ID: 1, Certificate: "-----BEGIN CERTIFICATE-----\nMIICertificate1\n-----END CERTIFICATE-----", TrustedIntermediate: true, SubjectName: "CN=Root CA 1", IssuerName: "CN=Root CA 1"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/ca_certificate/?limit=5&name__icontains=RootCA", mock.AnythingOfType("*config.CACertificateListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*CACertificateListResponse)
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
			result, err := service.ListCACertificates(t.Context(), tt.opts)

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

func TestService_GetCACertificate(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	keyID := "12:34:56:78:9a:bc:de:f0"
	issuerKeyID := "ab:cd:ef:12:34:56:78:9a"
	expectedCACertificate := &CACertificate{
		ID:                  1,
		Certificate:         "-----BEGIN CERTIFICATE-----\nMIITestCertificate\n-----END CERTIFICATE-----",
		TrustedIntermediate: true,
		SubjectName:         "CN=Test Root CA",
		SubjectHash:         "abcdef123456",
		RawSubject:          "CN=Test Root CA,O=Test Org,C=US",
		IssuerName:          "CN=Test Root CA",
		IssuerHash:          "abcdef123456",
		RawIssuer:           "CN=Test Root CA,O=Test Org,C=US",
		SerialNo:            "12345678901234567890",
		KeyID:               &keyID,
		IssuerKeyID:         &issuerKeyID,
		Text:                "Certificate details...",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/ca_certificate/1/", mock.AnythingOfType("*config.CACertificate")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*CACertificate)
		*result = *expectedCACertificate
	})

	service := New(client)
	result, err := service.GetCACertificate(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedCACertificate, result)
	client.AssertExpectations(t)
}

func TestService_CreateCACertificate(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &CACertificateCreateRequest{
		Certificate:         "-----BEGIN CERTIFICATE-----\nMIINewCertificate\n-----END CERTIFICATE-----",
		TrustedIntermediate: false,
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/ca_certificate/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/ca_certificate/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateCACertificate(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateCACertificate(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	trustedIntermediate := false
	updateRequest := &CACertificateUpdateRequest{
		Certificate:         "-----BEGIN CERTIFICATE-----\nMIIUpdatedCertificate\n-----END CERTIFICATE-----",
		TrustedIntermediate: &trustedIntermediate,
	}

	keyID := "12:34:56:78:9a:bc:de:f0"
	issuerKeyID := "ab:cd:ef:12:34:56:78:9a"
	expectedCACertificate := &CACertificate{
		ID:                  1,
		Certificate:         "-----BEGIN CERTIFICATE-----\nMIIUpdatedCertificate\n-----END CERTIFICATE-----",
		TrustedIntermediate: false,
		SubjectName:         "CN=Test Root CA",
		SubjectHash:         "abcdef123456",
		RawSubject:          "CN=Test Root CA,O=Test Org,C=US",
		IssuerName:          "CN=Test Root CA",
		IssuerHash:          "abcdef123456",
		RawIssuer:           "CN=Test Root CA,O=Test Org,C=US",
		SerialNo:            "12345678901234567890",
		KeyID:               &keyID,
		IssuerKeyID:         &issuerKeyID,
		Text:                "Certificate details...",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/ca_certificate/1/", updateRequest, mock.AnythingOfType("*config.CACertificate")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*CACertificate)
		*result = *expectedCACertificate
	})

	service := New(client)
	result, err := service.UpdateCACertificate(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedCACertificate, result)
	client.AssertExpectations(t)
}

func TestService_DeleteCACertificate(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/ca_certificate/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteCACertificate(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
