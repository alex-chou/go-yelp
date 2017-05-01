package yelp

import (
	"net/url"
	"time"
)

// Credentials stores a user's ClientID, ClientSecret, and AccessToken. When
// ExpiryDate is reached, a new AccessToken should be fetched.
type Credentials struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	AccessToken  string
	ExpiryDate   time.Time
}

const (
	// ccGrantType is the only grant type accepted by the Yelp authentication API
	ccGrantType = "client_credentials"
)

// IsValid returns true when ExpiryDate is set and is not set to a time one minute
// from now (so that the access token doesn't expire during a request).
func (c Credentials) IsValid() bool {
	return !(c.ExpiryDate.IsZero() || c.ExpiryDate.After(time.Now().Add(time.Second)))
}

// URLValues returns c as url.Values. `grant_type` is always set as `client_credentials`
// as that is the only supported value.
func (c Credentials) URLValues() url.Values {
	vals := url.Values{}
	vals.Add("client_id", c.ClientID)
	vals.Add("client_secret", c.ClientSecret)
	vals.Add("grant_type", ccGrantType)
	return vals
}
