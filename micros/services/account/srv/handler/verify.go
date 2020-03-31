package handler


import (
	"context"
	"fmt"

	Account "project/micros/models/account"
	DB "project/micros/dbs/account"
	proto "project/micros/proto/account"
	"project/utils"
)

func (this account)Verify(ctx context.Context, in *proto.VerifyAccountRequest, out *proto.VerifyAccountResponse) error {

	fmt.Println(in.GetNickname())
	fmt.Println(in.GetUsername())
	fmt.Println(in.GetPassword())
	fmt.Println(in.GetMethod())
	fmt.Println(in.GetVerifycode())

	/*
		逻辑
			根据in.GetUsername()为key去redis中寻找验证码与in.GetVerifycode()对比
			构建mysql语句,where nickname = in.GetNickname(),username = in.GetUsername() , status != ...
				修改status
			返回注册信息
	*/

	if !utils.GetRegistryStr(in.GetUsername(),in.GetVerifycode()){
		out.Code = 301
		out.Msg = "验证码不对"
		return nil
	}


	sql := "UPDATE user_account SET status = ? WHERE nickname = ? AND account_number = ? AND methor = ?"
	db := DB.DB.Exec(sql,Account.NormalAccountStatus,in.GetNickname(),in.GetUsername(),in.GetMethod())
	if !DB.TransactionOperation(db){
		out.Code = 501
		out.Msg = "注册失败"
		return nil
	}


	out.Code = 200
	out.Msg = in.GetNickname() + " registry success!"
	return nil
}
