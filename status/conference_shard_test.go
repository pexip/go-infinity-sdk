package status

import (
	"testing"

	mockClient "github.com/pexip/go-infinity-sdk/v38/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListConferenceShards(t *testing.T) {
	client := &mockClient.Client{}

	expectedResponse := &ConferenceShardListResponse{
		Meta: Meta{
			Limit:      100,
			Next:       "",
			Offset:     0,
			Previous:   "",
			TotalCount: 2,
		},
		Objects: []ConferenceShard{
			{
				ID:                 "shard-1",
				Conference:         "conf-1",
				GuestsMuted:        false,
				IsDirect:           false,
				IsLocked:           false,
				IsStarted:          false,
				Node:               "node-01",
				ResourceURI:        "/api/admin/status/v1/conference_shard/shard-1/",
				ServiceType:        "",
				StartTime:          nil,
				SystemLocation:     "",
				Tag:                "",
				TranscodingEnabled: false,
			},
			{
				ID:                 "shard-2",
				Conference:         "conf-1",
				GuestsMuted:        false,
				IsDirect:           false,
				IsLocked:           false,
				IsStarted:          false,
				Node:               "node-02",
				ResourceURI:        "/api/admin/status/v1/conference_shard/shard-2/",
				ServiceType:        "",
				StartTime:          nil,
				SystemLocation:     "",
				Tag:                "",
				TranscodingEnabled: false,
			},
		},
	}

	client.On("GetJSON", t.Context(), "status/v1/conference_shard/", mock.AnythingOfType("*status.ConferenceShardListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*ConferenceShardListResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ListConferenceShards(t.Context(), nil)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result.Objects))
	assert.Equal(t, "shard-1", result.Objects[0].ID)
	assert.Equal(t, "conf-1", result.Objects[0].Conference)
	assert.Equal(t, "node-01", result.Objects[0].Node)
	assert.Equal(t, "/api/admin/status/v1/conference_shard/shard-1/", result.Objects[0].ResourceURI)
	assert.Equal(t, "shard-2", result.Objects[1].ID)
	assert.Equal(t, "node-02", result.Objects[1].Node)
	assert.Equal(t, "/api/admin/status/v1/conference_shard/shard-2/", result.Objects[1].ResourceURI)
	client.AssertExpectations(t)
}

func TestService_GetConferenceShard(t *testing.T) {
	client := &mockClient.Client{}

	expectedShard := &ConferenceShard{
		ID:                 "shard-primary",
		Conference:         "conf-primary",
		GuestsMuted:        true,
		IsDirect:           false,
		IsLocked:           true,
		IsStarted:          true,
		Node:               "node-primary",
		ResourceURI:        "/api/admin/status/v1/conference_shard/shard-primary/",
		ServiceType:        "meeting",
		StartTime:          nil,
		SystemLocation:     "loc-1",
		Tag:                "tag-1",
		TranscodingEnabled: true,
	}

	client.On("GetJSON", t.Context(), "status/v1/conference_shard/shard-primary/", mock.AnythingOfType("*status.ConferenceShard")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*ConferenceShard)
		*result = *expectedShard
	})

	service := New(client)
	result, err := service.GetConferenceShard(t.Context(), "shard-primary")

	assert.NoError(t, err)
	assert.Equal(t, expectedShard, result)
	assert.Equal(t, "shard-primary", result.ID)
	assert.Equal(t, "conf-primary", result.Conference)
	assert.Equal(t, "node-primary", result.Node)
	assert.Equal(t, "meeting", result.ServiceType)
	assert.Equal(t, "loc-1", result.SystemLocation)
	assert.Equal(t, "tag-1", result.Tag)
	assert.True(t, result.GuestsMuted)
	assert.True(t, result.IsLocked)
	assert.True(t, result.IsStarted)
	assert.True(t, result.TranscodingEnabled)
	client.AssertExpectations(t)
}

func TestService_ListConferenceShards_WithOptions(t *testing.T) {
	client := &mockClient.Client{}
	service := New(client)

	opts := &ListOptions{
		Limit:  5,
		Offset: 0,
	}

	expectedResponse := &ConferenceShardListResponse{
		Meta: Meta{
			Limit:      5,
			Next:       "",
			Offset:     0,
			Previous:   "",
			TotalCount: 1,
		},
		Objects: []ConferenceShard{
			{
				ID:                 "shard-test",
				Conference:         "conf-test",
				GuestsMuted:        false,
				IsDirect:           false,
				IsLocked:           false,
				IsStarted:          false,
				Node:               "node-test",
				ResourceURI:        "/api/admin/status/v1/conference_shard/shard-test/",
				ServiceType:        "meeting",
				StartTime:          nil,
				SystemLocation:     "loc-test",
				Tag:                "tag-test",
				TranscodingEnabled: false,
			},
		},
	}

	client.On("GetJSON", t.Context(), mock.MatchedBy(func(endpoint string) bool {
		return endpoint != "status/v1/conference_shard/"
	}), mock.AnythingOfType("*status.ConferenceShardListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*ConferenceShardListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListConferenceShards(t.Context(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "shard-test", result.Objects[0].ID)
	assert.Equal(t, "conf-test", result.Objects[0].Conference)
	assert.Equal(t, "node-test", result.Objects[0].Node)
	assert.Equal(t, "meeting", result.Objects[0].ServiceType)
	assert.Equal(t, "loc-test", result.Objects[0].SystemLocation)
	assert.Equal(t, "tag-test", result.Objects[0].Tag)

	client.AssertExpectations(t)
}
