### Skeleton

```bash
# Export example config
go run main.go config init -o config.example.yaml

# Generate
go generate -x ./...

# Build
go build -o skeleton .

# Format
go fmt -x ./...
```

> 目录结构

```
├── app                     # 项目主目录
│   ├── commands            # 命令目录: 存放所有的命令定义
│   │   └── version.go      # 显示版本命令 (示例命令)
│   └── kernel              # 项目启动配置, 包括核心启动以及依赖注入等
│       ├── kernel.go       # 项目启动核心
│       ├── wire.go         # go wire 定义文件
│       └── wire_gen.go     # go wire 生成文件
├── bootstrap               # 启动目录: 项目与 Core 框架核心关联
│   ├── bootstrap.go        # 项目与 Core 的核心关联
│   └── root.go             # 项目根命令
├── config                  # 全局配置目录: 配置加载的配置信息, 路由等
│   └── app.go           # 配置模块加载与命令加载等
├── config.example.yaml     # 示例配置文件
├── config.yaml             # 配置文件
├── go.mod                  # go module
├── go.sum                  # go module
├── main.go                 # 项目入口
└── readme.md               # Readme
```
