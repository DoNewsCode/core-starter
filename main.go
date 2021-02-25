package main

import (
	"github.com/nfangxu/core-skeleton/app"
	"github.com/nfangxu/core-skeleton/bootstrap"
	"github.com/nfangxu/core-skeleton/cmd"
)

func main() {
	// Core Bootstrap
	root, shutdown := bootstrap.Bootstrap(
		app.Providers,
		app.Modules,
		cmd.NewVersionCmd,
	)

	// Shutdown
	defer shutdown()

	// Run
	(func() {
		_ = root.Execute()
	})()
}
