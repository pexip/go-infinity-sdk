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

// PostResponse represents a POST response that includes both the response body and location header containing the resource URI.
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

// PostResponseWithUUID represents a POST response that includes both the response body and location header containing the resource UUID.
// this is used by a small number of Infinity resources that use UUIDs instead of integer IDs.
type PostResponseWithUUID struct {
	Body         []byte
	ResourceUUID string
}

func (pr *PostResponseWithUUID) ResUUID() (string, error) {
	re := regexp.MustCompile(`/([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})/?$`)
	match := re.FindStringSubmatch(pr.ResourceUUID)

	if len(match) > 1 {
		return match[1], nil
	}
	return "", fmt.Errorf("failed to extract resource UUID from location URI: %s", pr.ResourceUUID)
}
