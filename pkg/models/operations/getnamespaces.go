package operations

import (
	"net/http"
)

type GetNamespacesQueryParams struct {
	ConnectionID int64 `queryParam:"style=form,explode=true,name=connectionID"`
}

type GetNamespacesRequest struct {
	QueryParams GetNamespacesQueryParams
}

type GetNamespaces200ApplicationJSON struct {
	Namespaces []string `json:"namespaces,omitempty"`
}

type GetNamespacesResponse struct {
	ContentType                           string
	StatusCode                            int
	RawResponse                           *http.Response
	GetNamespaces200ApplicationJSONObject *GetNamespaces200ApplicationJSON
}
