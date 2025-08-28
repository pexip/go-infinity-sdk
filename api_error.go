/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package infinity

import (
	"encoding/json"
	"fmt"
)

// APIError represents an error returned by the Infinity API
type APIError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Details    string `json:"details,omitempty"`
}

func (e *APIError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("API error %d: %s (%s)", e.StatusCode, e.Message, e.Details)
	}
	return fmt.Sprintf("API error %d: %s", e.StatusCode, e.Message)
}

// UnmarshalJSON implements the json.Unmarshaler interface for APIError
func (e *APIError) UnmarshalJSON(data []byte) error {
	// Define a temporary struct to handle various error response formats
	var errorResp struct {
		Error   string `json:"error"`
		Message string `json:"message"`
		Details string `json:"details"`
		Detail  string `json:"detail"`
	}

	if err := json.Unmarshal(data, &errorResp); err != nil {
		return err
	}

	// Prioritize fields based on common API error response formats
	if errorResp.Error != "" {
		e.Message = errorResp.Error
	} else if errorResp.Message != "" {
		e.Message = errorResp.Message
	}

	if errorResp.Details != "" {
		e.Details = errorResp.Details
	} else if errorResp.Detail != "" {
		e.Details = errorResp.Detail
	}

	return nil
}
