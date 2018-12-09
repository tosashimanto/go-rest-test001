package microservice_tranport

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/tosashimanto/go-rest-test001/microservice_tranport/gateway"
	"github.com/tosashimanto/go-rest-test001/microservice_tranport/users"
	"net/http"
	"os"
	"time"
)

func ConnectMicroService() {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// Service Proxy作成
	userproxy := gateway.NewUserService([]string{
		"localhost:8081",
	})
	userproxy = users.Logging(logger, "localhost:8000")(userproxy)

	r := mux.NewRouter()
	// microサービスのClientエンドポイント作成
	// 作成Clientエンドポイントを httptransport を使ってhandlerをbind
	r.Handle("/users/:name", httptransport.NewServer(
		// Endpoint作成 : Proxy指定
		users.NewEndpoints(userproxy).UserByName,
		// DecodeRequestFunc
		decodeUserByNameRequest,
		// EncodeResponseFunc
		encodeResponse),
	)

	//serverApiGateway := &http.Server{
	//	Handler: r,
	//	Addr:    "127.0.0.1:8000",
	//	// Good practice: enforce timeouts for servers you create!
	//	WriteTimeout: 15 * time.Second,
	//	ReadTimeout:  15 * time.Second,
	//}

	usrv := users.NewService()
	userinstances := []string{"localhost:8081"}
	for _, v := range userinstances {
		r := mux.NewRouter()
		srv := users.Logging(logger, v)(usrv)
		r.Handle("/:name", httptransport.NewServer(
			users.NewEndpoints(srv).UserByName,
			decodeUserByNameRequest,
			encodeResponse))
		s := &http.Server{
			Handler: r,
			Addr:    v,
			// Good practice: enforce timeouts for servers you create!
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}
		// user Service順次起動
		go func(ins string) {
			logger.Log("start", ins)
			logger.Log(s.ListenAndServe())
		}(v)
	}

	//// 127.0.0.1:8000 serverApiGateway起動
	//logger.Log(serverApiGateway.ListenAndServe())
}

func decodeUserByNameRequest(_ context.Context, r *http.Request) (interface{}, error) {
	if name, ok := mux.Vars(r)["name"]; ok {
		return name, nil
	}
	return nil, errors.New("name is required")
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
