package config

// TelehealthProfile represents a telehealth profile configuration
type TelehealthProfile struct {
	ID                                    int    `json:"id,omitempty"`
	Name                                  string `json:"name"`
	Description                           string `json:"description,omitempty"`
	UUID                                  string `json:"uuid"`
	TelehealthCallDomain                  string `json:"telehealth_call_domain"`
	TelehealthIntegrationBaseURL          string `json:"telehealth_integration_base_url"`
	TelehealthIntegrationOauth2BaseAPIURL string `json:"telehealth_integration_oauth2_base_api_url"`
	InfinityWebappServerBaseURL           string `json:"infinity_webapp_server_base_url"`
	EpicPatientAppClientID                string `json:"epic_patient_app_client_id"`
	EpicPatientAppClientSecret            string `json:"epic_patient_app_client_secret,omitempty"`
	EpicProviderAppClientID               string `json:"epic_provider_app_client_id"`
	EpicProviderAppClientSecret           string `json:"epic_provider_app_client_secret,omitempty"`
	EpicBackendOauth2AppClientID          string `json:"epic_backend_oauth2_app_client_id,omitempty"`
	EpicBackendOauth2AppPrivateKey        string `json:"epic_backend_oauth2_app_private_key,omitempty"`
	EpicBackendOauth2AppTokenEndpointURL  string `json:"epic_backend_oauth2_app_token_endpoint_url,omitempty"`
	EpicEncryptionAlgorithm               string `json:"epic_encryption_algorithm"`
	EpicEncryptionKey                     string `json:"epic_encryption_key,omitempty"`
	EpicEncryptionKeyType                 string `json:"epic_encryption_key_type"`
	PatientOauth2RedirectURL              string `json:"patient_oauth2_redirect_url"`
	ProviderOauth2RedirectURL             string `json:"provider_oauth2_redirect_url"`
	ServiceNameTemplate                   string `json:"service_name_template"`
	PatientAliasTemplate                  string `json:"patient_alias_template"`
	ProviderAliasTemplate                 string `json:"provider_alias_template"`
	PatientDisplayNameTemplate            string `json:"patient_display_name_template"`
	ProviderDisplayNameTemplate           string `json:"provider_display_name_template"`
	PatientWebJoinLinkTemplate            string `json:"patient_web_join_link_template"`
	ProviderWebJoinLinkTemplate           string `json:"provider_web_join_link_template"`
	LaunchErrorWebTemplate                string `json:"launch_error_web_template,omitempty"`
	ResourceURI                           string `json:"resource_uri,omitempty"`
}

// TelehealthProfileCreateRequest represents a request to create a telehealth profile
type TelehealthProfileCreateRequest struct {
	Name                                  string `json:"name"`
	Description                           string `json:"description,omitempty"`
	UUID                                  string `json:"uuid"`
	TelehealthCallDomain                  string `json:"telehealth_call_domain"`
	TelehealthIntegrationBaseURL          string `json:"telehealth_integration_base_url"`
	TelehealthIntegrationOauth2BaseAPIURL string `json:"telehealth_integration_oauth2_base_api_url"`
	InfinityWebappServerBaseURL           string `json:"infinity_webapp_server_base_url"`
	EpicPatientAppClientID                string `json:"epic_patient_app_client_id"`
	EpicPatientAppClientSecret            string `json:"epic_patient_app_client_secret,omitempty"`
	EpicProviderAppClientID               string `json:"epic_provider_app_client_id"`
	EpicProviderAppClientSecret           string `json:"epic_provider_app_client_secret,omitempty"`
	EpicBackendOauth2AppClientID          string `json:"epic_backend_oauth2_app_client_id,omitempty"`
	EpicBackendOauth2AppPrivateKey        string `json:"epic_backend_oauth2_app_private_key,omitempty"`
	EpicBackendOauth2AppTokenEndpointURL  string `json:"epic_backend_oauth2_app_token_endpoint_url,omitempty"`
	EpicEncryptionAlgorithm               string `json:"epic_encryption_algorithm"`
	EpicEncryptionKey                     string `json:"epic_encryption_key,omitempty"`
	EpicEncryptionKeyType                 string `json:"epic_encryption_key_type"`
	PatientOauth2RedirectURL              string `json:"patient_oauth2_redirect_url"`
	ProviderOauth2RedirectURL             string `json:"provider_oauth2_redirect_url"`
	ServiceNameTemplate                   string `json:"service_name_template"`
	PatientAliasTemplate                  string `json:"patient_alias_template"`
	ProviderAliasTemplate                 string `json:"provider_alias_template"`
	PatientDisplayNameTemplate            string `json:"patient_display_name_template"`
	ProviderDisplayNameTemplate           string `json:"provider_display_name_template"`
	PatientWebJoinLinkTemplate            string `json:"patient_web_join_link_template"`
	ProviderWebJoinLinkTemplate           string `json:"provider_web_join_link_template"`
	LaunchErrorWebTemplate                string `json:"launch_error_web_template,omitempty"`
}

// TelehealthProfileUpdateRequest represents a request to update a telehealth profile
type TelehealthProfileUpdateRequest struct {
	Name                                  string `json:"name,omitempty"`
	Description                           string `json:"description,omitempty"`
	UUID                                  string `json:"uuid,omitempty"`
	TelehealthCallDomain                  string `json:"telehealth_call_domain,omitempty"`
	TelehealthIntegrationBaseURL          string `json:"telehealth_integration_base_url,omitempty"`
	TelehealthIntegrationOauth2BaseAPIURL string `json:"telehealth_integration_oauth2_base_api_url,omitempty"`
	InfinityWebappServerBaseURL           string `json:"infinity_webapp_server_base_url,omitempty"`
	EpicPatientAppClientID                string `json:"epic_patient_app_client_id,omitempty"`
	EpicPatientAppClientSecret            string `json:"epic_patient_app_client_secret,omitempty"`
	EpicProviderAppClientID               string `json:"epic_provider_app_client_id,omitempty"`
	EpicProviderAppClientSecret           string `json:"epic_provider_app_client_secret,omitempty"`
	EpicBackendOauth2AppClientID          string `json:"epic_backend_oauth2_app_client_id,omitempty"`
	EpicBackendOauth2AppPrivateKey        string `json:"epic_backend_oauth2_app_private_key,omitempty"`
	EpicBackendOauth2AppTokenEndpointURL  string `json:"epic_backend_oauth2_app_token_endpoint_url,omitempty"`
	EpicEncryptionAlgorithm               string `json:"epic_encryption_algorithm,omitempty"`
	EpicEncryptionKey                     string `json:"epic_encryption_key,omitempty"`
	EpicEncryptionKeyType                 string `json:"epic_encryption_key_type,omitempty"`
	PatientOauth2RedirectURL              string `json:"patient_oauth2_redirect_url,omitempty"`
	ProviderOauth2RedirectURL             string `json:"provider_oauth2_redirect_url,omitempty"`
	ServiceNameTemplate                   string `json:"service_name_template,omitempty"`
	PatientAliasTemplate                  string `json:"patient_alias_template,omitempty"`
	ProviderAliasTemplate                 string `json:"provider_alias_template,omitempty"`
	PatientDisplayNameTemplate            string `json:"patient_display_name_template,omitempty"`
	ProviderDisplayNameTemplate           string `json:"provider_display_name_template,omitempty"`
	PatientWebJoinLinkTemplate            string `json:"patient_web_join_link_template,omitempty"`
	ProviderWebJoinLinkTemplate           string `json:"provider_web_join_link_template,omitempty"`
	LaunchErrorWebTemplate                string `json:"launch_error_web_template,omitempty"`
}

// TelehealthProfileListResponse represents the response from listing telehealth profiles
type TelehealthProfileListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []TelehealthProfile `json:"objects"`
}
