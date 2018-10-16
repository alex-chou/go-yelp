package yelp

import "testing"

func TestValidateLocale(t *testing.T) {
	t.Run("Valid locale", func(t *testing.T) {
		assert(t, ValidateLocale("zh_TW") == nil, "Valid locale should not error")
	})

	t.Run("Invalid locale", func(t *testing.T) {
		assert(t, ValidateLocale("Kanto") != nil, "Invalid locale should error")
	})
}
