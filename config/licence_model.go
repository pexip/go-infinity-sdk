package config

// Licence represents a licence configuration
type Licence struct {
	FulfillmentID        string                 `json:"fulfillment_id,omitempty"`
	EntitlementID        string                 `json:"entitlement_id"`
	FulfillmentType      string                 `json:"fulfillment_type"`
	ProductID            string                 `json:"product_id"`
	LicenseType          string                 `json:"license_type,omitempty"`
	Features             string                 `json:"features"`
	Concurrent           int                    `json:"concurrent"`
	ConcurrentOverdraft  int                    `json:"concurrent_overdraft"`
	Activatable          int                    `json:"activatable"`
	ActivatableOverdraft int                    `json:"activatable_overdraft"`
	Hybrid               int                    `json:"hybrid"`
	HybridOverdraft      int                    `json:"hybrid_overdraft"`
	StartDate            string                 `json:"start_date,omitempty"`
	ExpirationDate       string                 `json:"expiration_date,omitempty"`
	Status               string                 `json:"status,omitempty"`
	TrustFlags           int                    `json:"trust_flags"`
	Repair               int                    `json:"repair"`
	ServerChain          string                 `json:"server_chain,omitempty"`
	VendorDictionary     map[string]interface{} `json:"vendor_dictionary"`
	OfflineMode          bool                   `json:"offline_mode"`
	ResourceURI          string                 `json:"resource_uri,omitempty"`
}

// LicenceCreateRequest represents a request to create a licence
type LicenceCreateRequest struct {
	EntitlementID string `json:"entitlement_id"`
	OfflineMode   bool   `json:"offline_mode"`
}

// LicenceListResponse represents the response from listing licences
type LicenceListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []Licence `json:"objects"`
}
