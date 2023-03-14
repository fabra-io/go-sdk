package fabra

import (
	"github.com/fabra-io/go-sdk/pkg/models/shared"
	"github.com/fabra-io/go-sdk/pkg/utils"
	"net/http"
	"time"
)

var ServerList = []string{
	"https://api.fabra.io",
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

type Fabra struct {
	Connection  *connection
	Destination *destination
	LinkToken   *linkToken
	Object      *object
	Source      *source
	Sync        *sync

	// Non-idiomatic field names below are to namespace fields from the fields names above to avoid name conflicts
	_defaultClient  HTTPClient
	_securityClient HTTPClient
	_security       *shared.Security
	_serverURL      string
	_language       string
	_sdkVersion     string
	_genVersion     string
}

type SDKOption func(*Fabra)

func WithServerURL(serverURL string) SDKOption {
	return func(sdk *Fabra) {
		sdk._serverURL = serverURL
	}
}

func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Fabra) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk._serverURL = serverURL
	}
}

func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Fabra) {
		sdk._defaultClient = client
	}
}

func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *Fabra) {
		sdk._security = &security
	}
}

func New(opts ...SDKOption) *Fabra {
	sdk := &Fabra{
		_language:   "go",
		_sdkVersion: "0.6.2",
		_genVersion: "1.9.2",
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk._defaultClient == nil {
		sdk._defaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk._securityClient == nil {

		if sdk._security != nil {
			sdk._securityClient = utils.ConfigureSecurityClient(sdk._defaultClient, sdk._security)
		} else {
			sdk._securityClient = sdk._defaultClient
		}

	}

	if sdk._serverURL == "" {
		sdk._serverURL = ServerList[0]
	}

	sdk.Connection = newConnection(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Destination = newDestination(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.LinkToken = newLinkToken(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Object = newObject(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Source = newSource(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Sync = newSync(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	return sdk
}
