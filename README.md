# Fabra Go SDK

<div align="left">
   <p>Use the Fabra API to build customer-facing data warehouse integrations to let your customers start sending data to your application. Unblock your sales pipeline in days, not months.</p>
   <a href="https://github.com/fabra-io/go-sdk/actions"><img src="https://img.shields.io/github/actions/workflow/status/fabra-io/go-sdk/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
   <a href="https://www.fabra.io/#Email-Hero"><img src="https://img.shields.io/static/v1?label=Docs&message=Sign Up&color=2ca47c&style=for-the-badge" /></a>
</div>

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/fabra-io/go-sdk
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/fabra-io/go-sdk"
    "github.com/fabra-io/go-sdk/pkg/models/shared"
    "github.com/fabra-io/go-sdk/pkg/models/operations"
)

func main() {
    s := fabra.New(
        fabra.WithSecurity(shared.Security{
            APIKeyAuth: shared.SchemeAPIKeyAuth{
                APIKey: "YOUR_API_KEY_HERE",
            },
        }),
    )

    req := operations.GetNamespacesRequest{
        QueryParams: operations.GetNamespacesQueryParams{
            ConnectionID: 548814,
        },
    }

    ctx := context.Background()
    res, err := s.Connection.GetNamespaces(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Namespaces != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## SDK Available Operations


### Connection

* `GetNamespaces` - Get all namespaces
* `GetSchema` - Get schema for table
* `GetTables` - Get all tables

### Destination

* `CreateDestination` - Create a new destination
* `GetDestinations` - Get all destinations

### LinkToken

* `CreateLinkToken` - Create a new link token

### Object

* `CreateObject` - Create a new object
* `GetObjects` - Get all objects

### Source

* `CreateSource` - Create a new source
* `GetSources` - Get all sources

### Sync

* `CreateSync` - Create a new sync
* `GetSyncs` - Get all syncs
<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
