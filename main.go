package main

import (
	"github.com/tosashimanto/go-rest-test001/controller"
	"github.com/tosashimanto/go-rest-test001/log"
	"github.com/tosashimanto/go-rest-test001/service"
)

func main() {

	// 構成定義セット
	service.SetConfig()

	// 初期化してログ出力
	log.Init()
	log.AppLog.Info("サーバーが起動しました")

	// init server
	controller.Init()

	// run server
	controller.Server.Logger.Fatal(controller.Server.Start(":8080"))

	// microservice start
	// micro_service.ConnectMicroService()
}
