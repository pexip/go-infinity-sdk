/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package infinity

import (
	"context"
	"errors"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// RetryConfig defines the retry behavior for the client
type RetryConfig struct {
	MaxRetries   int           // Maximum number of retries (0 = no retries)
	BackoffMin   time.Duration // Minimum backoff duration
	BackoffMax   time.Duration // Maximum backoff duration
	Multiplier   float64       // Backoff multiplier for exponential backoff
	JitterFactor float64       // Jitter factor to add randomness (0.0-1.0)
}

// DefaultRetryConfig returns a sensible default retry configuration
func DefaultRetryConfig() *RetryConfig {
	return &RetryConfig{
		MaxRetries:   DefaultMaxRetries,
		BackoffMin:   DefaultBackoffMin,
		BackoffMax:   DefaultBackoffMax,
		Multiplier:   DefaultBackoffMultiplier,
		JitterFactor: DefaultJitterFactor,
	}
}

// IsRetriable determines if an error/status code should be retried
func (rc *RetryConfig) IsRetriable(statusCode int, err error) bool {
	if err != nil {
		// Don't retry on context cancellation
		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			return false
		}

		// Check for URL errors (network issues)
		var urlErr *url.Error
		if errors.As(err, &urlErr) {
			// Don't retry on context cancellation wrapped in url.Error
			if errors.Is(urlErr.Err, context.Canceled) || errors.Is(urlErr.Err, context.DeadlineExceeded) {
				return false
			}

			// Check for common retriable network errors
			errStr := strings.ToLower(urlErr.Err.Error())
			if strings.Contains(errStr, "connection refused") ||
				strings.Contains(errStr, "connection reset") ||
				strings.Contains(errStr, "no such host") ||
				strings.Contains(errStr, "timeout") ||
				strings.Contains(errStr, "temporary failure") ||
				urlErr.Temporary() || urlErr.Timeout() {
				return true
			}
		}

		// Retry on other network-related errors by default
		return true
	}

	// Retry on specific HTTP status codes
	switch statusCode {
	case http.StatusTooManyRequests, // 429
		http.StatusInternalServerError, // 500
		http.StatusBadGateway,          // 502
		http.StatusServiceUnavailable,  // 503
		http.StatusGatewayTimeout:      // 504
		return true
	default:
		return false
	}
}

// CalculateBackoff calculates the backoff duration for a given attempt
func (rc *RetryConfig) CalculateBackoff(attempt int) time.Duration {
	if attempt <= 0 {
		return 0
	}

	// Calculate exponential backoff
	backoff := float64(rc.BackoffMin) * math.Pow(rc.Multiplier, float64(attempt-1))

	// Apply maximum backoff limit
	if maxBackoff := float64(rc.BackoffMax); backoff > maxBackoff {
		backoff = maxBackoff
	}

	// Add jitter to prevent thundering herd
	if rc.JitterFactor > 0 {
		jitter := backoff * rc.JitterFactor * (rand.Float64()*2 - 1) // Random value between -jitter and +jitter
		backoff += jitter
	}

	// Ensure backoff is not negative
	if backoff < 0 {
		backoff = float64(rc.BackoffMin)
	}

	return time.Duration(backoff)
}
