package config

// LicenceRequest represents a licence request
type LicenceRequest struct {
	SequenceNumber string  `json:"sequence_number"`
	Reference      string  `json:"reference"`
	Actions        string  `json:"actions"`
	GenerationTime string  `json:"generation_time"`
	Status         string  `json:"status"`
	ResponseXML    *string `json:"response_xml,omitempty"`
	ResourceURI    string  `json:"resource_uri,omitempty"`
}

// LicenceRequestCreateRequest represents a request to create a licence request
type LicenceRequestCreateRequest struct {
	Reference string `json:"reference"`
	Actions   string `json:"actions"`
}

// LicenceRequestListResponse represents the response from listing licence requests
type LicenceRequestListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []LicenceRequest `json:"objects"`
}
