package main

import (
	"github.com/kirill0909/neurohacking-api"
	"log"
)

func main() {
	srv := new(server.Server)

	if err := srv.Run(":8000"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
