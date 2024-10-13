package main

import (
	"fmt"
	"task2/client"
	"task2/server"
)

func main() {
	newServer := server.NewServer(":8080")
	if err := newServer.Start(); err != nil {
		fmt.Println(err)
		return
	}

	newClient := client.NewClient("http://localhost:8080")
	body, err := newClient.GetVersion()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
