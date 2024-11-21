package main

import (
	"github.com/Kennedy-lsd/GoChat/cmd/api"
	"github.com/Kennedy-lsd/GoChat/handlers"
)

func main() {
	go handlers.HandleMessages()
	api.Api()
}
