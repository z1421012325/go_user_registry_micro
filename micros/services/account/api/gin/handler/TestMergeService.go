package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"project/micros/proto/account"
	"project/micros/services/account/api/gin/callsrv"
)




// 聚合api层,测试 登录之后再次call account服务中的test 测试token
func MergeService(c *gin.Context){

	var user account.LoginAccountRequest
	if err := c.ShouldBind(&user); err!= nil {
		c.JSON(301,nil)
		return
	}

	srvres,err := callsrv.AccountClient.Login(context.TODO(),&user)
	if err != nil {

	}

	testToken := account.TestAccountRequest{
		Token:                srvres.GetToken(),
	}
	testSrv ,err := callsrv.AccountClient.Test(context.TODO(),&testToken)
	if err != nil {

	}
	data := make(map[string]interface{})
	data["token"] = srvres.GetToken()
	data["istoken"] = testSrv.GetMsg()

	c.JSON(200,data)
}