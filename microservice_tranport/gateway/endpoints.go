package gateway

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/tosashimanto/go-rest-test001/microservice_tranport/users"
)

type Endpoints struct {
	UserByName endpoint.Endpoint
}

func NewEndpoints(serviceURL string) *Endpoints {
	return &Endpoints{
		UserByName: makeUserByNameEndpoint(serviceURL),
	}
}

func makeUserByNameEndpoint(proxyURL string) endpoint.Endpoint {
	url, err := url.Parse(proxyURL)
	if err != nil {
		panic(err)
	}

	return httptransport.NewClient(
		"GET",
		url,
		encodeUserByNameRequest,
		decodeUserByNameResponse,
	).Endpoint()
}
func encodeUserByNameRequest(ctx context.Context, req *http.Request, v interface{}) error {
	name := v.(string)
	req.URL.Path += "/" + name
	return nil
}

func decodeUserByNameResponse(ctx context.Context, resp *http.Response) (interface{}, error) {
	user := &users.User{}
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}
	return user, nil
}
