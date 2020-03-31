package handler

import (
	"context"

	"github.com/gin-gonic/gin"
	"project/micros/proto/account"
	"project/micros/services/account/api/gin/callsrv"
)

func RegistryAccount(c *gin.Context){

	var account account.RegistryAccountRequest
	if err := c.ShouldBind(&account); err != nil {
		c.JSON(301,"nil")
		return
	}

	srvres,err := callsrv.AccountClient.Registry(context.TODO(),&account)
	if err != nil {

	}

	c.JSON(200,srvres)
}
