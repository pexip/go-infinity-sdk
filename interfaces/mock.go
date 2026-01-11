/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package interfaces

import (
	"context"
	"io"
	"net/url"

	"github.com/pexip/go-infinity-sdk/v38/types"
	"github.com/stretchr/testify/mock"
)

// HTTPClientMock is a mock implementation of HTTPClient interface for testing
type HTTPClientMock struct {
	mock.Mock
}

// GetJSON performs a GET request and unmarshals the JSON response
func (m *HTTPClientMock) GetJSON(ctx context.Context, endpoint string, queryParams *url.Values, result interface{}) error {
	args := m.Called(ctx, endpoint, queryParams, result)
	return args.Error(0)
}

// PostJSON performs a POST request with JSON body and unmarshals the JSON response
func (m *HTTPClientMock) PostJSON(ctx context.Context, endpoint string, body interface{}, result interface{}) error {
	args := m.Called(ctx, endpoint, body, result)
	return args.Error(0)
}

// PutJSON performs a PUT request with JSON body and unmarshals the JSON response
func (m *HTTPClientMock) PutJSON(ctx context.Context, endpoint string, body interface{}, result interface{}) error {
	args := m.Called(ctx, endpoint, body, result)
	return args.Error(0)
}

// PatchJSON performs a PATCH request with JSON body and unmarshals the JSON response
func (m *HTTPClientMock) PatchJSON(ctx context.Context, endpoint string, body interface{}, result interface{}) error {
	args := m.Called(ctx, endpoint, body, result)
	return args.Error(0)
}

// DeleteJSON performs a DELETE request and unmarshals the JSON response
func (m *HTTPClientMock) DeleteJSON(ctx context.Context, endpoint string, result interface{}) error {
	args := m.Called(ctx, endpoint, result)
	return args.Error(0)
}

// PostWithResponse performs a POST request and returns both the response body and location header
func (m *HTTPClientMock) PostWithResponse(ctx context.Context, endpoint string, body interface{}, result interface{}) (*types.PostResponse, error) {
	args := m.Called(ctx, endpoint, body, result)
	resp := args.Get(0)
	if resp == nil {
		return nil, args.Error(1)
	}
	return resp.(*types.PostResponse), args.Error(1)
}

func (m *HTTPClientMock) PostMultipartFormWithFieldsAndResponse(ctx context.Context, endpoint string, fields map[string]string, fileFieldName string, filename string, fileContent io.Reader, result interface{}) (*types.PostResponse, error) {
	args := m.Called(ctx, endpoint, fields, fileFieldName, filename, fileContent, result)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*types.PostResponse), args.Error(1)
}

func (m *HTTPClientMock) PatchMultipartFormWithFieldsAndResponse(ctx context.Context, endpoint string, fields map[string]string, fileFieldName string, filename string, fileContent io.Reader, result interface{}) (*types.PostResponse, error) {
	args := m.Called(ctx, endpoint, fields, fileFieldName, filename, fileContent, result)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*types.PostResponse), args.Error(1)
}

func (m *HTTPClientMock) PostMultipartFormWithFieldsAndResponseUUID(ctx context.Context, endpoint string, fields map[string]string, fieldName string, filename string, file io.Reader, response interface{}) (*types.PostResponseWithUUID, error) {
    args := m.Called(ctx, endpoint, fields, fieldName, filename, file, response)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).(*types.PostResponseWithUUID), args.Error(1)
}

// NewHTTPClientMock creates a new mock HTTP client
func NewHTTPClientMock() *HTTPClientMock {
	return &HTTPClientMock{}
}
