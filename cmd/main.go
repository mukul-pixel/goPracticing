package main

import (
	"log"

	"example.com/go-practicing/cmd/api"

)

func main() {
	server := api.NewAPIServer(":8080",nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
