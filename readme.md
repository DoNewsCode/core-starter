### Skeleton

```bash
# Build
go build -o skeleton .

# Format
find ./ -name "*.go" | xargs gofmt -l -w
```

> 目录结构

```
├── app                     # 项目主目录
│   ├── kernel.go           # 项目基础核心配置
│   ├── wire.go             # go wire 定义文件
│   └── wire_gen.go         # go wire 生成文件
├── bootstrap               # 启动目录: 项目与 Core 框架核心关联
│   └── bootstrap.go        # 项目与 Core 的核心关联
├── cmd                     # 命令目录: 存放所有的命令定义
│   ├── root.go             # 项目根命令
│   └── version.go          # 显示版本命令
├── config.example.yaml     # 示例配置文件
├── config.yaml             # 配置文件
├── go.mod                  # go module
├── go.sum                  # go module
├── main.go                 # 项目入口
└── readme.md               # Readme
```
