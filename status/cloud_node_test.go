/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package status

import (
	"testing"
	"time"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	util "github.com/pexip/go-infinity-sdk/v38/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListCloudNodes(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	awsLaunchTime := time.Now().Add(-24 * time.Hour)
	cloudLaunchTime := time.Now().Add(-24 * time.Hour)

	expectedResponse := &CloudNodeListResponse{
		Meta: Meta{
			Limit:      100,
			Next:       "",
			Offset:     0,
			Previous:   "",
			TotalCount: 2,
		},
		Objects: []CloudNode{
			{
				AWSInstanceID:                     "i-1234567890abcdef0",
				AWSInstanceIP:                     "10.0.0.1",
				AWSInstanceLaunchTime:             &util.InfinityTime{Time: awsLaunchTime},
				AWSInstanceName:                   "cloud-node-1",
				AWSInstanceState:                  "running",
				CloudInstanceID:                   "cloud-1",
				CloudInstanceIP:                   "192.168.1.1",
				CloudInstanceLaunchTime:           &util.InfinityTime{Time: cloudLaunchTime},
				CloudInstanceName:                 "cloud-node-1",
				CloudInstanceState:                "running",
				MaxHDCalls:                        intPtr(10),
				MediaLoad:                         intPtr(20),
				ResourceURI:                       "/api/admin/status/v1/cloud_node/i-1234567890abcdef0/",
				WorkerVMConfigurationID:           intPtr(101),
				WorkerVMConfigurationLocationName: strPtr("us-west-2"),
				WorkerVMConfigurationName:         strPtr("worker-vm-1"),
			},
			{
				AWSInstanceID:                     "i-0987654321fedcba0",
				AWSInstanceIP:                     "10.0.0.2",
				AWSInstanceLaunchTime:             &util.InfinityTime{Time: awsLaunchTime},
				AWSInstanceName:                   "cloud-node-2",
				AWSInstanceState:                  "running",
				CloudInstanceID:                   "cloud-2",
				CloudInstanceIP:                   "192.168.1.2",
				CloudInstanceLaunchTime:           &util.InfinityTime{Time: cloudLaunchTime},
				CloudInstanceName:                 "cloud-node-2",
				CloudInstanceState:                "running",
				MaxHDCalls:                        intPtr(8),
				MediaLoad:                         intPtr(40),
				ResourceURI:                       "/api/admin/status/v1/cloud_node/i-0987654321fedcba0/",
				WorkerVMConfigurationID:           intPtr(102),
				WorkerVMConfigurationLocationName: strPtr("us-east-1"),
				WorkerVMConfigurationName:         strPtr("worker-vm-2"),
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
	assert.Equal(t, "i-1234567890abcdef0", result.Objects[0].AWSInstanceID)
	assert.Equal(t, "cloud-node-1", result.Objects[0].AWSInstanceName)
	assert.Equal(t, "running", result.Objects[0].AWSInstanceState)
	assert.Equal(t, "cloud-1", result.Objects[0].CloudInstanceID)
	assert.Equal(t, "us-west-2", derefString(result.Objects[0].WorkerVMConfigurationLocationName))
	assert.Equal(t, 10, derefInt(result.Objects[0].MaxHDCalls))
	assert.Equal(t, 20, derefInt(result.Objects[0].MediaLoad))
	assert.Equal(t, "/api/admin/status/v1/cloud_node/i-1234567890abcdef0/", result.Objects[0].ResourceURI)
	assert.Equal(t, "i-0987654321fedcba0", result.Objects[1].AWSInstanceID)
	assert.Equal(t, "cloud-node-2", result.Objects[1].AWSInstanceName)
	assert.Equal(t, "running", result.Objects[1].AWSInstanceState)
	assert.Equal(t, "cloud-2", result.Objects[1].CloudInstanceID)
	assert.Equal(t, "us-east-1", derefString(result.Objects[1].WorkerVMConfigurationLocationName))
	assert.Equal(t, 8, derefInt(result.Objects[1].MaxHDCalls))
	assert.Equal(t, 40, derefInt(result.Objects[1].MediaLoad))
	assert.Equal(t, "/api/admin/status/v1/cloud_node/i-0987654321fedcba0/", result.Objects[1].ResourceURI)
	client.AssertExpectations(t)
}

func TestService_GetCloudNode(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	awsLaunchTime := time.Now().Add(-48 * time.Hour)
	cloudLaunchTime := time.Now().Add(-48 * time.Hour)

	expectedNode := &CloudNode{
		AWSInstanceID:                     "i-1234567890abcdef0",
		AWSInstanceIP:                     "10.0.0.1",
		AWSInstanceLaunchTime:             &util.InfinityTime{Time: awsLaunchTime},
		AWSInstanceName:                   "cloud-node-primary",
		AWSInstanceState:                  "running",
		CloudInstanceID:                   "cloud-primary",
		CloudInstanceIP:                   "192.168.1.10",
		CloudInstanceLaunchTime:           &util.InfinityTime{Time: cloudLaunchTime},
		CloudInstanceName:                 "cloud-node-primary",
		CloudInstanceState:                "running",
		MaxHDCalls:                        intPtr(12),
		MediaLoad:                         intPtr(30),
		ResourceURI:                       "/api/admin/status/v1/cloud_node/i-1234567890abcdef0/",
		WorkerVMConfigurationID:           intPtr(201),
		WorkerVMConfigurationLocationName: strPtr("eu-west-1"),
		WorkerVMConfigurationName:         strPtr("worker-vm-primary"),
	}

	client.On("GetJSON", t.Context(), "status/v1/cloud_node/i-1234567890abcdef0/", mock.AnythingOfType("*status.CloudNode")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*CloudNode)
		*result = *expectedNode
	})

	service := New(client)
	result, err := service.GetCloudNode(t.Context(), "i-1234567890abcdef0")

	assert.NoError(t, err)
	assert.Equal(t, expectedNode, result)
	assert.Equal(t, "cloud-node-primary", result.AWSInstanceName)
	assert.Equal(t, "running", result.AWSInstanceState)
	assert.Equal(t, "cloud-primary", result.CloudInstanceID)
	assert.Equal(t, "eu-west-1", derefString(result.WorkerVMConfigurationLocationName))
	assert.Equal(t, 12, derefInt(result.MaxHDCalls))
	assert.Equal(t, 30, derefInt(result.MediaLoad))
	client.AssertExpectations(t)
}

func TestService_ListCloudNodes_WithOptions(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	service := New(client)

	opts := &ListOptions{
		Limit:  5,
		Offset: 10,
	}

	expectedResponse := &CloudNodeListResponse{
		Meta: Meta{
			Limit:      5,
			Next:       "",
			Offset:     10,
			Previous:   "",
			TotalCount: 1,
		},
		Objects: []CloudNode{
			{
				AWSInstanceID:                     "i-abc123def456",
				AWSInstanceIP:                     "10.0.0.3",
				AWSInstanceLaunchTime:             nil,
				AWSInstanceName:                   "cloud-node-test",
				AWSInstanceState:                  "pending",
				CloudInstanceID:                   "cloud-test",
				CloudInstanceIP:                   "192.168.1.3",
				CloudInstanceLaunchTime:           nil,
				CloudInstanceName:                 "cloud-node-test",
				CloudInstanceState:                "pending",
				MaxHDCalls:                        intPtr(5),
				MediaLoad:                         intPtr(10),
				ResourceURI:                       "/api/admin/status/v1/cloud_node/i-abc123def456/",
				WorkerVMConfigurationID:           intPtr(301),
				WorkerVMConfigurationLocationName: strPtr("ap-southeast-1"),
				WorkerVMConfigurationName:         strPtr("worker-vm-test"),
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
	assert.Equal(t, "cloud-node-test", result.Objects[0].AWSInstanceName)
	assert.Equal(t, "pending", result.Objects[0].AWSInstanceState)
	assert.Equal(t, "cloud-test", result.Objects[0].CloudInstanceID)
	assert.Equal(t, "ap-southeast-1", derefString(result.Objects[0].WorkerVMConfigurationLocationName))
	assert.Equal(t, 5, derefInt(result.Objects[0].MaxHDCalls))
	assert.Equal(t, 10, derefInt(result.Objects[0].MediaLoad))
	assert.Equal(t, "/api/admin/status/v1/cloud_node/i-abc123def456/", result.Objects[0].ResourceURI)

	client.AssertExpectations(t)
}
