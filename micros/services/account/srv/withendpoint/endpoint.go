package withendpoint
//
//import (
//	"github.com/micro/go-micro/api"
//	rapi "github.com/micro/go-micro/api/handler/api"
//	"github.com/micro/go-micro/server"
//)
//
//
//var (
//	Endpoint1 server.HandlerOption
//	Endpoint2 server.HandlerOption
//	Endpoint3 server.HandlerOption
//	Endpoint4 server.HandlerOption
//	Endpoint5 server.HandlerOption
//
//)
//
//func init(){
//
//	Endpoint1 = api.WithEndpoint(&api.Endpoint{
//		// 接口方法，一定要在proto接口中存在，不能是类的自有方法
//		Name: "account.registry",
//		// http请求路由，支持POSIX风格
//		Path: []string{"/registry"},
//		// 支持的方法类型
//		Method: []string{"POST", "GET"},
//		// 该接口使用的API转发模式
//		Handler: rapi.Handler,
//	})
//
//	Endpoint2 = api.WithEndpoint(&api.Endpoint{
//		// 接口方法，一定要在proto接口中存在，不能是类的自有方法
//		Name: "account.verify",
//		// http请求路由，支持POSIX风格
//		Path: []string{"/verify"},
//		// 支持的方法类型
//		Method: []string{"POST", "GET"},
//		// 该接口使用的API转发模式
//		Handler: rapi.Handler,
//	})
//
//	Endpoint3 = api.WithEndpoint(&api.Endpoint{
//		// 接口方法，一定要在proto接口中存在，不能是类的自有方法
//		Name: "account.login",
//		// http请求路由，支持POSIX风格
//		Path: []string{"/login"},
//		// 支持的方法类型
//		Method: []string{"POST", "GET"},
//		// 该接口使用的API转发模式
//		Handler: rapi.Handler,
//	})
//
//	Endpoint4 = api.WithEndpoint(&api.Endpoint{
//		// 接口方法，一定要在proto接口中存在，不能是类的自有方法
//		Name: "account.logout",
//		// http请求路由，支持POSIX风格
//		Path: []string{"/logout"},
//		// 支持的方法类型
//		Method: []string{"POST", "GET"},
//		// 该接口使用的API转发模式
//		Handler: rapi.Handler,
//	})
//
//	Endpoint5 = api.WithEndpoint(&api.Endpoint{
//		// 接口方法，一定要在proto接口中存在，不能是类的自有方法
//		Name: "account.test",
//		// http请求路由，支持POSIX风格
//		Path: []string{"/test"},
//		// 支持的方法类型
//		Method: []string{"POST", "GET"},
//		// 该接口使用的API转发模式
//		Handler: "http",
//	})
//
//}