package controller

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tosashimanto/go-rest-test001/controller/handler"
	"github.com/tosashimanto/go-rest-test001/micro_service/gateway"
	"github.com/tosashimanto/go-rest-test001/micro_service/users"
	"net/http"
	"os"
	"time"
)

var Server *echo.Echo

// init echo web server
func Init() {

	Server = echo.New()
	// Middlewareセット
	Server.Use(middleware.Logger())
	Server.Use(middleware.Recover())

	// CORSセット
	Server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	api := Server.Group("/rest_api/v1")

	// 1. トークン生成
	api.POST("/auth/token", handler.GetToken)

	// 2. 一覧取得
	api.GET("/list", handler.Constructions)
	// 3. 図面一覧取得
	api.GET("/list/:construction_id/drawings", handler.GetDrawings)
	// 4. Pile一覧取得
	api.GET("/list/:construction_id/piles", handler.Piles)
	// 5. 杭変更依頼
	// 6. 判定レコード作成
	api.POST("/judgements", handler.PostJudge)
	// 7. 判定履歴取得
	// 8. 判定実行
	api.PUT("/judgements/:judgement_id", handler.PutJudge)
	// 9. チェック結果登録
	api.PUT("/judgements/:judgement_id/check_results", handler.PostReJudge)
	// 10. チェック項目取得
	api.GET("/common/check_items", handler.GetCheckSheet)

	// 擬似S3 uplaod
	api.PUT("/upload", handler.UploadImage)

	// Server.Router().Add(http.MethodGet, "/users/:name",micro_service.TestServiceHandler)
	// api.GET("/users/:name", micro_service.TestServiceHandler)
	// Server.Use(micro_service.CustomMiddleware("Test"))
}

func TestMicroService() {

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
	r.Handle("/users/{name}", httptransport.NewServer(
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
		r.Handle("/{name}", httptransport.NewServer(
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
