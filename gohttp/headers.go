package gohttp

import (
	"net/http"

	"github.com/getmiranda/go-httpclient/gomime"
)

func getHeaders(headers ...http.Header) http.Header {
	if len(headers) > 0 {
		return headers[0]
	}
	return http.Header{}
}

func (c *httpClient) getRequestHeaders(requestHeaders http.Header) http.Header {
	result := make(http.Header)
	// Add default headers to the request
	for k, v := range c.builder.headers {
		if len(v) > 0 {
			result.Set(k, v[0])
		}
	}
	// Add custom headers to the request
	for k, v := range requestHeaders {
		if len(v) > 0 {
			result.Set(k, v[0])
		}
	}

	// Set User-Agent if it is defined and not there yet:
	if c.builder.userAgent != "" {
		if result.Get(gomime.HeaderUserAgent) != "" {
			return result
		}
		result.Set(gomime.HeaderUserAgent, c.builder.userAgent)
	}
	return result
}
