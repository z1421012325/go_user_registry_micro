package main

import (
	"github.com/micro/go-micro/web"
	"project/micros/proto/account"
	"project/micros/services/account/api/gin/callsrv"
	"project/micros/services/account/api/gin/router"
	"project/registry"
)

func main() {
	server := web.NewService(
		web.Name("go.micro.api.account"),
		web.Registry(registry.DefaultEtcdRegistry()),

		web.Handler(router.NewGinHandler()),  //  加载路由
		)

	_ = server.Init()

	callsrv.AccountClient = account.NewAccountService("go.micro.srv.account",server.Options().Service.Client())

	// 加载路由
	// server.Handle("/",router.NewGinHandler())

	err := server.Run()
	if err != nil {
		panic(err)
	}
}
