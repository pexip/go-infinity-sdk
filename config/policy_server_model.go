package config

// PolicyServer represents a policy server configuration
type PolicyServer struct {
	ID                                  int    `json:"id,omitempty"`
	Name                                string `json:"name"`
	Description                         string `json:"description,omitempty"`
	URL                                 string `json:"url,omitempty"`
	Username                            string `json:"username,omitempty"`
	Password                            string `json:"password,omitempty"`
	EnableServiceLookup                 bool   `json:"enable_service_lookup"`
	EnableParticipantLookup             bool   `json:"enable_participant_lookup"`
	EnableRegistrationLookup            bool   `json:"enable_registration_lookup"`
	EnableDirectoryLookup               bool   `json:"enable_directory_lookup"`
	EnableAvatarLookup                  bool   `json:"enable_avatar_lookup"`
	EnableMediaLocationLookup           bool   `json:"enable_media_location_lookup"`
	EnableInternalServicePolicy         bool   `json:"enable_internal_service_policy"`
	EnableInternalParticipantPolicy     bool   `json:"enable_internal_participant_policy"`
	EnableInternalMediaLocationPolicy   bool   `json:"enable_internal_media_location_policy"`
	InternalServicePolicyTemplate       string `json:"internal_service_policy_template,omitempty"`
	InternalParticipantPolicyTemplate   string `json:"internal_participant_policy_template,omitempty"`
	InternalMediaLocationPolicyTemplate string `json:"internal_media_location_policy_template,omitempty"`
	PreferLocalAvatarConfiguration      bool   `json:"prefer_local_avatar_configuration"`
	ResourceURI                         string `json:"resource_uri,omitempty"`
}

// PolicyServerCreateRequest represents a request to create a policy server
type PolicyServerCreateRequest struct {
	Name                                string `json:"name"`
	Description                         string `json:"description,omitempty"`
	URL                                 string `json:"url,omitempty"`
	Username                            string `json:"username,omitempty"`
	Password                            string `json:"password,omitempty"`
	EnableServiceLookup                 bool   `json:"enable_service_lookup"`
	EnableParticipantLookup             bool   `json:"enable_participant_lookup"`
	EnableRegistrationLookup            bool   `json:"enable_registration_lookup"`
	EnableDirectoryLookup               bool   `json:"enable_directory_lookup"`
	EnableAvatarLookup                  bool   `json:"enable_avatar_lookup"`
	EnableMediaLocationLookup           bool   `json:"enable_media_location_lookup"`
	EnableInternalServicePolicy         bool   `json:"enable_internal_service_policy"`
	EnableInternalParticipantPolicy     bool   `json:"enable_internal_participant_policy"`
	EnableInternalMediaLocationPolicy   bool   `json:"enable_internal_media_location_policy"`
	InternalServicePolicyTemplate       string `json:"internal_service_policy_template,omitempty"`
	InternalParticipantPolicyTemplate   string `json:"internal_participant_policy_template,omitempty"`
	InternalMediaLocationPolicyTemplate string `json:"internal_media_location_policy_template,omitempty"`
	PreferLocalAvatarConfiguration      bool   `json:"prefer_local_avatar_configuration"`
}

// PolicyServerUpdateRequest represents a request to update a policy server
type PolicyServerUpdateRequest struct {
	Name                                string `json:"name,omitempty"`
	Description                         string `json:"description,omitempty"`
	URL                                 string `json:"url,omitempty"`
	Username                            string `json:"username,omitempty"`
	Password                            string `json:"password,omitempty"`
	EnableServiceLookup                 *bool  `json:"enable_service_lookup,omitempty"`
	EnableParticipantLookup             *bool  `json:"enable_participant_lookup,omitempty"`
	EnableRegistrationLookup            *bool  `json:"enable_registration_lookup,omitempty"`
	EnableDirectoryLookup               *bool  `json:"enable_directory_lookup,omitempty"`
	EnableAvatarLookup                  *bool  `json:"enable_avatar_lookup,omitempty"`
	EnableMediaLocationLookup           *bool  `json:"enable_media_location_lookup,omitempty"`
	EnableInternalServicePolicy         *bool  `json:"enable_internal_service_policy,omitempty"`
	EnableInternalParticipantPolicy     *bool  `json:"enable_internal_participant_policy,omitempty"`
	EnableInternalMediaLocationPolicy   *bool  `json:"enable_internal_media_location_policy,omitempty"`
	InternalServicePolicyTemplate       string `json:"internal_service_policy_template,omitempty"`
	InternalParticipantPolicyTemplate   string `json:"internal_participant_policy_template,omitempty"`
	InternalMediaLocationPolicyTemplate string `json:"internal_media_location_policy_template,omitempty"`
	PreferLocalAvatarConfiguration      *bool  `json:"prefer_local_avatar_configuration,omitempty"`
}

// PolicyServerListResponse represents the response from listing policy servers
type PolicyServerListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []PolicyServer `json:"objects"`
}
