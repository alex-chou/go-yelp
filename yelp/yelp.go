package yelp

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	// apiHost is the base URL for the Yelp API
	apiHost = "https://api.yelp.com"

	// businessSearchPath is the path to search for businesses
	businessSearchPath = "/v3/businesses/search"
)

// Client defines the current available Yelp API requests that can be made.
type Client interface {
	BusinessSearch(*BusinessSearchOptions) (BusinessSearchResults, error)
}

// client implements the Client interface.
type client struct {
	*http.Client
	apiKey string
}

// New returns a new Yelp client.
func New(c *http.Client, apiKey string) *client {
	return &client{
		Client: c,
		apiKey: apiKey,
	}
}

// authedDo sets the Authorization header to the api key provided to the client .
// The response is decoded into v.
func (c *client) authedDo(method string, url string, body io.Reader, headers map[string]string, v interface{}) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	for key, val := range headers {
		req.Header.Set(key, val)
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.Do(req)
	if err != nil {
		return resp, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Print(err)
		}
	}()

	// return an error for non-2xx status codes
	if resp.StatusCode >= 300 {
		errString := fmt.Sprintf("%d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
		if respBytes, err := ioutil.ReadAll(resp.Body); err == nil {
			return nil, fmt.Errorf("%s: %s", errString, string(respBytes))
		}
		return nil, errors.New(errString)
	}

	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}

// postForm makes a POST request with form values and decodes the response body
// into v.
func (c *client) postForm(url string, data url.Values, v interface{}) (*http.Response, error) {
	resp, err := c.PostForm(url, data)
	if err != nil {
		return resp, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Print(err)
		}
	}()

	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
