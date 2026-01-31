/*
 * SPDX-FileCopyrightText: 2025 Pexip AS
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// MsExchangeConnector represents a Microsoft Exchange connector configuration
type MsExchangeConnector struct {
	ID                         int     `json:"id,omitempty"`                         // The primary key
	Name                       string  `json:"name"`                                 // The name used to refer to the Secure Scheduler for Exchange Integration. Maximum length: 250 characters.
	Description                string  `json:"description,omitempty"`                // An optional description of the Secure Scheduler for Exchange Integration. Maximum length: 250 characters. Default: ""
	RoomMailboxEmailAddress    *string `json:"room_mailbox_email_address,omitempty"` // The email address of the equipment resource or room resource that is to be used by the scheduling service. Maximum length: 100 characters.
	RoomMailboxName            string  `json:"room_mailbox_name,omitempty"`          // The name of the equipment resource or room resource that is to be used by the scheduling service. Maximum length: 250 characters. Default: ""
	URL                        string  `json:"url,omitempty"`                        // The URL used to connect to Exchange Web Services (EWS) on the Exchange server. Maximum length: 255 characters. Default: ""
	Username                   string  `json:"username,omitempty"`                   // The username of the service account to be used by the scheduling service. Maximum length: 100 characters. Default: ""
	Password                   string  `json:"password,omitempty"`                   // The password of the service account to be used by the scheduling service. Maximum length: 100 characters. Default: ""
	AuthenticationMethod       string  `json:"authentication_method"`                // The method used to authenticate to Exchange. Valid values: BASIC, NTLM, KERBEROS, OAUTH, APP_PERM. Default: BASIC
	AuthProvider               string  `json:"auth_provider"`                        // The method by which users will sign into the Outlook add-in. Valid values: ADFS, AZURE. Default: ADFS
	UUID                       string  `json:"uuid"`                                 // The unique identifier of the Secure Scheduler for Exchange Integration.
	ScheduledAliasPrefix       *string `json:"scheduled_alias_prefix,omitempty"`     // The prefix to use when generating aliases for scheduled conferences. Minimum length: 1 characters. Maximum length: 8 characters.
	ScheduledAliasDomain       string  `json:"scheduled_alias_domain,omitempty"`     // The domain to use when generating aliases for scheduled conferences. Maximum length: 192 characters. Default: ""
	ScheduledAliasSuffixLength int     `json:"scheduled_alias_suffix_length"`        // The length of the random number suffix part of aliases used for scheduled conferences. Range: 5 to 15. Default: 6
	MeetingBufferBefore        int     `json:"meeting_buffer_before"`                // The number of minutes before the meeting's scheduled start time that participants will be able to join the VMR. Range: 0 to 180. Default: 30
	MeetingBufferAfter         int     `json:"meeting_buffer_after"`                 // The number of minutes after the meeting's scheduled end of a conference participants will be able to join the VMR. Range: 0 to 180. Default: 60
	EnableDynamicVmrs          bool    `json:"enable_dynamic_vmrs"`                  // Enable this option to allow Outlook users to schedule meetings in single-use (randomly generated) VMRs. Default: true
	EnablePersonalVmrs         bool    `json:"enable_personal_vmrs"`                 // Enable this option to allow Outlook users to schedule meetings in their personal VMRs. Default: false
	AllowNewUsers              bool    `json:"allow_new_users"`                      // Disable this option to allow only those users with an existing User record to access the Outlook add-in. Default: true
	DisableProxy               bool    `json:"disable_proxy"`                        // Disable the usage of any web proxy which may have been configured on the Management Node by this Secure Scheduler for Exchange Integration. Default: false
	UseCustomAddInSources      bool    `json:"use_custom_add_in_sources"`            // Enable this to specify custom locations to serve add-in JavaScript and CSS from. This can be used to support offline deployments. Default: false
	EnableAddinDebugLogs       bool    `json:"enable_addin_debug_logs"`              // Enable this option to view debug logs within the add-in side pane. Note that these logs will appear for all users of this add-in. Default: false
	// OAuth fields
	OauthClientID      *string `json:"oauth_client_id,omitempty"`      // The Application ID which was generated when creating an App Registration in Azure Active Directory
	OauthClientSecret  string  `json:"oauth_client_secret,omitempty"`  // The OAuth Client Secret which was generated when creating an App Registration in Microsoft Entra. Default: ""
	OauthAuthEndpoint  string  `json:"oauth_auth_endpoint,omitempty"`  // The URI of the OAuth authorization endpoint. This should be copied from the 'Endpoints' section in Azure Active Directory App Registrations. Maximum length: 255 characters. Default: ""
	OauthTokenEndpoint string  `json:"oauth_token_endpoint,omitempty"` // The URI of the OAuth token endpoint. This should be copied from the 'Endpoints' section in Azure Active Directory App Registrations. Maximum length: 255 characters. Default: ""
	OauthRedirectURI   string  `json:"oauth_redirect_uri,omitempty"`   // The redirect URI you entered when creating an App Registration in Azure Active Directory. It should be in the format 'https://[Management Node Address]/admin/platform/msexchangeconnector/oauth_redirect/'. Maximum length: 255 characters. Default: ""
	OauthRefreshToken  string  `json:"oauth_refresh_token,omitempty"`  // The OAuth refresh token which is obtained after successfully signing in via the OAuth flow. Maximum length: 4096 characters. Default: ""
	OauthState         *string `json:"oauth_state,omitempty"`          // A unique state which is used during the OAuth sign-in flow.
	// Kerberos fields
	KerberosRealm                  string `json:"kerberos_realm,omitempty"`            // The Kerberos Realm, which is usually your domain in upper-case. Maximum length: 250 characters. Default: ""
	KerberosKdc                    string `json:"kerberos_kdc,omitempty"`              // The address of the Kerberos key distribution center (KDC). Maximum length: 255 characters. Default: ""
	KerberosKdcHttpsProxy          string `json:"kerberos_kdc_https_proxy,omitempty"`  // The URL of the Kerberos key distribution center (KDC) HTTPS proxy. Maximum length: 255 characters. Default: ""
	KerberosExchangeSpn            string `json:"kerberos_exchange_spn,omitempty"`     // The Exchange Service Principal Name (SPN). Maximum length: 255 characters. Default: ""
	KerberosEnableTls              bool   `json:"kerberos_enable_tls"`                 // If enabled, all communication to the KDC will go through an HTTPS proxy and all traffic to the KDC will be encrypted using TLS. Default: true
	KerberosAuthEveryRequest       bool   `json:"kerberos_auth_every_request"`         // When Kerberos authentication is enabled, send a Kerberos Authorization header in every request to the Exchange server. Default: false
	KerberosVerifyTlsUsingCustomCa bool   `json:"kerberos_verify_tls_using_custom_ca"` // If enabled, use the configured Root Trust CA Certificates to verify the KDC HTTPS proxy SSL certificate. If disabled, the HTTPS proxy SSL certificate is verified using the system-wide default set of trusted certificates. Default: false
	// Add-in fields
	AddinServerDomain           string  `json:"addin_server_domain"`                        // The FQDN of the reverse proxy or Conferencing Node that provides the add-in content. The FQDN must have a valid certificate. Maximum length: 192 characters. Default: ""
	AddinDisplayName            string  `json:"addin_display_name"`                         // The display name of the add-in. Maximum length: 250 characters. Default: "Pexip Scheduling Service"
	AddinDescription            string  `json:"addin_description"`                          // The description of the add-in. Maximum length: 250 characters. Default: "Turns meetings into Pexip meetings"
	AddinProviderName           string  `json:"addin_provider_name"`                        // The name of the organization which provides the add-in. Maximum length: 250 characters. Default: "Pexip"
	AddinButtonLabel            string  `json:"addin_button_label"`                         // The label for the add-in button on desktop clients. Maximum length: 250 characters. Default: "Create a Pexip meeting"
	AddinGroupLabel             string  `json:"addin_group_label"`                          // The name of the group in which to place the add-in button on desktop clients. Maximum length: 250 characters. Default: "Pexip meeting"
	AddinSupertipTitle          string  `json:"addin_supertip_title"`                       // The title of the supertip help text for the add-in button on desktop clients. Maximum length: 250 characters. Default: "Makes this a Pexip meeting"
	AddinSupertipDescription    string  `json:"addin_supertip_description"`                 // The text of the supertip for the add-in button on desktop clients. Maximum length: 250 characters. Default: "Turns this meeting into an audio or video conference hosted in a Pexip VMR. The meeting is not scheduled until you select Send."
	AddinApplicationID          *string `json:"addin_application_id,omitempty"`             // The Application (client) ID which was generated when creating the App Registration in Microsoft Entra for add-in authentication.
	AddinAuthorityURL           string  `json:"addin_authority_url,omitempty"`              // The Authority URL copied from the App Registration created in Microsoft Entra for add-in authentication. Maximum length: 255 characters. Default: ""
	AddinOidcMetadataURL        string  `json:"addin_oidc_metadata_url,omitempty"`          // The OpenID Connect metadata document copied from the App Registration created in Microsoft Entra for add-in authentication. Maximum length: 255 characters. Default: ""
	AddinAuthenticationMethod   string  `json:"addin_authentication_method"`                // The type of token the Outlook add-in uses to authenticate to Pexip. Valid values: EXCHANGE_USER_ID_TOKEN, SSO_TOKEN, NAA_TOKEN. Default: EXCHANGE_USER_ID_TOKEN
	AddinNaaWebApiApplicationID *string `json:"addin_naa_web_api_application_id,omitempty"` // The Application (client) ID for the NAA Web API which was generated when creating the App Registration in Microsoft Entra.
	// Add-in pane fields
	AddinPaneTitle                                   string `json:"addin_pane_title"`                                        // The title of the add-in on the side pane. Maximum length: 250 characters. Default: "Add a VMR"
	AddinPaneDescription                             string `json:"addin_pane_description"`                                  // The description of the add-in on the side pane. Maximum length: 250 characters. Default: "This assigns a Virtual Meeting Room for your meeting"
	AddinPaneButtonTitle                             string `json:"addin_pane_button_title"`                                 // The label of the button on the side pane used to add a single-use VMR. Maximum length: 250 characters. Default: "Add a Single-use VMR"
	AddinPaneSuccessHeading                          string `json:"addin_pane_success_heading"`                              // The heading that appears on the side pane when when an alias has been obtained successfully from the Management Node. Maximum length: 250 characters. Default: "Success"
	AddinPaneSuccessMessage                          string `json:"addin_pane_success_message"`                              // The message that appears on the side pane when when an alias has been obtained successfully from the Management Node. Maximum length: 250 characters. Default: "This meeting is now set up to be hosted as an audio or video conference in a Virtual Meeting Room. Please note this conference is not scheduled until you select Send."
	AddinPaneAlreadyVideoMeetingHeading              string `json:"addin_pane_already_video_meeting_heading"`                // The heading that appears on the side pane when the add-in is activated after an alias has already been obtained for the meeting. Maximum length: 250 characters. Default: "VMR already assigned"
	AddinPaneAlreadyVideoMeetingMessage              string `json:"addin_pane_already_video_meeting_message"`                // The message that appears on the side pane when the add-in is activated after an alias has already been obtained for the meeting. Maximum length: 250 characters. Default: "It looks like this meeting has already been set up to be hosted in a Virtual Meeting Room. If this is a new meeting, select Send to schedule the conference."
	AddinPaneGeneralErrorHeading                     string `json:"addin_pane_general_error_heading"`                        // The heading that appears on the side pane when an error occurs trying to add the joining instructions. Maximum length: 250 characters. Default: "Error"
	AddinPaneGeneralErrorMessage                     string `json:"addin_pane_general_error_message"`                        // The message that appears on the side pane when an error occurs trying to add the joining instructions of the single-use VMR. Maximum length: 250 characters. Default: "There was a problem adding the joining instructions. Please try again."
	AddinPaneManagementNodeDownHeading               string `json:"addin_pane_management_node_down_heading"`                 // The heading that appears on the side pane when the Management Node cannot be contacted to obtain an alias. Maximum length: 250 characters. Default: "Cannot assign a VMR right now"
	AddinPaneManagementNodeDownMessage               string `json:"addin_pane_management_node_down_message"`                 // The message that appears on the side pane when the Management Node cannot be contacted to obtain an alias. Maximum length: 250 characters. Default: "Sorry, we are unable to assign a Virtual Meeting Room at this time. Select Send to schedule the meeting, and all attendees will be sent joining instructions later."
	AddinPanePersonalVmrAddButton                    string `json:"addin_pane_personal_vmr_add_button"`                      // The label of the button on the side pane used to add a personal VMR. Maximum length: 250 characters. Default: "Add a Personal VMR"
	AddinPanePersonalVmrSignInButton                 string `json:"addin_pane_personal_vmr_sign_in_button"`                  // The label of the button on the side pane requesting users to sign in to obtain a list of their personal VMRs. Maximum length: 250 characters. Default: "Sign In"
	AddinPanePersonalVmrSelectMessage                string `json:"addin_pane_personal_vmr_select_message"`                  // The message that appears on the side pane requesting users to select a personal VMR to use for the meeting. Maximum length: 250 characters. Default: "Select the VMR you want to add to the meeting"
	AddinPanePersonalVmrNoneMessage                  string `json:"addin_pane_personal_vmr_none_message"`                    // The message that appears on the side pane when the user has no personal VMRs. Maximum length: 250 characters. Default: "You do not have any personal VMRs"
	AddinPanePersonalVmrErrorGettingMessage          string `json:"addin_pane_personal_vmr_error_getting_message"`           // The message that appears on the side pane when an error occurs trying to obtain a list of the user's personal VMRs. Maximum length: 250 characters. Default: "There was a problem getting your personal VMRs. Please try again."
	AddinPanePersonalVmrErrorSigningInMessage        string `json:"addin_pane_personal_vmr_error_signing_in_message"`        // The message that appears on the side pane when an error occurs trying to sign the user in. Maximum length: 250 characters. Default: "There was a problem signing you in. Please try again."
	AddinPanePersonalVmrErrorInsertingMeetingMessage string `json:"addin_pane_personal_vmr_error_inserting_meeting_message"` // The message that appears on the side pane when an error occurs trying to add the personal VMR details to the meeting. Maximum length: 250 characters. Default: "There was a problem adding the joining instructions. Please try again."
	// Personal VMR OAuth fields
	PersonalVmrOauthClientID                   *string `json:"personal_vmr_oauth_client_id,omitempty"`                     // The client ID of the OAuth application used to authenticate users when signing in to the Outlook add-in. Default: "4189c2b4-92ca-416c-b7ea-bc3cfab3d0f0"
	PersonalVmrOauthClientSecret               string  `json:"personal_vmr_oauth_client_secret,omitempty"`                 // The client secret of the OAuth application created for signing in users in the Outlook add-in. Default: ""
	PersonalVmrOauthAuthEndpoint               string  `json:"personal_vmr_oauth_auth_endpoint,omitempty"`                 // The authorization URI of the OAuth application used to authenticate users when signing in to the Outlook add-in. Maximum length: 255 characters. Default: ""
	PersonalVmrOauthTokenEndpoint              string  `json:"personal_vmr_oauth_token_endpoint,omitempty"`                // The token URI of the OAuth application used to authenticate users when signing in to the Outlook add-in. Maximum length: 255 characters. Default: ""
	PersonalVmrAdfsRelyingPartyTrustIdentifier string  `json:"personal_vmr_adfs_relying_party_trust_identifier,omitempty"` // The URL which identifies the OAuth 2.0 resource on AD FS. Maximum length: 255 characters. Default: ""
	// Template fields
	MeetingInstructionsTemplate         string `json:"meeting_instructions_template,omitempty"`           // A Jinja2 template that is used to generate the instructions added by the scheduling service to the body of the meeting request when a single-use VMR is being used. Maximum length: 12288 characters.
	PersonalVmrInstructionsTemplate     string `json:"personal_vmr_instructions_template,omitempty"`      // A Jinja2 template that is used to produce the joining instructions added by the scheduling service to the body of the meeting request when a personal VMR is being used. Maximum length: 12288 characters.
	PersonalVmrLocationTemplate         string `json:"personal_vmr_location_template,omitempty"`          // A Jinja2 template that is used to generate the text that will be inserted into the Location field of the meeting request when a personal VMR is being used. Maximum length: 12288 characters.
	PersonalVmrNameTemplate             string `json:"personal_vmr_name_template,omitempty"`              // A Jinja2 template that is used to generate the name of the personal VMR, as it appears on the button offered to users. Maximum length: 12288 characters. Default: "{{name}}"
	PersonalVmrDescriptionTemplate      string `json:"personal_vmr_description_template,omitempty"`       // A Jinja2 template that is used to generate the description of the personal VMR, shown to users when they hover over the button. Maximum length: 12288 characters. Default: "{{description}}"
	PlaceholderInstructionsTemplate     string `json:"placeholder_instructions_template,omitempty"`       // The text that is added by the scheduling service to email messages when the actual joining instructions cannot be obtained. Maximum length: 12288 characters.
	ConferenceNameTemplate              string `json:"conference_name_template,omitempty"`                // A Jinja2 template that is used to produce the name of scheduled conferences. Please note conference names must be unique so a random number may be appended if the name that is generated is already in use by another service. Maximum length: 12288 characters. Default: "{{subject}} ({{organizer_name}})"
	ConferenceDescriptionTemplate       string `json:"conference_description_template,omitempty"`         // A Jinja2 template that is used to produce the description of scheduled conferences. Maximum length: 12288 characters. Default: "Scheduled Conference booked by {{organizer_email}}"
	ConferenceSubjectTemplate           string `json:"conference_subject_template,omitempty"`             // A Jinja2 template that is used to produce the subject field of scheduled conferences. By default this will use the subject line of the meeting invitation but this field can be deleted or amended if you do not want the subject to be visible to administrators. Maximum length: 12288 characters. Default: "{{subject}}"
	ScheduledAliasDescriptionTemplate   string `json:"scheduled_alias_description_template,omitempty"`    // A Jinja2 template that is used to produce the description of scheduled conference aliases. Maximum length: 12288 characters. Default: "Scheduled Conference booked by {{organizer_email}}"
	AcceptNewSingleMeetingTemplate      string `json:"accept_new_single_meeting_template,omitempty"`      // A Jinja2 template that is used to produce the message sent to meeting organizers once the scheduling service successfully schedules a new single meeting. Maximum length: 12288 characters.
	AcceptNewRecurringSeriesTemplate    string `json:"accept_new_recurring_series_template,omitempty"`    // A Jinja2 template that is used to produce the message sent to meeting organizers once the scheduling service successfully schedules a new recurring meeting. Maximum length: 12288 characters.
	AcceptEditedSingleMeetingTemplate   string `json:"accept_edited_single_meeting_template,omitempty"`   // A Jinja2 template that is used to produce the message sent to meeting organizers once the scheduling service successfully schedules an edited single meeting. Maximum length: 12288 characters.
	AcceptEditedRecurringSeriesTemplate string `json:"accept_edited_recurring_series_template,omitempty"` // A Jinja2 template that is used to produce the message sent to meeting organizers once the scheduling service successfully schedules an edited recurring meeting. Maximum length: 12288 characters.
	AcceptEditedOccurrenceTemplate      string `json:"accept_edited_occurrence_template,omitempty"`       // A Jinja2 template that is used to produce the message sent to meeting organizers once the scheduling service successfully schedules an edited occurrence in a recurring series. Maximum length: 12288 characters.
	RejectGeneralErrorTemplate          string `json:"reject_general_error_template,omitempty"`           // A Jinja2 template that is used to produce the message sent to meeting organizers when the scheduling service fails to schedule a meeting because a general error occurred. Maximum length: 12288 characters.
	RejectAliasConflictTemplate         string `json:"reject_alias_conflict_template,omitempty"`          // A Jinja2 template that is used to produce the message sent to meeting organizers when the scheduling service fails to schedule a meeting because the alias conflicts with an existing alias. Maximum length: 12288 characters.
	RejectAliasDeletedTemplate          string `json:"reject_alias_deleted_template,omitempty"`           // The text that is sent to meeting organizers when the scheduling service fails to schedule a meeting because the alias for this meeting has been deleted. Maximum length: 12288 characters.
	RejectInvalidAliasIDTemplate        string `json:"reject_invalid_alias_id_template,omitempty"`        // The text that is sent to meeting organizers when the scheduling service fails to schedule a meeting because the alias ID in the meeting email is invalid. Maximum length: 12288 characters.
	RejectSingleMeetingPast             string `json:"reject_single_meeting_past,omitempty"`              // The text that is sent to meeting organizers when the scheduling service fails to schedule a meeting because it occurs in the past. Maximum length: 12288 characters.
	RejectRecurringSeriesPastTemplate   string `json:"reject_recurring_series_past_template,omitempty"`   // The text that is sent to meeting organizers when the scheduling service fails to schedule a recurring meeting because all occurrences occur in the past. Maximum length: 12288 characters.
	// JavaScript and CSS URLs
	OfficeJsURL                  string `json:"office_js_url"`                              // The URL used to download the Office.js JavaScript library. Maximum length: 255 characters. Default: "https://appsforoffice.microsoft.com/lib/1/hosted/office.js"
	MicrosoftFabricURL           string `json:"microsoft_fabric_url"`                       // The URL used to download the Microsoft Fabric CSS. Maximum length: 255 characters. Default: "https://appsforoffice.microsoft.com/fabric/1.0/fabric.min.css"
	MicrosoftFabricComponentsURL string `json:"microsoft_fabric_components_url"`            // The URL used to download the Microsoft Fabric Components CSS. Maximum length: 255 characters. Default: "https://appsforoffice.microsoft.com/fabric/1.0/fabric.components.min.css"
	AdditionalAddInScriptSources string `json:"additional_add_in_script_sources,omitempty"` // Optionally specify additional URLs to download JavaScript script files. Each URL must be entered on a separate line. Maximum length: 4096 characters. Default: ""
	// Related resources
	Domains                   *[]ExchangeDomain `json:"domains,omitempty"`                      // The Exchange Metadata Domains / URLs associated with this Secure Scheduler for Exchange Integration.
	HostIdentityProviderGroup *string           `json:"host_identity_provider_group,omitempty"` // The set of Identity Providers to use if participants are required to authenticate in order to join the scheduled conference. If this is blank, participants will not be required to authenticate.
	IvrTheme                  *string           `json:"ivr_theme,omitempty"`                    // The theme for use with this service.
	NonIdpParticipants        string            `json:"non_idp_participants,omitempty"`         // Determines whether participants attempting to join from devices other than the Infinity Connect apps (for example, SIP or H.323 endpoints) are permitted to join the conference when authentication is required. Disallow all: these devices may not join the conference. Allow if trusted: these devices may join the conference if they are locally registered. Default: "disallow_all"
	// Read-only fields
	PrivateKey  *string `json:"private_key,omitempty"`  // The private key used by this Secure Scheduler for Exchange Integration. Maximum length: 12288 characters.
	PublicKey   string  `json:"public_key,omitempty"`   // The public key used by this Secure Scheduler for Exchange Integration. Maximum length: 12288 characters.
	ResourceURI string  `json:"resource_uri,omitempty"` // The URI that identifies this resource.
}

// MsExchangeConnectorCreateRequest represents a request to create a Microsoft Exchange connector
type MsExchangeConnectorCreateRequest struct {
	Name                       string  `json:"name"`
	Description                string  `json:"description,omitempty"`
	RoomMailboxEmailAddress    *string `json:"room_mailbox_email_address,omitempty"`
	RoomMailboxName            string  `json:"room_mailbox_name,omitempty"`
	URL                        string  `json:"url,omitempty"`
	Username                   string  `json:"username,omitempty"`
	Password                   string  `json:"password,omitempty"`
	AuthenticationMethod       string  `json:"authentication_method"`
	AuthProvider               string  `json:"auth_provider"`
	UUID                       string  `json:"uuid,omitempty"`
	ScheduledAliasPrefix       *string `json:"scheduled_alias_prefix,omitempty"`
	ScheduledAliasDomain       string  `json:"scheduled_alias_domain,omitempty"`
	ScheduledAliasSuffixLength int     `json:"scheduled_alias_suffix_length"`
	MeetingBufferBefore        int     `json:"meeting_buffer_before"`
	MeetingBufferAfter         int     `json:"meeting_buffer_after"`
	EnableDynamicVmrs          bool    `json:"enable_dynamic_vmrs"`
	EnablePersonalVmrs         bool    `json:"enable_personal_vmrs"`
	AllowNewUsers              bool    `json:"allow_new_users"`
	DisableProxy               bool    `json:"disable_proxy"`
	UseCustomAddInSources      bool    `json:"use_custom_add_in_sources"`
	EnableAddinDebugLogs       bool    `json:"enable_addin_debug_logs"`
	// OAuth fields
	OauthClientID      *string `json:"oauth_client_id,omitempty"`
	OauthClientSecret  string  `json:"oauth_client_secret,omitempty"`
	OauthAuthEndpoint  string  `json:"oauth_auth_endpoint,omitempty"`
	OauthTokenEndpoint string  `json:"oauth_token_endpoint,omitempty"`
	OauthRedirectURI   string  `json:"oauth_redirect_uri,omitempty"`
	OauthRefreshToken  string  `json:"oauth_refresh_token,omitempty"`
	// Kerberos fields
	KerberosRealm                  string `json:"kerberos_realm,omitempty"`
	KerberosKdc                    string `json:"kerberos_kdc,omitempty"`
	KerberosKdcHttpsProxy          string `json:"kerberos_kdc_https_proxy,omitempty"`
	KerberosExchangeSpn            string `json:"kerberos_exchange_spn,omitempty"`
	KerberosEnableTls              bool   `json:"kerberos_enable_tls"`
	KerberosAuthEveryRequest       bool   `json:"kerberos_auth_every_request"`
	KerberosVerifyTlsUsingCustomCa bool   `json:"kerberos_verify_tls_using_custom_ca"`
	// Add-in fields
	AddinServerDomain           string  `json:"addin_server_domain"`
	AddinDisplayName            string  `json:"addin_display_name"`
	AddinDescription            string  `json:"addin_description"`
	AddinProviderName           string  `json:"addin_provider_name"`
	AddinButtonLabel            string  `json:"addin_button_label"`
	AddinGroupLabel             string  `json:"addin_group_label"`
	AddinSupertipTitle          string  `json:"addin_supertip_title"`
	AddinSupertipDescription    string  `json:"addin_supertip_description"`
	AddinApplicationID          *string `json:"addin_application_id,omitempty"`
	AddinAuthorityURL           string  `json:"addin_authority_url,omitempty"`
	AddinOidcMetadataURL        string  `json:"addin_oidc_metadata_url,omitempty"`
	AddinAuthenticationMethod   string  `json:"addin_authentication_method"`
	AddinNaaWebApiApplicationID *string `json:"addin_naa_web_api_application_id,omitempty"`
	// Add-in pane fields
	AddinPaneTitle                                   string `json:"addin_pane_title"`
	AddinPaneDescription                             string `json:"addin_pane_description"`
	AddinPaneButtonTitle                             string `json:"addin_pane_button_title"`
	AddinPaneSuccessHeading                          string `json:"addin_pane_success_heading"`
	AddinPaneSuccessMessage                          string `json:"addin_pane_success_message"`
	AddinPaneAlreadyVideoMeetingHeading              string `json:"addin_pane_already_video_meeting_heading"`
	AddinPaneAlreadyVideoMeetingMessage              string `json:"addin_pane_already_video_meeting_message"`
	AddinPaneGeneralErrorHeading                     string `json:"addin_pane_general_error_heading"`
	AddinPaneGeneralErrorMessage                     string `json:"addin_pane_general_error_message"`
	AddinPaneManagementNodeDownHeading               string `json:"addin_pane_management_node_down_heading"`
	AddinPaneManagementNodeDownMessage               string `json:"addin_pane_management_node_down_message"`
	AddinPanePersonalVmrAddButton                    string `json:"addin_pane_personal_vmr_add_button"`
	AddinPanePersonalVmrSignInButton                 string `json:"addin_pane_personal_vmr_sign_in_button"`
	AddinPanePersonalVmrSelectMessage                string `json:"addin_pane_personal_vmr_select_message"`
	AddinPanePersonalVmrNoneMessage                  string `json:"addin_pane_personal_vmr_none_message"`
	AddinPanePersonalVmrErrorGettingMessage          string `json:"addin_pane_personal_vmr_error_getting_message"`
	AddinPanePersonalVmrErrorSigningInMessage        string `json:"addin_pane_personal_vmr_error_signing_in_message"`
	AddinPanePersonalVmrErrorInsertingMeetingMessage string `json:"addin_pane_personal_vmr_error_inserting_meeting_message"`
	// Personal VMR OAuth fields
	PersonalVmrOauthClientID                   *string `json:"personal_vmr_oauth_client_id,omitempty"`
	PersonalVmrOauthClientSecret               string  `json:"personal_vmr_oauth_client_secret,omitempty"`
	PersonalVmrOauthAuthEndpoint               string  `json:"personal_vmr_oauth_auth_endpoint,omitempty"`
	PersonalVmrOauthTokenEndpoint              string  `json:"personal_vmr_oauth_token_endpoint,omitempty"`
	PersonalVmrAdfsRelyingPartyTrustIdentifier string  `json:"personal_vmr_adfs_relying_party_trust_identifier,omitempty"`
	// Template fields
	MeetingInstructionsTemplate         string `json:"meeting_instructions_template,omitempty"`
	PersonalVmrInstructionsTemplate     string `json:"personal_vmr_instructions_template,omitempty"`
	PersonalVmrLocationTemplate         string `json:"personal_vmr_location_template,omitempty"`
	PersonalVmrNameTemplate             string `json:"personal_vmr_name_template,omitempty"`
	PersonalVmrDescriptionTemplate      string `json:"personal_vmr_description_template,omitempty"`
	PlaceholderInstructionsTemplate     string `json:"placeholder_instructions_template,omitempty"`
	ConferenceNameTemplate              string `json:"conference_name_template,omitempty"`
	ConferenceDescriptionTemplate       string `json:"conference_description_template,omitempty"`
	ConferenceSubjectTemplate           string `json:"conference_subject_template,omitempty"`
	ScheduledAliasDescriptionTemplate   string `json:"scheduled_alias_description_template,omitempty"`
	AcceptNewSingleMeetingTemplate      string `json:"accept_new_single_meeting_template,omitempty"`
	AcceptNewRecurringSeriesTemplate    string `json:"accept_new_recurring_series_template,omitempty"`
	AcceptEditedSingleMeetingTemplate   string `json:"accept_edited_single_meeting_template,omitempty"`
	AcceptEditedRecurringSeriesTemplate string `json:"accept_edited_recurring_series_template,omitempty"`
	AcceptEditedOccurrenceTemplate      string `json:"accept_edited_occurrence_template,omitempty"`
	RejectGeneralErrorTemplate          string `json:"reject_general_error_template,omitempty"`
	RejectAliasConflictTemplate         string `json:"reject_alias_conflict_template,omitempty"`
	RejectAliasDeletedTemplate          string `json:"reject_alias_deleted_template,omitempty"`
	RejectInvalidAliasIDTemplate        string `json:"reject_invalid_alias_id_template,omitempty"`
	RejectSingleMeetingPast             string `json:"reject_single_meeting_past,omitempty"`
	RejectRecurringSeriesPastTemplate   string `json:"reject_recurring_series_past_template,omitempty"`
	// JavaScript and CSS URLs
	OfficeJsURL                  string `json:"office_js_url"`
	MicrosoftFabricURL           string `json:"microsoft_fabric_url"`
	MicrosoftFabricComponentsURL string `json:"microsoft_fabric_components_url"`
	AdditionalAddInScriptSources string `json:"additional_add_in_script_sources,omitempty"`
	// Related resources
	Domains                   *[]string `json:"domains,omitempty"`
	HostIdentityProviderGroup *string   `json:"host_identity_provider_group,omitempty"`
	IvrTheme                  *string   `json:"ivr_theme,omitempty"`
	NonIdpParticipants        string    `json:"non_idp_participants,omitempty"`
}

// MsExchangeConnectorUpdateRequest represents a request to update a Microsoft Exchange connector
type MsExchangeConnectorUpdateRequest struct {
	Name                       string  `json:"name,omitempty"`
	Description                string  `json:"description,omitempty"`
	RoomMailboxEmailAddress    *string `json:"room_mailbox_email_address,omitempty"`
	RoomMailboxName            string  `json:"room_mailbox_name,omitempty"`
	URL                        string  `json:"url,omitempty"`
	Username                   string  `json:"username,omitempty"`
	Password                   string  `json:"password,omitempty"`
	AuthenticationMethod       string  `json:"authentication_method,omitempty"`
	AuthProvider               string  `json:"auth_provider,omitempty"`
	UUID                       string  `json:"uuid,omitempty"`
	ScheduledAliasPrefix       *string `json:"scheduled_alias_prefix,omitempty"`
	ScheduledAliasDomain       string  `json:"scheduled_alias_domain,omitempty"`
	ScheduledAliasSuffixLength *int    `json:"scheduled_alias_suffix_length,omitempty"`
	MeetingBufferBefore        *int    `json:"meeting_buffer_before,omitempty"`
	MeetingBufferAfter         *int    `json:"meeting_buffer_after,omitempty"`
	EnableDynamicVmrs          *bool   `json:"enable_dynamic_vmrs,omitempty"`
	EnablePersonalVmrs         *bool   `json:"enable_personal_vmrs,omitempty"`
	AllowNewUsers              *bool   `json:"allow_new_users,omitempty"`
	DisableProxy               *bool   `json:"disable_proxy,omitempty"`
	UseCustomAddInSources      *bool   `json:"use_custom_add_in_sources,omitempty"`
	EnableAddinDebugLogs       *bool   `json:"enable_addin_debug_logs,omitempty"`
	// OAuth fields
	OauthClientID      *string `json:"oauth_client_id,omitempty"`
	OauthClientSecret  string  `json:"oauth_client_secret,omitempty"`
	OauthAuthEndpoint  string  `json:"oauth_auth_endpoint,omitempty"`
	OauthTokenEndpoint string  `json:"oauth_token_endpoint,omitempty"`
	OauthRedirectURI   string  `json:"oauth_redirect_uri,omitempty"`
	OauthRefreshToken  string  `json:"oauth_refresh_token,omitempty"`
	// Kerberos fields
	KerberosRealm                  string `json:"kerberos_realm,omitempty"`
	KerberosKdc                    string `json:"kerberos_kdc,omitempty"`
	KerberosKdcHttpsProxy          string `json:"kerberos_kdc_https_proxy,omitempty"`
	KerberosExchangeSpn            string `json:"kerberos_exchange_spn,omitempty"`
	KerberosEnableTls              *bool  `json:"kerberos_enable_tls,omitempty"`
	KerberosAuthEveryRequest       *bool  `json:"kerberos_auth_every_request,omitempty"`
	KerberosVerifyTlsUsingCustomCa *bool  `json:"kerberos_verify_tls_using_custom_ca,omitempty"`
	// Add-in fields
	AddinServerDomain           string  `json:"addin_server_domain,omitempty"`
	AddinDisplayName            string  `json:"addin_display_name,omitempty"`
	AddinDescription            string  `json:"addin_description,omitempty"`
	AddinProviderName           string  `json:"addin_provider_name,omitempty"`
	AddinButtonLabel            string  `json:"addin_button_label,omitempty"`
	AddinGroupLabel             string  `json:"addin_group_label,omitempty"`
	AddinSupertipTitle          string  `json:"addin_supertip_title,omitempty"`
	AddinSupertipDescription    string  `json:"addin_supertip_description,omitempty"`
	AddinApplicationID          *string `json:"addin_application_id,omitempty"`
	AddinAuthorityURL           string  `json:"addin_authority_url,omitempty"`
	AddinOidcMetadataURL        string  `json:"addin_oidc_metadata_url,omitempty"`
	AddinAuthenticationMethod   string  `json:"addin_authentication_method,omitempty"`
	AddinNaaWebApiApplicationID *string `json:"addin_naa_web_api_application_id,omitempty"`
	// Add-in pane fields
	AddinPaneTitle                                   string `json:"addin_pane_title,omitempty"`
	AddinPaneDescription                             string `json:"addin_pane_description,omitempty"`
	AddinPaneButtonTitle                             string `json:"addin_pane_button_title,omitempty"`
	AddinPaneSuccessHeading                          string `json:"addin_pane_success_heading,omitempty"`
	AddinPaneSuccessMessage                          string `json:"addin_pane_success_message,omitempty"`
	AddinPaneAlreadyVideoMeetingHeading              string `json:"addin_pane_already_video_meeting_heading,omitempty"`
	AddinPaneAlreadyVideoMeetingMessage              string `json:"addin_pane_already_video_meeting_message,omitempty"`
	AddinPaneGeneralErrorHeading                     string `json:"addin_pane_general_error_heading,omitempty"`
	AddinPaneGeneralErrorMessage                     string `json:"addin_pane_general_error_message,omitempty"`
	AddinPaneManagementNodeDownHeading               string `json:"addin_pane_management_node_down_heading,omitempty"`
	AddinPaneManagementNodeDownMessage               string `json:"addin_pane_management_node_down_message,omitempty"`
	AddinPanePersonalVmrAddButton                    string `json:"addin_pane_personal_vmr_add_button,omitempty"`
	AddinPanePersonalVmrSignInButton                 string `json:"addin_pane_personal_vmr_sign_in_button,omitempty"`
	AddinPanePersonalVmrSelectMessage                string `json:"addin_pane_personal_vmr_select_message,omitempty"`
	AddinPanePersonalVmrNoneMessage                  string `json:"addin_pane_personal_vmr_none_message,omitempty"`
	AddinPanePersonalVmrErrorGettingMessage          string `json:"addin_pane_personal_vmr_error_getting_message,omitempty"`
	AddinPanePersonalVmrErrorSigningInMessage        string `json:"addin_pane_personal_vmr_error_signing_in_message,omitempty"`
	AddinPanePersonalVmrErrorInsertingMeetingMessage string `json:"addin_pane_personal_vmr_error_inserting_meeting_message,omitempty"`
	// Personal VMR OAuth fields
	PersonalVmrOauthClientID                   *string `json:"personal_vmr_oauth_client_id,omitempty"`
	PersonalVmrOauthClientSecret               string  `json:"personal_vmr_oauth_client_secret,omitempty"`
	PersonalVmrOauthAuthEndpoint               string  `json:"personal_vmr_oauth_auth_endpoint,omitempty"`
	PersonalVmrOauthTokenEndpoint              string  `json:"personal_vmr_oauth_token_endpoint,omitempty"`
	PersonalVmrAdfsRelyingPartyTrustIdentifier string  `json:"personal_vmr_adfs_relying_party_trust_identifier,omitempty"`
	// Template fields
	MeetingInstructionsTemplate         string `json:"meeting_instructions_template,omitempty"`
	PersonalVmrInstructionsTemplate     string `json:"personal_vmr_instructions_template,omitempty"`
	PersonalVmrLocationTemplate         string `json:"personal_vmr_location_template,omitempty"`
	PersonalVmrNameTemplate             string `json:"personal_vmr_name_template,omitempty"`
	PersonalVmrDescriptionTemplate      string `json:"personal_vmr_description_template,omitempty"`
	PlaceholderInstructionsTemplate     string `json:"placeholder_instructions_template,omitempty"`
	ConferenceNameTemplate              string `json:"conference_name_template,omitempty"`
	ConferenceDescriptionTemplate       string `json:"conference_description_template,omitempty"`
	ConferenceSubjectTemplate           string `json:"conference_subject_template,omitempty"`
	ScheduledAliasDescriptionTemplate   string `json:"scheduled_alias_description_template,omitempty"`
	AcceptNewSingleMeetingTemplate      string `json:"accept_new_single_meeting_template,omitempty"`
	AcceptNewRecurringSeriesTemplate    string `json:"accept_new_recurring_series_template,omitempty"`
	AcceptEditedSingleMeetingTemplate   string `json:"accept_edited_single_meeting_template,omitempty"`
	AcceptEditedRecurringSeriesTemplate string `json:"accept_edited_recurring_series_template,omitempty"`
	AcceptEditedOccurrenceTemplate      string `json:"accept_edited_occurrence_template,omitempty"`
	RejectGeneralErrorTemplate          string `json:"reject_general_error_template,omitempty"`
	RejectAliasConflictTemplate         string `json:"reject_alias_conflict_template,omitempty"`
	RejectAliasDeletedTemplate          string `json:"reject_alias_deleted_template,omitempty"`
	RejectInvalidAliasIDTemplate        string `json:"reject_invalid_alias_id_template,omitempty"`
	RejectSingleMeetingPast             string `json:"reject_single_meeting_past,omitempty"`
	RejectRecurringSeriesPastTemplate   string `json:"reject_recurring_series_past_template,omitempty"`
	// JavaScript and CSS URLs
	OfficeJsURL                  string `json:"office_js_url,omitempty"`
	MicrosoftFabricURL           string `json:"microsoft_fabric_url,omitempty"`
	MicrosoftFabricComponentsURL string `json:"microsoft_fabric_components_url,omitempty"`
	AdditionalAddInScriptSources string `json:"additional_add_in_script_sources,omitempty"`
	// Related resources
	Domains                   *[]string `json:"domains,omitempty"`
	HostIdentityProviderGroup *string   `json:"host_identity_provider_group,omitempty"`
	IvrTheme                  *string   `json:"ivr_theme,omitempty"`
	NonIdpParticipants        string    `json:"non_idp_participants,omitempty"`
}

// MsExchangeConnectorListResponse represents the response from listing Microsoft Exchange connectors
type MsExchangeConnectorListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []MsExchangeConnector `json:"objects"`
}
