package tfregistry

import (
	"net/http"
	"net/url"
)

// newMockServer returns a new mock server for testing.
func newMockServer() (*http.ServeMux, *url.URL) { _ = "STUB: not implemented"; return nil, nil }

// newTestClient returns a new client for testing.
func newTestClient(mockServerURL *url.URL) *Client { _ = "STUB: not implemented"; return nil }
