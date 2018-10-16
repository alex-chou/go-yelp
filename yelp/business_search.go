package yelp

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// BusinessSearch makes a request given the options provided.
func (c *client) BusinessSearch(ctx context.Context, bso *BusinessSearchOptions) (*BusinessSearchResults, error) {
	if err := bso.Validate(); err != nil {
		return nil, err
	}
	var respBody BusinessSearchResults
	_, err := c.authedDo(ctx, http.MethodGet, businessSearchPath(bso), nil, nil, &respBody)
	return &respBody, err
}

// businessSearchPath returns the business search path with parameters.
func businessSearchPath(bso *BusinessSearchOptions) string {
	return fmt.Sprintf("/v3/businesses/search?%s", bso.URLValues().Encode())
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

// Validate returns an error with deteails when BusinessSearchOptions are not valid.
func (bso *BusinessSearchOptions) Validate() error {
	switch {
	case bso == nil:
		return errors.New("BusinessSearchOptions are unset")
	case (bso.Location == nil) == (bso.Coordinates == nil):
		return errors.New("BusinessSearchOptions must set either `Location` or `Coordinates`")
	case bso.OpenNow != nil && bso.OpenAt != nil:
		return errors.New("BusinessSearchOptions should not set both `OpenNow` and `OpenAt`")
	default:
		return nil
	}
}

// URLValues returns BusinessSearchOptions as url.Values.
func (bso *BusinessSearchOptions) URLValues() url.Values {
	if bso == nil {
		return nil
	}

	vals := url.Values{}
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
