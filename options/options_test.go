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
