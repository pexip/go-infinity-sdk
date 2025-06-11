package command

import (
	"context"
)

// CreateSnapshot creates a system diagnostic snapshot
func (s *Service) CreateSnapshot(ctx context.Context, req *SnapshotRequest) (*CommandResponse, error) {
	endpoint := "command/v1/snapshot/"

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// CreateSnapshotSimple creates a basic snapshot with default options
func (s *Service) CreateSnapshotSimple(ctx context.Context) (*CommandResponse, error) {
	req := &SnapshotRequest{}
	return s.CreateSnapshot(ctx, req)
}

// ImportCertificates imports SSL/TLS certificates
func (s *Service) ImportCertificates(ctx context.Context, bundle, privateKeyPassphrase string) (*CommandResponse, error) {
	endpoint := "command/v1/certificates/import/"

	req := &CertificatesImportRequest{
		Bundle:               bundle,
		PrivateKeyPassphrase: privateKeyPassphrase,
	}

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// ManageSoftwareBundle manages software bundle operations
func (s *Service) ManageSoftwareBundle(ctx context.Context, packageName string) (*CommandResponse, error) {
	endpoint := "command/v1/software/bundle/"

	req := &SoftwareBundleRequest{
		Package: packageName,
	}

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// UpgradeSystem upgrades the system
func (s *Service) UpgradeSystem(ctx context.Context, packageName string) (*CommandResponse, error) {
	endpoint := "command/v1/upgrade/"

	req := &UpgradeRequest{
		Package: packageName,
	}

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// StartCloudNode starts a cloud node instance
func (s *Service) StartCloudNode(ctx context.Context, instanceID string) (*CommandResponse, error) {
	endpoint := "command/v1/cloudnode/start/"

	req := &StartCloudNodeRequest{
		InstanceID: instanceID,
	}

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}

// Sync performs system synchronization
func (s *Service) Sync(ctx context.Context, conferenceSyncTemplateID string) (*CommandResponse, error) {
	endpoint := "command/v1/sync/"

	req := &SyncRequest{
		ConferenceSyncTemplateID: conferenceSyncTemplateID,
	}

	var result CommandResponse
	err := s.client.PostJSON(ctx, endpoint, req, &result)
	return &result, err
}
