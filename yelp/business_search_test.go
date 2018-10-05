package yelp

import (
	"errors"
	"net/http"
	"testing"
)

func TestBusinessSearch(t *testing.T) {
	mocks := &testMocks{}
	options := &BusinessSearchOptions{}
	t.Run("invalid options", func(t *testing.T) {
		client := New(nil, "API_KEY")
		_, err := client.BusinessSearch(options)
		if err == nil {
			t.Fatal("Expected an error when options are invalid")
		}
	})

	t.Run("failed request", func(t *testing.T) {
		options.Location = StringPointer("San Francisco")
		mocks.mockRequest(http.MethodGet, businessSearchURL(options), http.StatusInternalServerError, errors.New("Internal server error"))
		client := New(mocks.server.Client(), "API_KEY")

		_, err := client.BusinessSearch(options)
		if err == nil {
			t.Fatal("Expected an error when request fails.")
		}
	})
}
