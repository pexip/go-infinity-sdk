package status

func intPtr(i int) *int       { return &i }
func strPtr(s string) *string { return &s }

// Helper functions for pointer dereferencing
func derefString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func derefInt(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}
