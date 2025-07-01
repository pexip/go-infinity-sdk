package config

// LogLevel represents a log level configuration
type LogLevel struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Level       string `json:"level"`
	ResourceURI string `json:"resource_uri,omitempty"`
}

// LogLevelCreateRequest represents a request to create a log level
type LogLevelCreateRequest struct {
	Name  string `json:"name"`
	Level string `json:"level"`
}

// LogLevelUpdateRequest represents a request to update a log level
type LogLevelUpdateRequest struct {
	Name  string `json:"name,omitempty"`
	Level string `json:"level,omitempty"`
}

// LogLevelListResponse represents the response from listing log levels
type LogLevelListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []LogLevel `json:"objects"`
}
