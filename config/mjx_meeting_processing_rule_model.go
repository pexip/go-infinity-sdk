package config

// MjxMeetingProcessingRule represents a MJX meeting processing rule configuration
type MjxMeetingProcessingRule struct {
	ID                       int    `json:"id,omitempty"`
	Name                     string `json:"name"`
	Description              string `json:"description,omitempty"`
	Priority                 int    `json:"priority"`
	Enabled                  bool   `json:"enabled"`
	MeetingType              string `json:"meeting_type"`
	MjxIntegration           string `json:"mjx_integration"`
	MatchString              string `json:"match_string,omitempty"`
	ReplaceString            string `json:"replace_string,omitempty"`
	TransformRule            string `json:"transform_rule,omitempty"`
	CustomTemplate           string `json:"custom_template,omitempty"`
	Domain                   string `json:"domain,omitempty"`
	CompanyID                string `json:"company_id,omitempty"`
	IncludePin               bool   `json:"include_pin"`
	DefaultProcessingEnabled bool   `json:"default_processing_enabled"`
	ResourceURI              string `json:"resource_uri,omitempty"`
}

// MjxMeetingProcessingRuleCreateRequest represents a request to create a MJX meeting processing rule
type MjxMeetingProcessingRuleCreateRequest struct {
	Name                     string `json:"name"`
	Description              string `json:"description,omitempty"`
	Priority                 int    `json:"priority"`
	Enabled                  bool   `json:"enabled"`
	MeetingType              string `json:"meeting_type"`
	MjxIntegration           string `json:"mjx_integration"`
	MatchString              string `json:"match_string,omitempty"`
	ReplaceString            string `json:"replace_string,omitempty"`
	TransformRule            string `json:"transform_rule,omitempty"`
	CustomTemplate           string `json:"custom_template,omitempty"`
	Domain                   string `json:"domain,omitempty"`
	CompanyID                string `json:"company_id,omitempty"`
	IncludePin               bool   `json:"include_pin"`
	DefaultProcessingEnabled bool   `json:"default_processing_enabled"`
}

// MjxMeetingProcessingRuleUpdateRequest represents a request to update a MJX meeting processing rule
type MjxMeetingProcessingRuleUpdateRequest struct {
	Name                     string `json:"name,omitempty"`
	Description              string `json:"description,omitempty"`
	Priority                 *int   `json:"priority,omitempty"`
	Enabled                  *bool  `json:"enabled,omitempty"`
	MeetingType              string `json:"meeting_type,omitempty"`
	MjxIntegration           string `json:"mjx_integration,omitempty"`
	MatchString              string `json:"match_string,omitempty"`
	ReplaceString            string `json:"replace_string,omitempty"`
	TransformRule            string `json:"transform_rule,omitempty"`
	CustomTemplate           string `json:"custom_template,omitempty"`
	Domain                   string `json:"domain,omitempty"`
	CompanyID                string `json:"company_id,omitempty"`
	IncludePin               *bool  `json:"include_pin,omitempty"`
	DefaultProcessingEnabled *bool  `json:"default_processing_enabled,omitempty"`
}

// MjxMeetingProcessingRuleListResponse represents the response from listing MJX meeting processing rules
type MjxMeetingProcessingRuleListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []MjxMeetingProcessingRule `json:"objects"`
}
