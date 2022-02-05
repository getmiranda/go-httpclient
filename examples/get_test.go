package examples

import (
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/getmiranda/go-httpclient/gohttp_testing"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Tell the HTTP library to mock any further requests from here
	gohttp_testing.MockupServer.Start()
	os.Exit(m.Run())
}

func TestGetEndpoints(t *testing.T) {
	t.Run("TestErrorFetchingFromGithub", func(t *testing.T) {
		gohttp_testing.MockupServer.DeleteMocks()
		gohttp_testing.MockupServer.AddMock(&gohttp_testing.Mock{
			Method: "GET",
			URL:    "https://api.github.com",
			Error:  errors.New("Timeout fetching from github endpoints"),
		})

		endpoints, err := GetEndpoints()

		assert.Nil(t, endpoints)
		assert.NotNil(t, err)
		assert.EqualValues(t, "Timeout fetching from github endpoints", err.Error())
	})

	t.Run("TestErrorUnmarshalResponseBody", func(t *testing.T) {
		gohttp_testing.MockupServer.DeleteMocks()
		gohttp_testing.MockupServer.AddMock(&gohttp_testing.Mock{
			Method:             "GET",
			URL:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": 123}`,
			ResponseHeaders:    make(http.Header),
		})

		endpoints, err := GetEndpoints()

		assert.Nil(t, endpoints)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "json: cannot unmarshal number into Go struct field")
	})

	t.Run("TestNoError", func(t *testing.T) {
		gohttp_testing.MockupServer.DeleteMocks()
		gohttp_testing.MockupServer.AddMock(&gohttp_testing.Mock{
			Method:             "GET",
			URL:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": "https://api.github.com/users/"}`,
		})

		endpoints, err := GetEndpoints()

		assert.Nil(t, err)
		assert.NotNil(t, endpoints)
		assert.EqualValues(t, "https://api.github.com/users/", endpoints.CurrentUserUrl)
	})
}
