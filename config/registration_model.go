/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// Registration represents the registration configuration (singleton resource)
type Registration struct {
	ID                         int    `json:"id,omitempty"`
	Enable                     bool   `json:"enable"`
	RefreshStrategy            string `json:"refresh_strategy"`
	AdaptiveMinRefresh         int    `json:"adaptive_min_refresh"`
	AdaptiveMaxRefresh         int    `json:"adaptive_max_refresh"`
	MaximumMinRefresh          int    `json:"maximum_min_refresh"`
	MaximumMaxRefresh          int    `json:"maximum_max_refresh"`
	NattedMinRefresh           int    `json:"natted_min_refresh"`
	NattedMaxRefresh           int    `json:"natted_max_refresh"`
	RouteViaRegistrar          bool   `json:"route_via_registrar"`
	EnablePushNotifications    bool   `json:"enable_push_notifications"`
	EnableGoogleCloudMessaging bool   `json:"enable_google_cloud_messaging"`
	PushToken                  string `json:"push_token,omitempty"`
	ResourceURI                string `json:"resource_uri,omitempty"`
}

// RegistrationUpdateRequest represents a request to update registration configuration
type RegistrationUpdateRequest struct {
	Enable                     *bool  `json:"enable,omitempty"`
	RefreshStrategy            string `json:"refresh_strategy,omitempty"`
	AdaptiveMinRefresh         *int   `json:"adaptive_min_refresh,omitempty"`
	AdaptiveMaxRefresh         *int   `json:"adaptive_max_refresh,omitempty"`
	MaximumMinRefresh          *int   `json:"maximum_min_refresh,omitempty"`
	MaximumMaxRefresh          *int   `json:"maximum_max_refresh,omitempty"`
	NattedMinRefresh           *int   `json:"natted_min_refresh,omitempty"`
	NattedMaxRefresh           *int   `json:"natted_max_refresh,omitempty"`
	RouteViaRegistrar          *bool  `json:"route_via_registrar,omitempty"`
	EnablePushNotifications    *bool  `json:"enable_push_notifications,omitempty"`
	EnableGoogleCloudMessaging *bool  `json:"enable_google_cloud_messaging,omitempty"`
	PushToken                  string `json:"push_token,omitempty"`
}
