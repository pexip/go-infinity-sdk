package infinity_go

type ConferencesResponse struct {
	Meta    map[string]interface{} `json:"meta"`
	Objects []Conference           `json:"objects"`
}

type Conference struct {
	StartTime   string `json:"start_time"`
	ResourceURI string `json:"resource_uri"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	ServiceType string `json:"service_type"`
	IsLocked    bool   `json:"is_locked"`
	IsStarted   bool   `json:"is_started"`
	GuestsMuted bool   `json:"guests_muted"`
	Tag         string `json:"tag"`
}

type ParticipantsResponse struct {
	Meta    map[string]interface{} `json:"meta"`
	Objects []Participant          `json:"objects"`
}

type Participant struct {
	Bandwidth               int    `json:"bandwidth"`
	CallDirection           string `json:"call_direction"`
	CallQuality             string `json:"call_quality"`
	CallTag                 string `json:"call_tag"`
	CallUUID                string `json:"call_uuid"`
	Conference              string `json:"conference"`
	ConnectTime             string `json:"connect_time"`
	ConversationID          string `json:"conversation_id"`
	DestinationAlias        string `json:"destination_alias"`
	DisplayName             string `json:"display_name"`
	Encryption              string `json:"encryption"`
	HasMedia                bool   `json:"has_media"`
	ID                      string `json:"id"`
	IdpUUID                 string `json:"idp_uuid"`
	IsClientMuted           bool   `json:"is_client_muted"`
	IsDirect                bool   `json:"is_direct"`
	IsDisconnectSupported   bool   `json:"is_disconnect_supported"`
	IsIdpAuthenticated      bool   `json:"is_idp_authenticated"`
	IsMuteSupported         bool   `json:"is_mute_supported"`
	IsMuted                 bool   `json:"is_muted"`
	IsOnHold                bool   `json:"is_on_hold"`
	IsPresentationSupported bool   `json:"is_presentation_supported"`
	IsPresenting            bool   `json:"is_presenting"`
	IsRecording             bool   `json:"is_recording"`
	IsStreaming             bool   `json:"is_streaming"`
	IsTranscribing          bool   `json:"is_transcribing"`
	IsTransferSupported     bool   `json:"is_transfer_supported"`
	LicenseCount            int    `json:"license_count"`
	LicenseType             string `json:"license_type"`
	MediaNode               string `json:"media_node"`
	ParentID                string `json:"parent_id"`
	ParticipantAlias        string `json:"participant_alias"`
	Protocol                string `json:"protocol"`
	ProxyNode               string `json:"proxy_node"`
	RemoteAddress           string `json:"remote_address"`
	RemotePort              int    `json:"remote_port"`
	ResourceURI             string `json:"resource_uri"`
	Role                    string `json:"role"`
	RxBandwidth             int    `json:"rx_bandwidth"`
	ServiceTag              string `json:"service_tag"`
	ServiceType             string `json:"service_type"`
	SignallingNode          string `json:"signalling_node"`
	SourceAlias             string `json:"source_alias"`
	SystemLocation          string `json:"system_location"`
	TxBandwidth             int    `json:"tx_bandwidth"`
	Vendor                  string `json:"vendor"`
}

type LocationResponse struct {
	Meta    map[string]interface{} `json:"meta"`
	Objects []Location             `json:"objects"`
}

type Location struct {
	FreeHDCalls int     `json:"free_hd_calls"`
	ID          int     `json:"id"`
	MaxHDCalls  int     `json:"max_hd_calls"`
	MediaLoad   float64 `json:"media_load"`
	Name        string  `json:"name"`
	ResourceURI string  `json:"resource_uri"`
}

type Node struct {
	BootTime              string `json:"boot_time"`
	ConfigurationID       int    `json:"configuration_id"`
	CpuCapabilities       string `json:"cpu_capabilities"`
	CpuCount              int    `json:"cpu_count"`
	CpuModel              string `json:"cpu_model"`
	DeployStatus          string `json:"deploy_status"`
	Hypervisor            string `json:"hypervisor"`
	ServerID              int    `json:"id"`
	LastAttemptedContact  string `json:"last_attempted_contact"`
	LastReported          string `json:"last_reported"`
	LastUpdated           string `json:"last_updated"`
	MaintenanceMode       bool   `json:"maintenance_mode"`
	MaintenanceModeReason string `json:"maintenance_mode_reason"`
	MaxAudioCalls         int    `json:"max_audio_calls"`
	MaxDirectParticipants int    `json:"max_direct_participants"`
	MaxFullHdCalls        int    `json:"max_full_hd_calls"`
	MaxHDCalls            int    `json:"max_hd_calls"`
	MaxMediaTokens        int    `json:"max_media_tokens"`
	MaxSdCalls            int    `json:"max_sd_calls"`
	MediaLoad             int    `json:"media_load"`
	MediaTokensUsed       int    `json:"media_tokens_used"`
	Name                  string `json:"name"`
	NodeType              string `json:"node_type"`
	ResourceURI           string `json:"resource_uri"`
	SignalingCount        int    `json:"signaling_count"`
	SyncStatus            string `json:"sync_status"`
	SystemLocation        string `json:"system_location"`
	TotalRam              int    `json:"total_ram"`
	UpgradeStatus         string `json:"upgrade_status"`
	Version               string `json:"version"`
}

type NodesResponse struct {
	Meta    map[string]interface{} `json:"meta"`
	Objects []Node                 `json:"objects"`
}

type LicenseUsage struct {
	AudioCount          int  `json:"audio_count"`
	AudioTotal          int  `json:"audio_total"`
	CustomLayoutsActive bool `json:"customlayouts_active"`
	GhmCount            int  `json:"ghm_count"`
	GhmTotal            int  `json:"ghm_total"`
	OtjCount            int  `json:"otj_count"`
	OtjTotal            int  `json:"otj_total"`
	PortCount           int  `json:"port_count"`
	PortTotal           int  `json:"port_total"`
	SchedulingCount     int  `json:"scheduling_count"`
	SchedulingTotal     int  `json:"scheduling_total"`
	SystemCount         int  `json:"system_count"`
	SystemTotal         int  `json:"system_total"`
	TeamsCount          int  `json:"teams_count"`
	TeamsTotal          int  `json:"teams_total"`
	TelehealthCount     int  `json:"telehealth_count"`
	TelehealthTotal     int  `json:"telehealth_total"`
	VmrCount            int  `json:"vmr_count"`
	VmrTotal            int  `json:"vmr_total"`
}

type LicenseResponse struct {
	Meta    map[string]interface{} `json:"meta"`
	Objects []LicenseUsage         `json:"objects"`
}
