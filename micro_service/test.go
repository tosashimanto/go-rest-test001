package micro_service

import (
	"fmt"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/labstack/echo"
	"github.com/tosashimanto/go-rest-test001/micro_service/gateway"
	"github.com/tosashimanto/go-rest-test001/micro_service/users"
	"net/http"
	"os"
	"time"
)

func CustomMiddleware(name string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

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
			userproxy = users.Logging(logger, "localhost:8080")(userproxy)

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

			defer fmt.Printf("middleware-%s: defer\n", name)
			fmt.Printf("middleware-%s: before\n", name)
			err := next(c)
			fmt.Printf("middleware-%s: after\n", name)
			return err
		}
	}
}

func TestServiceHandler(c echo.Context) error {

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
	// userproxy = users.Logging(logger, "localhost:8080")(userproxy)

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

	return c.JSON(http.StatusOK, nil)
}

//func Test() {
//
//	var logger log.Logger
//	{
//		logger = log.NewLogfmtLogger(os.Stdout)
//		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
//		logger = log.With(logger, "caller", log.DefaultCaller)
//	}
//
//	usrv := users.NewService()
//	srv := users.Logging(logger, v)(usrv)
//
//	ser := httptransport.NewServer(
//		users.NewEndpoints(srv).UserByName,
//		decodeUserByNameRequest,
//		encodeResponse)
//
//}
