package config

import "github.com/pexip/go-infinity-sdk/v38/util"

// SystemBackup represents a system backup configuration
type SystemBackup struct {
	Filename    string             `json:"filename"`
	Date        *util.InfinityTime `json:"date,omitempty"`
	Build       string             `json:"build,omitempty"`
	Version     string             `json:"version,omitempty"`
	Size        int                `json:"size"`
	ResourceURI string             `json:"resource_uri,omitempty"`
}

// SystemBackupListResponse represents the response from listing system backups
type SystemBackupListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []SystemBackup `json:"objects"`
}
