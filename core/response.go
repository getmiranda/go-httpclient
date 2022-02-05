package core

import (
	"encoding/json"
	"net/http"
	"net/http/httputil"
)

type Response struct {
	BodyBytes []byte
	*http.Response
}

// Bytes return the Response Body as bytes.
func (r *Response) Bytes() []byte {
	return r.BodyBytes
}

// String return the Respnse Body as a String.
func (r *Response) String() string {
	return string(r.BodyBytes)
}

// UnmarshalJson set the *target* parameter with the corresponding JSON response.
// target could be `struct` or `map[string]interface{}`
func (r *Response) UnmarshalJson(target interface{}) error {
	return json.Unmarshal(r.Bytes(), target)
}

// Debug let any request/response to be dumped, showing how the request/response
// went through the wire.
func (r *Response) Debug() string {
	var strReq, strResp string

	if req, err := httputil.DumpRequest(r.Request, true); err != nil {
		strReq = err.Error()
	} else {
		strReq = string(req)
	}

	if resp, err := httputil.DumpResponse(r.Response, false); err != nil {
		strResp = err.Error()
	} else {
		strResp = string(resp)
	}

	const separator = "--------\n"

	dump := separator
	dump += "REQUEST\n"
	dump += separator
	dump += strReq
	dump += "\n" + separator
	dump += "RESPONSE\n"
	dump += separator
	dump += strResp
	dump += r.String() + "\n"

	return dump

}
