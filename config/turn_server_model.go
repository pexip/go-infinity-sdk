package config

// TURNServer represents a TURN server configuration
type TURNServer struct {
	ID            int    `json:"id,omitempty"`
	Name          string `json:"name"`
	Description   string `json:"description,omitempty"`
	Address       string `json:"address"`
	Port          *int   `json:"port,omitempty"`
	Username      string `json:"username,omitempty"`
	Password      string `json:"password,omitempty"`
	SecretKey     string `json:"secret_key,omitempty"`
	ServerType    string `json:"server_type"`
	TransportType string `json:"transport_type"`
	ResourceURI   string `json:"resource_uri,omitempty"`
}

// TURNServerCreateRequest represents a request to create a TURN server
type TURNServerCreateRequest struct {
	Name          string `json:"name"`
	Description   string `json:"description,omitempty"`
	Address       string `json:"address"`
	Port          *int   `json:"port,omitempty"`
	Username      string `json:"username,omitempty"`
	Password      string `json:"password,omitempty"`
	SecretKey     string `json:"secret_key,omitempty"`
	ServerType    string `json:"server_type"`
	TransportType string `json:"transport_type"`
}

// TURNServerUpdateRequest represents a request to update a TURN server
type TURNServerUpdateRequest struct {
	Name          string `json:"name,omitempty"`
	Description   string `json:"description,omitempty"`
	Address       string `json:"address,omitempty"`
	Port          *int   `json:"port,omitempty"`
	Username      string `json:"username,omitempty"`
	Password      string `json:"password,omitempty"`
	SecretKey     string `json:"secret_key,omitempty"`
	ServerType    string `json:"server_type,omitempty"`
	TransportType string `json:"transport_type,omitempty"`
}

// TURNServerListResponse represents the response from listing TURN servers
type TURNServerListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []TURNServer `json:"objects"`
}
