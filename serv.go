package main

import "net/http"

// docker run -p 5672:5672 -d --hostname storage-rabbit --name storage-rabbit rabbitmq:3
// docker run -p 5430:5432 --name postgres89  -e POSTGRES_PASSWORD=password -d postgres
// go run serv.go database.go models.go rpc.go api.go -- r
func main() {
	ConnectToDatabase()
	SetupRPC()
	SetupHTTPRouting()
	http.ListenAndServe(":9000", nil)
}
