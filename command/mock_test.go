/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package command

import (
	"context"

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

// PostJSON performs a POST request with JSON body and unmarshals the JSON response
func (m *mockHTTPClient) PostJSON(ctx context.Context, endpoint string, body interface{}, result interface{}) error {
	args := m.Called(ctx, endpoint, body, result)
	return args.Error(0)
}
