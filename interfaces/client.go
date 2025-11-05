/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

// Package interfaces defines shared interfaces used across the go-infinity-sdk packages
// to enable dependency injection and avoid import cycles.
package interfaces

import (
	"context"
	"io"
	"net/url"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// HTTPClient defines the interface for HTTP operations used by the SDK services.
// This interface allows for dependency injection and testing without import cycles.
type HTTPClient interface {
	// GetJSON performs a GET request and unmarshals the JSON response
	GetJSON(ctx context.Context, endpoint string, queryParams *url.Values, result interface{}) error

	// PostJSON performs a POST request with JSON body and unmarshals the JSON response
	PostJSON(ctx context.Context, endpoint string, body interface{}, result interface{}) error

	// PutJSON performs a PUT request with JSON body and unmarshals the JSON response
	PutJSON(ctx context.Context, endpoint string, body interface{}, result interface{}) error

	// PatchJSON performs a PATCH request with JSON body and unmarshals the JSON response
	PatchJSON(ctx context.Context, endpoint string, body interface{}, result interface{}) error

	// DeleteJSON performs a DELETE request and unmarshals the JSON response
	DeleteJSON(ctx context.Context, endpoint string, result interface{}) error

	// PostFile performs a POST request with a file upload
	PostFile(ctx context.Context, endpoint string, fieldName, filename string, file io.Reader, result interface{}) error

	// PutFile performs a PUT request with a file upload
	PutFile(ctx context.Context, endpoint string, fieldName, filename string, file io.Reader, result interface{}) error

	// PatchFile performs a PATCH request with a file upload
	PatchFile(ctx context.Context, endpoint string, fieldName, filename string, file io.Reader, result interface{}) error

	// PostWithResponse performs a POST request and returns both the response body and location header
	PostWithResponse(ctx context.Context, endpoint string, body interface{}, result interface{}) (*types.PostResponse, error)
}
