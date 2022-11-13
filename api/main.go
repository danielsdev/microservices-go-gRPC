package main

import (
	"log"

	"github.com/danielsdev/microservices-go-gRPC/api/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Panic("Erro ao carregar env")
	}

	routes.HandleRequests()
}
