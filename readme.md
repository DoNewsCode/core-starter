### CoreStarter

[DoNewsCode/core](https://github.com/DoNewsCode/core) is a service container that elegantly bootstrap and coordinate twelve-factor apps in Go.
This is a starter template using the [DoNewsCode/core](https://github.com/DoNewsCode/core) framework.

[![Go Report Card](https://goreportcard.com/badge/github.com/DoNewsCode/core-starter)](https://goreportcard.com/report/github.com/DoNewsCode/core-starter)
[![Core Release](https://img.shields.io/github/release/DoNewsCode/core.svg)](https://github.com/DoNewsCode/core/releases/latest)

#### Layouts

```
├── app
│   ├── commands
│   │   └── example.go
│   └── module.go               # App module
├── docs
│   ├── docsify                 # Docsify: {hostname:port}/docs/docsify
│   ├── swagger                 # Swagger Api: {hostname:port}/docs/swagger
│   └── module.go               # Docs module
├── internal
│   ├── bootstrap
│   │   ├── bootstrap.go
│   │   ├── bootstrap_test.go
│   ├── cmd
│   │   └── root.go
│   └── config
│       ├── option.go
│       └── register.go         # Register global module or dependency
├── .gitignore
├── config.example.yaml         # config example file
├── config.yaml                 # config file
├── go.mod
├── go.sum
├── main.go
└── readme.md
```

#### Generate

You can use the [fx](https://github.com/nfangxu/tools) command to quickly create a project

```bash
fx create helloworld
# Or
fx create https://github.com/DoNewsCode/core-starter.git helloworld
# Or
fx create https://github.com/DoNewsCode/core-starter.git github.com/nfangxu/helloworld
```

#### Example

[Gin Http Example](https://github.com/DoNewsCode/core-starter/tree/gin-http)

[Go Kit Example](https://github.com/DoNewsCode/core-starter/tree/go-kit)

[Kratos Example](https://github.com/DoNewsCode/core-starter/tree/kratos)

#### Help

```bash
# Export example config
go run main.go config init -o config.example.yaml

# Config
cp config.example.yaml config.yaml

# Generate
go generate -x ./...

# Build
go build -o starter .

# Format
go fmt -x ./...

# Test
go test ./...

# lint
golint ./...
```
