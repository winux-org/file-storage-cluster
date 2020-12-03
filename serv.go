package main

import (
	"log"
)

// docker run -p 5672:5672 -d --hostname storage-rabbit --name storage-rabbit rabbitmq:3
// docker run -p 5430:5432 --name postgres89  -e POSTGRES_PASSWORD=password -d postgres

func main() {
	ConnectToDatabase()
	SetupRPC()

	log.Printf(" [*] File cluster started")
	WaitForEnd()
}
