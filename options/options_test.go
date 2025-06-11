package options

import (
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBaseListOptions_ToURLValues(t *testing.T) {
	tests := []struct {
		name     string
		opts     BaseListOptions
		expected url.Values
	}{
		{
			name: "with limit and offset",
			opts: BaseListOptions{Limit: 10, Offset: 5},
			expected: url.Values{
				"limit":  []string{"10"},
				"offset": []string{"5"},
			},
		},
		{
			name:     "with zero values",
			opts:     BaseListOptions{Limit: 0, Offset: 0},
			expected: url.Values{},
		},
		{
			name: "with only limit",
			opts: BaseListOptions{Limit: 20, Offset: 0},
			expected: url.Values{
				"limit": []string{"20"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.opts.ToURLValues()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSearchableListOptions_ToURLValues(t *testing.T) {
	tests := []struct {
		name     string
		opts     SearchableListOptions
		expected url.Values
	}{
		{
			name: "with all fields",
			opts: SearchableListOptions{
				BaseListOptions: BaseListOptions{Limit: 10, Offset: 5},
				Search:          "test",
			},
			expected: url.Values{
				"limit":           []string{"10"},
				"offset":          []string{"5"},
				"name__icontains": []string{"test"},
			},
		},
		{
			name: "without search",
			opts: SearchableListOptions{
				BaseListOptions: BaseListOptions{Limit: 10, Offset: 5},
				Search:          "",
			},
			expected: url.Values{
				"limit":  []string{"10"},
				"offset": []string{"5"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.opts.ToURLValues()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestTimeFilteredListOptions_ToURLValues(t *testing.T) {
	startTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC)

	tests := []struct {
		name     string
		opts     TimeFilteredListOptions
		expected url.Values
	}{
		{
			name: "with all fields",
			opts: TimeFilteredListOptions{
				SearchableListOptions: SearchableListOptions{
					BaseListOptions: BaseListOptions{Limit: 10, Offset: 5},
					Search:          "test",
				},
				StartTime: &startTime,
				EndTime:   &endTime,
			},
			expected: url.Values{
				"limit":           []string{"10"},
				"offset":          []string{"5"},
				"name__icontains": []string{"test"},
				"start_time__gte": []string{"2023-01-01T00:00:00Z"},
				"end_time__lt":    []string{"2023-12-31T23:59:59Z"},
			},
		},
		{
			name: "without time filters",
			opts: TimeFilteredListOptions{
				SearchableListOptions: SearchableListOptions{
					BaseListOptions: BaseListOptions{Limit: 10, Offset: 5},
					Search:          "test",
				},
				StartTime: nil,
				EndTime:   nil,
			},
			expected: url.Values{
				"limit":           []string{"10"},
				"offset":          []string{"5"},
				"name__icontains": []string{"test"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.opts.ToURLValues()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestTimeFilteredListOptions_ToURLValuesWithSearchField(t *testing.T) {
	startTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC)

	opts := TimeFilteredListOptions{
		SearchableListOptions: SearchableListOptions{
			BaseListOptions: BaseListOptions{Limit: 10, Offset: 5},
			Search:          "john",
		},
		StartTime: &startTime,
		EndTime:   &endTime,
	}

	result := opts.ToURLValuesWithSearchField("display_name__icontains")

	expected := url.Values{
		"limit":                   []string{"10"},
		"offset":                  []string{"5"},
		"display_name__icontains": []string{"john"},
		"start_time__gte":         []string{"2023-01-01T00:00:00Z"},
		"end_time__lt":            []string{"2023-12-31T23:59:59Z"},
	}

	assert.Equal(t, expected, result)
}

func TestTimeFilteredListOptions_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		opts     TimeFilteredListOptions
		expected url.Values
	}{
		{
			name: "negative limit and offset",
			opts: TimeFilteredListOptions{
				SearchableListOptions: SearchableListOptions{
					BaseListOptions: BaseListOptions{Limit: -1, Offset: -5},
				},
			},
			expected: url.Values{}, // Negative values should be ignored
		},
		{
			name: "very large values",
			opts: TimeFilteredListOptions{
				SearchableListOptions: SearchableListOptions{
					BaseListOptions: BaseListOptions{Limit: 999999, Offset: 888888},
				},
			},
			expected: url.Values{
				"limit":  []string{"999999"},
				"offset": []string{"888888"},
			},
		},
		{
			name: "search with unicode and special characters",
			opts: TimeFilteredListOptions{
				SearchableListOptions: SearchableListOptions{
					BaseListOptions: BaseListOptions{Limit: 10},
					Search:          "tést@example.com & special chars",
				},
			},
			expected: url.Values{
				"limit":           []string{"10"},
				"name__icontains": []string{"tést@example.com & special chars"},
			},
		},
		{
			name: "time with nanoseconds",
			opts: TimeFilteredListOptions{
				StartTime: func() *time.Time {
					t := time.Date(2023, 6, 15, 14, 30, 45, 123456789, time.UTC)
					return &t
				}(),
				EndTime: func() *time.Time {
					t := time.Date(2023, 6, 15, 18, 45, 30, 987654321, time.UTC)
					return &t
				}(),
			},
			expected: url.Values{
				"start_time__gte": []string{"2023-06-15T14:30:45Z"}, // Nanoseconds are truncated
				"end_time__lt":    []string{"2023-06-15T18:45:30Z"},
			},
		},
		{
			name: "only start time",
			opts: TimeFilteredListOptions{
				StartTime: func() *time.Time {
					t := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
					return &t
				}(),
				EndTime: nil,
			},
			expected: url.Values{
				"start_time__gte": []string{"2023-01-01T00:00:00Z"},
			},
		},
		{
			name: "only end time",
			opts: TimeFilteredListOptions{
				StartTime: nil,
				EndTime: func() *time.Time {
					t := time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC)
					return &t
				}(),
			},
			expected: url.Values{
				"end_time__lt": []string{"2023-12-31T23:59:59Z"},
			},
		},
		{
			name: "time in different timezone",
			opts: TimeFilteredListOptions{
				StartTime: func() *time.Time {
					loc, _ := time.LoadLocation("America/New_York")
					t := time.Date(2023, 6, 15, 10, 30, 0, 0, loc)
					return &t
				}(),
			},
			expected: url.Values{
				"start_time__gte": []string{"2023-06-15T10:30:00-04:00"}, // Preserves timezone
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.opts.ToURLValues()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestTimeFilteredListOptions_ToURLValuesWithSearchField_EdgeCases(t *testing.T) {
	tests := []struct {
		name        string
		opts        TimeFilteredListOptions
		searchField string
		expected    url.Values
	}{
		{
			name: "empty search field",
			opts: TimeFilteredListOptions{
				SearchableListOptions: SearchableListOptions{
					Search: "test",
				},
			},
			searchField: "",
			expected: url.Values{
				"": []string{"test"}, // Empty search field still gets populated
			},
		},
		{
			name: "search field with special characters",
			opts: TimeFilteredListOptions{
				SearchableListOptions: SearchableListOptions{
					Search: "test",
				},
			},
			searchField: "field__with__underscores",
			expected: url.Values{
				"field__with__underscores": []string{"test"},
			},
		},
		{
			name: "empty search string with custom field",
			opts: TimeFilteredListOptions{
				SearchableListOptions: SearchableListOptions{
					Search: "",
				},
			},
			searchField: "custom_field",
			expected:    url.Values{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.opts.ToURLValuesWithSearchField(tt.searchField)
			assert.Equal(t, tt.expected, result)
		})
	}
}
