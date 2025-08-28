/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package infinity

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRetryConfig_IsRetriable(t *testing.T) {
	config := DefaultRetryConfig()

	tests := []struct {
		name       string
		statusCode int
		err        error
		want       bool
	}{
		{
			name:       "429 Too Many Requests",
			statusCode: http.StatusTooManyRequests,
			err:        nil,
			want:       true,
		},
		{
			name:       "500 Internal Server Error",
			statusCode: http.StatusInternalServerError,
			err:        nil,
			want:       true,
		},
		{
			name:       "502 Bad Gateway",
			statusCode: http.StatusBadGateway,
			err:        nil,
			want:       true,
		},
		{
			name:       "503 Service Unavailable",
			statusCode: http.StatusServiceUnavailable,
			err:        nil,
			want:       true,
		},
		{
			name:       "504 Gateway Timeout",
			statusCode: http.StatusGatewayTimeout,
			err:        nil,
			want:       true,
		},
		{
			name:       "400 Bad Request (not retriable)",
			statusCode: http.StatusBadRequest,
			err:        nil,
			want:       false,
		},
		{
			name:       "401 Unauthorized (not retriable)",
			statusCode: http.StatusUnauthorized,
			err:        nil,
			want:       false,
		},
		{
			name:       "404 Not Found (not retriable)",
			statusCode: http.StatusNotFound,
			err:        nil,
			want:       false,
		},
		{
			name:       "Context canceled (not retriable)",
			statusCode: 0,
			err:        context.Canceled,
			want:       false,
		},
		{
			name:       "Context deadline exceeded (not retriable)",
			statusCode: 0,
			err:        context.DeadlineExceeded,
			want:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := config.IsRetriable(tt.statusCode, tt.err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRetryConfig_CalculateBackoff(t *testing.T) {
	config := &RetryConfig{
		BackoffMin:   1 * time.Second,
		BackoffMax:   30 * time.Second,
		Multiplier:   2.0,
		JitterFactor: 0,
	}

	tests := []struct {
		name     string
		attempt  int
		expected time.Duration
	}{
		{
			name:     "attempt 0",
			attempt:  0,
			expected: 0,
		},
		{
			name:     "attempt 1",
			attempt:  1,
			expected: 1 * time.Second,
		},
		{
			name:     "attempt 2",
			attempt:  2,
			expected: 2 * time.Second,
		},
		{
			name:     "attempt 3",
			attempt:  3,
			expected: 4 * time.Second,
		},
		{
			name:     "attempt 4",
			attempt:  4,
			expected: 8 * time.Second,
		},
		{
			name:     "attempt 5",
			attempt:  5,
			expected: 16 * time.Second,
		},
		{
			name:     "attempt 6 (should be capped at max)",
			attempt:  6,
			expected: 30 * time.Second,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := config.CalculateBackoff(tt.attempt)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestRetryConfig_CalculateBackoffWithJitter(t *testing.T) {
	config := &RetryConfig{
		BackoffMin:   1 * time.Second,
		BackoffMax:   30 * time.Second,
		Multiplier:   2.0,
		JitterFactor: 0.1,
	}

	// Test that jitter produces different values
	backoff1 := config.CalculateBackoff(2)
	backoff2 := config.CalculateBackoff(2)

	// With jitter, we should get slightly different values
	// (though there's a small chance they could be the same)
	expected := 2 * time.Second
	tolerance := time.Duration(float64(expected) * 0.1) // 10% jitter tolerance

	assert.InDelta(t, float64(expected), float64(backoff1), float64(tolerance))
	assert.InDelta(t, float64(expected), float64(backoff2), float64(tolerance))
}

func TestClient_RetryOnServerError(t *testing.T) {
	attemptCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attemptCount++
		if attemptCount < 3 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status": "success"}`))
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
	require.NoError(t, err)

	var result map[string]interface{}
	err = client.GetJSON(context.Background(), "test", &result)

	assert.NoError(t, err)
	assert.Equal(t, "success", result["status"])
	assert.Equal(t, 3, attemptCount) // Should have tried 3 times
}

func TestClient_RetryOnNetworkError(t *testing.T) {
	// Create a server that immediately closes connections to simulate network errors
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// This should not be reached
		w.WriteHeader(http.StatusOK)
	}))
	serverURL := server.URL
	server.Close() // Close immediately to cause connection errors

	client, err := New(
		WithBaseURL(serverURL),
		WithRetryConfig(&RetryConfig{
			MaxRetries:   2,
			BackoffMin:   1 * time.Millisecond,
			BackoffMax:   10 * time.Millisecond,
			Multiplier:   2.0,
			JitterFactor: 0,
		}),
	)
	require.NoError(t, err)

	var result map[string]interface{}
	err = client.GetJSON(context.Background(), "test", &result)

	// Should fail after exhausting retries
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to perform HTTP request after 3 attempts")
}

func TestClient_NoRetryOnClientError(t *testing.T) {
	attemptCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attemptCount++
		w.WriteHeader(http.StatusBadRequest) // 400 - should not retry
		_, _ = w.Write([]byte(`{"error": "Bad Request"}`))
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
	require.NoError(t, err)

	var result map[string]interface{}
	err = client.GetJSON(context.Background(), "test", &result)

	assert.Error(t, err)
	assert.Equal(t, 1, attemptCount) // Should only try once
}

func TestClient_RetryExhaustion(t *testing.T) {
	attemptCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attemptCount++
		w.WriteHeader(http.StatusInternalServerError) // Always fail
	}))
	defer server.Close()

	client, err := New(
		WithBaseURL(server.URL),
		WithRetryConfig(&RetryConfig{
			MaxRetries:   2,
			BackoffMin:   1 * time.Millisecond,
			BackoffMax:   10 * time.Millisecond,
			Multiplier:   2.0,
			JitterFactor: 0,
		}),
	)
	require.NoError(t, err)

	var result map[string]interface{}
	err = client.GetJSON(context.Background(), "test", &result)

	assert.Error(t, err)
	assert.Equal(t, 3, attemptCount) // Should try 3 times total (1 initial + 2 retries)
}

func TestClient_ContextCancellation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Delay to ensure context cancellation happens
		time.Sleep(100 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client, err := New(
		WithBaseURL(server.URL),
		WithRetryConfig(&RetryConfig{
			MaxRetries:   3,
			BackoffMin:   1 * time.Second,
			BackoffMax:   10 * time.Second,
			Multiplier:   2.0,
			JitterFactor: 0,
		}),
	)
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	var result map[string]interface{}
	err = client.GetJSON(ctx, "test", &result)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "context deadline exceeded")
}

func TestClient_WithNoRetries(t *testing.T) {
	attemptCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attemptCount++
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client, err := New(
		WithBaseURL(server.URL),
		WithNoRetries(),
	)
	require.NoError(t, err)

	var result map[string]interface{}
	err = client.GetJSON(context.Background(), "test", &result)

	assert.Error(t, err)
	assert.Equal(t, 1, attemptCount) // Should only try once
}

func TestClient_WithMaxRetries(t *testing.T) {
	attemptCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attemptCount++
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client, err := New(
		WithBaseURL(server.URL),
		WithMaxRetries(1),
	)
	require.NoError(t, err)

	var result map[string]interface{}
	err = client.GetJSON(context.Background(), "test", &result)

	assert.Error(t, err)
	assert.Equal(t, 2, attemptCount) // Should try 2 times total (1 initial + 1 retry)
}

func TestDefaultRetryConfig(t *testing.T) {
	config := DefaultRetryConfig()

	assert.Equal(t, DefaultMaxRetries, config.MaxRetries)
	assert.Equal(t, DefaultBackoffMin, config.BackoffMin)
	assert.Equal(t, DefaultBackoffMax, config.BackoffMax)
	assert.Equal(t, DefaultBackoffMultiplier, config.Multiplier)
	assert.Equal(t, DefaultJitterFactor, config.JitterFactor)
}

func TestWithRetryConfig_NilConfig(t *testing.T) {
	_, err := New(WithRetryConfig(nil))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "retry config cannot be nil")
}

func TestWithMaxRetries_NegativeValue(t *testing.T) {
	_, err := New(WithMaxRetries(-1))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "max retries cannot be negative")
}

func TestRetryConfig_IsRetriable_URLErrors(t *testing.T) {
	config := DefaultRetryConfig()

	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "connection refused",
			err:  &url.Error{Op: "dial", URL: "http://example.com", Err: errors.New("connection refused")},
			want: true,
		},
		{
			name: "connection reset",
			err:  &url.Error{Op: "read", URL: "http://example.com", Err: errors.New("connection reset by peer")},
			want: true,
		},
		{
			name: "no such host",
			err:  &url.Error{Op: "dial", URL: "http://nonexistent.com", Err: errors.New("no such host")},
			want: true,
		},
		{
			name: "timeout",
			err:  &url.Error{Op: "dial", URL: "http://example.com", Err: errors.New("i/o timeout")},
			want: true,
		},
		{
			name: "temporary failure",
			err:  &url.Error{Op: "dial", URL: "http://example.com", Err: errors.New("temporary failure in name resolution")},
			want: true,
		},
		{
			name: "context canceled in url error",
			err:  &url.Error{Op: "dial", URL: "http://example.com", Err: context.Canceled},
			want: false,
		},
		{
			name: "context deadline exceeded in url error",
			err:  &url.Error{Op: "dial", URL: "http://example.com", Err: context.DeadlineExceeded},
			want: false,
		},
		{
			name: "other url error",
			err:  &url.Error{Op: "parse", URL: "invalid", Err: errors.New("invalid URL")},
			want: true,
		},
		{
			name: "non-url error",
			err:  errors.New("some other error"),
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := config.IsRetriable(0, tt.err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRetryConfig_CalculateBackoff_EdgeCases(t *testing.T) {
	tests := []struct {
		name    string
		config  *RetryConfig
		attempt int
		check   func(t *testing.T, duration time.Duration)
	}{
		{
			name: "zero jitter factor",
			config: &RetryConfig{
				BackoffMin:   100 * time.Millisecond,
				BackoffMax:   1000 * time.Millisecond,
				Multiplier:   2.0,
				JitterFactor: 0,
			},
			attempt: 2,
			check: func(t *testing.T, duration time.Duration) {
				assert.Equal(t, 200*time.Millisecond, duration)
			},
		},
		{
			name: "large jitter creates variance",
			config: &RetryConfig{
				BackoffMin:   100 * time.Millisecond,
				BackoffMax:   1000 * time.Millisecond,
				Multiplier:   2.0,
				JitterFactor: 1.0, // 100% jitter
			},
			attempt: 2,
			check: func(t *testing.T, duration time.Duration) {
				// With 100% jitter, result can range from 0 to 400ms
				assert.GreaterOrEqual(t, duration, time.Duration(0))
				assert.LessOrEqual(t, duration, 400*time.Millisecond)
			},
		},
		{
			name: "jitter can create negative value",
			config: &RetryConfig{
				BackoffMin:   100 * time.Millisecond,
				BackoffMax:   1000 * time.Millisecond,
				Multiplier:   2.0,
				JitterFactor: 2.0, // Very large jitter
			},
			attempt: 1,
			check: func(t *testing.T, duration time.Duration) {
				// Should not be negative, jitter negative cases fall back to BackoffMin
				assert.GreaterOrEqual(t, duration, time.Duration(0), "Duration should never be negative")
				// With very large jitter factor, we can get wide variance but always >= 0
			},
		},
		{
			name: "very small multiplier",
			config: &RetryConfig{
				BackoffMin:   100 * time.Millisecond,
				BackoffMax:   1000 * time.Millisecond,
				Multiplier:   0.5, // Decreasing backoff
				JitterFactor: 0,
			},
			attempt: 3,
			check: func(t *testing.T, duration time.Duration) {
				// 100 * 0.5^2 = 25ms
				assert.Equal(t, 25*time.Millisecond, duration)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			duration := tt.config.CalculateBackoff(tt.attempt)
			tt.check(t, duration)
		})
	}
}

func TestClient_RetryWithContextCancellationDuringBackoff(t *testing.T) {
	attemptCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attemptCount++
		w.WriteHeader(http.StatusInternalServerError) // Always fail
	}))
	defer server.Close()

	client, err := New(
		WithBaseURL(server.URL),
		WithRetryConfig(&RetryConfig{
			MaxRetries:   3,
			BackoffMin:   100 * time.Millisecond, // Long enough for cancellation
			BackoffMax:   1000 * time.Millisecond,
			Multiplier:   2.0,
			JitterFactor: 0,
		}),
	)
	require.NoError(t, err)

	// Cancel context after first failure but before retry backoff completes
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	var result map[string]interface{}
	err = client.GetJSON(ctx, "test", &result)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "context deadline exceeded")
	// Should only have made one attempt because context was canceled during backoff
	assert.Equal(t, 1, attemptCount)
}
