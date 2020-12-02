# go-common/app/service/go-micro
提供RPC内部服务
主要技术点：
1.go-micro
2.etcd
3.gin
4.hystrix-go 插件应用（提供熔断器使用、服务降级)
对应插件:"github.com/afex/hystrix-go/hystrix
5.wrapper 【装饰器模式使用】
6.调用http api的正确姿势
处理参数模型中的 json tag插件protoc-go-inject-tag）的使用
对应插件
github.com/favadi/protoc-go-inject-tag v1.1.0 // indirect
7.api  GateWay(网关)
8.测试gprc 方法
8.1 利用postman
8.2 利用 micro 命令
例如 micro --registry=etcd --registry_address=192.168.188.141:2379 list service
8.3 利用api  GateWay


