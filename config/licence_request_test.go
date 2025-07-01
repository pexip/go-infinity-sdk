package config

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListLicenceRequests(t *testing.T) {
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
				responseXML := "<?xml version=\"1.0\" encoding=\"UTF-8\"?><licence>test</licence>"
				expectedResponse := &LicenceRequestListResponse{
					Objects: []LicenceRequest{
						{SequenceNumber: "REQ-001", Reference: "Annual License Request", Actions: "ISSUE", GenerationTime: "2023-01-01T00:00:00Z", Status: "PENDING", ResponseXML: &responseXML},
						{SequenceNumber: "REQ-002", Reference: "Upgrade License Request", Actions: "UPGRADE", GenerationTime: "2023-01-02T00:00:00Z", Status: "COMPLETED", ResponseXML: nil},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/licence_request/", mock.AnythingOfType("*config.LicenceRequestListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*LicenceRequestListResponse)
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
				Search: "Annual",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				responseXML := "<?xml version=\"1.0\" encoding=\"UTF-8\"?><licence>test</licence>"
				expectedResponse := &LicenceRequestListResponse{
					Objects: []LicenceRequest{
						{SequenceNumber: "REQ-001", Reference: "Annual License Request", Actions: "ISSUE", GenerationTime: "2023-01-01T00:00:00Z", Status: "PENDING", ResponseXML: &responseXML},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/licence_request/?limit=5&name__icontains=Annual", mock.AnythingOfType("*config.LicenceRequestListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*LicenceRequestListResponse)
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
			result, err := service.ListLicenceRequests(t.Context(), tt.opts)

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

func TestService_GetLicenceRequest(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	responseXML := "<?xml version=\"1.0\" encoding=\"UTF-8\"?><licence>test response</licence>"
	expectedLicenceRequest := &LicenceRequest{
		SequenceNumber: "REQ-123",
		Reference:      "Test License Request",
		Actions:        "ISSUE",
		GenerationTime: "2023-01-01T12:00:00Z",
		Status:         "COMPLETED",
		ResponseXML:    &responseXML,
	}

	client.On("GetJSON", t.Context(), "configuration/v1/licence_request/REQ-123/", mock.AnythingOfType("*config.LicenceRequest")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*LicenceRequest)
		*result = *expectedLicenceRequest
	})

	service := New(client)
	result, err := service.GetLicenceRequest(t.Context(), "REQ-123")

	assert.NoError(t, err)
	assert.Equal(t, expectedLicenceRequest, result)
	client.AssertExpectations(t)
}

func TestService_CreateLicenceRequest(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &LicenceRequestCreateRequest{
		Reference: "New License Request",
		Actions:   "ISSUE",
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/licence_request/REQ-456/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/licence_request/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateLicenceRequest(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}
