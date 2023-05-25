package server

import (
	"fmt"
	"log"

	"sample-go-app/config"
)

func Init() {
	config := config.GetConfig()
	router := NewRouter()
	port := config.GetString("server.port")
	router.Run(fmt.Sprintf(":%s", port))
	log.Println("Server started on port", port)
}
