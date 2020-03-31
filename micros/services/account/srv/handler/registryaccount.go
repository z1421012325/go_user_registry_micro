package handler

import (
	"context"
	"fmt"
	DB "project/micros/dbs/account"
	model "project/micros/models/account"
	proto "project/micros/proto/account"
	"project/utils"
)

func (this account)Registry(ctx context.Context, in *proto.RegistryAccountRequest, out *proto.RegistryAccountResponse) error {


	fmt.Println(in.GetNickname())
	fmt.Println(in.GetUsername())
	fmt.Println(in.GetPassword())
	fmt.Println(in.GetMethod())

	/*
		逻辑:
			查询重复注册
			保存为未确认注册账户(密码加密)
			根据GetMethod选择短信或者email 调用sms短信发送或者sendemail,并在redis中存入验证码(账号为key,code为value)

			返回对应信息
	 */

	var account model.UserAccount
	DB.DB.Where("nickname = ? AND account_number = ?",in.GetNickname(),in.GetUsername()).First(&account)
	if account.NickName == in.GetNickname() {
		out.Code = 301
		out.Msg = "昵称已存在"
		return nil
	}else if account.AccountNumber == in.GetUsername() {
		out.Code = 301
		out.Msg = "账号已注册"
		return nil
	}

	RegistrationMethor := VerifyRegistryMethor(in)
	if RegistrationMethor == 0 {
		out.Code = 301
		out.Msg = "注册方式异常"
		return nil
	}

	nowAccount := model.UserAccount{
		//NickName:           in.GetNickname(),
		//AccountNumber:      in.GetUsername(),
		Password:           in.GetPassword(),
		//RegistrationMethor: RegistrationMethor,
		//Status:             model.DefaultAccountStatus,
	}

	// 加密
	if !nowAccount.CipherEncryption(){
		out.Code = 301
		out.Msg = "加密异常"
		return nil
	}

	sql := "insert into user_account (nickname,account_number,password,methor,status) values (?,?,?,?,?)"
	err := DB.DB.Exec(sql,in.GetNickname(),in.GetUsername(),nowAccount.Password,RegistrationMethor,model.DefaultAccountStatus).Error
	//err := DB.DB.Create(&nowAccount).Error
	// 异常 Incorrect datetime value: '0000-00-00 00:00:00' for column 'login_time' at row 1  直接使用sql语句来save
	if err != nil {
		out.Code = 401
		out.Msg = "注册异常"
		return nil
	}

	/*
		todo 邮件或者email
		生成短信码
		redis保存 k : 注册账号 v : 短信码
		异步发送 返回
	 */

	// 假设为email发送
	RegistryStrCode := utils.Random(6)
	go utils.SendRegistryEmail(in.GetUsername(),RegistryStrCode)
	utils.SaveRegistryStr(in.GetUsername(),RegistryStrCode)

	out.Code = 200
	out.Msg = "短信或者email已发送,请查收"

	return nil
}






func VerifyRegistryMethor(in *proto.RegistryAccountRequest) (status int){
	if in.GetMethod() == model.MobilePhone {
		status = model.DefaultRegistrationMethorToMobilePhone
		return
	}else if in.GetMethod() == model.WeChat {
		status = model.DefaultRegistrationMethorToWeChat
		return
	}
	return 0
}

