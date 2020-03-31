package main


import (
	"github.com/micro/go-plugins/micro/cors"
	"github.com/micro/micro/plugin"
	"github.com/micro/micro/cmd"
)

func init(){
	// 注册跨域插件
	if err := plugin.Register(cors.NewPlugin()); err != nil {
		panic(err)
	}
}

func main() {
	cmd.Init()
}

/*
go run main.go
--cors-allowed-headers="Content-Type,X-Token"
--cors-allowed-origins="*"
--cors-allowed-methods="OPTIONS,DELETE,GET,POST" api
 */