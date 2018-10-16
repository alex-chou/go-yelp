package yelp

import (
	"errors"
	"fmt"
	"net/http"
)

// GetBusiness makes a request given the options provided.
func (c *client) GetBusiness(gbo *GetBusinessOptions) (*Business, error) {
	if err := gbo.Validate(); err != nil {
		return nil, nil
	}
	var respBody Business
	_, err := c.authedDo(http.MethodGet, getBusinessPath(gbo), nil, nil, &respBody)
	return &respBody, err
}

// getBusinessPath returns the business details path.
func getBusinessPath(gbo *GetBusinessOptions) string {
	return fmt.Sprintf("/v3/businesses/%s", gbo.ID)
}

// GetBusinessOptions contains the available parameters for the Get Business API.
type GetBusinessOptions struct {
	ID     string
	Locale *string
}

// Validate returns an error with details when GetBusinessOptions is not valid.
func (gbo *GetBusinessOptions) Validate() error {
	switch {
	case gbo.ID == "":
		return errors.New("GetBusinessOptions `ID` is not set")
	case gbo.Locale != nil:
		if _, ok := ValidLocales[*gbo.Locale]; !ok {
			return fmt.Errorf("GetBusinessOptions `Locale` is invalid: %s", *gbo.Locale)
		}
		// Note: make sure to check other conditions if locale is valid
	}
	return nil
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
	Hours     []Hours  `json:"hours"`
	Alias     *string  `json:"alias"`
	IsClaimed *bool    `json:"is_claimed"`
	Photos    []string `json:"photos"`

	// The fields below are only available for Yelp Fusion VIP clients.
	Attributes map[string]interface{} `json:"attributes"`
}
