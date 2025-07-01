package config

// Autobackup represents the automatic backup configuration (singleton resource)
type Autobackup struct {
	ID                       int    `json:"id,omitempty"`
	AutobackupEnabled        bool   `json:"autobackup_enabled"`
	AutobackupInterval       int    `json:"autobackup_interval"`
	AutobackupPassphrase     string `json:"autobackup_passphrase,omitempty"`
	AutobackupStartHour      int    `json:"autobackup_start_hour"`
	AutobackupUploadURL      string `json:"autobackup_upload_url,omitempty"`
	AutobackupUploadUsername string `json:"autobackup_upload_username,omitempty"`
	AutobackupUploadPassword string `json:"autobackup_upload_password,omitempty"`
	ResourceURI              string `json:"resource_uri,omitempty"`
}

// AutobackupUpdateRequest represents a request to update autobackup configuration
type AutobackupUpdateRequest struct {
	AutobackupEnabled        *bool  `json:"autobackup_enabled,omitempty"`
	AutobackupInterval       *int   `json:"autobackup_interval,omitempty"`
	AutobackupPassphrase     string `json:"autobackup_passphrase,omitempty"`
	AutobackupStartHour      *int   `json:"autobackup_start_hour,omitempty"`
	AutobackupUploadURL      string `json:"autobackup_upload_url,omitempty"`
	AutobackupUploadUsername string `json:"autobackup_upload_username,omitempty"`
	AutobackupUploadPassword string `json:"autobackup_upload_password,omitempty"`
}
