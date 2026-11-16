/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"context"
	"io"
	"net/url"

	"github.com/pexip/go-infinity-sdk/v42/types"
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

// PutJSON performs a PUT request with JSON body and unmarshals the JSON response
func (m *mockHTTPClient) PutJSON(ctx context.Context, endpoint string, body interface{}, result interface{}) error {
	args := m.Called(ctx, endpoint, body, result)
	return args.Error(0)
}

// PatchJSON performs a PATCH request with JSON body and unmarshals the JSON response
func (m *mockHTTPClient) PatchJSON(ctx context.Context, endpoint string, body interface{}, result interface{}) error {
	args := m.Called(ctx, endpoint, body, result)
	return args.Error(0)
}

// DeleteJSON performs a DELETE request and unmarshals the JSON response
func (m *mockHTTPClient) DeleteJSON(ctx context.Context, endpoint string, result interface{}) error {
	args := m.Called(ctx, endpoint, result)
	return args.Error(0)
}

// PostWithResponse performs a POST request and returns both the response body and location header
func (m *mockHTTPClient) PostWithResponse(ctx context.Context, endpoint string, body interface{}, result interface{}) (*types.PostResponse, error) {
	args := m.Called(ctx, endpoint, body, result)
	resp := args.Get(0)
	if resp == nil {
		return nil, args.Error(1)
	}
	return resp.(*types.PostResponse), args.Error(1)
}

func (m *mockHTTPClient) PostMultipartFormWithFieldsAndResponse(ctx context.Context, endpoint string, fields map[string]string, fileFieldName string, filename string, fileContent io.Reader, result interface{}) (*types.PostResponse, error) {
	args := m.Called(ctx, endpoint, fields, fileFieldName, filename, fileContent, result)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*types.PostResponse), args.Error(1)
}

func (m *mockHTTPClient) PatchMultipartFormWithFieldsAndResponse(ctx context.Context, endpoint string, fields map[string]string, fileFieldName string, filename string, fileContent io.Reader, result interface{}) (*types.PostResponse, error) {
	args := m.Called(ctx, endpoint, fields, fileFieldName, filename, fileContent, result)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*types.PostResponse), args.Error(1)
}

func (m *mockHTTPClient) PostMultipartFormWithFieldsAndResponseUUID(ctx context.Context, endpoint string, fields map[string]string, fieldName string, filename string, file io.Reader, response interface{}) (*types.PostResponseWithUUID, error) {
	args := m.Called(ctx, endpoint, fields, fieldName, filename, file, response)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*types.PostResponseWithUUID), args.Error(1)
}
