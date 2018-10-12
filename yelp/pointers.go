package yelp

// Helper functions to create and extract values from primtive pointers.

// Int64Pointer returns a pointer to the input.
func Int64Pointer(i int64) *int64 {
	return &i
}

// Int64Value extracts the value of the input.
// Default: 0
func Int64Value(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}

// Float64Pointer returns a pointer to the input.
func Float64Pointer(f float64) *float64 {
	return &f
}

// Float64Value extracts the value of the input.
// Default: 0.0
func Float64Value(f *float64) float64 {
	if f == nil {
		return 0.0
	}
	return *f
}

// BoolPointer returns a pointer to the input.
func BoolPointer(b bool) *bool {
	return &b
}

// BoolValue extracts the value of the input.
// Default: false
func BoolValue(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

// StringPointer returns a pointer to the input.
func StringPointer(s string) *string {
	return &s
}

// StringValue extracts the value of the inputer
// Default: ""
func StringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
