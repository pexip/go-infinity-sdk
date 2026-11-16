/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

// Package command provides access to the Pexip Infinity command API.
// It allows real-time control of conferences and participants including operations like
// muting, spotlighting, transferring, and sending messages.
package command

import "context"

// HTTPClient is the subset of HTTP operations required by the command service.
type HTTPClient interface {
	// PostJSON performs a POST request with JSON body and unmarshals the JSON response
	PostJSON(ctx context.Context, endpoint string, body interface{}, result interface{}) error
}

// Service handles command API endpoints
type Service struct {
	client HTTPClient
}

// New creates a new command API service
func New(client HTTPClient) *Service {
	return &Service{
		client: client,
	}
}
