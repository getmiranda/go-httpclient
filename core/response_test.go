package core

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	resp := &Response{
		BodyBytes: []byte(`{"message":"Hello World"}`),
		Response: &http.Response{
			StatusCode: 200,
			Request: &http.Request{
				Method: "GET",
				URL: &url.URL{
					Scheme: "http",
					Host:   "localhost",
					Path:   "/",
				},
			},
		},
	}

	type TestResponse struct {
		Message string `json:"message"`
	}

	assert.EqualValues(t, http.StatusOK, resp.StatusCode)
	assert.EqualValues(t, `{"message":"Hello World"}`, resp.String())
	assert.EqualValues(t, []byte(`{"message":"Hello World"}`), resp.Bytes())
	assert.Contains(t, resp.Debug(), resp.String())

	var response TestResponse
	err := resp.UnmarshalJson(&response)

	assert.Nil(t, err)
	assert.EqualValues(t, "Hello World", response.Message)
}
