package gohttp

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetRequestBody(t *testing.T) {
	client := &httpClient{}
	t.Run("NoBodyNilResponse", func(t *testing.T) {
		requestBody, err := client.getRequestBody("", nil)

		assert.Nil(t, err)
		assert.Nil(t, requestBody)
	})

	t.Run("BodyWithJson", func(t *testing.T) {
		requestBody := []string{"one", "two", "three"}
		body, err := client.getRequestBody("application/json", requestBody)

		assert.Nil(t, err)
		assert.NotNil(t, body)
		assert.EqualValues(t, `["one","two","three"]`, string(body))
	})

	t.Run("BodyWithXml", func(t *testing.T) {
		requestBody := []string{"one", "two", "three"}
		body, err := client.getRequestBody("application/xml", requestBody)

		assert.Nil(t, err)
		assert.NotNil(t, body)
		assert.EqualValues(t, `<string>one</string><string>two</string><string>three</string>`, string(body))
	})

	t.Run("BodyWithJsonAsDefault", func(t *testing.T) {
		requestBody, err := client.getRequestBody("", "")

		assert.Nil(t, err)
		assert.NotNil(t, requestBody)
	})
}

func TestGetMaxIdleConnections(t *testing.T) {
	t.Run("DefaultMaxIdleConnections", func(t *testing.T) {
		client := &httpClient{builder: &clientBuilder{}}

		result := client.getMaxIdleConnections()

		assert.EqualValues(t, defaultMaxIdleConnections, result)
	})

	t.Run("CustomMaxIdleConnections", func(t *testing.T) {
		client := &httpClient{builder: &clientBuilder{maxIdleConnections: 50}}

		result := client.getMaxIdleConnections()

		assert.EqualValues(t, 50, result)
	})
}

func TestGetResponseTimeout(t *testing.T) {
	t.Run("DefaultGetResponseTimeout", func(t *testing.T) {
		client := &httpClient{builder: &clientBuilder{}}

		result := client.getResponseTimeout()

		assert.EqualValues(t, defaultResponseTimeout, result)
	})

	t.Run("CustomGetResponseTimeout", func(t *testing.T) {
		client := &httpClient{
			builder: &clientBuilder{responseTimeout: 50 * time.Second},
		}

		result := client.getResponseTimeout()

		assert.EqualValues(t, 50*time.Second, result)
	})

	t.Run("DisableGetResponseTimeout", func(t *testing.T) {
		client := &httpClient{builder: &clientBuilder{disableTimeouts: true}}

		result := client.getResponseTimeout()

		assert.EqualValues(t, 0, result)
	})
}

func TestGetConnectionTimeout(t *testing.T) {
	t.Run("DefaultGetConnectionTimeout", func(t *testing.T) {
		client := &httpClient{builder: &clientBuilder{}}

		result := client.getConnectionTimeout()

		assert.EqualValues(t, defaultConnectionTimeout, result)
	})

	t.Run("CustomGetConnectionTimeout", func(t *testing.T) {
		client := &httpClient{
			builder: &clientBuilder{connectionTimeout: 50 * time.Second},
		}

		result := client.getConnectionTimeout()

		assert.EqualValues(t, 50*time.Second, result)
	})

	t.Run("DisableGetConnectionTimeout", func(t *testing.T) {
		client := &httpClient{builder: &clientBuilder{disableTimeouts: true}}

		result := client.getConnectionTimeout()

		assert.EqualValues(t, 0, result)
	})
}
