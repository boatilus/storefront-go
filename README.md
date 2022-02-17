# storefront-go

> **IMPORTANT NOTE**
> This library is highly experimental and currently very minimally tested. It's being dogfooded in a current commercial project, but isn't production-worthy, and is probably heavily broken.

`storefront-go` is a Go client for Shopify's [Storefront](https://shopify.dev/api/storefront#top) API. It's focused on the ergonomics around querying rather than on mutation, using schema introspection and code generation to provide types, but otherwise providing a pretty bog standard HTTP client interface.

Documentation is "best effort" given the existing documentation within the schema and doesn't conform to Go best standards with respect to formatting.

## Requirements

`shopify-go` uses generics to streamline some aspects of working with the API, and therefore requires Go 1.18 (the current beta version [beta 2] works). There is no plan to support Go < 1.18.

## Introspecting the Schema

The current schema, `2022-01`, is provided in the `schema/2022-01` directory. To use a different API version, use the [Apollo CLI](https://www.apollographql.com/docs/devtools/cli/) to introspect your store's Storefront endpoint with your access token provided as the `X-Shopify-Storefront-Access-Token` header. Code generation uses the JSON version:

```bash
npm install -g graphql apollo
# or yarn global graphql apollo

apollo schema:download schema/<API_VERSION>/schema.json \
  --endpoint=https://<DOMAIN>/api/<API_VERSION>/graphql.json \
  --header="X-Shopify-Storefront-Access-Token: <API_KEY>"
```

For reference, to grab the native GraphQL schema, use [Rover](https://www.apollographql.com/docs/rover/) instead:

````bash
npm install -g rover
# or yarn global rover

rover graph introspect https://<DOMAIN>/api/<API_VERSION>/graphql.json \
  --header "X-Shopify-Storefront-Access-Token: <API_KEY>" \
  > schema/<API_VERSION>/schema.graphqls```
````

## Generating Types

Types from introspecting the current stable `2022-01` schema are already implemented in [`types.go`](types.go). If, as above, you need to target a different API version, manually run `scripts/parse.go`, pointing to the JSON schema file you'd like to use (using `go generate` directly won't be a good option until 1.18 final). So, presently:

```bash
go1.18beta2 run scripts/parse.go <PATH_TO_SCHEMA>
```

This will overwrite the existing `types.go`.

## Basic Usage

```go
import "log"

import "github.com/boatilus/storefront-go"

sf := storefront.NewClient("DOMAIN", "API_KEY")

var set storefront.Set
if err := sf.Query(`{
    collections(first: 1) {
      edges {
        node {
					title
        }
      }
    }
  }`, &set); err != nil {
    // Handle
  }

log.Print(set.Data.Collections.Edges[0].Node.Title)
// Outputs: Example Product Title
```
