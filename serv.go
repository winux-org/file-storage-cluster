package main

import (
	"log"
)

func main() {
	//ConnectToDatabase()
	SetupRPC()

	log.Printf(" [*] File cluster started")
	WaitForEnd()
}
