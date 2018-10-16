package yelp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// Client defines the current available Yelp API requests that can be made.
type Client interface {
	BusinessSearch(context.Context, *BusinessSearchOptions) (*BusinessSearchResults, error)
	GetBusiness(context.Context, *GetBusinessOptions) (*Business, error)
}

// client implements the Client interface.
type client struct {
	*http.Client
	apiKey string
	host   string
}

// New returns a new Yelp client. The default host is https://api.yelp.com.
func New(c *http.Client, apiKey string) Client {
	return &client{
		Client: c,
		apiKey: apiKey,
		host:   "https://api.yelp.com",
	}
}

// authedDo sets the Authorization header to the api key provided to the client .
// The response is decoded into v.
func (c *client) authedDo(ctx context.Context, method string, path string, body io.Reader, headers map[string]string, v interface{}) (*http.Response, error) {
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.host, path), body)
	if err != nil {
		return nil, err
	}

	for key, val := range headers {
		req.Header.Set(key, val)
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	req.WithContext(ctx)

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
