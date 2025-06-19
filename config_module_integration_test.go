//go:build integration

package infinity

import (
	"crypto/tls"
	"github.com/pexip/go-infinity-sdk/v38/config"
	"github.com/stretchr/testify/require"
	"net/http"
	"os"
	"testing"
)

var (
	infinityURL      = os.Getenv("INFINITY_URL")
	infinityUsername = os.Getenv("INFINITY_USERNAME")
	infinityPassword = os.Getenv("INFINITY_PASSWORD")
)

func TestRegisterWorkWithInfinity(t *testing.T) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client, err := New(
		WithBaseURL(infinityURL),
		WithBasicAuth(infinityUsername, infinityPassword),
		WithMaxRetries(2),
		WithTransport(transport),
	)
	require.NoError(t, err)

	locationsResp, err := client.Config.ListSystemLocations(t.Context(), nil)
	require.NoError(t, err)
	require.NotNil(t, locationsResp)
	require.Greater(t, len(locationsResp.Objects), 0, "locationsResp.Objects is empty")

	location := locationsResp.Objects[0]
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

	_, VMresp, err := client.Config.CreateWorkerVMWithResponse(t.Context(), req)
	require.NoError(t, err)
	require.NotNil(t, VMresp)

	id, err := VMresp.ResourceID()
	require.NoError(t, err)

	err = client.Config.DeleteWorkerVM(t.Context(), id)
	require.NoError(t, err, "Failed to delete worker VM after test")
}
