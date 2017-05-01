package yelp

// Helper functions to create and extract values from primtive pointers.

// Int64Ptr returns a pointer to the input.
func Int64Ptr(i int64) *int64 {
	return &i
}

// Int64Val extracts the value of the input.
// Default: 0
func Int64Val(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}

// Float64Ptr returns a pointer to the input.
func Float64Ptr(f float64) *float64 {
	return &f
}

// Float64Val extracts the value of the input.
// Default: 0.0
func Float64Val(f *float64) float64 {
	if f == nil {
		return 0.0
	}
	return *f
}

// BoolPtr returns a pointer to the input.
func BoolPtr(b bool) *bool {
	return &b
}

// BoolVal extracts the value of the input.
// Default: false
func BoolVal(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

// StringPtr returns a pointer to the input.
func StringPtr(s string) *string {
	return &s
}

// StringVal extracts the value of the inputer
// Default: ""
func StringVal(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
