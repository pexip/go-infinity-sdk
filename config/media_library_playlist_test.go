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

func TestService_ListMediaLibraryPlaylists(t *testing.T) {
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
				expectedResponse := &MediaLibraryPlaylistListResponse{
					Objects: []MediaLibraryPlaylist{
						{ID: 1, Name: "welcome-playlist", Description: "Welcome and onboarding videos", Loop: true, Shuffle: false, PlaylistEntries: []string{"/api/admin/configuration/v1/media_library_playlist_entry/1/", "/api/admin/configuration/v1/media_library_playlist_entry/2/"}},
						{ID: 2, Name: "hold-music-playlist", Description: "Music on hold collection", Loop: true, Shuffle: true, PlaylistEntries: []string{"/api/admin/configuration/v1/media_library_playlist_entry/3/", "/api/admin/configuration/v1/media_library_playlist_entry/4/"}},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/media_library_playlist/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.MediaLibraryPlaylistListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*MediaLibraryPlaylistListResponse)
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
				expectedResponse := &MediaLibraryPlaylistListResponse{
					Objects: []MediaLibraryPlaylist{
						{ID: 1, Name: "welcome-playlist", Description: "Welcome and onboarding videos", Loop: true, Shuffle: false, PlaylistEntries: []string{"/api/admin/configuration/v1/media_library_playlist_entry/1/", "/api/admin/configuration/v1/media_library_playlist_entry/2/"}},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/media_library_playlist/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.MediaLibraryPlaylistListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(3).(*MediaLibraryPlaylistListResponse)
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
			result, err := service.ListMediaLibraryPlaylists(t.Context(), tt.opts)

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

func TestService_GetMediaLibraryPlaylist(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedMediaLibraryPlaylist := &MediaLibraryPlaylist{
		ID:          1,
		Name:        "test-playlist",
		Description: "Test media playlist",
		Loop:        true,
		Shuffle:     false,
		PlaylistEntries: []string{
			"/api/admin/configuration/v1/media_library_playlist_entry/1/",
			"/api/admin/configuration/v1/media_library_playlist_entry/2/",
			"/api/admin/configuration/v1/media_library_playlist_entry/3/",
		},
	}

	client.On("GetJSON", t.Context(), "configuration/v1/media_library_playlist/1/", mock.AnythingOfType("*url.Values"), mock.AnythingOfType("*config.MediaLibraryPlaylist")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*MediaLibraryPlaylist)
		*result = *expectedMediaLibraryPlaylist
	})

	service := New(client)
	result, err := service.GetMediaLibraryPlaylist(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedMediaLibraryPlaylist, result)
	client.AssertExpectations(t)
}

func TestService_CreateMediaLibraryPlaylist(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &MediaLibraryPlaylistCreateRequest{
		Name:        "new-playlist",
		Description: "New media playlist",
		Loop:        false,
		Shuffle:     true,
		PlaylistEntries: []string{
			"/api/admin/configuration/v1/media_library_playlist_entry/5/",
			"/api/admin/configuration/v1/media_library_playlist_entry/6/",
		},
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/media_library_playlist/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/media_library_playlist/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateMediaLibraryPlaylist(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateMediaLibraryPlaylist(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	shuffle := true
	updateRequest := &MediaLibraryPlaylistUpdateRequest{
		Description: "Updated media playlist",
		Shuffle:     &shuffle,
		PlaylistEntries: []string{
			"/api/admin/configuration/v1/media_library_playlist_entry/1/",
			"/api/admin/configuration/v1/media_library_playlist_entry/2/",
			"/api/admin/configuration/v1/media_library_playlist_entry/7/",
		},
	}

	expectedMediaLibraryPlaylist := &MediaLibraryPlaylist{
		ID:          1,
		Name:        "test-playlist",
		Description: "Updated media playlist",
		Loop:        true,
		Shuffle:     true,
		PlaylistEntries: []string{
			"/api/admin/configuration/v1/media_library_playlist_entry/1/",
			"/api/admin/configuration/v1/media_library_playlist_entry/2/",
			"/api/admin/configuration/v1/media_library_playlist_entry/7/",
		},
	}

	client.On("PutJSON", t.Context(), "configuration/v1/media_library_playlist/1/", updateRequest, mock.AnythingOfType("*config.MediaLibraryPlaylist")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*MediaLibraryPlaylist)
		*result = *expectedMediaLibraryPlaylist
	})

	service := New(client)
	result, err := service.UpdateMediaLibraryPlaylist(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedMediaLibraryPlaylist, result)
	client.AssertExpectations(t)
}

func TestService_DeleteMediaLibraryPlaylist(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/media_library_playlist/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteMediaLibraryPlaylist(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
