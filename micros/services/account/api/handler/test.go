package handler

import (
	"context"
	"encoding/json"
	"fmt"

	proto "project/micros/proto/account"

	api "github.com/micro/go-micro/api/proto"
)

// 测试gin和micro给我们的api能否一样执行
func (this Account)Test(ctx context.Context,req *api.Request,res *api.Response)  {

	fmt.Println("micro api request ~~~")

	fmt.Println(req.Method)

	srvreq := &proto.RegistryAccountRequest{
		Nickname:             req.Get["nickname"].Values[0],
		Username:             req.Get["username"].Values[0],
		Password:             req.Get["password"].Values[0],
		Method:               req.Get["method"].Values[0],
	}

	srvres ,err := this.Client.Registry(context.TODO(),srvreq)
	if err != nil {
	}

	data ,_ := json.Marshal(srvres)

	res.StatusCode = 200
	res.Body = string(data)
}