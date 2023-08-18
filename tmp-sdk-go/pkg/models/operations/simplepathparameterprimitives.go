// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type SimplePathParameterPrimitivesRequest struct {
	BoolParam bool    `pathParam:"style=simple,explode=false,name=boolParam"`
	IntParam  int64   `pathParam:"style=simple,explode=false,name=intParam"`
	NumParam  float64 `pathParam:"style=simple,explode=false,name=numParam"`
	StrParam  string  `pathParam:"style=simple,explode=false,name=strParam"`
}

func (o *SimplePathParameterPrimitivesRequest) GetBoolParam() bool {
	if o == nil {
		return false
	}
	return o.BoolParam
}

func (o *SimplePathParameterPrimitivesRequest) GetIntParam() int64 {
	if o == nil {
		return 0
	}
	return o.IntParam
}

func (o *SimplePathParameterPrimitivesRequest) GetNumParam() float64 {
	if o == nil {
		return 0.0
	}
	return o.NumParam
}

func (o *SimplePathParameterPrimitivesRequest) GetStrParam() string {
	if o == nil {
		return ""
	}
	return o.StrParam
}

// SimplePathParameterPrimitivesRes - OK
type SimplePathParameterPrimitivesRes struct {
	URL string `json:"url"`
}

func (o *SimplePathParameterPrimitivesRes) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

type SimplePathParameterPrimitivesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	Res *SimplePathParameterPrimitivesRes
}

func (o *SimplePathParameterPrimitivesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SimplePathParameterPrimitivesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SimplePathParameterPrimitivesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *SimplePathParameterPrimitivesResponse) GetRes() *SimplePathParameterPrimitivesRes {
	if o == nil {
		return nil
	}
	return o.Res
}
