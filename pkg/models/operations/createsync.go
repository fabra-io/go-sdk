// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/fabra-io/go-sdk/pkg/models/shared"
	"net/http"
)

// CreateSyncResponseBody - Successfully created sync
type CreateSyncResponseBody struct {
	Sync *shared.Sync `json:"sync,omitempty"`
}

func (o *CreateSyncResponseBody) GetSync() *shared.Sync {
	if o == nil {
		return nil
	}
	return o.Sync
}

type CreateSyncResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully created sync
	Object *CreateSyncResponseBody
}

func (o *CreateSyncResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateSyncResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateSyncResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateSyncResponse) GetObject() *CreateSyncResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
