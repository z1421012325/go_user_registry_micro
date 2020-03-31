package main

import (
	"log"
	"net/http"
	"project/utils/token"
	"strconv"
	"strings"

	// 监控网关
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/micro/cli"
	"github.com/micro/micro/api"
	"github.com/micro/micro/cmd"
	"github.com/micro/micro/plugin"
)


func init() {
	namespace := ""
	subsystem := ""
	hidden := false
	// plugin.Register()
	_ = api.Register(
		plugin.NewPlugin(
			// 网关名称
			plugin.WithName("gateway"),

			// 网关option选项 添加了hidden选项
			plugin.WithFlag(
				cli.BoolFlag{
					Name:"hidden",
					Usage:"Hidden log",
				}),

			// 是否运行网关是选择该参数  比如 micro api --hidden 存在该参数,将hidden变为true
			plugin.WithInit(func(c *cli.Context) error {
				hidden = c.Bool("hidden")
				return nil
			}),
			plugin.WithCommand(),

			// http hand;er 基本网关重要功能都在这里实现
			plugin.WithHandler(func(h http.Handler) http.Handler {

				// return handler之前都属于三方库prometheus监控网关运行状态
				md := make(map[string]string)
				reqTotalCounter := prometheus.NewCounterVec(
					prometheus.CounterOpts{
						Namespace: namespace,
						Subsystem: subsystem,
						Name:      "request_total",
						Help:      "Total request count.",
					},
					[]string{"host", "status"},
				)
				reg := prometheus.NewRegistry()
				wrapreg := prometheus.WrapRegistererWith(md, reg)
				wrapreg.MustRegister(
					prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}),
					prometheus.NewGoCollector(),
					reqTotalCounter,
				)
				prometheus.DefaultGatherer = reg
				prometheus.DefaultRegisterer = wrapreg


				// return h
				return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {

					log.Println("            \n")
					log.Println("request url path : ",r.URL.Path)

					//cors
					w.Header().Set("Access-Control-Allow-Origin", "*")
					w.Header().Set("Access-Control-Allow-Credentials", "true")
					w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
					w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

					// todo 如果请求路径是login或者注册registry 那么直接通过
					paths := strings.Split(r.URL.Path,"/")
					for _, path := range paths {
						if path == "login" || path == "registry" || path == "register" {
							h.ServeHTTP(w,r)
							return
						}
						// 拦截metrics path，默认"/metrics"  进入promhttp监控
						if path == "metrics" {
							promhttp.Handler().ServeHTTP(w, r)
							return
						}
					}

					// todo 鉴权   请求中header和get与post请求中求得token
					tokenStr := r.Header.Get("Authorization")
					//data,ok := utils.CheckToken(tokenStr)
					if len(tokenStr) <= 0 {
						_ = r.ParseForm()
						tokenStr = r.FormValue("token")
					}
					_,ok := token.CheckToken(tokenStr)
					if !ok {
						w.WriteHeader(http.StatusUnauthorized)
						return
					}
					// r.Header.Set("data",data)

					if hidden {
						h.ServeHTTP(w,r)
						return
					}

					log.Println(r.URL.Path)

					ww := WrapWriter{ResponseWriter: w}
					h.ServeHTTP(&ww, r)

					reqTotalCounter.WithLabelValues(r.Host, strconv.Itoa(ww.StatusCode)).Inc()

				})
			}),
			))
}

//https://github.com/micro-in-cn/learning-videos/blob/master/docs/micro-api/example/main_02.go
func main() {
	cmd.Init()
}



type WrapWriter struct {
	StatusCode  int
	Size        int64
	wroteHeader bool

	http.ResponseWriter
}




func (ww *WrapWriter) WriteHeader(statusCode int) {
	ww.wroteHeader = true
	ww.StatusCode = statusCode
	ww.ResponseWriter.WriteHeader(statusCode)
}

func (ww *WrapWriter) Write(b []byte) (n int, err error) {
	// 默认200
	if !ww.wroteHeader {
		ww.WriteHeader(http.StatusOK)
	}

	n, err = ww.ResponseWriter.Write(b)
	ww.Size += int64(n)
	return
}

// RequestSize returns the size of request object.
func RequestSize(r *http.Request) float64 {
	size := 0
	if r.URL != nil {
		size = len(r.URL.String())
	}

	size += len(r.Method)
	size += len(r.Proto)

	for name, values := range r.Header {
		size += len(name)
		for _, value := range values {
			size += len(value)
		}
	}
	size += len(r.Host)

	// r.Form and r.MultipartForm are assumed to be included in r.URL.
	if r.ContentLength != -1 {
		size += int(r.ContentLength)
	}
	return float64(size)
}

func cors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
}