// Package options provides shared option types for list operations across all API modules.
// It includes base types for pagination, search, and time filtering that are used
// throughout the Configuration, Status, and History APIs.
package options

import (
	"net/url"
	"strconv"
	"time"
)

// BaseListOptions provides common pagination parameters for all list operations
type BaseListOptions struct {
	Limit  int
	Offset int
}

// ToURLValues converts BaseListOptions to url.Values for query parameters
func (opts *BaseListOptions) ToURLValues() url.Values {
	params := url.Values{}
	if opts.Limit > 0 {
		params.Set("limit", strconv.Itoa(opts.Limit))
	}
	if opts.Offset > 0 {
		params.Set("offset", strconv.Itoa(opts.Offset))
	}
	return params
}

// SearchableListOptions extends BaseListOptions with search functionality
type SearchableListOptions struct {
	BaseListOptions
	Search string
}

// ToURLValues converts SearchableListOptions to url.Values for query parameters
func (opts *SearchableListOptions) ToURLValues() url.Values {
	params := opts.BaseListOptions.ToURLValues()
	if opts.Search != "" {
		params.Set("name__icontains", opts.Search)
	}
	return params
}

// TimeFilteredListOptions extends SearchableListOptions with time filtering
type TimeFilteredListOptions struct {
	SearchableListOptions
	StartTime *time.Time
	EndTime   *time.Time
}

// ToURLValues converts TimeFilteredListOptions to url.Values for query parameters
func (opts *TimeFilteredListOptions) ToURLValues() url.Values {
	params := opts.SearchableListOptions.ToURLValues()
	if opts.StartTime != nil {
		params.Set("start_time__gte", opts.StartTime.Format(time.RFC3339))
	}
	if opts.EndTime != nil {
		params.Set("end_time__lt", opts.EndTime.Format(time.RFC3339))
	}
	return params
}

// ToURLValuesWithSearchField converts TimeFilteredListOptions to url.Values with custom search field
func (opts *TimeFilteredListOptions) ToURLValuesWithSearchField(searchField string) url.Values {
	params := opts.BaseListOptions.ToURLValues()
	if opts.Search != "" {
		params.Set(searchField, opts.Search)
	}
	if opts.StartTime != nil {
		params.Set("start_time__gte", opts.StartTime.Format(time.RFC3339))
	}
	if opts.EndTime != nil {
		params.Set("end_time__lt", opts.EndTime.Format(time.RFC3339))
	}
	return params
}
