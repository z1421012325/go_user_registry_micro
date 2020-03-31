package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	proto "project/micros/proto/account"

	api "github.com/micro/go-micro/api/proto"
)


// 测试gin和micro给我们的api能否一样执行
func (this Account)Verify(ctx context.Context,req *api.Request,res *api.Response)  {
	fmt.Println("micro api request ~~~",req.Method)

	verifycode,err := strconv.ParseInt(req.Get["code"].Values[0],10,32)

	srvreq := &proto.VerifyAccountRequest{
		Nickname:             req.Get["nickname"].Values[0],
		Username:             req.Get["username"].Values[0],
		Password:             req.Get["password"].Values[0],
		Method:               req.Get["method"].Values[0],
		Verifycode:           int32(verifycode),
	}

	srvres ,err := this.Client.Verify(context.TODO(),srvreq)
	if err != nil {
	}

	data ,_ := json.Marshal(srvres)

	res.StatusCode = 200
	res.Body = string(data)
}
