package config

// UserGroup represents a user group configuration
type UserGroup struct {
	ID                      int      `json:"id,omitempty"`
	Name                    string   `json:"name"`
	Description             string   `json:"description,omitempty"`
	Users                   []string `json:"users,omitempty"`
	UserGroupEntityMappings []string `json:"user_group_entity_mappings,omitempty"`
	ResourceURI             string   `json:"resource_uri,omitempty"`
}

// UserGroupCreateRequest represents a request to create a user group
type UserGroupCreateRequest struct {
	Name                    string   `json:"name"`
	Description             string   `json:"description,omitempty"`
	Users                   []string `json:"users,omitempty"`
	UserGroupEntityMappings []string `json:"user_group_entity_mappings,omitempty"`
}

// UserGroupUpdateRequest represents a request to update a user group
type UserGroupUpdateRequest struct {
	Name                    string   `json:"name,omitempty"`
	Description             string   `json:"description,omitempty"`
	Users                   []string `json:"users,omitempty"`
	UserGroupEntityMappings []string `json:"user_group_entity_mappings,omitempty"`
}

// UserGroupListResponse represents the response from listing user groups
type UserGroupListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []UserGroup `json:"objects"`
}
