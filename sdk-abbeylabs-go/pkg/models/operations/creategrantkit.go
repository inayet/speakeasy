// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"openapi/pkg/models/shared"
)

type CreateGrantKitResponse struct {
	ContentType string
	// Request Failed
	Error *shared.Error
	// Created
	GrantKit    *shared.GrantKit
	StatusCode  int
	RawResponse *http.Response
}

func (o *CreateGrantKitResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateGrantKitResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *CreateGrantKitResponse) GetGrantKit() *shared.GrantKit {
	if o == nil {
		return nil
	}
	return o.GrantKit
}

func (o *CreateGrantKitResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateGrantKitResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
