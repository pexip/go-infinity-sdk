package infinity

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestRetryIntegration demonstrates the complete retry mechanism in action
func TestRetryIntegration(t *testing.T) {
	var callCount int64

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		count := atomic.AddInt64(&callCount, 1)

		switch {
		case count <= 2:
			// First two calls fail with 503 Service Unavailable (retriable)
			w.WriteHeader(http.StatusServiceUnavailable)
			_, _ = w.Write([]byte(`{"error": "Service temporarily unavailable"}`))
		case count == 3:
			// Third call succeeds
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"status": "success", "message": "Service restored"}`))
		default:
			// Subsequent calls also succeed
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"status": "success", "message": "All good"}`))
		}
	}))
	defer server.Close()

	// Configure client with retry
	client, err := New(
		WithBaseURL(server.URL),
		WithRetryConfig(&RetryConfig{
			MaxRetries:   3,
			BackoffMin:   10 * time.Millisecond,
			BackoffMax:   100 * time.Millisecond,
			Multiplier:   2.0,
			JitterFactor: 0,
		}),
	)
	require.NoError(t, err)

	// Test successful retry scenario
	t.Run("successful_after_retries", func(t *testing.T) {
		atomic.StoreInt64(&callCount, 0) // Reset counter

		var result map[string]interface{}
		err := client.GetJSON(context.Background(), "test", &result)

		assert.NoError(t, err)
		assert.Equal(t, "success", result["status"])
		assert.Equal(t, "Service restored", result["message"])
		assert.Equal(t, int64(3), atomic.LoadInt64(&callCount))
	})

	// Test no retry on client errors
	t.Run("no_retry_on_client_error", func(t *testing.T) {
		atomic.StoreInt64(&callCount, 0) // Reset counter

		// Create a server that always returns 400 Bad Request
		badServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			atomic.AddInt64(&callCount, 1)
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(`{"error": "Bad Request"}`))
		}))
		defer badServer.Close()

		badClient, err := New(
			WithBaseURL(badServer.URL),
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
		err = badClient.GetJSON(context.Background(), "test", &result)

		assert.Error(t, err)
		assert.Equal(t, int64(1), atomic.LoadInt64(&callCount)) // Should only try once
	})
}

// TestRetryBackoffTiming verifies that backoff timing works correctly
func TestRetryBackoffTiming(t *testing.T) {
	config := &RetryConfig{
		MaxRetries:   3,
		BackoffMin:   50 * time.Millisecond,
		BackoffMax:   500 * time.Millisecond,
		Multiplier:   2.0,
		JitterFactor: 0, // No jitter for predictable timing
	}

	var callTimes []time.Time
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callTimes = append(callTimes, time.Now())
		w.WriteHeader(http.StatusInternalServerError) // Always fail
	}))
	defer server.Close()

	client, err := New(
		WithBaseURL(server.URL),
		WithRetryConfig(config),
	)
	require.NoError(t, err)

	start := time.Now()
	var result map[string]interface{}
	err = client.GetJSON(context.Background(), "test", &result)

	assert.Error(t, err)
	assert.Len(t, callTimes, 4) // 1 initial + 3 retries

	// Verify backoff timing (with some tolerance for execution overhead)
	tolerance := 20 * time.Millisecond

	// First retry should be after ~50ms
	firstRetryDelay := callTimes[1].Sub(callTimes[0])
	assert.InDelta(t, 50*time.Millisecond, firstRetryDelay, float64(tolerance))

	// Second retry should be after ~100ms
	secondRetryDelay := callTimes[2].Sub(callTimes[1])
	assert.InDelta(t, 100*time.Millisecond, secondRetryDelay, float64(tolerance))

	// Third retry should be after ~200ms
	thirdRetryDelay := callTimes[3].Sub(callTimes[2])
	assert.InDelta(t, 200*time.Millisecond, thirdRetryDelay, float64(tolerance))

	// Total time should be roughly 350ms + tolerance
	totalTime := time.Since(start)
	expectedTotal := 350 * time.Millisecond
	assert.True(t, totalTime >= expectedTotal)
	assert.True(t, totalTime <= expectedTotal+4*tolerance)
}

// TestRetryJitter verifies that jitter adds randomness to backoff timing
func TestRetryJitter(t *testing.T) {
	config := &RetryConfig{
		MaxRetries:   2,
		BackoffMin:   100 * time.Millisecond,
		BackoffMax:   1000 * time.Millisecond,
		Multiplier:   2.0,
		JitterFactor: 0.5, // 50% jitter
	}

	// Run multiple iterations to test jitter variance
	var backoffDurations []time.Duration

	for i := 0; i < 10; i++ {
		backoff := config.CalculateBackoff(1) // First retry
		backoffDurations = append(backoffDurations, backoff)
	}

	// With 50% jitter, we should see variance in the backoff durations
	// Base duration is 100ms, so with 50% jitter we expect range of 50ms-150ms
	minExpected := 50 * time.Millisecond
	maxExpected := 150 * time.Millisecond

	hasVariance := false
	for i, duration := range backoffDurations {
		assert.True(t, duration >= minExpected,
			"Duration %d (%v) should be >= %v", i, duration, minExpected)
		assert.True(t, duration <= maxExpected,
			"Duration %d (%v) should be <= %v", i, duration, maxExpected)

		// Check if we have variance (not all durations are the same)
		if i > 0 && duration != backoffDurations[0] {
			hasVariance = true
		}
	}

	assert.True(t, hasVariance, "Jitter should create variance in backoff durations")
}

// ExampleRetryConfig demonstrates how to configure retry behavior
func ExampleRetryConfig() {
	// Default retry configuration
	client1, _ := New(
		WithBaseURL("https://api.example.com"),
		WithBasicAuth("user", "pass"),
		// Uses DefaultRetryConfig() automatically
	)

	// Custom retry configuration
	customConfig := &RetryConfig{
		MaxRetries:   5,
		BackoffMin:   1 * time.Second,
		BackoffMax:   30 * time.Second,
		Multiplier:   1.5,
		JitterFactor: 0.1,
	}

	client2, _ := New(
		WithBaseURL("https://api.example.com"),
		WithBasicAuth("user", "pass"),
		WithRetryConfig(customConfig),
	)

	// Disable retries
	client3, _ := New(
		WithBaseURL("https://api.example.com"),
		WithBasicAuth("user", "pass"),
		WithNoRetries(),
	)

	// Set only max retries
	client4, _ := New(
		WithBaseURL("https://api.example.com"),
		WithBasicAuth("user", "pass"),
		WithMaxRetries(1),
	)

	// Use the clients...
	_ = client1
	_ = client2
	_ = client3
	_ = client4

	fmt.Println("Retry configurations applied successfully")
	// Output: Retry configurations applied successfully
}
