# storefront-go

![build](https://github.com/boatilus/storefront-go/actions/workflows/go.yml/badge.svg) [![Go Reference](https://pkg.go.dev/badge/github.com/boatilus/storefront-go.svg)](https://pkg.go.dev/github.com/boatilus/storefront-go)

`storefront-go` is a Go client for [Shopify's Storefront API](https://shopify.dev/api/storefront#top). It's focused on the ergonomics around querying rather than on mutation, using schema introspection and code generation to provide types, but otherwise providing a pretty bog standard HTTP client interface.

Documentation is "best effort" given the existing documentation within the schema and doesn't conform to Go best standards with respect to formatting.

> ⚠️ **IMPORTANT NOTE**: This library is highly experimental and currently very minimally tested. It's being dogfooded in a current commercial project, but isn't production-worthy, and is probably heavily broken.

## Requirements

`storefront-go` uses generics to streamline some aspects of working with the API, and therefore requires Go 1.18 (the current beta version [beta 2] works). There is no plan to support Go < 1.18.

## Introspecting the Schema

The current stable schema, `2022-01`, is provided in the `schema/2022-01` directory. To use a different API version, use the [Apollo CLI](https://www.apollographql.com/docs/devtools/cli/) to introspect your store's Storefront API endpoint with your access token provided as the `X-Shopify-Storefront-Access-Token` header. Code generation uses the JSON version:

```bash
npm install -g graphql apollo
# or yarn global graphql apollo

apollo schema:download schema/<API_VERSION>/schema.json \
  --endpoint=https://<DOMAIN>/api/<API_VERSION>/graphql.json \
  --header="X-Shopify-Storefront-Access-Token: <ACCESS_TOKEN>"
```

For reference, to grab the native GraphQL schema, use [Rover](https://www.apollographql.com/docs/rover/) instead:

````bash
npm install -g rover
# or yarn global rover

rover graph introspect https://<DOMAIN>/api/<API_VERSION>/graphql.json \
  --header "X-Shopify-Storefront-Access-Token: <ACCESS_TOKEN>" \
  > schema/<API_VERSION>/schema.graphqls```
````

## Generating Types

Types from introspecting the current stable `2022-01` schema are already implemented in [`types.go`](types.go). If, as above, you need to target a different API version, manually run `scripts/parse.go`, pointing to the schema file you'd like to use (using `go generate` directly won't be a good option until 1.18 final). So, presently:

```bash
go1.18beta2 run scripts/parse.go <PATH_TO_SCHEMA>
```

This will overwrite the existing `types.go`.

## Installing

```bash
go1.18beta2 get github.com/boatilus/storefront-go
```

## Basic Usage

```go
import (
    "log"
    
    "github.com/boatilus/storefront-go"
)

sf := storefront.NewClient("<DOMAIN>", "<ACCESS_TOKEN>")

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

A convenience function, `LoadQuery`, is also provided to make reading queries from the filesystem a bit easier. There is no syntax validation or otherwise -- it merely reads the file contents and returns a string you can then pass into `Query`.

```go
import "github.com/boatilus/storefront-go"

sf := storefront.NewClient("<DOMAIN>", "<ACCESS_TOKEN>")

q, err := storefront.LoadQuery("path/to/query.graphql")
if err != nil {
  // Handle
}

var set storefront.Set
if err := sf.Query(q, &set); err != nil {
  // Handle
}
```

## Running the Tests

This is a little complicated.

Currently, Shopify doesn't provide a container that'd let us spin up a local test store/Storefront endpoint, so to ensure types are being generated correctly, we really have to rely on end-to-end tests against real (or _real demo_) stores.

To set up the environment for passing tests, do the following:

### Specify Credentials in .env

Create a .env in the project root:

```env
SHOPIFY_DOMAIN=<YOUR_STORE_DOMAIN>
SHOPIFY_STOREFRONT_ACCESS_TOKEN=<YOUR_ACCESS_TOKEN>
```

### Specify Expected Values in expected.json

Create an `expected.json` in the `test` directory, providing values from your store. A [test/expected.schema.json](schema) for `expected.json` is included for reference. Minimally, `expected.json` should contain something like:

```json
{
  "collections": {
    "title": "Home page",
    "description": "Basic description"
  },
  "shop": {
    "name": "Shop name"
  }
}
```

### Run the Test Command

You can then run `go test` using something like [godotenv](https://github.com/joho/godotenv) to inject `.env` into the environment (or specify them in whatever way you prefer):

```bash
godotenv go test
# or, currently:
# godotenv go1.18beta2 test
```
