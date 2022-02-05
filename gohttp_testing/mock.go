package gohttp_testing

import (
	"fmt"
	"net/http"

	"github.com/getmiranda/go-httpclient/core"
)

// Mock structure provides a clean way to configure HTTP mocks based on
// the combination between request method, URL and request body.
//
// All requests will be sent to the mockup server if mockup is activated.
// To activate the mockup *environment* you have to programmatically start the mockup server
// 	gohttp_testing.MockupServer.Start()
type Mock struct {
	Method      string
	URL         string
	RequestBody string

	Error              error
	ResponseBody       string
	ResponseStatusCode int
	ResponseHeaders    http.Header
}

// GetResponse returns a Response object based on the mock configuration.
func (m *Mock) GetResponse() (*core.Response, error) {
	if m.Error != nil {
		return nil, m.Error
	}

	// Fill response object with current mock details:
	response := &core.Response{
		BodyBytes: []byte(m.ResponseBody),
		Response: &http.Response{
			Status:     fmt.Sprintf("%d %s", m.ResponseStatusCode, http.StatusText(m.ResponseStatusCode)),
			StatusCode: m.ResponseStatusCode,
			Header:     make(http.Header),
		},
	}

	// Make sure each mocked response header is present in the final response object:
	for header := range m.ResponseHeaders {
		response.Header.Set(header, m.ResponseHeaders.Get(header))
	}
	return response, nil
}
