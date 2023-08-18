// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

var RequestBodyPostApplicationJSONMapOfMapOfPrimitiveServerList = []string{
	"http://localhost:35456",
}

type RequestBodyPostApplicationJSONMapOfMapOfPrimitiveResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	Res map[string]map[string]string
}

func (o *RequestBodyPostApplicationJSONMapOfMapOfPrimitiveResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RequestBodyPostApplicationJSONMapOfMapOfPrimitiveResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RequestBodyPostApplicationJSONMapOfMapOfPrimitiveResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *RequestBodyPostApplicationJSONMapOfMapOfPrimitiveResponse) GetRes() map[string]map[string]string {
	if o == nil {
		return nil
	}
	return o.Res
}
