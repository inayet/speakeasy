// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type ConflictingParamsRequest struct {
	StrPathParameter  string `pathParam:"style=simple,explode=false,name=str"`
	StrQueryParameter string `queryParam:"style=form,explode=true,name=str"`
}

func (o *ConflictingParamsRequest) GetStrPathParameter() string {
	if o == nil {
		return ""
	}
	return o.StrPathParameter
}

func (o *ConflictingParamsRequest) GetStrQueryParameter() string {
	if o == nil {
		return ""
	}
	return o.StrQueryParameter
}

// ConflictingParamsRes - OK
type ConflictingParamsRes struct {
	Args map[string]string `json:"args"`
	URL  string            `json:"url"`
}

func (o *ConflictingParamsRes) GetArgs() map[string]string {
	if o == nil {
		return map[string]string{}
	}
	return o.Args
}

func (o *ConflictingParamsRes) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

type ConflictingParamsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	Res *ConflictingParamsRes
}

func (o *ConflictingParamsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ConflictingParamsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ConflictingParamsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ConflictingParamsResponse) GetRes() *ConflictingParamsRes {
	if o == nil {
		return nil
	}
	return o.Res
}
