package config

// LdapRole represents an LDAP role configuration
type LdapRole struct {
	ID          int      `json:"id,omitempty"`
	Name        string   `json:"name"`
	LdapGroupDN string   `json:"ldap_group_dn"`
	Roles       []string `json:"roles,omitempty"`
	ResourceURI string   `json:"resource_uri,omitempty"`
}

// LdapRoleCreateRequest represents a request to create an LDAP role
type LdapRoleCreateRequest struct {
	Name        string   `json:"name"`
	LdapGroupDN string   `json:"ldap_group_dn"`
	Roles       []string `json:"roles,omitempty"`
}

// LdapRoleUpdateRequest represents a request to update an LDAP role
type LdapRoleUpdateRequest struct {
	Name        string   `json:"name,omitempty"`
	LdapGroupDN string   `json:"ldap_group_dn,omitempty"`
	Roles       []string `json:"roles,omitempty"`
}

// LdapRoleListResponse represents the response from listing LDAP roles
type LdapRoleListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []LdapRole `json:"objects"`
}
