// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"openapi/pkg/models/shared"
)

// StronglyTypedOneOfPostRes - OK
type StronglyTypedOneOfPostRes struct {
	JSON shared.StronglyTypedOneOfObject `json:"json"`
}

func (o *StronglyTypedOneOfPostRes) GetJSON() shared.StronglyTypedOneOfObject {
	if o == nil {
		return shared.StronglyTypedOneOfObject{}
	}
	return o.JSON
}

func (o *StronglyTypedOneOfPostRes) GetJSONSimpleObjectWithType() *shared.SimpleObjectWithType {
	return o.GetJSON().SimpleObjectWithType
}

func (o *StronglyTypedOneOfPostRes) GetJSONDeepObjectWithType() *shared.DeepObjectWithType {
	return o.GetJSON().DeepObjectWithType
}

type StronglyTypedOneOfPostResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	Res *StronglyTypedOneOfPostRes
}

func (o *StronglyTypedOneOfPostResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *StronglyTypedOneOfPostResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *StronglyTypedOneOfPostResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *StronglyTypedOneOfPostResponse) GetRes() *StronglyTypedOneOfPostRes {
	if o == nil {
		return nil
	}
	return o.Res
}
