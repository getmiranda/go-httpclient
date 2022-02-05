package gohttp

import (
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

// ClientBuilder is the interface for building a client.
type ClientBuilder interface {
	SetHeaders(headers http.Header) ClientBuilder
	SetConnectionTimeout(timeout time.Duration) ClientBuilder
	SetResponseTimeout(timeout time.Duration) ClientBuilder
	SetMaxIdleConnections(connections int) ClientBuilder
	DisableTimeouts(disable bool) ClientBuilder
	SetHttpClient(client *http.Client) ClientBuilder
	SetUserAgent(userAgent string) ClientBuilder
	SetBaseUrl(baseUrl string) ClientBuilder
	SetRateLimiter(r rate.Limit, requests int) ClientBuilder
	Build() Client
}

type clientBuilder struct {
	headers            http.Header
	maxIdleConnections int
	connectionTimeout  time.Duration
	responseTimeout    time.Duration
	disableTimeouts    bool
	baseUrl            string
	client             *http.Client
	userAgent          string
	rateLimiter        *rate.Limiter
}

// NewBuilder creates a new client builder.
func NewBuilder() ClientBuilder {
	return &clientBuilder{}
}

// Build builds the client.
func (c *clientBuilder) Build() Client {
	return &httpClient{
		builder: c,
	}
}

// SetHeaders sets the common headers to be sent with the request.
func (c *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	c.headers = headers
	return c
}

// SetConnectionTimeout set the maximum amount of time a dial will
// wait for a connect to complete.
//
// If zero, the default is defaultConnectionTimeout.
//
// With or without a timeout, the operating system may impose
// its own earlier timeout. For instance, TCP timeouts are
// often around 3 minutes.
func (c *clientBuilder) SetConnectionTimeout(timeout time.Duration) ClientBuilder {
	c.connectionTimeout = timeout
	return c
}

// SetResponseTimeout, if non-zero, specifies the amount of
// time to wait for a server's response headers after fully
// writing the request (including its body, if any). This
// time does not include the time to read the response body.
func (c *clientBuilder) SetResponseTimeout(timeout time.Duration) ClientBuilder {
	c.responseTimeout = timeout
	return c
}

// SetMaxIdleConnections, if non-zero, controls the maximum idle
// (keep-alive) connections to keep per-host. If zero,
// defaultMaxIdleConnections is used.
func (c *clientBuilder) SetMaxIdleConnections(i int) ClientBuilder {
	c.maxIdleConnections = i
	return c
}

// DisableTimeouts, if true, disables both the connection and
// read timeouts.
func (c *clientBuilder) DisableTimeouts(disable bool) ClientBuilder {
	c.disableTimeouts = disable
	return c
}

// SetHttpClient, if non-nil, will be used instead of creating
// a new client.
func (c *clientBuilder) SetHttpClient(client *http.Client) ClientBuilder {
	c.client = client
	return c
}

// SetUserAgent, set the User-Agent header.
func (c *clientBuilder) SetUserAgent(userAgent string) ClientBuilder {
	c.userAgent = userAgent
	return c
}

// SetBaseUrl, set the base url for the http client.
func (c *clientBuilder) SetBaseUrl(baseUrl string) ClientBuilder {
	c.baseUrl = baseUrl
	return c
}

// SetRateLimiter, sets the rate limiter.
//
// If nil, the default is rate.NewLimiter(rate.Inf, 0).
func (c *clientBuilder) SetRateLimiter(r rate.Limit, requests int) ClientBuilder {
	c.rateLimiter = rate.NewLimiter(r, requests)
	return c
}
