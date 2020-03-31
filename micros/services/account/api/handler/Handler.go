package handler

import (
	"encoding/json"
	"project/utils/token"
	"strings"

	"project/micros/proto/account"
	"project/micros/services/account/srv/handler"

	api "github.com/micro/go-micro/api/proto"
)


type Account struct {
	Client account.AccountService
}


func GetTokenData(req *api.Request) handler.LoginToToken {

	var ReqtokenStr string

	// 得到header和post中的token信息
	headerMap := req.GetHeader()
	for _, k := range headerMap {
		if strings.ToLower(k.GetKey()) == "authorization" {

			if len(k.GetValues()[0]) == 0 {
				postMap := req.GetPost()
				for _, v := range postMap {
					if strings.ToLower(v.GetKey()) == "token" {
						ReqtokenStr = v.GetValues()[0]
					}else {
						ReqtokenStr = ""
					}
				}
			}

			ReqtokenStr = k.GetValues()[0]
		}
	}

	tokenStr ,ok := token.CheckToken(ReqtokenStr)
	if !ok {
		return handler.LoginToToken{}
	}

	var data handler.LoginToToken
	err := json.Unmarshal([]byte(tokenStr),&data)
	if err != nil {
		return handler.LoginToToken{}
	}
	return data
}


