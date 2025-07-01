package config

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListPexipStreamingCredentials(t *testing.T) {
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
				expectedResponse := &PexipStreamingCredentialListResponse{
					Objects: []PexipStreamingCredential{
						{
							ID:        1,
							Kid:       "primary-streaming-key",
							PublicKey: "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA...\n-----END PUBLIC KEY-----",
						},
						{
							ID:        2,
							Kid:       "backup-streaming-key",
							PublicKey: "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA...\n-----END PUBLIC KEY-----",
						},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/pexip_streaming_credential/", mock.AnythingOfType("*config.PexipStreamingCredentialListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*PexipStreamingCredentialListResponse)
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
				Search: "primary",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				expectedResponse := &PexipStreamingCredentialListResponse{
					Objects: []PexipStreamingCredential{
						{
							ID:        1,
							Kid:       "primary-streaming-key",
							PublicKey: "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA...\n-----END PUBLIC KEY-----",
						},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/pexip_streaming_credential/?limit=5&name__icontains=primary", mock.AnythingOfType("*config.PexipStreamingCredentialListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*PexipStreamingCredentialListResponse)
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
			result, err := service.ListPexipStreamingCredentials(t.Context(), tt.opts)

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

func TestService_GetPexipStreamingCredential(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	expectedCredential := &PexipStreamingCredential{
		ID:        1,
		Kid:       "test-streaming-key",
		PublicKey: "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxyz1lCKE8rMQS7MdJvzz\nGtx3Ip8K7OqN9ZKF2eODJzCq3ZOD4oBnVtJOv3V8XqOWLm8HRgfEKVkZtQz2pMJ3\nJ5dT4V6L9QxzB7wX8oGjWzF9QyX8LmN2VfZpJzOKvQyF6qRz3tJ9QyNkD8G7X2Lf\nL8QV9rOq3ZdFG7wX2qVzPyGjNkT4V9LrOz3QyF8V2Lm3J9ZFq7wXzGjNkL8Vr2Qy\nF9LmOz3V7wXGjN4T8V2Lm3QyF9ZrOzVf7wXGjNkT8V2LmQyF9ZOz3V7wGjN4TLmQ\nyVZrOz3F7wXGjNkT8V2LmQyFZOzV7wXGjN4TLmQyVZOzF7wGjNkTLmQyVZF7wGjN\nQIDAQAB\n-----END PUBLIC KEY-----",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/pexip_streaming_credential/1/", mock.AnythingOfType("*config.PexipStreamingCredential")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*PexipStreamingCredential)
		*result = *expectedCredential
	})

	service := New(client)
	result, err := service.GetPexipStreamingCredential(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedCredential, result)
	client.AssertExpectations(t)
}

func TestService_CreatePexipStreamingCredential(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &PexipStreamingCredentialCreateRequest{
		Kid:       "new-streaming-key",
		PublicKey: "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAyGJnVkL3m8qOv4V9XqF2\nR7oZ8KdJ3LnM4V5TqG9WzO8P3mFj6K2LnV8XqZ9pT4V7wGjNkL8V2QyF9ZrOz3F7\nwXGjNkT8V2LmQyFZOzV7wXGjN4TLmQyVZOzF7wGjNkTLmQyVZF7wGjNkLmQyVZOz\nF7wXGjNkT8V2LmQyFZOzV7wXGjN4TLmQyVZOzF7wGjNkTLmQyVZF7wGjNkLmQyVZ\nOzF7wXGjNkT8V2LmQyFZOzV7wXGjN4TLmQyVZOzF7wGjNkTLmQyVZF7wGjNkLmQy\nVZOzF7wXGjNkT8V2LmQyFZOzV7wXGjN4TLmQyVZOzF7wGjNkTLmQyVZF7wGjNkLm\nQIDAQAB\n-----END PUBLIC KEY-----",
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/pexip_streaming_credential/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/pexip_streaming_credential/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreatePexipStreamingCredential(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdatePexipStreamingCredential(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	updateRequest := &PexipStreamingCredentialUpdateRequest{
		PublicKey: "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAupdated1lCKE8rMQS7Md\nJvzzGtx3Ip8K7OqN9ZKF2eODJzCq3ZOD4oBnVtJOv3V8XqOWLm8HRgfEKVkZtQz2\npMJ3J5dT4V6L9QxzB7wX8oGjWzF9QyX8LmN2VfZpJzOKvQyF6qRz3tJ9QyNkD8G7\nX2LfL8QV9rOq3ZdFG7wX2qVzPyGjNkT4V9LrOz3QyF8V2Lm3J9ZFq7wXzGjNkL8V\nr2QyF9LmOz3V7wXGjN4T8V2Lm3QyF9ZrOzVf7wXGjNkT8V2LmQyF9ZOz3V7wGjN4\nTLmQyVZrOz3F7wXGjNkT8V2LmQyFZOzV7wXGjN4TLmQyVZOzF7wGjNkTLmQyVZF7\nwGjNQIDAQAB\n-----END PUBLIC KEY-----",
	}

	expectedCredential := &PexipStreamingCredential{
		ID:        1,
		Kid:       "test-streaming-key",
		PublicKey: "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAupdated1lCKE8rMQS7Md\nJvzzGtx3Ip8K7OqN9ZKF2eODJzCq3ZOD4oBnVtJOv3V8XqOWLm8HRgfEKVkZtQz2\npMJ3J5dT4V6L9QxzB7wX8oGjWzF9QyX8LmN2VfZpJzOKvQyF6qRz3tJ9QyNkD8G7\nX2LfL8QV9rOq3ZdFG7wX2qVzPyGjNkT4V9LrOz3QyF8V2Lm3J9ZFq7wXzGjNkL8V\nr2QyF9LmOz3V7wXGjN4T8V2Lm3QyF9ZrOzVf7wXGjNkT8V2LmQyF9ZOz3V7wGjN4\nTLmQyVZrOz3F7wXGjNkT8V2LmQyFZOzV7wXGjN4TLmQyVZOzF7wGjNkTLmQyVZF7\nwGjNQIDAQAB\n-----END PUBLIC KEY-----",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/pexip_streaming_credential/1/", updateRequest, mock.AnythingOfType("*config.PexipStreamingCredential")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*PexipStreamingCredential)
		*result = *expectedCredential
	})

	service := New(client)
	result, err := service.UpdatePexipStreamingCredential(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedCredential, result)
	client.AssertExpectations(t)
}

func TestService_DeletePexipStreamingCredential(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/pexip_streaming_credential/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeletePexipStreamingCredential(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
