package core

import (
	"net/http"
)

// HttpClient is the interface that wraps the Do method.
type HttpClient interface {
	Do(request *http.Request) (*http.Response, error)
}
