package gohttp

import (
	"net/http"
	"sync"

	"github.com/getmiranda/go-httpclient/core"
)

type httpClient struct {
	builder *clientBuilder

	client     *http.Client
	clientOnce sync.Once
}

// Client is the interface used to interact with the HTTP client.
type Client interface {
	Get(url string, headers ...http.Header) (*core.Response, error)
	Post(url string, body interface{}, headers ...http.Header) (*core.Response, error)
	Put(url string, body interface{}, headers ...http.Header) (*core.Response, error)
	Patch(url string, body interface{}, headers ...http.Header) (*core.Response, error)
	Delete(url string, headers ...http.Header) (*core.Response, error)
	Head(url string, headers ...http.Header) (*core.Response, error)
	Options(url string, headers ...http.Header) (*core.Response, error)
}

// Get issues a GET HTTP verb to the specified URL.
//
// In Restful, GET is used for "reading" or retrieving a resource.
// Client should expect a response status code of 200(OK) if resource exists,
// 404(Not Found) if it doesn't, or 400(Bad Request).
func (c *httpClient) Get(url string, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodGet, url, getHeaders(headers...), nil)
}

// Post issues a POST HTTP verb to the specified URL.
//
// In Restful, POST is used for "creating" a resource.
// Client should expect a response status code of 201(Created), 400(Bad Request),
// 404(Not Found), or 409(Conflict) if resource already exist.
//
// Body could be any of the form: string, []byte, struct & map.
func (c *httpClient) Post(url string, body interface{}, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodPost, url, getHeaders(headers...), body)
}

// Put issues a PUT HTTP verb to the specified URL.
//
// In Restful, PUT is used for "updating" a resource.
// Client should expect a response status code of of 200(OK), 404(Not Found),
// or 400(Bad Request). 200(OK) could be also 204(No Content)
//
// Body could be any of the form: string, []byte, struct & map.
func (c *httpClient) Put(url string, body interface{}, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodPost, url, getHeaders(headers...), body)
}

// Patch issues a PATCH HTTP verb to the specified URL
//
// In Restful, PATCH is used for "partially updating" a resource.
// Client should expect a response status code of of 200(OK), 404(Not Found),
// or 400(Bad Request). 200(OK) could be also 204(No Content)
//
// Body could be any of the form: string, []byte, struct & map.
func (c *httpClient) Patch(url string, body interface{}, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodPost, url, getHeaders(headers...), body)
}

// Delete issues a DELETE HTTP verb to the specified URL
//
// In Restful, DELETE is used to "delete" a resource.
// Client should expect a response status code of of 200(OK), 404(Not Found),
// or 400(Bad Request).
func (c *httpClient) Delete(url string, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodPost, url, getHeaders(headers...), nil)
}

// Head issues a HEAD HTTP verb to the specified URL
//
// In Restful, HEAD is used to "read" a resource headers only.
// Client should expect a response status code of 200(OK) if resource exists,
// 404(Not Found) if it doesn't, or 400(Bad Request).
func (c *httpClient) Head(url string, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodHead, url, getHeaders(headers...), nil)
}

// Options issues a OPTIONS HTTP verb to the specified URL
//
// In Restful, OPTIONS is used to get information about the resource
// and supported HTTP verbs.
// Client should expect a response status code of 200(OK) if resource exists,
// 404(Not Found) if it doesn't, or 400(Bad Request).
func (c *httpClient) Options(url string, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodOptions, url, getHeaders(headers...), nil)
}
