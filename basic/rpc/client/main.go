package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	var reply string
	err = client.Call("Server.Hello", "Go Rpc", &reply)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
