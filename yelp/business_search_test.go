package yelp

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestBusinessSearch(t *testing.T) {
	ctx := context.Background()
	mocks := &testMocks{}
	options := &BusinessSearchOptions{}
	t.Run("invalid options", func(t *testing.T) {
		client := newTestClient(nil, "API_KEY", mocks)
		_, err := client.BusinessSearch(ctx, options)
		assert(t, err != nil, "Expected an error when options are invalid")
	})

	t.Run("failed request", func(t *testing.T) {
		options.Location = StringPointer("Unova")
		mocks.mockRequest(http.MethodGet, businessSearchPath(options), http.StatusInternalServerError, errors.New("Internal server error"))
		client := newTestClient(mocks.server.Client(), "API_KEY", mocks)

		_, err := client.BusinessSearch(ctx, options)
		assert(t, err != nil, "Expected an error when request fails")
	})

	t.Run("successful request", func(t *testing.T) {
		expected := BusinessSearchResults{}
		options.Location = StringPointer("Sevii")
		mocks.mockRequest(http.MethodGet, businessSearchPath(options), http.StatusOK, expected)
		client := newTestClient(mocks.server.Client(), "API_KEY", mocks)

		results, err := client.BusinessSearch(ctx, options)
		assert(t, err == nil, "Expected no error (%v) when request succeeds", err)
		assert(t, reflect.DeepEqual(*results, expected), "Results (%v) did not match expected (%v)", results, expected)
	})
}

func TestValidate(t *testing.T) {
	t.Run("Location and Coordinates are set", func(t *testing.T) {
		options := BusinessSearchOptions{
			Location: StringPointer("Kanto"),
			Coordinates: &Coordinates{
				Longitude: 31.54,
				Latitude:  3.22,
			},
		}
		assert(t, options.Validate() != nil, "Location and Coordinates set should error")
	})

	t.Run("OpenNow and OpenAt are set", func(t *testing.T) {
		options := BusinessSearchOptions{
			Location: StringPointer("Hoenn"),
			OpenNow:  BoolPointer(true),
			OpenAt:   Int64Pointer(int64(time.Now().Second())),
		}
		assert(t, options.Validate() != nil, "OpenNow and OpenAt set should error")
	})

	t.Run("Only Location is set", func(t *testing.T) {
		options := BusinessSearchOptions{
			Location: StringPointer("Johto"),
		}
		assert(t, options.Validate() == nil, "Location set should not error")
	})

	t.Run("Only Coordinates is set", func(t *testing.T) {
		options := BusinessSearchOptions{
			Coordinates: &Coordinates{
				Longitude: 20.17,
				Latitude:  3.14159,
			},
		}
		assert(t, options.Validate() == nil, "Coordinates set should not error")
	})

	t.Run("OpenNow is set", func(t *testing.T) {
		options := BusinessSearchOptions{
			Location: StringPointer("Unova"),
			OpenNow:  BoolPointer(true),
		}
		assert(t, options.Validate() == nil, "OpenNow set should not error")
	})

	t.Run("Only OpenAt is set", func(t *testing.T) {
		options := BusinessSearchOptions{
			Location: StringPointer("Alola"),
			OpenAt:   Int64Pointer(int64(time.Now().Second())),
		}
		assert(t, options.Validate() == nil, "OpenAt set should not error")
	})
}

func TestURLValues(t *testing.T) {
	var options *BusinessSearchOptions
	t.Run("BusinessSearchOptions is nil", func(t *testing.T) {
		options = nil
		assert(t, len(options.URLValues()) == 0, "Nil options should return empty url values")
	})

	t.Run("Only location is set", func(t *testing.T) {
		options = &BusinessSearchOptions{
			Location: StringPointer("Kalos"),
		}
		location := options.URLValues().Get("location")
		assert(t, location == "Kalos", "Location: Expected \"%s\" to equal Kalos", location)
	})

	t.Run("Only open_at is set", func(t *testing.T) {
		options = &BusinessSearchOptions{
			OpenAt: Int64Pointer(int64(time.Now().Second())),
		}
		openAt := options.URLValues().Get("open_at")
		assert(t, openAt == IntString(*options.OpenAt), "Location: Expected \"%s\" to equal %d", openAt, *options.OpenAt)
	})

	t.Run("All url.Values are set correctly", func(t *testing.T) {
		options = &BusinessSearchOptions{
			Coordinates: &Coordinates{
				Longitude: 132.231,
				Latitude:  123.57,
			},
			Term:       StringPointer("restaurants"),
			Radius:     Int64Pointer(1987),
			Categories: StringPointer("chinese"),
			Locale:     StringPointer("en_US"),
			Limit:      Int64Pointer(23),
			Offset:     Int64Pointer(13),
			SortBy:     StringPointer("rating"),
			Price:      StringPointer("1,4"),
			OpenNow:    BoolPointer(false),
			OpenAt:     Int64Pointer(int64(time.Now().Second())),
			Attributes: StringPointer("hot_and_new"),
		}

		vals := options.URLValues()
		for k, _ := range vals {
			v := vals.Get(k)
			switch k {
			case "longitude":
				assert(t, v == "132.231", "Longitude: Expected %s to equal 132.231", v)
			case "latitude":
				if v != "123.57" {
					t.Fatalf("Latitude: Expected %s to equal 123.57", v)
				}
			case "location":
				if v != "" {
					t.Fatalf("Location: Expected \"%s\" to equal \"\"", v)
				}
			case "term":
				if v != "restaurants" {
					t.Fatalf("Term: Expected \"%s\" to equal \"restaurants\"", v)
				}
			}
		}
	})
}
