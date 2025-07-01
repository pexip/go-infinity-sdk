package config

// DiagnosticGraph represents a diagnostic graph configuration
type DiagnosticGraph struct {
	ID          int      `json:"id,omitempty"`
	Title       string   `json:"title"`
	Order       int      `json:"order"`
	Datasets    []string `json:"datasets,omitempty"`
	ResourceURI string   `json:"resource_uri,omitempty"`
}

// DiagnosticGraphCreateRequest represents a request to create a diagnostic graph
type DiagnosticGraphCreateRequest struct {
	Title    string   `json:"title"`
	Order    int      `json:"order"`
	Datasets []string `json:"datasets,omitempty"`
}

// DiagnosticGraphUpdateRequest represents a request to update a diagnostic graph
type DiagnosticGraphUpdateRequest struct {
	Title    string   `json:"title,omitempty"`
	Order    *int     `json:"order,omitempty"`
	Datasets []string `json:"datasets,omitempty"`
}

// DiagnosticGraphListResponse represents the response from listing diagnostic graphs
type DiagnosticGraphListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []DiagnosticGraph `json:"objects"`
}
