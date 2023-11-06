// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/fabra-io/go-sdk/pkg/models/shared"
	"net/http"
)

// GetSyncs200ApplicationJSON - Successfully fetched syncs
type GetSyncs200ApplicationJSON struct {
	Syncs []shared.Sync `json:"syncs,omitempty"`
}

func (o *GetSyncs200ApplicationJSON) GetSyncs() []shared.Sync {
	if o == nil {
		return nil
	}
	return o.Syncs
}

type GetSyncsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully fetched syncs
	GetSyncs200ApplicationJSONObject *GetSyncs200ApplicationJSON
}

func (o *GetSyncsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSyncsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSyncsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSyncsResponse) GetGetSyncs200ApplicationJSONObject() *GetSyncs200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetSyncs200ApplicationJSONObject
}
