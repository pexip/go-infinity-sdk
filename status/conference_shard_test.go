package status

import (
	"testing"

	mockClient "github.com/pexip/go-infinity-sdk/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListConferenceShards(t *testing.T) {
	client := &mockClient.Client{}

	expectedResponse := &ConferenceShardListResponse{
		Objects: []ConferenceShard{
			{
				ID:               "shard-1",
				ConferenceName:   "Large Conference",
				ShardNumber:      1,
				NodeID:           "node-01",
				ParticipantCount: 45,
				ResourceURI:      "/api/admin/status/v1/conference_shard/shard-1/",
			},
			{
				ID:               "shard-2",
				ConferenceName:   "Large Conference",
				ShardNumber:      2,
				NodeID:           "node-02",
				ParticipantCount: 38,
				ResourceURI:      "/api/admin/status/v1/conference_shard/shard-2/",
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
	assert.Equal(t, "Large Conference", result.Objects[0].ConferenceName)
	assert.Equal(t, 1, result.Objects[0].ShardNumber)
	assert.Equal(t, "node-01", result.Objects[0].NodeID)
	assert.Equal(t, 45, result.Objects[0].ParticipantCount)
	assert.Equal(t, "shard-2", result.Objects[1].ID)
	assert.Equal(t, 2, result.Objects[1].ShardNumber)
	assert.Equal(t, 38, result.Objects[1].ParticipantCount)
	client.AssertExpectations(t)
}

func TestService_GetConferenceShard(t *testing.T) {
	client := &mockClient.Client{}

	expectedShard := &ConferenceShard{
		ID:               "shard-primary",
		ConferenceName:   "Executive Meeting",
		ShardNumber:      1,
		NodeID:           "node-primary",
		ParticipantCount: 25,
		ResourceURI:      "/api/admin/status/v1/conference_shard/shard-primary/",
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
	assert.Equal(t, "Executive Meeting", result.ConferenceName)
	assert.Equal(t, 1, result.ShardNumber)
	assert.Equal(t, "node-primary", result.NodeID)
	assert.Equal(t, 25, result.ParticipantCount)
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
		Objects: []ConferenceShard{
			{
				ID:               "shard-test",
				ConferenceName:   "Test Conference",
				ShardNumber:      1,
				NodeID:           "node-test",
				ParticipantCount: 12,
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
	assert.Equal(t, "Test Conference", result.Objects[0].ConferenceName)
	assert.Equal(t, 12, result.Objects[0].ParticipantCount)

	client.AssertExpectations(t)
}