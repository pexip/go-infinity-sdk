package status

import (
	"testing"
	"time"

	mockClient "github.com/pexip/go-infinity-sdk/internal/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListCloudNodes(t *testing.T) {
	client := &mockClient.Client{}

	launchTime := time.Now().Add(-24 * time.Hour)
	lastContactTime := time.Now().Add(-5 * time.Minute)

	expectedResponse := &CloudNodeListResponse{
		Objects: []CloudNode{
			{
				ID:                  "i-1234567890abcdef0",
				Name:                "cloud-node-1",
				Status:              "running",
				InstanceType:        "m5.large",
				Region:              "us-west-2",
				LaunchTime:          &launchTime,
				LastContactTime:     &lastContactTime,
				CPU:                 45.5,
				Memory:              68.2,
				ActiveConferences:   3,
				ActiveParticipants:  15,
				ResourceURI:         "/api/admin/status/v1/cloud_node/i-1234567890abcdef0/",
			},
			{
				ID:                  "i-0987654321fedcba0",
				Name:                "cloud-node-2",
				Status:              "running",
				InstanceType:        "m5.xlarge",
				Region:              "us-east-1",
				LaunchTime:          &launchTime,
				LastContactTime:     &lastContactTime,
				CPU:                 72.1,
				Memory:              84.3,
				ActiveConferences:   5,
				ActiveParticipants:  25,
				ResourceURI:         "/api/admin/status/v1/cloud_node/i-0987654321fedcba0/",
			},
		},
	}

	client.On("GetJSON", t.Context(), "status/v1/cloud_node/", mock.AnythingOfType("*status.CloudNodeListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*CloudNodeListResponse)
		*result = *expectedResponse
	})

	service := New(client)
	result, err := service.ListCloudNodes(t.Context(), nil)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result.Objects))
	assert.Equal(t, "cloud-node-1", result.Objects[0].Name)
	assert.Equal(t, "running", result.Objects[0].Status)
	assert.Equal(t, "m5.large", result.Objects[0].InstanceType)
	assert.Equal(t, "us-west-2", result.Objects[0].Region)
	assert.Equal(t, 3, result.Objects[0].ActiveConferences)
	assert.Equal(t, 15, result.Objects[0].ActiveParticipants)
	client.AssertExpectations(t)
}

func TestService_GetCloudNode(t *testing.T) {
	client := &mockClient.Client{}
	launchTime := time.Now().Add(-48 * time.Hour)
	lastContactTime := time.Now().Add(-2 * time.Minute)

	expectedNode := &CloudNode{
		ID:                  "i-1234567890abcdef0",
		Name:                "cloud-node-primary",
		Status:              "running",
		InstanceType:        "m5.2xlarge",
		Region:              "eu-west-1",
		LaunchTime:          &launchTime,
		LastContactTime:     &lastContactTime,
		CPU:                 85.7,
		Memory:              76.4,
		ActiveConferences:   8,
		ActiveParticipants:  42,
		ResourceURI:         "/api/admin/status/v1/cloud_node/i-1234567890abcdef0/",
	}

	client.On("GetJSON", t.Context(), "status/v1/cloud_node/i-1234567890abcdef0/", mock.AnythingOfType("*status.CloudNode")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*CloudNode)
		*result = *expectedNode
	})

	service := New(client)
	result, err := service.GetCloudNode(t.Context(), "i-1234567890abcdef0")

	assert.NoError(t, err)
	assert.Equal(t, expectedNode, result)
	assert.Equal(t, "cloud-node-primary", result.Name)
	assert.Equal(t, "running", result.Status)
	assert.Equal(t, "m5.2xlarge", result.InstanceType)
	assert.Equal(t, "eu-west-1", result.Region)
	assert.Equal(t, 8, result.ActiveConferences)
	assert.Equal(t, 42, result.ActiveParticipants)
	client.AssertExpectations(t)
}

func TestService_ListCloudNodes_WithOptions(t *testing.T) {
	client := &mockClient.Client{}
	service := New(client)

	opts := &ListOptions{
		Limit:  5,
		Offset: 10,
	}

	expectedResponse := &CloudNodeListResponse{
		Objects: []CloudNode{
			{
				ID:                 "i-abc123def456",
				Name:               "cloud-node-test",
				Status:             "pending",
				InstanceType:       "t3.medium",
				Region:             "ap-southeast-1",
				CPU:                25.0,
				Memory:             40.0,
				ActiveConferences:  1,
				ActiveParticipants: 5,
			},
		},
	}

	client.On("GetJSON", t.Context(), mock.MatchedBy(func(endpoint string) bool {
		return endpoint != "status/v1/cloud_node/"
	}), mock.AnythingOfType("*status.CloudNodeListResponse")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*CloudNodeListResponse)
		*result = *expectedResponse
	})

	result, err := service.ListCloudNodes(t.Context(), opts)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Objects))
	assert.Equal(t, "cloud-node-test", result.Objects[0].Name)
	assert.Equal(t, "pending", result.Objects[0].Status)
	assert.Equal(t, "t3.medium", result.Objects[0].InstanceType)

	client.AssertExpectations(t)
}