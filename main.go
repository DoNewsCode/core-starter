package main

import (
	"github.com/nfangxu/core-skeleton/app"
	"github.com/nfangxu/core-skeleton/bootstrap"
	"github.com/nfangxu/core-skeleton/cmd"
)

func main() {
	// 启动基础, 将项目与 Core 连接
	root, c := bootstrap.Bootstrap()

	// 注册项目所需的一些东西
	app.Modules(c)
	app.Providers(c)

	// 添加一些 Command
	root.AddCommand(cmd.NewVersionCmd(c))

	// 运行
	(func() {
		defer c.Shutdown()
		_ = root.Execute()
	})()
}
