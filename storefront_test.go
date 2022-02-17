package storefront

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClient_Do(t *testing.T) {
	assert := assert.New(t)

	c := NewClient(
		os.Getenv("SHOPIFY_DOMAIN"), os.Getenv("SHOPIFY_STOREFRONT_ACCESS_TOKEN"),
	)

	var set Set

	assert.NoError(c.Query(`{
    collections(first: 1) {
      edges {
        node {
          id
					title
					description
					updatedAt
					metafields(namespace: "fields", first: 1) {
						edges {
							node {
								namespace
								key
								value
							}
						}
					}
        }
      }
    }
  }`, &set))

	assert.NotNil(set)

	edges := set.Data.Collections.Edges
	assert.GreaterOrEqual(len(edges), 1)

	node := edges[0].Node
	assert.NotEmpty(node.Id)
	assert.NotEmpty(node.Description)
	assert.NotEmpty(node.Title)
	assert.False(time.Time{}.Equal(node.UpdatedAt))

	metafieldEdges := node.Metafields.Edges
	assert.Len(metafieldEdges, 1)
	assert.Equal(metafieldEdges[0].Node.Namespace, "fields")
	assert.Equal(metafieldEdges[0].Node.Key, "primary")
}
