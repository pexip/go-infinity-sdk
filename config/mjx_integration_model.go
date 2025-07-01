package config

// MjxIntegration represents a MJX integration configuration
type MjxIntegration struct {
	ID                          int      `json:"id,omitempty"`
	Name                        string   `json:"name"`
	Description                 string   `json:"description,omitempty"`
	DisplayUpcomingMeetings     int      `json:"display_upcoming_meetings"`
	EnableNonVideoMeetings      bool     `json:"enable_non_video_meetings"`
	EnablePrivateMeetings       bool     `json:"enable_private_meetings"`
	EndBuffer                   int      `json:"end_buffer"`
	StartBuffer                 int      `json:"start_buffer"`
	EPUsername                  string   `json:"ep_username,omitempty"`
	EPPassword                  string   `json:"ep_password,omitempty"`
	EPUseHTTPS                  bool     `json:"ep_use_https"`
	EPVerifyCertificate         bool     `json:"ep_verify_certificate"`
	ExchangeDeployment          *string  `json:"exchange_deployment,omitempty"`
	GoogleDeployment            *string  `json:"google_deployment,omitempty"`
	GraphDeployment             *string  `json:"graph_deployment,omitempty"`
	ProcessAliasPrivateMeetings bool     `json:"process_alias_private_meetings"`
	ReplaceEmptySubject         bool     `json:"replace_empty_subject"`
	ReplaceSubjectType          string   `json:"replace_subject_type"`
	ReplaceSubjectTemplate      string   `json:"replace_subject_template,omitempty"`
	UseWebex                    bool     `json:"use_webex"`
	WebexAPIDomain              string   `json:"webex_api_domain,omitempty"`
	WebexClientID               *string  `json:"webex_client_id,omitempty"`
	WebexClientSecret           *string  `json:"webex_client_secret,omitempty"`
	WebexOAuthState             *string  `json:"webex_oauth_state,omitempty"`
	WebexRedirectURI            *string  `json:"webex_redirect_uri,omitempty"`
	WebexRefreshToken           *string  `json:"webex_refresh_token,omitempty"`
	EndpointGroups              []string `json:"endpoint_groups,omitempty"`
	ResourceURI                 string   `json:"resource_uri,omitempty"`
}

// MjxIntegrationCreateRequest represents a request to create a MJX integration
type MjxIntegrationCreateRequest struct {
	Name                        string  `json:"name"`
	Description                 string  `json:"description,omitempty"`
	DisplayUpcomingMeetings     int     `json:"display_upcoming_meetings"`
	EnableNonVideoMeetings      bool    `json:"enable_non_video_meetings"`
	EnablePrivateMeetings       bool    `json:"enable_private_meetings"`
	EndBuffer                   int     `json:"end_buffer"`
	StartBuffer                 int     `json:"start_buffer"`
	EPUsername                  string  `json:"ep_username,omitempty"`
	EPPassword                  string  `json:"ep_password,omitempty"`
	EPUseHTTPS                  bool    `json:"ep_use_https"`
	EPVerifyCertificate         bool    `json:"ep_verify_certificate"`
	ExchangeDeployment          *string `json:"exchange_deployment,omitempty"`
	GoogleDeployment            *string `json:"google_deployment,omitempty"`
	GraphDeployment             *string `json:"graph_deployment,omitempty"`
	ProcessAliasPrivateMeetings bool    `json:"process_alias_private_meetings"`
	ReplaceEmptySubject         bool    `json:"replace_empty_subject"`
	ReplaceSubjectType          string  `json:"replace_subject_type"`
	ReplaceSubjectTemplate      string  `json:"replace_subject_template,omitempty"`
	UseWebex                    bool    `json:"use_webex"`
	WebexAPIDomain              string  `json:"webex_api_domain,omitempty"`
	WebexClientID               *string `json:"webex_client_id,omitempty"`
	WebexClientSecret           *string `json:"webex_client_secret,omitempty"`
	WebexRedirectURI            *string `json:"webex_redirect_uri,omitempty"`
	WebexRefreshToken           *string `json:"webex_refresh_token,omitempty"`
}

// MjxIntegrationUpdateRequest represents a request to update a MJX integration
type MjxIntegrationUpdateRequest struct {
	Name                        string  `json:"name,omitempty"`
	Description                 string  `json:"description,omitempty"`
	DisplayUpcomingMeetings     *int    `json:"display_upcoming_meetings,omitempty"`
	EnableNonVideoMeetings      *bool   `json:"enable_non_video_meetings,omitempty"`
	EnablePrivateMeetings       *bool   `json:"enable_private_meetings,omitempty"`
	EndBuffer                   *int    `json:"end_buffer,omitempty"`
	StartBuffer                 *int    `json:"start_buffer,omitempty"`
	EPUsername                  string  `json:"ep_username,omitempty"`
	EPPassword                  string  `json:"ep_password,omitempty"`
	EPUseHTTPS                  *bool   `json:"ep_use_https,omitempty"`
	EPVerifyCertificate         *bool   `json:"ep_verify_certificate,omitempty"`
	ExchangeDeployment          *string `json:"exchange_deployment,omitempty"`
	GoogleDeployment            *string `json:"google_deployment,omitempty"`
	GraphDeployment             *string `json:"graph_deployment,omitempty"`
	ProcessAliasPrivateMeetings *bool   `json:"process_alias_private_meetings,omitempty"`
	ReplaceEmptySubject         *bool   `json:"replace_empty_subject,omitempty"`
	ReplaceSubjectType          string  `json:"replace_subject_type,omitempty"`
	ReplaceSubjectTemplate      string  `json:"replace_subject_template,omitempty"`
	UseWebex                    *bool   `json:"use_webex,omitempty"`
	WebexAPIDomain              string  `json:"webex_api_domain,omitempty"`
	WebexClientID               *string `json:"webex_client_id,omitempty"`
	WebexClientSecret           *string `json:"webex_client_secret,omitempty"`
	WebexRedirectURI            *string `json:"webex_redirect_uri,omitempty"`
	WebexRefreshToken           *string `json:"webex_refresh_token,omitempty"`
}

// MjxIntegrationListResponse represents the response from listing MJX integrations
type MjxIntegrationListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []MjxIntegration `json:"objects"`
}
