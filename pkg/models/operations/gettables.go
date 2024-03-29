// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetTablesRequest struct {
	ConnectionID int64  `queryParam:"style=form,explode=true,name=connectionID"`
	Namespace    string `queryParam:"style=form,explode=true,name=namespace"`
}

// GetTables200ApplicationJSON - Successfully fetched tables
type GetTables200ApplicationJSON struct {
	Tables []string `json:"tables,omitempty"`
}

type GetTablesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successfully fetched tables
	GetTables200ApplicationJSONObject *GetTables200ApplicationJSON
}
