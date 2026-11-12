/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package status

import (
	"context"
	"net/url"

	"github.com/stretchr/testify/mock"
)

// mockHTTPClient is a testify mock implementing HTTPClient
type mockHTTPClient struct {
	mock.Mock
}

var _ HTTPClient = (*mockHTTPClient)(nil)

// newMockHTTPClient creates a new mock HTTP client
func newMockHTTPClient() *mockHTTPClient {
	return &mockHTTPClient{}
}

// GetJSON performs a GET request and unmarshals the JSON response
func (m *mockHTTPClient) GetJSON(ctx context.Context, endpoint string, queryParams *url.Values, result interface{}) error {
	args := m.Called(ctx, endpoint, queryParams, result)
	return args.Error(0)
}
