package gohttp

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/ajg/form"
	"github.com/getmiranda/go-httpclient/core"
	"github.com/getmiranda/go-httpclient/gohttp_testing"
	"github.com/getmiranda/go-httpclient/gomime"
	"golang.org/x/time/rate"
)

const (
	defaultMaxIdleConnections = 10
	defaultResponseTimeout    = time.Second * 30
	defaultConnectionTimeout  = time.Second * 30
)

type request struct {
	method  string
	url     string
	headers http.Header
	body    interface{}
	req     *http.Request
}

func (c *httpClient) do(request *request) (*core.Response, error) {
	req, err := c.getRequest(request)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	if err := c.getRateLimit().Wait(ctx); err != nil { // This is a blocking call. Honors the rate limit
		return nil, err
	}

	response, err := c.getHttpClient().Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	finalResponse := &core.Response{
		BodyBytes: responseBody,
		Response:  response,
	}
	return finalResponse, nil
}

func (c *httpClient) getRequest(request *request) (*http.Request, error) {
	if request.req != nil {
		return request.req, nil
	}

	fullHeaders := c.getRequestHeaders(request.headers)

	requestBody, err := c.getRequestBody(fullHeaders.Get(gomime.HeaderContentType), request.body)
	if err != nil {
		return nil, err
	}

	url := c.builder.baseUrl + request.url

	req, err := http.NewRequest(request.method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	req.Header = fullHeaders

	return req, nil
}

func (c *httpClient) getHttpClient() core.HttpClient {
	if gohttp_testing.MockupServer.IsEnabled() {
		return gohttp_testing.MockupServer.GetMockedClient()
	}
	c.clientOnce.Do(func() {
		if c.builder.client != nil {
			c.client = c.builder.client
			return
		}
		c.client = &http.Client{
			Timeout: c.getResponseTimeout() + c.getConnectionTimeout(),
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   c.getMaxIdleConnections(),
				ResponseHeaderTimeout: c.getResponseTimeout(),
				DialContext: (&net.Dialer{
					Timeout: c.getConnectionTimeout(),
				}).DialContext,
				DisableKeepAlives: c.builder.disableKeepAlives,
			},
		}
	})
	return c.client
}

func (c *httpClient) getRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	switch strings.ToLower(contentType) {
	case gomime.ContentTypeJson:
		return json.Marshal(body)
	case gomime.ContentTypeXml:
		return xml.Marshal(body)
	case gomime.ContentTypeFormUrlEncoded:
		v, err := form.EncodeToString(body)
		if err != nil {
			return nil, err
		}
		return []byte(v), nil
	default:
		return json.Marshal(body)
	}
}

func (c *httpClient) getMaxIdleConnections() int {
	if c.builder.maxIdleConnections > 0 {
		return c.builder.maxIdleConnections
	}
	return defaultMaxIdleConnections
}

func (c *httpClient) getResponseTimeout() time.Duration {
	if c.builder.disableTimeouts {
		return 0
	}
	if c.builder.responseTimeout > 0 {
		return c.builder.responseTimeout
	}
	return defaultResponseTimeout
}

func (c *httpClient) getConnectionTimeout() time.Duration {
	if c.builder.disableTimeouts {
		return 0
	}
	if c.builder.connectionTimeout > 0 {
		return c.builder.connectionTimeout
	}
	return defaultConnectionTimeout
}

func (c *httpClient) getRateLimit() *rate.Limiter {
	if c.builder.rateLimiter != nil {
		return c.builder.rateLimiter
	}
	return rate.NewLimiter(rate.Inf, 0)
}
