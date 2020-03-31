package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	proto "project/micros/proto/account"

	api "github.com/micro/go-micro/api/proto"
)


// 测试gin和micro给我们的api能否一样执行
func (this Account)Logout(ctx context.Context,req *api.Request,res *api.Response)  {
	fmt.Println("micro api request ~~~",req.Method)

	var tokenStr string
	for _, v := range req.Header {
		if strings.ToLower(v.GetKey()) == "authorization" {
			tokenStr = v.GetValues()[0]
		}
	}
	if len(tokenStr) == 0 {
		tokenStr = req.Get["token"].Values[0]
	}

	srvreq := &proto.LogoutAccounRequest{
		Token:                tokenStr,
	}

	srvres ,err := this.Client.Logout(context.TODO(),srvreq)
	if err != nil {
	}

	data ,_ := json.Marshal(srvres)

	res.StatusCode = 200
	res.Body = string(data)
}
