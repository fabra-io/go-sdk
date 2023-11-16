// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/fabra-io/go-sdk/pkg/models/shared"
	"net/http"
)

type GetNamespacesRequest struct {
	ConnectionID int64 `queryParam:"style=form,explode=true,name=connectionID"`
}

func (o *GetNamespacesRequest) GetConnectionID() int64 {
	if o == nil {
		return 0
	}
	return o.ConnectionID
}

type GetNamespacesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successfully fetched namespaces
	Namespaces *shared.Namespaces
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetNamespacesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetNamespacesResponse) GetNamespaces() *shared.Namespaces {
	if o == nil {
		return nil
	}
	return o.Namespaces
}

func (o *GetNamespacesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetNamespacesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
