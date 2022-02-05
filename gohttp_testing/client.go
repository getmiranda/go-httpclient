package gohttp_testing

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type httpClientMock struct{}

func (c *httpClientMock) Do(request *http.Request) (*http.Response, error) {
	requestBody, err := request.GetBody()
	if err != nil {
		return nil, err
	}
	defer requestBody.Close()

	body, err := io.ReadAll(requestBody)
	if err != nil {
		return nil, err
	}

	var response http.Response

	mock := MockupServer.mocks[MockupServer.getMockKey(request.Method, request.URL.String(), string(body))]
	if mock != nil {
		if mock.Error != nil {
			return nil, mock.Error
		}
		response.StatusCode = mock.ResponseStatusCode
		response.Body = ioutil.NopCloser(strings.NewReader(mock.ResponseBody))
		response.ContentLength = int64(len(mock.ResponseBody))
		response.Request = request
		return &response, nil
	}
	bodyEncode := base64.StdEncoding.EncodeToString(body)
	return nil, fmt.Errorf("no mock matching %s from '%s' with given body '%s'", request.Method, request.URL.String(), bodyEncode)
}
