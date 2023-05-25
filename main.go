package main

import (
	"sample-go-app/config"
	"sample-go-app/server"
)

func main() {
	config.Init("env")
	server.Init()
}
