### CoreStarter

[DoNewsCode/core](https://github.com/DoNewsCode/core) is a service container that elegantly bootstrap and coordinate twelve-factor apps in Go.
This is a starter template using the [DoNewsCode/core](https://github.com/DoNewsCode/core) framework.

[![Go Report Card](https://goreportcard.com/badge/github.com/DoNewsCode/core-starter)](https://goreportcard.com/report/github.com/DoNewsCode/core-starter)
[![Core Release](https://img.shields.io/github/release/DoNewsCode/core.svg)](https://github.com/DoNewsCode/core/releases/latest)

#### Usage

- 0x01: Generate

You can use the [fx](https://github.com/nfangxu/tools) command to quickly create a project.

```bash
fx create helloworld
```

Or

> When you use `Git` to generate a project, you need to modify the package name of the project yourself.

```bash
git clone https://github.com/DoNewsCode/core-starter.git helloworld
```

- 0x02: Config

When you use `go run main.go` to run the project, the starter uses `config.yaml` by default.
You can use the `--config` parameter to specify the config file, or use the following command to generate:

```bash
cp config.example.yaml config.yaml
```

When you add other modules, you can use the `go run main.go config init -o config.example.yaml` re-build the sample config file.

- 0x03: Happy to coding.

> You will see the following directory structure.

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

#### Examples of using other frameworks

[Gin](https://github.com/DoNewsCode/core-starter/tree/gin-http)

[Go Kit](https://github.com/DoNewsCode/core-starter/tree/go-kit)

[Kratos](https://github.com/DoNewsCode/core-starter/tree/kratos)
