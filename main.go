package main

import (
	"github.com/nfangxu/core-skeleton/bootstrap"
)

func main() {
	// Core Bootstrap
	root, shutdown := bootstrap.Bootstrap()

	// Shutdown
	defer shutdown()

	// Run
	(func() {
		_ = root.Execute()
	})()
}
