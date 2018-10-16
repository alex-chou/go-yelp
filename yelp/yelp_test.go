package yelp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// testMocks contains the mocks created for testing.
type testMocks struct {
	server *httptest.Server
}

func (m *testMocks) mockRequest(method, path string, status int, response interface{}) {
	m.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == method && r.URL.String() == path {
			b, err := json.Marshal(response)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(status)
			w.Write(b)
			return
		}
		http.Error(w, fmt.Sprintf("%s %s was not mocked (%s %s was)", r.Method, r.URL.String(), method, path), http.StatusInternalServerError)
	}))
}

func newTestClient(c *http.Client, apiKey string, m *testMocks) Client {
	var host string
	if m != nil && m.server != nil {
		host = m.server.URL
	}
	return &client{
		Client: c,
		apiKey: apiKey,
		host:   host,
	}
}

func assert(t *testing.T, condition bool, assertionFormat string, values ...interface{}) {
	if !condition {
		t.Fatalf(assertionFormat, values...)
	}
}
