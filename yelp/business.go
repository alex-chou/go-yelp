package yelp

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// GetBusinessOptions contains the available parameters for the Get Business API.
type GetBusinessOptions struct {
	ID     string
	Locale *string
}

// Business defines a business returned by the Yelp API.
type Business struct {
	Categories   []Category  `json:"categories"`
	Coodinates   Coordinates `json:"coordinates"`
	DisplayPhone string      `json:"display_phone"`
	Distance     float64     `json:"distance"`
	ID           string      `json:"id"`
	ImageURL     string      `json:"image_url"`
	IsClosed     bool        `json:"is_closed"`
	Location     Location    `json:"location"`
	Name         string      `json:"name"`
	Phone        string      `json:"phone"`
	Price        string      `json:"price"`
	Rating       float64     `json:"rating"`
	ReviewCount  int64       `json:"review_count"`
	URL          string      `json:"url"`
	Transactions []string    `json:"transactions"`

	// The fields below are only returned by the Get Business API.
	Hours     []Hours  `json:"hours,omitempty"`
	Alias     *string  `json:"alias,omitempty"`
	IsClaimed *bool    `json:"is_claimed,omitempty"`
	Photos    []string `json:"photos,omitempty"`

	// The fields below are only available for Yelp Fusion VIP clients.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}

// Category describes a business.
type Category struct {
	Alias string `json:"alias"`
	Title string `json:"title"`
}

// Hours has opening business hours for each day in a week.
type Hours struct {
	IsOpenNow bool   `json:"is_open_now"`
	HoursType string `json:"hours_type"`
	Open      []Open `json:"open"`
}

// Open has detailed opening business hours for each day in a week.
type Open struct {
	Day         int    `json:"day"`
	Start       string `json:"start"`
	End         string `json:"end"`
	IsOvernight bool   `json:"is_overnight"`
}

// GetBusiness makes a request given the options provided.
func (c *client) GetBusiness(ctx context.Context, gbo *GetBusinessOptions) (*Business, error) {
	if err := gbo.Validate(); err != nil {
		return nil, err
	}
	var respBody Business
	_, err := c.authedDo(ctx, http.MethodGet, getBusinessPath(gbo), nil, nil, &respBody)
	return &respBody, err
}

// getBusinessPath returns the business details path.
func getBusinessPath(gbo *GetBusinessOptions) string {
	return fmt.Sprintf("/v3/businesses/%s?%s", gbo.ID, gbo.URLValues().Encode())
}

// Validate returns an error with details when GetBusinessOptions are not valid.
func (gbo *GetBusinessOptions) Validate() error {
	switch {
	case gbo == nil:
		return errors.New("GetBusinessOptions are unset")
	case gbo.ID == "":
		return errors.New("GetBusinessOptions `ID` is not set")
	case gbo.Locale != nil && ValidateLocale(*gbo.Locale) != nil:
		return fmt.Errorf("GetBusinessOptions `Locale` is invalid: %s", *gbo.Locale)
	default:
		return nil
	}
}

// URLValues returns GetBusinessOptions as url.Values.
func (gbo *GetBusinessOptions) URLValues() url.Values {
	if gbo == nil {
		return nil
	}

	vals := url.Values{}
	if gbo.Locale != nil {
		vals.Add("locale", *gbo.Locale)
	}
	return vals
}
