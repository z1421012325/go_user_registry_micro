
```shell script
启动etcd注册中心
将代码中etcd地址改为etcd启动宿主机ip
```

```shell script
build-run

go build ./micros/services/srv/main.go -o srv
chmod +x srv

./srv 
or 命令行参数指定注册中心启动
../srv/main.go 中注释 micro.Registry(reg.DefaultEtcdRegistry())
./srv --registry=etcd  --registry_address=xxx.xx.xx.xx:2379 
```

```shell script
build-run

go build ./micros/services/api/main.go -o api
chmod +x api

./api
or 命令行参数指定注册中心启动
../api/main.go 中注释 micro.Registry(reg.DefaultEtcdRegistry())
./api --registry=etcd  --registry_address=xxx.xx.xx.xx:2379 
```

```shell script
build-run gatweay

go build gatweay.go gatweay_plugins.go -o gatweay

网关指定注册中心
gateway --registry=etcd  --registry_address=xxx.xx.xx.xx:2379 api --handler=api
```



