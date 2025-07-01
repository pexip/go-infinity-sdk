package config

// RecurringConference represents a recurring conference configuration
type RecurringConference struct {
	ID             int     `json:"id,omitempty"`
	Conference     string  `json:"conference"`
	CurrentIndex   int     `json:"current_index"`
	EWSItemID      string  `json:"ews_item_id"`
	IsDepleted     bool    `json:"is_depleted"`
	Subject        string  `json:"subject,omitempty"`
	ScheduledAlias *string `json:"scheduled_alias,omitempty"`
	ResourceURI    string  `json:"resource_uri,omitempty"`
}

// RecurringConferenceCreateRequest represents a request to create a recurring conference
type RecurringConferenceCreateRequest struct {
	Conference     string  `json:"conference"`
	CurrentIndex   int     `json:"current_index"`
	EWSItemID      string  `json:"ews_item_id"`
	IsDepleted     bool    `json:"is_depleted"`
	Subject        string  `json:"subject,omitempty"`
	ScheduledAlias *string `json:"scheduled_alias,omitempty"`
}

// RecurringConferenceUpdateRequest represents a request to update a recurring conference
type RecurringConferenceUpdateRequest struct {
	Conference     string  `json:"conference,omitempty"`
	CurrentIndex   *int    `json:"current_index,omitempty"`
	EWSItemID      string  `json:"ews_item_id,omitempty"`
	IsDepleted     *bool   `json:"is_depleted,omitempty"`
	Subject        string  `json:"subject,omitempty"`
	ScheduledAlias *string `json:"scheduled_alias,omitempty"`
}

// RecurringConferenceListResponse represents the response from listing recurring conferences
type RecurringConferenceListResponse struct {
	Meta struct {
		Limit      int    `json:"limit"`
		Next       string `json:"next"`
		Offset     int    `json:"offset"`
		Previous   string `json:"previous"`
		TotalCount int    `json:"total_count"`
	} `json:"meta"`
	Objects []RecurringConference `json:"objects"`
}
