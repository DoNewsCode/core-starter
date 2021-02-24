package main

import (
	"github.com/nfangxu/core-skeleton/app"
	"github.com/nfangxu/core-skeleton/bootstrap"
	"github.com/nfangxu/core-skeleton/cmd"
)

func main() {
	// 启动基础, 将项目与 Core 连接
	root, c := bootstrap.Bootstrap()

	// setup global dependencies
	app.Providers(c)

	// setup global modules
	app.Modules(c)

	// register global commands from modules
	c.ApplyRootCommand(root)

	// register custom command
	root.AddCommand(cmd.NewVersionCmd(c))

	// 运行
	(func() {
		defer c.Shutdown()
		_ = root.Execute()
	})()
}
