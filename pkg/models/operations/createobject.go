// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/fabra-io/go-sdk/pkg/models/shared"
	"net/http"
)

// CreateObject200ApplicationJSON - Successfully created object
type CreateObject200ApplicationJSON struct {
	Object *shared.Object `json:"object,omitempty"`
}

func (o *CreateObject200ApplicationJSON) GetObject() *shared.Object {
	if o == nil {
		return nil
	}
	return o.Object
}

type CreateObjectResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully created object
	CreateObject200ApplicationJSONObject *CreateObject200ApplicationJSON
}

func (o *CreateObjectResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateObjectResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateObjectResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateObjectResponse) GetCreateObject200ApplicationJSONObject() *CreateObject200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.CreateObject200ApplicationJSONObject
}
