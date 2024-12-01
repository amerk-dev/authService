package main

import (
	"authService/internal/frameworks"
	"authService/internal/store"
)

func main() {
	store.DataBaseInit()

	server := frameworks.NewServer()
	server.Start()
}
