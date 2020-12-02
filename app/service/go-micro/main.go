package main

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"hym-go-comm/app/service/go-micro/services/proto"
	servicesImpl"hym-go-comm/app/service/go-micro/servicesImpl"

)

func  main()  {
	// 定义 etcd 的注册
	// 其实这里不需要在代码里写注册 etcd 的信息，只需要命令行添加即可 --registry=etcd --registry_address=127.0.0.1:2379
	reg := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.188.141:2379",
		}
	})
	service := micro.NewService(
		micro.Address("127.0.0.1:9901"), // 其实这里不需要在代码里写的，只需要命令行添加即可 --server_address=127.0.0.1:9900
		micro.Registry(reg),
		micro.Name("TestService.call.server"), // 当前微服务服务名
		micro.Version("v1.0.0"),

	)

	service.Init()

	err := proto.RegisterTestServiceHandler(service.Server(),new(servicesImpl.TestService))
	if err != nil {
		fmt.Println("RegisterTestServiceHandler:", err)
		return
	}

	err = service.Run()
	if err != nil {
		fmt.Println("Run:", err)
	}

}
