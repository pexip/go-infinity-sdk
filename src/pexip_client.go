package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ManagerMetrics struct {
	ConferenceCount         map[string]int  `json:"conference_count"`
	ParticipantByTagCount   map[string]int  `json:"participant_by_tag_count"`
	Nodes                   map[string]Node `json:"nodes"`
	LicenseUsage            LicenseUsage    `json:"license_usage"`
	Status                  bool            `json:"status"`
	ParticipantByAliasCount map[string]int  `json:"participant_by_alias_count"`
	Locations               []Location      `json:"location_map"`
}

type Manager struct {
	Host     string         `yaml:"host"`
	Port     string         `yaml:"port"`
	Username string         `yaml:"username"`
	Password string         `yaml:"password"`
	Scheme   string         `yaml:"scheme"`
	Metrics  ManagerMetrics `yaml:"manager_metrics"`
	Interval int            `yaml:"interval"`
}

func basicAuth(username string, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func NewPexipClient(s *Server, m *Manager, apiPath string) (*http.Response, error) {
	// Create a new HTTP request
	// Use the scheme, host, port, and path from the Manager struct
	req, err := http.NewRequest("GET", fmt.Sprintf("%s://%s:%s%s", m.Scheme, m.Host, m.Port, apiPath), nil)
	if err != nil {
		s.Log.Error("Error creating request: ", "error", err)
		return nil, err
	}

	if m.Username != "" && m.Password != "" {
		// Add the Authorization header
		req.Header.Add("Authorization", "Basic "+basicAuth(m.Username, m.Password))
	}

	req.Header.Add("User-Agent", "Pexip Metric Exporter")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		s.Log.Error("Error making request: ", "error", err)
		m.Metrics.Status = false
		return nil, err
	}
	m.Metrics.Status = true // Assuming if we reach here, the manager is up
	return resp, nil
}

type ConferencesResponse struct {
	Meta    map[string]interface{} `json:"meta"`
	Objects []Conference           `json:"objects"`
}

type Conference struct {
	// StartTime   string `json:"start_time"`
	// ResourceURI string `json:"resource_uri"`
	// ID          string `json:"id"`
	Name string `json:"name"`
	// ServiceType string `json:"service_type"`
	// IsLocked    bool   `json:"is_locked"`
	// IsStarted   bool   `json:"is_started"`
	// GuestsMuted bool   `json:"guests_muted"`
	Tag string `json:"tag"`
}

func CheckConferences(s *Server) {
	s.Log.Debug("Running conference check")
	if s.Managers != nil {
		for i := range *s.Managers {
			m := &(*s.Managers)[i]
			// Create a map to count conferences by tag
			conferenceCount := make(map[string]int)
			s.Log.Debug("Manager", "host", m.Host)

			// Make a GET request to the Pexip API
			resp, err := NewPexipClient(s, m, "/api/admin/status/v1/conference/?limit=5000")
			if err != nil {
				s.Log.Error("Error creating request: ", "error", err)
				continue
			}
			if resp == nil {
				s.Log.Error("No response received from server.")
				continue
			}

			if resp.StatusCode != http.StatusOK {
				s.Log.Error("HTTP Status Error: ", "status", resp.Status)
				continue
			} else {
				s.Log.Debug("Response status: ", "status", resp.Status)
			}
			s.Log.Debug("Response status code: ", "status_code", resp.StatusCode)
			s.Log.Debug("Response headers: ", "headers", resp.Header)

			body, err := io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				s.Log.Error("Failed to read response body", "error", err)
				continue
			}
			if len(body) <= 0 {
				s.Log.Error("Empty response body")
				continue
			}

			var confResp ConferencesResponse
			if err := json.Unmarshal(body, &confResp); err != nil {
				s.Log.Error("Failed to parse response body", "error", err)
				continue
			}
			s.Log.Debug("Parsed conferences", "count", len(confResp.Objects))
			// Optionally log or process each conference
			for _, conf := range confResp.Objects {
				s.Log.Debug("Conference", "name", conf.Name, "tag", conf.Tag)
				conferenceCount[conf.Tag]++
			}

			m.Metrics.ConferenceCount = conferenceCount
		}
	}
	s.Log.Debug("Finished conference check")
}

type ParticipantsResponse struct {
	Meta    map[string]interface{} `json:"meta"`
	Objects []Participant          `json:"objects"`
}

type Participant struct {
	// Bandwidth              int    `json:"bandwidth"`
	// CallDirection          string `json:"call_direction"`
	// CallQuality            string `json:"call_quality"`
	// CallTag                string `json:"call_tag"`
	// CallUUID               string `json:"call_uuid"`
	// Conference             string `json:"conference"`
	// ConnectTime            string `json:"connect_time"`
	// ConversationID         string `json:"conversation_id"`
	// DestinationAlias       string `json:"destination_alias"`
	DisplayName string `json:"display_name"`
	// Encryption             string `json:"encryption"`
	// HasMedia               bool   `json:"has_media"`
	// ID                     string `json:"id"`
	// IdpUUID                string `json:"idp_uuid"`
	// IsClientMuted          bool   `json:"is_client_muted"`
	// IsDirect               bool   `json:"is_direct"`
	// IsDisconnectSupported  bool   `json:"is_disconnect_supported"`
	// IsIdpAuthenticated     bool   `json:"is_idp_authenticated"`
	// IsMuteSupported        bool   `json:"is_mute_supported"`
	// IsMuted                bool   `json:"is_muted"`
	// IsOnHold               bool   `json:"is_on_hold"`
	// IsPresentationSupported bool  `json:"is_presentation_supported"`
	// IsPresenting           bool   `json:"is_presenting"`
	// IsRecording            bool   `json:"is_recording"`
	// IsStreaming            bool   `json:"is_streaming"`
	// IsTranscribing         bool   `json:"is_transcribing"`
	// IsTransferSupported    bool   `json:"is_transfer_supported"`
	// LicenseCount           int    `json:"license_count"`
	// LicenseType            string `json:"license_type"`
	// MediaNode              string `json:"media_node"`
	// ParentID               string `json:"parent_id"`
	ParticipantAlias string `json:"participant_alias"`
	// Protocol               string `json:"protocol"`
	// ProxyNode              string `json:"proxy_node"`
	// RemoteAddress          string `json:"remote_address"`
	// RemotePort             int    `json:"remote_port"`
	// ResourceURI            string `json:"resource_uri"`
	// Role                   string `json:"role"`
	// RxBandwidth            int    `json:"rx_bandwidth"`
	ServiceTag string `json:"service_tag"`
	// ServiceType            string `json:"service_type"`
	// SignallingNode         string `json:"signalling_node"`
	// SourceAlias            string `json:"source_alias"`
	// SystemLocation         string `json:"system_location"`
	// TxBandwidth            int    `json:"tx_bandwidth"`
	// Vendor                 string `json:"vendor"`
}

func CheckParticipants(s *Server) {
	s.Log.Debug("Running participant check")
	if s.Managers != nil {
		for i := range *s.Managers {
			m := &(*s.Managers)[i]
			s.Log.Debug("Manager", "host", m.Host)

			// Make a GET request to the Pexip API
			resp, err := NewPexipClient(s, m, "/api/admin/status/v1/participant/?limit=5000")
			if err != nil {
				s.Log.Error("Error creating request: ", "error", err)
				continue
			}
			if resp == nil {
				s.Log.Error("No response received from server.")
				continue
			}

			if resp.StatusCode != http.StatusOK {
				s.Log.Error("HTTP Status Error: ", "status", resp.Status)
				continue
			} else {
				//s.Log.Debug("Response status: ", "status", resp.Status)
			}
			//s.Log.Debug("Response status code: ", "status_code", resp.StatusCode)
			//s.Log.Debug("Response headers: ", "headers", resp.Header)

			body, err := io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				s.Log.Error("Failed to read response body", "error", err)
				continue
			}
			if len(body) <= 0 {
				s.Log.Error("Empty response body")
				continue
			}

			var partResp ParticipantsResponse
			if err := json.Unmarshal(body, &partResp); err != nil {
				s.Log.Error("Failed to parse response body", "error", err)
				continue
			}
			s.Log.Debug("Parsed conferences", "count", len(partResp.Objects))
			// Optionally log or process each conference
			// Create a map to count conferences by tag
			participantByTagCount := make(map[string]int)
			participantByAliasCount := make(map[string]int)
			m.Metrics.ParticipantByTagCount = participantByTagCount
			m.Metrics.ParticipantByAliasCount = participantByAliasCount
			for _, part := range partResp.Objects {
				s.Log.Debug("Participant", "participant_alias", part.ParticipantAlias, "tag", part.ServiceTag)
				participantByTagCount[part.ServiceTag]++
				participantByAliasCount[part.ParticipantAlias]++
			}

			m.Metrics.ParticipantByTagCount = participantByTagCount
			m.Metrics.ParticipantByAliasCount = participantByAliasCount
		}
	}
	s.Log.Debug("Finished participant check")
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

func GetLocations(s *Server) map[string][]Location {
	s.Log.Debug("Running location check")
	locationMap := make(map[string][]Location)
	if s.Managers != nil {
		for i := range *s.Managers {
			m := &(*s.Managers)[i]
			s.Log.Debug("Manager", "host", m.Host)

			// Make a GET request to the Pexip API
			resp, err := NewPexipClient(s, m, "/api/admin/status/v1/system_location/?limit=5000")
			if err != nil {
				s.Log.Error("Error creating request: ", "error", err)
				continue
			}
			if resp == nil {
				s.Log.Error("No response received from server.")
				continue
			}

			if resp.StatusCode != http.StatusOK {
				s.Log.Error("HTTP Status Error: ", "status", resp.Status)
				continue
			} else {
				//s.Log.Debug("Response status: ", "status", resp.Status)
			}
			//s.Log.Debug("Response status code: ", "status_code", resp.StatusCode)
			//s.Log.Debug("Response headers: ", "headers", resp.Header)

			body, err := io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				s.Log.Error("Failed to read response body", "error", err)
				continue
			}
			if len(body) <= 0 {
				s.Log.Error("Empty response body")
				continue
			}

			var locResp LocationResponse
			if err := json.Unmarshal(body, &locResp); err != nil {
				s.Log.Error("Failed to parse response body", "error", err)
				continue
			}
			s.Log.Debug("Parsed locations", "count", len(locResp.Objects))
			// Optionally log or process each node
			for _, loc := range locResp.Objects {
				s.Log.Debug("Location", "free_hd_calls", loc.FreeHDCalls, "id", loc.ID, "max_hd_calls", loc.MaxHDCalls, "media_load", loc.MediaLoad, "name", loc.Name, "resource_uri", loc.ResourceURI)
			}
			locationMap[m.Host] = locResp.Objects
		}
		s.Log.Debug("Finished getting locations.")
	}
	return locationMap
}

type Node struct {
	// BootTime           string `json:"boot_time"`
	// ConfigurationID    int    `json:"configuration_id"`
	// CpuCapabilities    string `json:"cpu_capabilities"`
	CpuCount int `json:"cpu_count"`
	// CpuModel           string `json:"cpu_model"`
	// DeployStatus       string `json:"deploy_status"`
	// Hypervisor         string `json:"hypervisor"`
	ServerID int `json:"id"`
	// LastAttemptedContact string `json:"last_attempted_contact"`
	// LastReported         string `json:"last_reported"`
	LastUpdated string `json:"last_updated"`
	// MaintenanceMode      bool   `json:"maintenance_mode"`
	// MaintenanceModeReason string `json:"maintenance_mode_reason"`
	MaxAudioCalls int `json:"max_audio_calls"`
	// MaxDirectParticipants int   `json:"max_direct_participants"`
	// MaxFullHdCalls int `json:"max_full_hd_calls"`
	MaxHDCalls int `json:"max_hd_calls"`
	// MaxMediaTokens       int    `json:"max_media_tokens"`
	// MaxSdCalls           int    `json:"max_sd_calls"`
	MediaLoad int `json:"media_load"`
	// MediaTokensUsed      int    `json:"media_tokens_used"`
	Name     string `json:"name"`
	NodeType string `json:"node_type"`
	// ResourceURI          string `json:"resource_uri"`
	// SignalingCount       int    `json:"signaling_count"`
	// SyncStatus           string `json:"sync_status"`
	SystemLocation string `json:"system_location"`
	TotalRam       int    `json:"total_ram"`
	UpgradeStatus  string `json:"upgrade_status"`
	Version        string `json:"version"`
	BurstOverflow  bool   `json:"burst_overflow,omitempty"` // This field is not in the API response, but we add it for our logic
}

type NodesResponse struct {
	Meta    map[string]interface{} `json:"meta"`
	Objects []Node                 `json:"objects"`
}

func CheckNodes(s *Server) {
	s.Log.Debug("Running node check")
	locations := GetLocations(s)
	if s.Managers != nil {
		for i := range *s.Managers {
			m := &(*s.Managers)[i]
			s.Log.Debug("Manager", "host", m.Host)

			// Make a GET request to the Pexip API
			resp, err := NewPexipClient(s, m, "/api/admin/status/v1/worker_vm/?limit=5000")
			if err != nil {
				s.Log.Error("Error creating request: ", "error", err)
				continue
			}
			if resp == nil {
				s.Log.Error("No response received from server.")
				continue
			}

			if resp.StatusCode != http.StatusOK {
				s.Log.Error("HTTP Status Error: ", "status", resp.Status)
				continue
			} else {
				//s.Log.Debug("Response status: ", "status", resp.Status)
			}
			//s.Log.Debug("Response status code: ", "status_code", resp.StatusCode)
			//s.Log.Debug("Response headers: ", "headers", resp.Header)

			body, err := io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				s.Log.Error("Failed to read response body", "error", err)
				continue
			}
			if len(body) <= 0 {
				s.Log.Error("Empty response body")
				continue
			}

			var nodesResp NodesResponse
			if err := json.Unmarshal(body, &nodesResp); err != nil {
				s.Log.Error("Failed to parse response body", "error", err)
				continue
			}
			s.Log.Debug("Parsed nodes", "count", len(nodesResp.Objects))
			// Optionally log or process each node
			m.Metrics.Nodes = make(map[string]Node)
			for _, node := range nodesResp.Objects {
				found := false
				for _, loc := range locations[m.Host] {
					if node.SystemLocation == loc.Name {
						found = true
						break
					}
				}
				node.BurstOverflow = found

				s.Log.Debug("Node", "name", node.Name, "location", node.SystemLocation, "max_hd_calls", node.MaxHDCalls, "burst_overflow", node.BurstOverflow, "media_load", node.MediaLoad, "node_type", node.NodeType)
				m.Metrics.Nodes[node.Name] = node
				m.Metrics.Locations = locations[m.Host]
			}
		}
		s.Log.Debug("Finished node check")
	}
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

func CheckLicenseUsage(s *Server) {
	s.Log.Debug("Running license usage check")
	if s.Managers != nil {
		for i := range *s.Managers {
			m := &(*s.Managers)[i]
			s.Log.Debug("Manager", "host", m.Host)

			// Make a GET request to the Pexip API
			resp, err := NewPexipClient(s, m, "/api/admin/status/v1/licensing/?limit=5000")
			if err != nil {
				s.Log.Error("Error creating request: ", "error", err)
				continue
			}
			if resp == nil {
				s.Log.Error("No response received from server.")
				continue
			}

			if resp.StatusCode != http.StatusOK {
				s.Log.Error("HTTP Status Error: ", "status", resp.Status)
				continue
			} else {
				//s.Log.Debug("Response status: ", "status", resp.Status)
			}
			//s.Log.Debug("Response status code: ", "status_code", resp.StatusCode)
			//s.Log.Debug("Response headers: ", "headers", resp.Header)

			body, err := io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				s.Log.Error("Failed to read response body", "error", err)
				continue
			}
			if len(body) <= 0 {
				s.Log.Error("Empty response body")
				continue
			}

			var licResp LicenseResponse
			m.Metrics.LicenseUsage = LicenseUsage{}
			if err := json.Unmarshal(body, &licResp); err != nil {
				s.Log.Error("Failed to parse response body", "error", err)
				continue
			}
			s.Log.Debug("Parsed licences", "count", len(licResp.Objects))

			if len(licResp.Objects) > 0 {
				// If there are multiple objects, we take the first one
				m.Metrics.LicenseUsage = licResp.Objects[0]
			} else {
				s.Log.Error("No license usage objects found in response")
			}
		}
		s.Log.Debug("Finished license usage check.")
	}
}
