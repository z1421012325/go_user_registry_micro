package handler



type account struct {}


func NewAccount () account{
	return account{}
}


// 保存入token中的用户信息
type LoginToToken struct {
	Id 				int				`json:"id"`				// 用户id
	Nickname 		string			`json:"nickname"`		// 用户昵称
	AccountUser 	string			`json:"account_user"`	// 用户账号
}