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

func TestService_ListMediaLibraryEntries(t *testing.T) {
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
				expectedResponse := &MediaLibraryEntryListResponse{
					Objects: []MediaLibraryEntry{
						{ID: 1, Name: "welcome-video", Description: "Welcome video for conferences", UUID: "123e4567-e89b-12d3-a456-426614174000", FileName: "welcome.mp4", MediaType: "video", MediaFormat: "video/mp4", MediaSize: 1048576, MediaFile: "/media/welcome.mp4"},
						{ID: 2, Name: "hold-music", Description: "Music on hold", UUID: "123e4567-e89b-12d3-a456-426614174001", FileName: "hold.mp3", MediaType: "audio", MediaFormat: "audio/mpeg", MediaSize: 512000, MediaFile: "/media/hold.mp3"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/media_library_entry/", mock.AnythingOfType("*config.MediaLibraryEntryListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*MediaLibraryEntryListResponse)
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
				Search: "welcome",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				expectedResponse := &MediaLibraryEntryListResponse{
					Objects: []MediaLibraryEntry{
						{ID: 1, Name: "welcome-video", Description: "Welcome video for conferences", UUID: "123e4567-e89b-12d3-a456-426614174000", FileName: "welcome.mp4", MediaType: "video", MediaFormat: "video/mp4", MediaSize: 1048576, MediaFile: "/media/welcome.mp4"},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/media_library_entry/?limit=5&name__icontains=welcome", mock.AnythingOfType("*config.MediaLibraryEntryListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*MediaLibraryEntryListResponse)
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
			result, err := service.ListMediaLibraryEntries(t.Context(), tt.opts)

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

func TestService_GetMediaLibraryEntry(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedMediaLibraryEntry := &MediaLibraryEntry{
		ID:          1,
		Name:        "test-media",
		Description: "Test media file",
		UUID:        "123e4567-e89b-12d3-a456-426614174000",
		FileName:    "test.mp4",
		MediaType:   "video",
		MediaFormat: "video/mp4",
		MediaSize:   2048000,
		MediaFile:   "/media/test.mp4",
	}

	client.On("GetJSON", t.Context(), "configuration/v1/media_library_entry/1/", mock.AnythingOfType("*config.MediaLibraryEntry")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*MediaLibraryEntry)
		*result = *expectedMediaLibraryEntry
	})

	service := New(client)
	result, err := service.GetMediaLibraryEntry(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedMediaLibraryEntry, result)
	client.AssertExpectations(t)
}

func TestService_CreateMediaLibraryEntry(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &MediaLibraryEntryCreateRequest{
		Name:        "new-media",
		Description: "New media file",
		UUID:        "123e4567-e89b-12d3-a456-426614174002",
		MediaFile:   "/uploads/new.mp4",
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/media_library_entry/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/media_library_entry/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateMediaLibraryEntry(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateMediaLibraryEntry(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	updateRequest := &MediaLibraryEntryUpdateRequest{
		Description: "Updated media file",
		MediaFile:   "/media/updated.mp4",
	}

	expectedMediaLibraryEntry := &MediaLibraryEntry{
		ID:          1,
		Name:        "test-media",
		Description: "Updated media file",
		UUID:        "123e4567-e89b-12d3-a456-426614174000",
		FileName:    "updated.mp4",
		MediaType:   "video",
		MediaFormat: "video/mp4",
		MediaSize:   2048000,
		MediaFile:   "/media/updated.mp4",
	}

	client.On("PutJSON", t.Context(), "configuration/v1/media_library_entry/1/", updateRequest, mock.AnythingOfType("*config.MediaLibraryEntry")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*MediaLibraryEntry)
		*result = *expectedMediaLibraryEntry
	})

	service := New(client)
	result, err := service.UpdateMediaLibraryEntry(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedMediaLibraryEntry, result)
	client.AssertExpectations(t)
}

func TestService_DeleteMediaLibraryEntry(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/media_library_entry/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteMediaLibraryEntry(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
