package breaker

import (
	"context"
	"log"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/client"
)



type clientWrapper struct {
	client.Client
}


func (c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {

	// hystrix 修改默认值
	hystrix.DefaultVolumeThreshold = 3
	hystrix.DefaultErrorPercentThreshold = 75
	hystrix.DefaultTimeout = 500
	hystrix.DefaultSleepWindow = 3500

	return hystrix.Do(req.Service()+"."+req.Method(), func() error {
		return c.Client.Call(ctx, req, rsp, opts...)
	}, func(err error) error {
		log.Printf("fallback error!!!!!  %v", err)

		// 设置一个通用返回,比如将res转换为通用返回struct结构体,制作另一个函数返回该结构体的数据,res = 该数据
		/*
		ErrRes := res.(XXXX.xxx)
			ErrRes = defaultXXXXX
		 */
		return err
	})
}

func NewHystrixClientWrapper() client.Wrapper {
	return func(c client.Client) client.Client {
		return &clientWrapper{c}
	}
}
