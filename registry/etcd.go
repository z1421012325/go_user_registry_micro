package registry

import (

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"

)

func DefaultEtcdRegistry() registry.Registry{
	reg := etcd.NewRegistry(func(o *registry.Options) {
		o.Addrs = []string{"localhost:2379"}
	})
	return reg
}

func InEtcdRegistry(addres ...string) registry.Registry {

	var address []string
	for _,v := range addres {
		address = append(address,v)
	}

	reg := etcd.NewRegistry(func(o *registry.Options) {
		o.Addrs = address
	})
	return reg
}
