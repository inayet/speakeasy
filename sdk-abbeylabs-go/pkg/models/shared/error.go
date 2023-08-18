// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Error - Authentication Failed
type Error struct {
	Code    string  `json:"code"`
	DocURL  *string `json:"doc_url,omitempty"`
	Message string  `json:"message"`
	Param   *string `json:"param,omitempty"`
	Type    string  `json:"type"`
}

func (o *Error) GetCode() string {
	if o == nil {
		return ""
	}
	return o.Code
}

func (o *Error) GetDocURL() *string {
	if o == nil {
		return nil
	}
	return o.DocURL
}

func (o *Error) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *Error) GetParam() *string {
	if o == nil {
		return nil
	}
	return o.Param
}

func (o *Error) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}
