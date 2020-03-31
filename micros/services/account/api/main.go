package main

import (
	"log"
	"project/micros/services/account/api/handler"
	"time"

	reg "project/registry"
	proto "project/micros/proto/account"
	"project/wrapper/breaker"
	"project/wrapper/ratelimit"

	"github.com/micro/go-micro"
)


func main() {
	server := micro.NewService(
		micro.Name("go.micro.api.test"),
		micro.RegisterTTL(time.Second*20),
		micro.RegisterInterval(time.Second*15),
		micro.Registry(reg.DefaultEtcdRegistry()), 						// etcd 注册中心

		micro.WrapClient(breaker.NewHystrixClientWrapper()),		    // hystrix 熔断机制
		micro.WrapClient(ratelimit.NewRatelimitWrap(500,1000,false)),  //ratelimit 限流机制
		)

	server.Init()

	server.Server().Handle(
		server.Server().NewHandler(
			&handler.Account{Client: proto.NewAccountService("go.micro.srv.account", server.Client())},
		),
	)

	err := server.Run()
	if err != nil {
		log.Println(err)
	}
}


