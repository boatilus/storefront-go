package storefront

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
)

const API_VERSION = "2022-01"

// Connection describes a type for paginating through multiple objects. It
// essentially encompasses a list of edges, whose edges contain a node, which
// thereby contains the object data.
type Connection[T any] struct {
	// Edges is a list edges.
	Edges []Edge[T] `json:"edges,omitempty"`
	// PageInfo is a set of data to aid in pagination.
	PageInfo map[string]interface{} `json:"pageInfo,omitempty"`
}

// Edge describes a node/cursor pair.
type Edge[T any] struct {
	// Cursor is a cursor for use in pagination.
	Cursor string `json:"cursor,omitempty"`
	// Node is the item at the end of edge.
	Node T `json:"node,omitempty"`
}

// Set describes the full result set from a query. It encloses multiple result
// sets within an enclosed data property.
type Set struct {
	Data QueryRoot `json:"data,omitempty"`
}

// Client describes an HTTP client for the Storefront API with credentials.
type Client struct {
	endpoint    string
	accessToken string
	HTTPClient  *http.Client
}

// NewClient constructs a new instance of a Storefront client given the store
// domain and access token.
func NewClient(domain, accessToken string) *Client {
	endpoint := url.URL{
		Scheme: "https",
		Host:   domain,
		Path:   path.Join("api", API_VERSION, "graphql.json"),
	}

	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	return &Client{
		endpoint:    endpoint.String(),
		accessToken: accessToken,
		HTTPClient:  httpClient,
	}
}

func (c *Client) Query(q string, out interface{}) error {
	reader := strings.NewReader(q)

	req, err := http.NewRequest(http.MethodPost, c.endpoint, reader)
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/graphql")
	req.Header.Add("X-Shopify-Storefront-Access-Token", c.accessToken)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	bs, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bs, &out); err != nil {
		return err
	}

	return nil
}

var ErrInvalidFilename = errors.New("an invalid filename was specified")

// LoadQuery opens the specify file and returns a string value for the query.
func LoadQuery(filename string) (string, error) {
	if filename == "" {
		return "", ErrInvalidFilename
	}

	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(bs), err
}
