package main

import (
	_ "github.com/micro/go-plugins/registry/consul"
	_ "github.com/micro/go-plugins/registry/etcd"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	_ "github.com/micro/go-plugins/registry/nats"

	_ "github.com/micro/go-plugins/transport/grpc"
	_ "github.com/micro/go-plugins/transport/http"
	_ "github.com/micro/go-plugins/transport/tcp"
)
