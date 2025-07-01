package config

// LdapSyncSource represents an LDAP sync source configuration
type LdapSyncSource struct {
	ID                   int    `json:"id,omitempty"`
	Name                 string `json:"name"`
	Description          string `json:"description,omitempty"`
	LdapServer           string `json:"ldap_server,omitempty"`
	LdapBaseDN           string `json:"ldap_base_dn,omitempty"`
	LdapBindUsername     string `json:"ldap_bind_username,omitempty"`
	LdapBindPassword     string `json:"ldap_bind_password,omitempty"`
	LdapUseGlobalCatalog bool   `json:"ldap_use_global_catalog"`
	LdapPermitNoTLS      bool   `json:"ldap_permit_no_tls"`
	ResourceURI          string `json:"resource_uri,omitempty"`
}

// LdapSyncSourceCreateRequest represents a request to create an LDAP sync source
type LdapSyncSourceCreateRequest struct {
	Name                 string `json:"name"`
	Description          string `json:"description,omitempty"`
	LdapServer           string `json:"ldap_server,omitempty"`
	LdapBaseDN           string `json:"ldap_base_dn,omitempty"`
	LdapBindUsername     string `json:"ldap_bind_username,omitempty"`
	LdapBindPassword     string `json:"ldap_bind_password,omitempty"`
	LdapUseGlobalCatalog bool   `json:"ldap_use_global_catalog"`
	LdapPermitNoTLS      bool   `json:"ldap_permit_no_tls"`
}

// LdapSyncSourceUpdateRequest represents a request to update an LDAP sync source
type LdapSyncSourceUpdateRequest struct {
	Name                 string `json:"name,omitempty"`
	Description          string `json:"description,omitempty"`
	LdapServer           string `json:"ldap_server,omitempty"`
	LdapBaseDN           string `json:"ldap_base_dn,omitempty"`
	LdapBindUsername     string `json:"ldap_bind_username,omitempty"`
	LdapBindPassword     string `json:"ldap_bind_password,omitempty"`
	LdapUseGlobalCatalog *bool  `json:"ldap_use_global_catalog,omitempty"`
	LdapPermitNoTLS      *bool  `json:"ldap_permit_no_tls,omitempty"`
}

// LdapSyncSourceListResponse represents the response from listing LDAP sync sources
type LdapSyncSourceListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []LdapSyncSource `json:"objects"`
}
