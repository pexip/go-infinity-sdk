/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	infinity "github.com/pexip/go-infinity-sdk/v38"
)

func main() {
	// Example 1: Using WithTransport for proxy configuration
	proxyExample()

	// Example 2: Using WithTransport for custom TLS settings
	tlsExample()

	// Example 3: Using WithTransport for connection pooling
	connectionPoolingExample()
}

func proxyExample() {
	fmt.Println("=== Proxy Configuration Example ===")

	// Configure proxy
	proxyURL, err := url.Parse("http://proxy.company.com:8080")
	if err != nil {
		log.Printf("Failed to parse proxy URL: %v", err)
		return
	}

	// Create custom transport with proxy settings
	proxyTransport := &http.Transport{
		Proxy:               http.ProxyURL(proxyURL),
		MaxIdleConns:        50,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     30 * time.Second,
	}

	// Create client with proxy transport
	client, err := infinity.New(
		infinity.WithBaseURL("https://your-pexip-server.com"),
		infinity.WithTransport(proxyTransport),
		infinity.WithBasicAuth("admin", "password"),
	)
	if err != nil {
		log.Printf("Failed to create client: %v", err)
		return
	}

	fmt.Printf("ClientMock created with proxy transport: %+v\n", client)
}

func tlsExample() {
	fmt.Println("\n=== TLS Configuration Example ===")

	// Create custom TLS transport (for testing environments)
	tlsTransport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // WARNING: Only for testing!
			MinVersion:         tls.VersionTLS12,
		},
		MaxIdleConns:        30,
		MaxIdleConnsPerHost: 5,
		IdleConnTimeout:     60 * time.Second,
	}

	// Create client with TLS transport
	client, err := infinity.New(
		infinity.WithBaseURL("https://your-pexip-server.com"),
		infinity.WithTransport(tlsTransport),
		infinity.WithTokenAuth("your-api-token"),
	)
	if err != nil {
		log.Printf("Failed to create client: %v", err)
		return
	}

	// Example usage - get system status
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// This would make a real API call if the server was accessible
	_, err = client.Status().GetSystemStatus(ctx)
	if err != nil {
		fmt.Printf("API call failed (expected in example): %v\n", err)
	} else {
		fmt.Println("Successfully connected with custom TLS settings")
	}
}

func connectionPoolingExample() {
	fmt.Println("\n=== Connection Pooling Example ===")

	// Create transport optimized for high-throughput scenarios
	poolingTransport := &http.Transport{
		MaxIdleConns:        100,              // Total pool size
		MaxIdleConnsPerHost: 20,               // Per-host pool size
		MaxConnsPerHost:     50,               // Max concurrent connections per host
		IdleConnTimeout:     90 * time.Second, // Keep connections alive longer
		DisableCompression:  false,            // Enable compression
		ForceAttemptHTTP2:   true,             // Use HTTP/2 when possible
	}

	// Create client with optimized transport
	client, err := infinity.New(
		infinity.WithBaseURL("https://your-pexip-server.com"),
		infinity.WithTransport(poolingTransport),
		infinity.WithBasicAuth("admin", "password"),
		infinity.WithMaxRetries(5), // More retries for high-availability
	)
	if err != nil {
		log.Printf("Failed to create client: %v", err)
		return
	}

	fmt.Printf("ClientMock created with optimized connection pooling\n")
	fmt.Printf("Max idle connections: %d\n", poolingTransport.MaxIdleConns)
	fmt.Printf("Max idle connections per host: %d\n", poolingTransport.MaxIdleConnsPerHost)
	fmt.Printf("Max connections per host: %d\n", poolingTransport.MaxConnsPerHost)
	fmt.Printf("Transport configured for client: %T\n", client.HttpClient().Transport)
}
