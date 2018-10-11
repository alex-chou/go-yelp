package yelp

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// BusinessSearch makes a request given the options passed in.
func (c *client) BusinessSearch(bso *BusinessSearchOptions) (*BusinessSearchResults, error) {
	if !bso.IsValid() {
		return nil, errors.New("BusinessSearchOptions provided is not valid. Please see yelp/business_search.go for more details.")
	}
	var respBody BusinessSearchResults
	_, err := c.authedDo(http.MethodGet, businessSearchURL(bso), nil, nil, &respBody)
	return &respBody, err
}

// businessSearchURL returns the business search URL.
func businessSearchURL(bso *BusinessSearchOptions) string {
	return fmt.Sprintf("%s%s?%s", apiHost, businessSearchPath, bso.URLValues().Encode())
}

// BusinessSearchOptions contains the available parameters for the Business Search API.
type BusinessSearchOptions struct {
	Term        *string
	Location    *string
	Coordinates *Coordinates
	Radius      *int64
	Categories  *string
	Locale      *string
	Limit       *int64
	Offset      *int64
	SortBy      *string
	Price       *string
	OpenNow     *bool
	OpenAt      *int64
	Attributes  *string
}

// BusinessSearchResults reflects the JSON returned by the Business Search API.
type BusinessSearchResults struct {
	Total      int64      `json:"total"`
	Businesses []Business `json:"businesses"`
	Region     Region     `json:"region"`
}

// IsValid returns true when SearchOptions is not nil, either Location or Coordinates
// is set, and OpenNow and OpenAt are not both set.
func (bso *BusinessSearchOptions) IsValid() bool {
	return bso != nil && ((bso.Location != nil) != (bso.Coordinates != nil)) &&
		!(bso.OpenNow != nil && bso.OpenAt != nil)
}

// URLValues returns SearchOptions as url.Values.
func (bso *BusinessSearchOptions) URLValues() url.Values {
	vals := url.Values{}
	if bso == nil {
		return vals
	}

	if bso.Coordinates != nil {
		vals = bso.Coordinates.URLValues()
	} else if bso.Location != nil {
		vals.Add("location", *bso.Location)
	}

	if bso.Term != nil {
		vals.Add("term", *bso.Term)
	}
	if bso.Radius != nil {
		vals.Add("radius", IntString(*bso.Radius))
	}
	if bso.Categories != nil {
		vals.Add("categories", *bso.Categories)
	}
	if bso.Locale != nil {
		vals.Add("locale", *bso.Locale)
	}
	if bso.Limit != nil {
		vals.Add("limit", IntString(*bso.Limit))
	}
	if bso.Offset != nil {
		vals.Add("offset", IntString(*bso.Offset))
	}
	if bso.SortBy != nil {
		vals.Add("bsort_by", *bso.SortBy)
	}
	if bso.Price != nil {
		vals.Add("price", *bso.Price)
	}
	if bso.OpenNow != nil {
		vals.Add("open_now", BoolString(*bso.OpenNow))
	} else if bso.OpenAt != nil {
		vals.Add("open_at", IntString(*bso.OpenAt))
	}
	if bso.Attributes != nil {
		vals.Add("attributes", *bso.Attributes)
	}
	return vals
}
