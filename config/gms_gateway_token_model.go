package config

// GMSGatewayToken represents the Google Meet gateway token configuration (singleton resource)
type GMSGatewayToken struct {
	ID                      int     `json:"id,omitempty"`
	Certificate             string  `json:"certificate,omitempty"`
	IntermediateCertificate *string `json:"intermediate_certificate,omitempty"`
	LeafCertificate         *string `json:"leaf_certificate,omitempty"`
	PrivateKey              *string `json:"private_key,omitempty"`
	SupportsDirectGuestJoin *bool   `json:"supports_direct_guest_join,omitempty"`
	ResourceURI             string  `json:"resource_uri,omitempty"`
}

// GMSGatewayTokenUpdateRequest represents a request to update Google Meet gateway token configuration
type GMSGatewayTokenUpdateRequest struct {
	Certificate             string  `json:"certificate,omitempty"`
	IntermediateCertificate *string `json:"intermediate_certificate,omitempty"`
	LeafCertificate         *string `json:"leaf_certificate,omitempty"`
	PrivateKey              *string `json:"private_key,omitempty"`
	SupportsDirectGuestJoin *bool   `json:"supports_direct_guest_join,omitempty"`
}
