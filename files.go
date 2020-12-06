package main

import "fmt"

func InitFileRPC(body []byte) []byte {
	fmt.Print("receiived RPC request")
	return []byte{1}
}
