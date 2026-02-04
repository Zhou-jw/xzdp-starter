```
xzdp-go-starter/
├── api/                # 接口层：API定义、请求/响应结构体、路由注册
│   └── handler/        # 接口处理器：处理HTTP请求，调用服务层
├── cmd/                # 程序入口：项目启动主函数
│   └── server/         # 服务入口：main.go（唯一启动入口）
├── conf/               # 配置层：环境隔离配置（开发/测试/生产）
│   ├── dev/            # 开发环境配置
│   ├── test/           # 测试环境配置
│   └── prod/           # 生产环境配置
├── data/               # 持久化数据目录：mysql/redis数据（.gitignore忽略）
│   ├── mysql/          # MySQL持久化数据（映射容器内/var/lib/mysql）
│   └── redis/          # Redis持久化数据（映射容器内/data）
├── internal/           # 内部业务层：项目核心逻辑（外部不可引用）
│   ├── dao/            # 数据访问层：操作MySQL/Redis，封装数据库逻辑
│   ├── model/          # 数据模型层：数据库表结构体、常量定义
│   ├── service/        # 业务服务层：核心业务逻辑，协调DAO和其他服务
│   └── util/           # 内部工具类：项目通用工具（加密、校验、日志等）
├── middleware/         # 中间件层：全局中间件（限流、日志、跨域、鉴权等）
├── pkg/                # 公共包：跨项目可复用工具（外部可引用，如Redis/MySQL客户端）
│   ├── mysql/          # MySQL通用客户端：初始化、连接池配置
│   └── redis/          # Redis通用客户端：初始化、连接配置
├── script/             # 脚本目录：初始化、启动、部署脚本
│   └── init.sql        # MySQL初始化脚本：创建库、表结构
├── docker-compose.yaml # Docker Compose配置：一键启动MySQL+Redis（绑定data目录）
├── .gitignore          # Git忽略文件：忽略data、编译文件、配置密文等
├── go.mod              # Go模块依赖
└── go.sum              # 依赖版本锁定
```