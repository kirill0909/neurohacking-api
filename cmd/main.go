package main

import (
	"github.com/kirill0909/neurohacking-api"
	"github.com/kirill0909/neurohacking-api/pkg/handler"
	"log"
)

func main() {
	srv := new(server.Server)
	handlers := new(handler.Handler)

	if err := srv.Run(":8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
