package config

import (
	"context"
	"fmt"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListLdapRoles retrieves a list of LDAP roles
func (s *Service) ListLdapRoles(ctx context.Context, opts *ListOptions) (*LdapRoleListResponse, error) {
	endpoint := "configuration/v1/ldap_role/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result LdapRoleListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetLdapRole retrieves a specific LDAP role by ID
func (s *Service) GetLdapRole(ctx context.Context, id int) (*LdapRole, error) {
	endpoint := fmt.Sprintf("configuration/v1/ldap_role/%d/", id)

	var result LdapRole
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateLdapRole creates a new LDAP role
func (s *Service) CreateLdapRole(ctx context.Context, req *LdapRoleCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/ldap_role/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateLdapRole updates an existing LDAP role
func (s *Service) UpdateLdapRole(ctx context.Context, id int, req *LdapRoleUpdateRequest) (*LdapRole, error) {
	endpoint := fmt.Sprintf("configuration/v1/ldap_role/%d/", id)

	var result LdapRole
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteLdapRole deletes an LDAP role
func (s *Service) DeleteLdapRole(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/ldap_role/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
