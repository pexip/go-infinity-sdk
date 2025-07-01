package config

import (
	"context"
	"fmt"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListManagementVMs retrieves a list of management VMs
func (s *Service) ListManagementVMs(ctx context.Context, opts *ListOptions) (*ManagementVMListResponse, error) {
	endpoint := "configuration/v1/management_vm/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result ManagementVMListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetManagementVM retrieves a specific management VM by ID
func (s *Service) GetManagementVM(ctx context.Context, id int) (*ManagementVM, error) {
	endpoint := fmt.Sprintf("configuration/v1/management_vm/%d/", id)

	var result ManagementVM
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateManagementVM creates a new management VM
func (s *Service) CreateManagementVM(ctx context.Context, req *ManagementVMCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/management_vm/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// DeleteManagementVM deletes a management VM
func (s *Service) DeleteManagementVM(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/management_vm/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
