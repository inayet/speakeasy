// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"openapi/pkg/models/shared"
)

type CreateTaskResponse struct {
	ContentType string
	// Request Failed
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
	// Created
	Task *shared.Task
}

func (o *CreateTaskResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateTaskResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *CreateTaskResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateTaskResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateTaskResponse) GetTask() *shared.Task {
	if o == nil {
		return nil
	}
	return o.Task
}
