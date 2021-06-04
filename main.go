package main

import (
	"github.com/DoNewsCode/core-starter/internal/bootstrap"
	"log"
	"math/rand"
	"time"
)

func main() {
	// Recover panic
	defer func() {
		if err := recover(); err != nil {
			log.Printf("panic(%v) \n", err)
		}
	}()

	// Setup rand seeder
	rand.Seed(time.Now().UnixNano())

	// Core Bootstrap
	root, shutdown := bootstrap.Bootstrap()

	// Shutdown
	defer shutdown()

	// Run
	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
