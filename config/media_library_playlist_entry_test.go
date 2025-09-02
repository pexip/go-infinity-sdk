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

func TestService_ListMediaLibraryPlaylistEntries(t *testing.T) {
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
				media1 := "/api/admin/configuration/v1/media_library_entry/1/"
				media2 := "/api/admin/configuration/v1/media_library_entry/2/"
				playlist1 := "/api/admin/configuration/v1/media_library_playlist/1/"
				expectedResponse := &MediaLibraryPlaylistEntryListResponse{
					Objects: []MediaLibraryPlaylistEntry{
						{ID: 1, EntryType: "media", Media: &media1, Position: 1, Playcount: 0},
						{ID: 2, EntryType: "media", Media: &media2, Position: 2, Playcount: 5},
						{ID: 3, EntryType: "playlist", Playlist: &playlist1, Position: 3, Playcount: 2},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/media_library_playlist_entry/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.MediaLibraryPlaylistEntryListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*MediaLibraryPlaylistEntryListResponse)
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
				Search: "media",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				media1 := "/api/admin/configuration/v1/media_library_entry/1/"
				media2 := "/api/admin/configuration/v1/media_library_entry/2/"
				expectedResponse := &MediaLibraryPlaylistEntryListResponse{
					Objects: []MediaLibraryPlaylistEntry{
						{ID: 1, EntryType: "media", Media: &media1, Position: 1, Playcount: 0},
						{ID: 2, EntryType: "media", Media: &media2, Position: 2, Playcount: 5},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/media_library_playlist_entry/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.MediaLibraryPlaylistEntryListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*MediaLibraryPlaylistEntryListResponse)
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
			result, err := service.ListMediaLibraryPlaylistEntries(t.Context(), tt.opts)

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

func TestService_GetMediaLibraryPlaylistEntry(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	media := "/api/admin/configuration/v1/media_library_entry/1/"
	expectedMediaLibraryPlaylistEntry := &MediaLibraryPlaylistEntry{
		ID:        1,
		EntryType: "media",
		Media:     &media,
		Position:  1,
		Playcount: 10,
	}

	client.On("GetJSON", t.Context(), "configuration/v1/media_library_playlist_entry/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.MediaLibraryPlaylistEntry")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*MediaLibraryPlaylistEntry)
		*result = *expectedMediaLibraryPlaylistEntry
	})

	service := New(client)
	result, err := service.GetMediaLibraryPlaylistEntry(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedMediaLibraryPlaylistEntry, result)
	client.AssertExpectations(t)
}

func TestService_CreateMediaLibraryPlaylistEntry(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	media := "/api/admin/configuration/v1/media_library_entry/3/"
	createRequest := &MediaLibraryPlaylistEntryCreateRequest{
		EntryType: "media",
		Media:     &media,
		Position:  5,
		Playcount: 0,
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/media_library_playlist_entry/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/media_library_playlist_entry/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateMediaLibraryPlaylistEntry(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateMediaLibraryPlaylistEntry(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	position := 3
	playcount := 15
	updateRequest := &MediaLibraryPlaylistEntryUpdateRequest{
		Position:  &position,
		Playcount: &playcount,
	}

	media := "/api/admin/configuration/v1/media_library_entry/1/"
	expectedMediaLibraryPlaylistEntry := &MediaLibraryPlaylistEntry{
		ID:        1,
		EntryType: "media",
		Media:     &media,
		Position:  3,
		Playcount: 15,
	}

	client.On("PutJSON", t.Context(), "configuration/v1/media_library_playlist_entry/1/", updateRequest, mock.AnythingOfType("*config.MediaLibraryPlaylistEntry")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*MediaLibraryPlaylistEntry)
		*result = *expectedMediaLibraryPlaylistEntry
	})

	service := New(client)
	result, err := service.UpdateMediaLibraryPlaylistEntry(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedMediaLibraryPlaylistEntry, result)
	client.AssertExpectations(t)
}

func TestService_DeleteMediaLibraryPlaylistEntry(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/media_library_playlist_entry/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteMediaLibraryPlaylistEntry(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
