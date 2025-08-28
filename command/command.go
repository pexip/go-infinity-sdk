/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

// Package command provides access to the Pexip Infinity command API.
// It allows real-time control of conferences and participants including operations like
// muting, spotlighting, transferring, and sending messages.
package command

import (
	"github.com/pexip/go-infinity-sdk/v38/interfaces"
)

// Service handles command API endpoints
type Service struct {
	client interfaces.HTTPClient
}

// New creates a new command API service
func New(client interfaces.HTTPClient) *Service {
	return &Service{
		client: client,
	}
}
