package server

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"os"
	"testing"
)

type mockRoundTripper struct {
	response *http.Response
	err      error
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.response, m.err
}

// Helper to allow custom logic in mock RoundTripper
type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestCheckConferences(t *testing.T) {
	mockBody := `{
        "meta": {},
        "objects": [
            {"name": "VMR_1", "tag": "prod"},
            {"name": "VMR_2", "tag": "prod"},
            {"name": "VMR_3", "tag": "dev"}
        ]
    }`

	mockResp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBufferString(mockBody)),
		Header:     make(http.Header),
	}

	client := &http.Client{
		Transport: &mockRoundTripper{response: mockResp},
	}

	managers := []Manager{
		{Host: "test", Port: "443", Scheme: "https", Metrics: ManagerMetrics{}},
	}
	s := &Server{
		Client:   client,
		Managers: &managers,
		Log:      slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}

	CheckConferences(s)

	if managers[0].Metrics.ConferenceCount["prod"] != 2 {
		t.Errorf("expected 2 prod conferences, got %d", managers[0].Metrics.ConferenceCount["prod"])
	}
	if managers[0].Metrics.ConferenceCount["dev"] != 1 {
		t.Errorf("expected 1 dev conference, got %d", managers[0].Metrics.ConferenceCount["dev"])
	}
}

func TestCheckParticipants(t *testing.T) {
	mockBody := `{
        "meta": {},
        "objects": [
            {"display_name": "Alice", "service_tag": "prod"},
            {"display_name": "Bob", "service_tag": "prod"},
            {"display_name": "Charlie", "service_tag": "dev"}
        ]
    }`

	mockResp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBufferString(mockBody)),
		Header:     make(http.Header),
	}

	client := &http.Client{
		Transport: &mockRoundTripper{response: mockResp},
	}

	managers := []Manager{
		{Host: "test",
			Port:   "443",
			Scheme: "https",
			Metrics: ManagerMetrics{
				ConferenceCount:         make(map[string]int),
				ParticipantByTagCount:   make(map[string]int),
				ParticipantByAliasCount: make(map[string]int),
				Nodes:                   make(map[string]Node),
				LicenseUsage:            LicenseUsage{},
			},
		},
	}
	s := &Server{
		Client:   client,
		Managers: &managers,
		Log:      slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}

	CheckParticipants(s)

	if managers[0].Metrics.ParticipantByTagCount["prod"] != 2 {
		t.Errorf("expected 2 prod participants, got %d", managers[0].Metrics.ParticipantByTagCount["prod"])
	}
	if managers[0].Metrics.ParticipantByTagCount["dev"] != 1 {
		t.Errorf("expected 1 dev participant, got %d", managers[0].Metrics.ParticipantByTagCount["dev"])
	}
}

// --- Test for GetCloudMonitoredLocations ---
func TestGetCloudMonitoredLocations(t *testing.T) {
	mockBody := `{
        "meta": {},
        "objects": [
            {"free_hd_calls": 10, "id": 1, "max_hd_calls": 20, "media_load": 5, "name": "London", "resource_uri": "/api/admin/status/v1/cloud_monitored_location/1/"},
            {"free_hd_calls": 5, "id": 2, "max_hd_calls": 10, "media_load": 2, "name": "Paris", "resource_uri": "/api/admin/status/v1/cloud_monitored_location/2/"}
        ]
    }`

	mockResp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBufferString(mockBody)),
		Header:     make(http.Header),
	}

	client := &http.Client{
		Transport: &mockRoundTripper{response: mockResp},
	}

	managers := []Manager{
		{Host: "test", Port: "443", Scheme: "https", Metrics: ManagerMetrics{}},
	}
	s := &Server{
		Client:   client,
		Managers: &managers,
		Log:      slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}

	locationMap := GetLocations(s)
	if len(locationMap["test"]) != 2 {
		t.Errorf("expected 2 locations, got %d", len(locationMap["test"]))
	}
	if locationMap["test"][0].Name != "London" {
		t.Errorf("expected first location to be London, got %s", locationMap["test"][0].Name)
	}
}

// --- Test for CheckNodes ---
func TestCheckNodes(t *testing.T) {
	// Mock for cloud monitored locations
	locationBody := `{
        "meta": {},
        "objects": [
            {"free_hd_calls": 10, "id": 1, "max_hd_calls": 20, "media_load": 5, "name": "London", "resource_uri": "/api/admin/status/v1/cloud_monitored_location/1/"}
        ]
    }`
	// Mock for nodes
	nodesBody := `{
        "meta": {},
        "objects": [
            {"id": 1, "max_hd_calls": 20, "media_load": 5, "name": "lon1", "node_type": "CONFERENCING", "system_location": "London"}
        ]
    }`

	// The first call returns locations, the second call returns nodes
	callCount := 0
	client := &http.Client{
		Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
			callCount++
			if callCount == 1 {
				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewBufferString(locationBody)),
					Header:     make(http.Header),
				}, nil
			}
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewBufferString(nodesBody)),
				Header:     make(http.Header),
			}, nil
		}),
	}

	managers := []Manager{
		{Host: "test", Port: "443", Scheme: "https", Metrics: ManagerMetrics{Nodes: make(map[string]Node)}},
	}
	s := &Server{
		Client:   client,
		Managers: &managers,
		Log:      slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}

	CheckNodes(s)

	node, ok := managers[0].Metrics.Nodes["lon1"]
	if !ok {
		t.Fatalf("expected node 'lon1' in metrics")
	}
	if node.MaxHDCalls != 20 {
		t.Errorf("expected MaxHDCalls 20, got %d", node.MaxHDCalls)
	}
	if !node.BurstOverflow {
		t.Errorf("expected BurstOverflow true, got false")
	}
}

func TestCheckLicenseUsage(t *testing.T) {
	mockBody := `{
        "meta": {},
        "objects": [
            {
                "audio_count": 1,
                "audio_total": 10,
                "customlayouts_active": true,
                "ghm_count": 2,
                "ghm_total": 20,
                "otj_count": 3,
                "otj_total": 30,
                "port_count": 4,
                "port_total": 40,
                "scheduling_count": 5,
                "scheduling_total": 50,
                "system_count": 6,
                "system_total": 60,
                "teams_count": 7,
                "teams_total": 70,
                "telehealth_count": 8,
                "telehealth_total": 80,
                "vmr_count": 9,
                "vmr_total": 90
            }
        ]
    }`

	mockResp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBufferString(mockBody)),
		Header:     make(http.Header),
	}

	client := &http.Client{
		Transport: &mockRoundTripper{response: mockResp},
	}

	managers := []Manager{
		{Host: "test", Port: "443", Scheme: "https", Metrics: ManagerMetrics{}},
	}
	s := &Server{
		Client:   client,
		Managers: &managers,
		Log:      slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}

	CheckLicenseUsage(s)

	lic := managers[0].Metrics.LicenseUsage
	if lic.PortCount != 4 {
		t.Errorf("expected PortCount 4, got %d", lic.PortCount)
	}
	if lic.PortTotal != 40 {
		t.Errorf("expected PortTotal 40, got %d", lic.PortTotal)
	}
	if !lic.CustomLayoutsActive {
		t.Errorf("expected CustomLayoutsActive true, got false")
	}
}
