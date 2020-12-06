package main

import "fmt"
import "context"
import "github.com/georgysavva/scany/pgxscan"

func InitFileRPC(body []byte) []byte {
	fmt.Print("receiived RPC request")

	var files []*File
	var db = DBpool()

	
	ctx := context.Background()
	rows, _ := db.Query(ctx, "seleft * from files")

	fmt.Println("Start querring")

	if err := pgxscan.ScanAll(&files, rows); err != nil {
		// Handle rows processing error.
		fmt.Println("error")
	}

	fmt.Println(len(files))

	return []byte{1}
}

func DeleteFileRPC(body []byte) []byte {
	fmt.Print("receiived RPC request")
	return []byte{1}
}