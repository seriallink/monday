package monday

import (
	"context"
	"net/http"
	"reflect"

	"github.com/fatih/structs"
	"github.com/hasura/go-graphql-client"
)

const endpointUrl = "https://api.monday.com/v2"

const (
	ApiVersion202307 = "2023-07"
	ApiVersion202310 = "2023-10"
	ApoVersion202401 = "2024-01"
)

type Client struct {
	token  string
	client *graphql.Client
}

// NewClient creates a graphql client with the monday.com endpoint.
func NewClient(token, version string) *Client {
	return &Client{
		token: token,
		client: graphql.NewClient(endpointUrl, nil).WithRequestModifier(func(r *http.Request) {
			r.Header.Add("Cache-Control", "no-cache")
			r.Header.Add("Content-Type", "application/json")
			r.Header.Add("API-Version", version)
			r.Header.Add("Authorization", token)
		}),
	}
}

// RunQuery executes request and decodes response into response param (address of object)
func (c *Client) RunQuery(response interface{}, arguments Arguments) error {
	err := c.client.Query(context.Background(), response, arguments)
	return err
}

// ArgumentsToMap converts arguments struct to map[string]interface{} with nil values converted to *(T)(nil)
func ArgumentsToMap(arguments interface{}) map[string]interface{} {
	if reflect.ValueOf(arguments).Kind() == reflect.Ptr && reflect.ValueOf(arguments).IsNil() {
		arguments = reflect.New(reflect.TypeOf(arguments).Elem()).Interface()
	}
	return structs.Map(arguments)
}
