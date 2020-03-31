package handler

import (
	"context"
	"fmt"
	"project/utils/token"

	DB "project/micros/dbs/account"
	Account "project/micros/models/account"
	proto "project/micros/proto/account"
)

func (this account)Login(ctx context.Context, in *proto.LoginAccountRequest, out *proto.LoginAccountResponse) error {

	fmt.Println(in.GetUsername())
	fmt.Println(in.GetPassword())
	fmt.Println(in.GetMethod())

	/*
		逻辑
			查询该用户in.GetUsername()并为以注册用户,是否为空
				-为空:返回err,账号密码不对
				-不为空:查询的密码与请求的密码verify
			除密码之外的信息来生成token
			****redis保存,key为用户id,value为token -- 确保唯一性
				不确定做不做,唯一性会对redis产生大量请求
					每个请求经过网关,网关会对token解析是否过期,在根据token中的payload信息去redis检测),
						增加了响应时间)
	 */

	var user Account.UserAccount

	methor := user.IsRegistrationMethor(in.GetMethod())
	if methor == 0 {
		out.Code = 301
		out.Msg = "登录方式不正确"
		out.Token = ""
		return nil
	}

	DB.DB.Where("account_number = ? AND methor = ? AND status = ? ",in.GetUsername(),methor,Account.NormalAccountStatus).First(&user)
	if !user.CheckPassword(in.GetPassword()){
		out.Code = 301
		out.Msg = "账号或者密码不正确"
		out.Token = ""
		return nil
	}

	user.Password = ""
	if tokenString,ok := token.NewToken(user); !ok{
		out.Code = 301
		out.Msg = "token 获取错误"
		out.Token = ""
		return nil
	}else {
		out.Code = 200
		out.Msg = "get token is success"
		out.Token = tokenString
		return nil
	}
}
