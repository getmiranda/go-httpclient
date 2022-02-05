package gohttp

import (
	"net/http"
	"testing"

	"github.com/getmiranda/go-httpclient/gomime"
	"github.com/stretchr/testify/assert"
)

func TestGetRequestHeaders(t *testing.T) {
	t.Run("SetUserAgent", func(t *testing.T) {
		client := &httpClient{}
		commenHeaders := make(http.Header)
		commenHeaders.Set("Content-Type", "application/json")
		client.builder = &clientBuilder{
			headers:   commenHeaders,
			userAgent: "Example",
		}

		finalheaders := client.getRequestHeaders(make(http.Header))

		t.Log(finalheaders)
		assert.EqualValues(t, 2, len(finalheaders), "Should have 2 header")
		assert.EqualValues(t, "application/json", finalheaders.Get(gomime.HeaderContentType))
		assert.EqualValues(t, "Example", finalheaders.Get(gomime.HeaderUserAgent))
	})

	t.Run("", func(t *testing.T) {
		client := &httpClient{}
		commenHeaders := make(http.Header)
		commenHeaders.Set(gomime.HeaderContentType, "application/json")
		commenHeaders.Set(gomime.HeaderUserAgent, "cool-agent")
		client.builder = &clientBuilder{
			headers: commenHeaders,
		}

		requestHeaders := make(http.Header)
		requestHeaders.Set("X-Request-Id", "ABC-123")

		finalheaders := client.getRequestHeaders(requestHeaders)

		assert.EqualValues(t, 3, len(finalheaders), "Should have 3 headers")
		assert.EqualValues(t, "ABC-123", finalheaders.Get("X-Request-Id"))
		assert.EqualValues(t, "application/json", finalheaders.Get("Content-Type"))
		assert.EqualValues(t, "cool-agent", finalheaders.Get("User-Agent"))
	})
}

func TestGetHeaders(t *testing.T) {
	t.Run("SetCustomHeader", func(t *testing.T) {
		header1 := make(http.Header)
		header1.Set("Content-Type", "application/json")
		header2 := make(http.Header)
		header2.Set("X-Request-Id", "ABC-123")

		header := getHeaders(header1, header2)

		assert.EqualValues(t, 1, len(header), "Should have 1 header")
		assert.EqualValues(t, "application/json", header.Get("Content-Type"))
	})

	t.Run("SetDefaultHeader", func(t *testing.T) {
		header := getHeaders()

		assert.NotNil(t, header)
		assert.Equal(t, http.Header{}, header)
	})
}
