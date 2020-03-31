package handler

import (
	"context"
	"fmt"
	proto "project/micros/proto/account"
)

func (this account)Logout(ctx context.Context, in *proto.LogoutAccounRequest, out *proto.LogoutAccountResponse) error {

	fmt.Println(in.GetToken())

	/*
		逻辑
			解密token得到其中用户信息,用户id去redis查询该key并删除
				2020.2.21 不做logout逻辑,唯一性对redis会产生大量请求 增加了响应时间
	*/


	out.Code = 200
	out.Msg = "logout account success!"

	return nil
}
