学习笔记
```
|- api（API协议定义目录，protobuf文件，以及生成的go文件。与internal找中的server进行交互）
|- configs（配置文件模板或默认配置，推荐使用ymal文件）
|- test（额外外部测试应用程序和测试数据）
    |- testdata
|- cmd（cmd目录负责程序的启动、关闭、配置初始化。App服务类型分为四类：interface、service、job、admin）
    |- myapp-admin（面向运营侧的服务）
        |- server
            |- http
            |- grpc
    |- myapp-interface（对外的BFF）
    |- myapp-job（流式任务处理的服务）
    |- myapp-service（对内微服务，仅接受内部服务或网关的调用）
    |- myapp-task（定时任务，类似cronjob，部署到task托管平台）
|- internal
    |- biz（业务逻辑的组装，类似DDD的domain层，data类似DDD的repo，repo的接口定义在这里，使用依赖倒置注入）
    |- data（业务数据访问，包含cache、db等封装，实现biz接口）
    |- service（实现api定义的服务层，类似DDD的application层，实现protobuf的接口，处理DTO到DO的转换，协同各类biz交互，但不处理复杂逻辑）
```

作业结构
```shell
.
├── MakeFile
├── README.md
├── api
│   └── userapp
│       └── v1
│           ├── userapp.pb.go
│           └── userapp.proto
├── cmd
│   └── userapp-admin
│       ├── main.go
│       ├── wire.go
│       └── wire_gen.go
├── configs
│   └── config.yaml
├── go.mod
├── go.sum
└── internal
    └── userapp
        ├── biz
        │   └── userapp.go
        ├── data
        │   └── userapp.go
        ├── server
        │   └── userapp.go
        └── service
            └── userapp.go

```