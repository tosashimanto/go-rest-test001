package controller

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tosashimanto/go-rest-test001/controller/handler"
	"github.com/tosashimanto/go-rest-test001/microservice_tranport"
	"github.com/tosashimanto/go-rest-test001/util"
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
	// 5. 変更依頼
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

	// micro service
	// Server.Router().Add(http.MethodGet, "/users/:name",microservice_tranport.TestServiceHandler)
	api.GET("/users/:name", microservice_tranport.TestServiceHandler)
	// Server.Use(microservice_tranport.CustomMiddleware("Test"))

	/**
	 * Route確認
	 */
	data, _ := json.MarshalIndent(Server.Routes(), "", "  ")
	util.JSONFormatOut(data)
}
