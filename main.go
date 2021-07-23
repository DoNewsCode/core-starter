package main

import (
	"github.com/DoNewsCode/core-starter/internal/bootstrap"
	"log"
	"math/rand"
	"time"
)

var (
	Name    = "starter"
	Version = "v1.0.0"
)

func main() {
	// setup rand seeder
	rand.Seed(time.Now().UnixNano())

	// Core Bootstrap
	root, shutdown := bootstrap.Bootstrap(Name, Version)

	// Shutdown
	defer shutdown()

	// Run
	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
