/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package infinity

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func BenchmarkClient_SuccessfulRequestNoRetry(b *testing.B) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	}))
	defer server.Close()

	client, err := New(
		WithBaseURL(server.URL),
		WithRetryConfig(&RetryConfig{
			MaxRetries:   3,
			BackoffMin:   1 * time.Millisecond,
			BackoffMax:   10 * time.Millisecond,
			Multiplier:   2.0,
			JitterFactor: 0,
		}),
	)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var result map[string]interface{}
		err := client.GetJSON(context.Background(), "test", &result)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkClient_SuccessfulRequestAfterRetries(b *testing.B) {
	attemptCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attemptCount++
		if attemptCount%3 != 0 { // Fail on first 2 attempts, succeed on 3rd
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	}))
	defer server.Close()

	client, err := New(
		WithBaseURL(server.URL),
		WithRetryConfig(&RetryConfig{
			MaxRetries:   3,
			BackoffMin:   1 * time.Millisecond,
			BackoffMax:   10 * time.Millisecond,
			Multiplier:   2.0,
			JitterFactor: 0,
		}),
	)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var result map[string]interface{}
		err := client.GetJSON(context.Background(), "test", &result)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkRetryConfig_CalculateBackoff(b *testing.B) {
	config := DefaultRetryConfig()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = config.CalculateBackoff(i%10 + 1) // Test with attempts 1-10
	}
}

func BenchmarkRetryConfig_IsRetriable(b *testing.B) {
	config := DefaultRetryConfig()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = config.IsRetriable(500, nil)
	}
}
