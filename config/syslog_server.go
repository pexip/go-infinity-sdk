package config

import (
	"context"
	"fmt"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListSyslogServers retrieves a list of syslog servers
func (s *Service) ListSyslogServers(ctx context.Context, opts *ListOptions) (*SyslogServerListResponse, error) {
	endpoint := "configuration/v1/syslog_server/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result SyslogServerListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetSyslogServer retrieves a specific syslog server by ID
func (s *Service) GetSyslogServer(ctx context.Context, id int) (*SyslogServer, error) {
	endpoint := fmt.Sprintf("configuration/v1/syslog_server/%d/", id)

	var result SyslogServer
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateSyslogServer creates a new syslog server
func (s *Service) CreateSyslogServer(ctx context.Context, req *SyslogServerCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/syslog_server/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateSyslogServer updates an existing syslog server
func (s *Service) UpdateSyslogServer(ctx context.Context, id int, req *SyslogServerUpdateRequest) (*SyslogServer, error) {
	endpoint := fmt.Sprintf("configuration/v1/syslog_server/%d/", id)

	var result SyslogServer
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteSyslogServer deletes a syslog server
func (s *Service) DeleteSyslogServer(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/syslog_server/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
