/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

// Package types provides common types used across the Pexip Infinity SDK.
package types

import (
	"fmt"
	"regexp"
	"strconv"
)

// PostResponse represents a POST response that includes both the response body and location header
type PostResponse struct {
	Body        []byte
	ResourceURI string
}

func (pr *PostResponse) ResourceID() (int, error) {
	re := regexp.MustCompile(`/(\d+)/?$`)
	match := re.FindStringSubmatch(pr.ResourceURI)

	if len(match) > 1 {
		return strconv.Atoi(match[1])
	}
	return 0, fmt.Errorf("failed to extract resource ID from location URI: %s", pr.ResourceURI)
}
