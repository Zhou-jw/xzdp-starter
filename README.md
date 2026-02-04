### 项目结构
```
xzdp-go-starter/
├── config/               # 配置层：环境隔离配置（开发/测试/生产）
│   ├── dev/            # 开发环境配置
│   ├── test/           # 测试环境配置
│   └── prod/           # 生产环境配置
├── data/               # docker挂载目录
│   ├── mysql/          # MySQL持久化数据（映射容器内/var/lib/mysql）
│   ├── redis/          # Redis持久化数据（映射容器内/data）
│   └── nginx/          # nginx 运行日志
├── docs/               # 项目文档/笔记
├── frontend/           # 前端代码/nginx 配置目录
├── scripts/            # 脚本目录：初始化、启动、部署脚本
│   ├── xzdp.sql/       # Redis持久化数据（映射容器内/data）
│   └── init_go.sh      # go 环境初始化脚本
├── docker-compose.yaml # Docker Compose配置：一键启动MySQL+Redis（绑定data目录）
├── .gitignore          # Git忽略文件：忽略data、编译文件、配置密文等
├── go.mod              # Go模块依赖
├── go.sum              # 依赖版本锁定
└── README.md           # 项目说明jk
```

### 快速开始
1. 使用wsl2 + docker快速部署
```shell
# -d 代表后台启动
docker-compose up -d

# 查看运行状态
docker-compose ps

# 查看运行日志
docker-compose logs -f mysql
```

启动成功之后，浏览器可以访问 `localhost:8080` 来查看页面

修改 `nginx.conf` 文件之后，需要热重启`nginx`
```
docker exec -it xzdp-nginx nginx -s reload
```

2. 安装herzt框架

```shell
go env -w GOPROXY=https://goproxy.cn
go install github.com/cloudwego/hertz/cmd/hz@latest
```
把 `$GOPATH/bin` 添加到环境变量
```shell
cd src
# 初始化项目代码
hz new 
# 更新mod, go.mod 中添加需要忽略的文件夹 ignore ./data
go mod tidy
```