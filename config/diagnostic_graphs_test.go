package config

import (
	"testing"

	"github.com/pexip/go-infinity-sdk/v38/interfaces"
	"github.com/pexip/go-infinity-sdk/v38/options"
	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_ListDiagnosticGraphs(t *testing.T) {
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
				expectedResponse := &DiagnosticGraphListResponse{
					Objects: []DiagnosticGraph{
						{ID: 1, Title: "CPU Usage", Order: 1, Datasets: []string{"cpu_usage", "cpu_idle"}},
						{ID: 2, Title: "Memory Usage", Order: 2, Datasets: []string{"memory_used", "memory_free"}},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/diagnostic_graphs/", mock.AnythingOfType("*config.DiagnosticGraphListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*DiagnosticGraphListResponse)
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
				Search: "CPU",
			},
			setup: func(m *interfaces.HTTPClientMock) {
				expectedResponse := &DiagnosticGraphListResponse{
					Objects: []DiagnosticGraph{
						{ID: 1, Title: "CPU Usage", Order: 1, Datasets: []string{"cpu_usage", "cpu_idle"}},
					},
				}
				m.On("GetJSON", t.Context(), "configuration/v1/diagnostic_graphs/?limit=5&name__icontains=CPU", mock.AnythingOfType("*config.DiagnosticGraphListResponse")).Return(nil).Run(func(args mock.Arguments) {
					result := args.Get(2).(*DiagnosticGraphListResponse)
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
			result, err := service.ListDiagnosticGraphs(t.Context(), tt.opts)

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

func TestService_GetDiagnosticGraph(t *testing.T) {
	client := interfaces.NewHTTPClientMock()
	expectedDiagnosticGraph := &DiagnosticGraph{
		ID:       1,
		Title:    "Network Traffic",
		Order:    3,
		Datasets: []string{"network_in", "network_out", "packet_loss"},
	}

	client.On("GetJSON", t.Context(), "configuration/v1/diagnostic_graphs/1/", mock.AnythingOfType("*config.DiagnosticGraph")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(2).(*DiagnosticGraph)
		*result = *expectedDiagnosticGraph
	})

	service := New(client)
	result, err := service.GetDiagnosticGraph(t.Context(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedDiagnosticGraph, result)
	client.AssertExpectations(t)
}

func TestService_CreateDiagnosticGraph(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	createRequest := &DiagnosticGraphCreateRequest{
		Title:    "Disk Usage",
		Order:    4,
		Datasets: []string{"disk_used", "disk_free", "disk_io"},
	}

	expectedResponse := &types.PostResponse{
		Body:        []byte{},
		ResourceURI: "/api/admin/configuration/v1/diagnostic_graphs/123/",
	}

	client.On("PostWithResponse", t.Context(), "configuration/v1/diagnostic_graphs/", createRequest, nil).Return(expectedResponse, nil)

	service := New(client)
	result, err := service.CreateDiagnosticGraph(t.Context(), createRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
	client.AssertExpectations(t)
}

func TestService_UpdateDiagnosticGraph(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	order := 5
	updateRequest := &DiagnosticGraphUpdateRequest{
		Title:    "Updated Network Traffic",
		Order:    &order,
		Datasets: []string{"network_in", "network_out", "packet_loss", "bandwidth"},
	}

	expectedDiagnosticGraph := &DiagnosticGraph{
		ID:       1,
		Title:    "Updated Network Traffic",
		Order:    5,
		Datasets: []string{"network_in", "network_out", "packet_loss", "bandwidth"},
	}

	client.On("PutJSON", t.Context(), "configuration/v1/diagnostic_graphs/1/", updateRequest, mock.AnythingOfType("*config.DiagnosticGraph")).Return(nil).Run(func(args mock.Arguments) {
		result := args.Get(3).(*DiagnosticGraph)
		*result = *expectedDiagnosticGraph
	})

	service := New(client)
	result, err := service.UpdateDiagnosticGraph(t.Context(), 1, updateRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedDiagnosticGraph, result)
	client.AssertExpectations(t)
}

func TestService_DeleteDiagnosticGraph(t *testing.T) {
	client := interfaces.NewHTTPClientMock()

	client.On("DeleteJSON", t.Context(), "configuration/v1/diagnostic_graphs/1/", mock.Anything).Return(nil)

	service := New(client)
	err := service.DeleteDiagnosticGraph(t.Context(), 1)

	assert.NoError(t, err)
	client.AssertExpectations(t)
}
