package config

import (
	"context"
	"fmt"

	"github.com/pexip/go-infinity-sdk/v38/types"
)

// ListEventSinks retrieves a list of event sinks
func (s *Service) ListEventSinks(ctx context.Context, opts *ListOptions) (*EventSinkListResponse, error) {
	endpoint := "configuration/v1/event_sink/"

	if opts != nil {
		params := opts.ToURLValues()
		if len(params) > 0 {
			endpoint += "?" + params.Encode()
		}
	}

	var result EventSinkListResponse
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// GetEventSink retrieves a specific event sink by ID
func (s *Service) GetEventSink(ctx context.Context, id int) (*EventSink, error) {
	endpoint := fmt.Sprintf("configuration/v1/event_sink/%d/", id)

	var result EventSink
	err := s.client.GetJSON(ctx, endpoint, &result)
	return &result, err
}

// CreateEventSink creates a new event sink
func (s *Service) CreateEventSink(ctx context.Context, req *EventSinkCreateRequest) (*types.PostResponse, error) {
	endpoint := "configuration/v1/event_sink/"
	return s.client.PostWithResponse(ctx, endpoint, req, nil)
}

// UpdateEventSink updates an existing event sink
func (s *Service) UpdateEventSink(ctx context.Context, id int, req *EventSinkUpdateRequest) (*EventSink, error) {
	endpoint := fmt.Sprintf("configuration/v1/event_sink/%d/", id)

	var result EventSink
	err := s.client.PutJSON(ctx, endpoint, req, &result)
	return &result, err
}

// DeleteEventSink deletes an event sink
func (s *Service) DeleteEventSink(ctx context.Context, id int) error {
	endpoint := fmt.Sprintf("configuration/v1/event_sink/%d/", id)
	return s.client.DeleteJSON(ctx, endpoint, nil)
}
