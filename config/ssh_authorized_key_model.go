package config

// SSHAuthorizedKey represents an SSH authorized key configuration
type SSHAuthorizedKey struct {
	ID          int      `json:"id,omitempty"`
	Keytype     string   `json:"keytype"`
	Key         string   `json:"key"`
	Comment     string   `json:"comment,omitempty"`
	Nodes       []string `json:"nodes,omitempty"`
	ResourceURI string   `json:"resource_uri,omitempty"`
}

// SSHAuthorizedKeyCreateRequest represents a request to create an SSH authorized key
type SSHAuthorizedKeyCreateRequest struct {
	Keytype string   `json:"keytype"`
	Key     string   `json:"key"`
	Comment string   `json:"comment,omitempty"`
	Nodes   []string `json:"nodes,omitempty"`
}

// SSHAuthorizedKeyUpdateRequest represents a request to update an SSH authorized key
type SSHAuthorizedKeyUpdateRequest struct {
	Keytype string   `json:"keytype,omitempty"`
	Key     string   `json:"key,omitempty"`
	Comment string   `json:"comment,omitempty"`
	Nodes   []string `json:"nodes,omitempty"`
}

// SSHAuthorizedKeyListResponse represents the response from listing SSH authorized keys
type SSHAuthorizedKeyListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []SSHAuthorizedKey `json:"objects"`
}
