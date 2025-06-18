//go:build integration

package infinity

import (
	"crypto/tls"
	"github.com/pexip/go-infinity-sdk/v38/config"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestRegisterWorkWithInfinity(t *testing.T) {
	infinityURL := "https://34.13.196.28"
	infinityUsername := "pi99admin"
	infinityPassword := "Pexme123web"

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // For testing purposes only, do not use in production
		},
	}

	client, err := New(
		WithBaseURL(infinityURL),
		WithBasicAuth(infinityUsername, infinityPassword),
		WithMaxRetries(2),
		WithTransport(transport),
	)
	require.NoError(t, err)

	resp, err := client.Config.ListSystemLocations(t.Context(), nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	location := resp.Objects[0]

	req := &config.WorkerVMCreateRequest{
		Name:            "test-worker",
		Hostname:        "testworker",
		Domain:          "example.com",
		Address:         "10.0.0.10",
		Netmask:         "255.255.255.0",
		Gateway:         "10.0.0.1",
		NodeType:        "CONFERENCING",
		Transcoding:     true,
		Password:        "securepassword",
		MaintenanceMode: false,
		SystemLocation:  location.ResourceURI,
	}

	r, err := client.Config.CreateWorkerVM(t.Context(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)

	id, err := r.GetID()
	require.NoError(t, err)

	err = client.Config.DeleteWorkerVM(t.Context(), id)
	require.NoError(t, err, "Failed to delete worker VM after test")
}
