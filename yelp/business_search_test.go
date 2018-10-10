package yelp

import (
	"errors"
	"net/http"
	"testing"
	"time"
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
		options.Location = StringPointer("Unova")
		mocks.mockRequest(http.MethodGet, businessSearchURL(options), http.StatusInternalServerError, errors.New("Internal server error"))
		client := New(mocks.server.Client(), "API_KEY")

		_, err := client.BusinessSearch(options)
		if err == nil {
			t.Fatal("Expected an error when request fails.")
		}
	})
}

func TestIsValid(t *testing.T) {
	t.Run("Location and Coordinates are set", func(t *testing.T) {
		options := BusinessSearchOptions{
			Location: StringPointer("Kanto"),
			Coordinates: &Coordinates{
				Longitude: 31.54,
				Latitude:  3.22,
			},
		}
		if options.IsValid() {
			t.Fatal("This should not be valid.")
		}
	})

	t.Run("OpenNow and OpenAt are set", func(t *testing.T) {
		options := BusinessSearchOptions{
			Location: StringPointer("Hoenn"),
			OpenNow:  BoolPointer(true),
			OpenAt:   Int64Pointer(int64(time.Now().Second())),
		}
		if options.IsValid() {
			t.Fatal("This should not be valid.")
		}
	})

	t.Run("Only Location is set", func(t *testing.T) {
		options := BusinessSearchOptions{
			Location: StringPointer("Johto"),
		}
		if !options.IsValid() {
			t.Fatal("This should be valid.")
		}
	})

	t.Run("Only Coordinates is set", func(t *testing.T) {
		options := BusinessSearchOptions{
			Coordinates: &Coordinates{
				Longitude: 20.17,
				Latitude:  3.14159,
			},
		}
		if !options.IsValid() {
			t.Fatal("This should be valid.")
		}
	})

	t.Run("OpenNow is set", func(t *testing.T) {
		options := BusinessSearchOptions{
			Location: StringPointer("Unova"),
			OpenNow:  BoolPointer(true),
		}
		if !options.IsValid() {
			t.Fatal("This should be valid.")
		}
	})

	t.Run("Only OpenAt is set", func(t *testing.T) {
		options := BusinessSearchOptions{
			Location: StringPointer("Alola"),
			OpenAt:   Int64Pointer(int64(time.Now().Second())),
		}
		if !options.IsValid() {
			t.Fatal("This should be valid.")
		}
	})
}

func TestURLValues(t *testing.T) {
	var options *BusinessSearchOptions
	t.Run("BusinessSearchOptions is nil", func(t *testing.T) {
		options = nil
		if len(options.URLValues()) != 0 {
			t.Fatal("Nil options should return empty url values.")
		}
	})

	t.Run("Only location is set", func(t *testing.T) {
		options = &BusinessSearchOptions{
			Location: StringPointer("Kalos"),
		}
		if location := options.URLValues().Get("location"); location != "Kalos" {
			t.Fatalf("Location: Expected \"%s\" to equal Kalos", location)
		}
	})

	t.Run("Only open_at is set", func(t *testing.T) {
		options = &BusinessSearchOptions{
			OpenAt: Int64Pointer(int64(time.Now().Second())),
		}
		if openAt := options.URLValues().Get("open_at"); openAt != IntString(*options.OpenAt) {
			t.Fatalf("Location: Expected \"%s\" to equal %d", openAt, *options.OpenAt)
		}
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
				if v != "132.231" {
					t.Fatalf("Longitude: Expected %s to equal 132.231", v)
				}
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
