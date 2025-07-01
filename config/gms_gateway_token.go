package config

import (
	"context"
)

// GetGMSGatewayToken retrieves the Google Meet gateway token configuration (singleton resource)
func (s *Service) GetGMSGatewayToken(ctx context.Context) (*GMSGatewayToken, error) {
	endpoint := "configuration/v1/gms_gateway_token/1/"

	var result GMSGatewayToken
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// UpdateGMSGatewayToken updates the Google Meet gateway token configuration (singleton resource)
func (s *Service) UpdateGMSGatewayToken(ctx context.Context, req *GMSGatewayTokenUpdateRequest) (*GMSGatewayToken, error) {
	endpoint := "configuration/v1/gms_gateway_token/1/"

	var result GMSGatewayToken
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}
