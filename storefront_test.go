package storefront

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// _expected describes the structure of the test/expected.json file.
type _expected struct {
	Collections struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"collections"`
	Shop struct {
		Name string `json:"name"`
	}
}

var expected _expected

func init() {
	bs, err := ioutil.ReadFile(path.Join("test", "expected.json"))
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(bs, &expected); err != nil {
		log.Fatal(err)
	}
}

func TestNewClient(t *testing.T) {
	assert := assert.New(t)

	domain := "DOMAIN"
	accessToken := "API_KEY"

	t.Run("DefaultHTTPClient", func(t *testing.T) {
		c := NewClient(domain, accessToken)

		assert.NotNil(c)
		assert.Equal(
			fmt.Sprintf(
				"https://%s/api/%s/graphql.json",
				domain,
				APIVersion,
			),
			c.endpoint,
		)
		assert.Equal(accessToken, c.accessToken)
	})

	t.Run("CustomHTTPClient", func(t *testing.T) {
		httpClient := http.DefaultClient

		c := NewClient(domain, accessToken, httpClient)
		assert.NotNil(c)
		assert.Equal(
			fmt.Sprintf(
				"https://%s/api/%s/graphql.json",
				domain,
				APIVersion,
			),
			c.endpoint,
		)
		assert.Equal(accessToken, c.accessToken)
		assert.Equal(httpClient, c.HTTPClient)
	})
}

func TestClient_Query(t *testing.T) {
	assert := assert.New(t)

	c := NewClient(
		os.Getenv("SHOPIFY_DOMAIN"), os.Getenv("SHOPIFY_STOREFRONT_ACCESS_TOKEN"),
	)

	t.Run("Collections", func(t *testing.T) {
		q, err := LoadQuery(path.Join("test", "collections_query.graphql"))
		if err != nil {
			t.Fatal(err)
		}

		var set Set

		assert.NoError(c.Query(q, &set))
		assert.NotNil(set)

		edges := set.Data.Collections.Edges
		assert.GreaterOrEqual(len(edges), 1)

		node := edges[0].Node
		assert.NotEmpty(node.Id)
		assert.Equal(expected.Collections.Description, node.Description)
		assert.Equal(expected.Collections.Title, node.Title)
		assert.False(time.Time{}.Equal(node.UpdatedAt))
	})

	t.Run("Shop", func(t *testing.T) {
		q, err := LoadQuery(path.Join("test", "shop_query.graphql"))
		if err != nil {
			t.Fatal(err)
		}

		var set Set

		assert.NoError(c.Query(q, &set))
		assert.NotNil(set)
		assert.Equal(expected.Shop.Name, set.Data.Shop.Name)
	})
}

func TestLoadQuery(t *testing.T) {
	assert := assert.New(t)

	t.Run("OK", func(t *testing.T) {
		filename := path.Join("test", "collections_query.graphql")

		q, err := LoadQuery(filename)
		assert.NotEmpty(q)
		assert.NoError(err)
	})

	t.Run("ErrEmptyFilename", func(t *testing.T) {
		_, err := LoadQuery("")
		assert.ErrorIs(err, ErrEmptyFilename)
	})

	t.Run("ReadFileErr", func(t *testing.T) {
		_, err := LoadQuery("nonexistent-file")
		assert.Error(err)
	})
}
