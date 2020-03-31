package handler

import (
	"context"
	"fmt"
	"project/utils/token"

	proto "project/micros/proto/account"
)

func (this account)Test(ctx context.Context, in *proto.TestAccountRequest, out *proto.TstAccountResponse) error {

	fmt.Println(in.GetToken())

	/*
		逻辑
			test测试
			加密解密token,并打印出来
	 */

	TokenData,ok := token.CheckToken(in.GetToken())
	if !ok{
		out.Code = 301
		out.Msg = "test work faild"
		return nil
	}

	fmt.Println(TokenData)

	out.Code = 200
	out.Msg = "test work success"
	return nil
}