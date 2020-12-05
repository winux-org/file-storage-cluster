package main

import "net/http"

// docker run -p 5672:5672 -d --hostname storage-rabbit --name storage-rabbit rabbitmq:3
// docker run -p 5432:5432 --name postgres89  -e POSTGRES_PASSWORD=password -d postgres
// go run serv.go database.go models.go rpc.go api.go -- r

// docker exec -it b7a9f5eb6b85 sh
// docker exec -it pg-db sh
// docker attach --sig-proxy=false pg-db

func main() {
	ConnectToDatabase()
	SetupRPC()
	SetupHTTPRouting()
	http.ListenAndServe(":9000", nil)
}
