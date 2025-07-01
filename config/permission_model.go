package config

// Permission represents a permission (read-only)
type Permission struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Codename    string `json:"codename"`
	ResourceURI string `json:"resource_uri,omitempty"`
}

// PermissionListResponse represents the response from listing permissions
type PermissionListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []Permission `json:"objects"`
}
