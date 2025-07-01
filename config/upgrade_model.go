package config

// UpgradeCreateRequest represents a request to perform a system upgrade
type UpgradeCreateRequest struct {
	Package *string `json:"package,omitempty"`
}
