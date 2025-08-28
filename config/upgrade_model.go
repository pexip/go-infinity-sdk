/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// UpgradeCreateRequest represents a request to perform a system upgrade
type UpgradeCreateRequest struct {
	Package *string `json:"package,omitempty"`
}
