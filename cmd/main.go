package main

import (
	"log"

	"github.com/SicParv1sMagna/NetworkingTransportLayer/internal/app"
)

func main() {
	log.Println("Transport Layer Started")

	application, err := app.New()
	if err != nil {
		log.Println(err)
	}

	application.StartServer()

	log.Println("Transport Layer Shutting Down")
}
