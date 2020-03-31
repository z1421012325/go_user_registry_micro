package account

import (
	"time"
	"project/utils"
)

/*
   用户id
   昵称
   账号  手机号码或者email
   密码
   注册手段 比如 手机号码,email
   注册状态
   注册时间
   注销时间(软删除)


     mysql创建语句:
 		 create table user_account (
		`id` int PRIMARY KEY AUTO_INCREMENT COMMENT  "用户id",
		`nickname` varchar(20) not null COMMENT  "昵称",
		`account_number` varchar(20) COMMENT  "账号",
		`password` text COMMENT  "密码",
		`methor` int COMMENT "注册方式or登录方式  手机号码为1 email为2",
		`status` int DEFAULT 0 COMMENT "注册状态 未注册0 注册1 删除2",
		`create_at` datetime DEFAULT now() COMMENT "注册时间",
		`delete_at` datetime COMMENT  "删除时间",
		INDEX only_account_number (account_number),
		INDEX only_nickname (nickname)
		)ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8;
 */

const (
	DefaultRegistrationMethorToMobilePhone = 1		// 手机号码注册
	MobilePhone							   = "phone"
	DefaultRegistrationMethorToWeChat 	   = 2		// 微信号码注册
	WeChat								   = "wx"
)

const (
	DefaultAccountStatus = 0				// 默认注册用户未激活
	NormalAccountStatus  = 1				// 激活用户
	DeleteAccountStatus  = 2				// 封禁用户
)

const DefaultEncryptionLevel  = 12			// max = 15



type UserAccount struct {
	Id 					int			`gorm:"column:id" json:"id"`
	NickName 			string		`gorm:"column:nickname" json:"nk"`
	AccountNumber 		string		`gorm:"column:account_number" json:"an"`
	Password 			string		`gorm:"column:password" json:"pswd"`
	RegistrationMethor 	int			`gorm:"column:methor" json:"mt"`
	Status 				int			`gorm:"column:status" json:"status"`
	Create_at 			time.Time	`gorm:"column:create_at" json:"ct"`
	Delete_at 			time.Time	`gorm:"column:delete_at" json:"dt"`
}

func (this UserAccount)TableName() string{
	return "user_account"
}



func (this *UserAccount)CipherEncryption() bool {
	/*
		加密
	 */
	EncryptedStr := utils.Encryption(this.Password,DefaultEncryptionLevel)
	if EncryptedStr == "" {
		return false
	}

	this.Password = EncryptedStr
	return true
}

func (this *UserAccount)CheckPassword(out string) bool {
	if !utils.CheckEncryption(this.Password,out){
		return false
	}
	return true
}


func (this *UserAccount)PassPassword(out string)  {
	this.Password = ""
}

func (this *UserAccount)IsRegistrationMethor(inputmethor string) int {
	if inputmethor == MobilePhone {
		return DefaultRegistrationMethorToMobilePhone
	}else if inputmethor == WeChat {
		return DefaultRegistrationMethorToWeChat
	}
	return 0
}