package util

import (
	"strings"
	"time"
)

// InfinityTime is a custom time type that can handle both UTC and local time formats
type InfinityTime struct {
	time.Time
}

func (ft *InfinityTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	var err error
	// Try with timezone
	ft.Time, err = time.Parse("2006-01-02T15:04:05.000000Z07:00", s)
	if err == nil {
		return nil
	}
	// Try without timezone
	ft.Time, err = time.Parse("2006-01-02T15:04:05.000000", s)
	return err
}
