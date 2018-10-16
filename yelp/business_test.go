package yelp

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestGetBusiness(t *testing.T) {
	ctx := context.Background()
	mocks := &testMocks{}
	options := &GetBusinessOptions{}
	t.Run("invalid options", func(t *testing.T) {
		client := newTestClient(nil, "API_KEY", mocks)
		_, err := client.GetBusiness(ctx, options)
		assert(t, err != nil, "Expected an error when options are invalid")
	})

	t.Run("failed request", func(t *testing.T) {
		options.ID = "test_ID_0"
		mocks.mockRequest(http.MethodGet, getBusinessPath(options), http.StatusInternalServerError, errors.New("Internal server error"))
		client := newTestClient(mocks.server.Client(), "API_KEY", mocks)

		_, err := client.GetBusiness(ctx, options)
		assert(t, err != nil, "Expected an error when request fails")
	})

	t.Run("successful request", func(t *testing.T) {
		expected := Business{}
		options.ID = "test_ID_1"
		mocks.mockRequest(http.MethodGet, getBusinessPath(options), http.StatusOK, expected)
		client := newTestClient(mocks.server.Client(), "API_KEY", mocks)

		results, err := client.GetBusiness(ctx, options)
		assert(t, err == nil, "Expected no error (%v) when request succeeds", err)
		assert(t, reflect.DeepEqual(*results, expected), "Results (%v) did not match expected (%v)", results, expected)
	})
}

func TestGetBusinessOptions(t *testing.T) {
	t.Run("Validate", func(t *testing.T) {
		var options *GetBusinessOptions
		t.Run("GetBusinessOptions are unset", func(t *testing.T) {
			options = nil
			assert(t, options.Validate() != nil, "Empty options should error")
		})

		t.Run("Locale is invalid", func(t *testing.T) {
			options = &GetBusinessOptions{
				ID:     "test_ID_2",
				Locale: StringPointer("en"),
			}
			assert(t, options.Validate() != nil, "Invalid Locale should error")
		})

		t.Run("ID is unset", func(t *testing.T) {
			options = &GetBusinessOptions{
				ID: "",
			}
			assert(t, options.Validate() != nil, "ID set should error")
		})

		t.Run("All options set properly", func(t *testing.T) {
			options = &GetBusinessOptions{
				ID:     "test_ID_3",
				Locale: StringPointer("en_US"),
			}
			assert(t, options.Validate() == nil, "Valid options should not error")
		})
	})

	t.Run("URLValues", func(t *testing.T) {
		var options *GetBusinessOptions
		t.Run("GetBusinessOptions is nil", func(t *testing.T) {
			options = nil
			assert(t, len(options.URLValues()) == 0, "Nil options should return empty url values")
		})

		t.Run("Only Locale is set", func(t *testing.T) {
			options = &GetBusinessOptions{
				Locale: StringPointer("en_US"),
			}
			locale := options.URLValues().Get("locale")
			assert(t, locale == "en_US", "Locale: Expected \"%s\" to equal en_US", locale)
		})
	})
}
