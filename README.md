# Iris 通用Web开发模板

``` 
.
├── README.md
├── config   // 配置层
│   └── config.go
├── datamodels  # 模型层
├── datasource  # 数据源
│   ├── pgsql_conn     # pgsql 链接池
│   │   └── pgsql.go
│   └── redis_conn     # redis 连接池
│       └── redis.go
├── go.mod
├── go.sum
├── repositories # 操作层
├── services     # 服务层
└── web          # web 层
    ├── config.yml  # 配置文件
    ├── controllers # 控制器
    ├── main.go     # 入口
    ├── middleware  # 中间件
    │   └── global.go  # 全局中间件
    ├── register    # 注册器
    │   └── iris_register.go
    ├── router      # 路由层
    │   └── app.go      # 主路由
    └── views       # 视图层
```