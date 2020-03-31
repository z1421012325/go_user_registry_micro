package ratelimit

import (
	"time"

	"github.com/juju/ratelimit"
	"github.com/micro/go-micro/client"
	micro_ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/ratelimit"
)


// default 每秒qps=500 容量cap=1000 wait容量满之后是否返回或者等待
func NewRatelimitWrap(qps int,cap int,wait bool) client.Wrapper {
	ratelimitBuckut := ratelimit.NewBucketWithRate(float64(qps),int64(cap))
	// false 当请求超过中cap不为空则直接失败,返回429 状态码
	return micro_ratelimit.NewClientWrapper(ratelimitBuckut,wait)
}

// 根据时间间隔来确定每次限流,比如3秒之内3000个请求限制
func NewTimeRatelimitWrap(t time.Duration,cap int,wait bool) client.Wrapper {
	ratelimitBuckut := ratelimit.NewBucket(t,int64(cap))
	return micro_ratelimit.NewClientWrapper(ratelimitBuckut,wait)
}