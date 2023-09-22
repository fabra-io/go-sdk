// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetTablesRequest struct {
	ConnectionID int64  `queryParam:"style=form,explode=true,name=connectionID"`
	Namespace    string `queryParam:"style=form,explode=true,name=namespace"`
}

func (o *GetTablesRequest) GetConnectionID() int64 {
	if o == nil {
		return 0
	}
	return o.ConnectionID
}

func (o *GetTablesRequest) GetNamespace() string {
	if o == nil {
		return ""
	}
	return o.Namespace
}

// GetTables200ApplicationJSON - Successfully fetched tables
type GetTables200ApplicationJSON struct {
	Tables []string `json:"tables,omitempty"`
}

func (o *GetTables200ApplicationJSON) GetTables() []string {
	if o == nil {
		return nil
	}
	return o.Tables
}

type GetTablesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successfully fetched tables
	GetTables200ApplicationJSONObject *GetTables200ApplicationJSON
}

func (o *GetTablesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTablesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTablesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTablesResponse) GetGetTables200ApplicationJSONObject() *GetTables200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetTables200ApplicationJSONObject
}
