package users

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

/**
 * Endpoint定義
 */
type Endpoints struct {
	UserByName endpoint.Endpoint
}

/**
 * UserServiceのEndpoint作成
 */
func NewEndpoints(srv UserService) *Endpoints {
	return &Endpoints{
		UserByName: makeUserByNameEndpoint(srv),
	}
}

/**
 * Endpoint作成
 * サービスをhttpハンドラにする
 *
 * 戻り値 Endpoint
 */
func makeUserByNameEndpoint(svc UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(string)
		user, err := svc.UserByName(ctx, req)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
}
