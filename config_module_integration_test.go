//go:build integration

/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

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

	locationsResp, err := client.config.ListSystemLocations(t.Context(), nil)
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

	VMresp, err := client.config.CreateWorkerVM(t.Context(), req)
	require.NoError(t, err)
	require.NotNil(t, VMresp)

	id, err := VMresp.ResourceID()
	require.NoError(t, err)

	err = client.config.DeleteWorkerVM(t.Context(), id)
	require.NoError(t, err, "Failed to delete worker VM after test")
}

func TestRegisterDNSServerWithInfinity(t *testing.T) {
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

	req := &config.DNSServerCreateRequest{
		Address:     "1.1.1.1",
		Description: "Test DNS Server",
	}

	dnsResp, err := client.config.CreateDNSServer(t.Context(), req)
	require.NoError(t, err)
	require.NotNil(t, dnsResp)

	id, err := dnsResp.ResourceID()
	require.NoError(t, err)

	err = client.config.DeleteDNSServer(t.Context(), id)
	require.NoError(t, err, "Failed to delete DNS server after test")
}
