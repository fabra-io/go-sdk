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

import(
	"context"
	"log"
	gosdk "github.com/fabra-io/go-sdk"
	"github.com/fabra-io/go-sdk/pkg/models/shared"
	"github.com/fabra-io/go-sdk/pkg/models/operations"
)

func main() {
    s := gosdk.New(
        gosdk.WithSecurity(shared.Security{
            APIKeyAuth: "",
        }),
    )
    connectionID := 995455

    ctx := context.Background()
    res, err := s.Connection.GetNamespaces(ctx, connectionID)
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
## Available Resources and Operations


### [Connection](docs/sdks/connection/README.md)

* [GetNamespaces](docs/sdks/connection/README.md#getnamespaces) - Get all namespaces
* [GetSchema](docs/sdks/connection/README.md#getschema) - Get schema for table
* [GetTables](docs/sdks/connection/README.md#gettables) - Get all tables

### [CustomerData](docs/sdks/customerdata/README.md)

* [QueryObject](docs/sdks/customerdata/README.md#queryobject) - Query object record for customer

### [Destination](docs/sdks/destination/README.md)

* [CreateDestination](docs/sdks/destination/README.md#createdestination) - Create a new destination
* [GetDestinations](docs/sdks/destination/README.md#getdestinations) - Get all destinations

### [LinkToken](docs/sdks/linktoken/README.md)

* [CreateLinkToken](docs/sdks/linktoken/README.md#createlinktoken) - Create a new link token

### [Object](docs/sdks/object/README.md)

* [CreateObject](docs/sdks/object/README.md#createobject) - Create a new object
* [GetObjects](docs/sdks/object/README.md#getobjects) - Get all objects

### [Source](docs/sdks/source/README.md)

* [CreateSource](docs/sdks/source/README.md#createsource) - Create a new source
* [GetSources](docs/sdks/source/README.md#getsources) - Get all sources

### [Sync](docs/sdks/sync/README.md)

* [CreateSync](docs/sdks/sync/README.md#createsync) - Create a new sync
* [GetSyncs](docs/sdks/sync/README.md#getsyncs) - Get all syncs
<!-- End SDK Available Operations -->



<!-- Start Dev Containers -->

<!-- End Dev Containers -->



<!-- Start Pagination -->
# Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
<!-- End Pagination -->



<!-- Start Go Types -->

<!-- End Go Types -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
