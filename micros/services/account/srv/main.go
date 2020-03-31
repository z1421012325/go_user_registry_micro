package main

import (
	"log"
	"time"

	proto "project/micros/proto/account"
	"project/micros/services/account/srv/handler"
	reg "project/registry"

	_ "project/db/redis"           // 公用redis初始化
	_ "project/init"               // 配置文件env初始化		配置文件优先init
	_ "project/micros/dbs/account" // 初始化mysql pool

	"github.com/micro/go-micro"
	_ "github.com/micro/go-micro/registry/etcd"
)


func main() {
	server := micro.NewService(
		micro.Name("go.micro.srv.account"),
		micro.RegisterTTL(time.Second*20),
		micro.RegisterInterval(time.Second*15),
		micro.Registry(reg.DefaultEtcdRegistry()),
		)

	server.Init()

	_ = proto.RegisterAccountHandler(server.Server(),handler.NewAccount())

	err := server.Run()
	if err != nil {
		log.Println(err)
	}
}
