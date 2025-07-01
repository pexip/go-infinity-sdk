package config

// MjxEndpoint represents a MJX endpoint configuration
type MjxEndpoint struct {
	ID                             int     `json:"id,omitempty"`
	Name                           string  `json:"name"`
	Description                    string  `json:"description,omitempty"`
	EndpointType                   string  `json:"endpoint_type"`
	RoomResourceEmail              string  `json:"room_resource_email,omitempty"`
	MjxEndpointGroup               *string `json:"mjx_endpoint_group,omitempty"`
	APIAddress                     *string `json:"api_address,omitempty"`
	APIPort                        *int    `json:"api_port,omitempty"`
	APIUsername                    *string `json:"api_username,omitempty"`
	APIPassword                    *string `json:"api_password,omitempty"`
	UseHTTPS                       string  `json:"use_https"`
	VerifyCert                     string  `json:"verify_cert"`
	PolyUsername                   *string `json:"poly_username,omitempty"`
	PolyPassword                   *string `json:"poly_password,omitempty"`
	PolyRaiseAlarmsForThisEndpoint bool    `json:"poly_raise_alarms_for_this_endpoint"`
	WebexDeviceID                  *string `json:"webex_device_id,omitempty"`
	ResourceURI                    string  `json:"resource_uri,omitempty"`
}

// MjxEndpointCreateRequest represents a request to create a MJX endpoint
type MjxEndpointCreateRequest struct {
	Name                           string  `json:"name"`
	Description                    string  `json:"description,omitempty"`
	EndpointType                   string  `json:"endpoint_type"`
	RoomResourceEmail              string  `json:"room_resource_email,omitempty"`
	MjxEndpointGroup               *string `json:"mjx_endpoint_group,omitempty"`
	APIAddress                     *string `json:"api_address,omitempty"`
	APIPort                        *int    `json:"api_port,omitempty"`
	APIUsername                    *string `json:"api_username,omitempty"`
	APIPassword                    *string `json:"api_password,omitempty"`
	UseHTTPS                       string  `json:"use_https"`
	VerifyCert                     string  `json:"verify_cert"`
	PolyUsername                   *string `json:"poly_username,omitempty"`
	PolyPassword                   *string `json:"poly_password,omitempty"`
	PolyRaiseAlarmsForThisEndpoint bool    `json:"poly_raise_alarms_for_this_endpoint"`
	WebexDeviceID                  *string `json:"webex_device_id,omitempty"`
}

// MjxEndpointUpdateRequest represents a request to update a MJX endpoint
type MjxEndpointUpdateRequest struct {
	Name                           string  `json:"name,omitempty"`
	Description                    string  `json:"description,omitempty"`
	EndpointType                   string  `json:"endpoint_type,omitempty"`
	RoomResourceEmail              string  `json:"room_resource_email,omitempty"`
	MjxEndpointGroup               *string `json:"mjx_endpoint_group,omitempty"`
	APIAddress                     *string `json:"api_address,omitempty"`
	APIPort                        *int    `json:"api_port,omitempty"`
	APIUsername                    *string `json:"api_username,omitempty"`
	APIPassword                    *string `json:"api_password,omitempty"`
	UseHTTPS                       string  `json:"use_https,omitempty"`
	VerifyCert                     string  `json:"verify_cert,omitempty"`
	PolyUsername                   *string `json:"poly_username,omitempty"`
	PolyPassword                   *string `json:"poly_password,omitempty"`
	PolyRaiseAlarmsForThisEndpoint *bool   `json:"poly_raise_alarms_for_this_endpoint,omitempty"`
	WebexDeviceID                  *string `json:"webex_device_id,omitempty"`
}

// MjxEndpointListResponse represents the response from listing MJX endpoints
type MjxEndpointListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []MjxEndpoint `json:"objects"`
}
