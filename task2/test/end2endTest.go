package main

import (
	"context"
	"fmt"
	"task2/client"
	"task2/server"
	"time"
)

func main() {
	srv := server.NewServer(":8080")
	if err := srv.Start(); err != nil {
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

	decodedString, err := newClient.PostDecode("aGVsbG8gd29ybGQ=") // hello world
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(decodedString)

	status, code, err := newClient.GetHardOp()
	if err != nil {
		fmt.Println(err)
		return
	}
	if status {
		fmt.Printf("%t, %d\n", status, code)
		return
	}
	fmt.Printf("%t\n", status)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Server shutdown failed: %s\n", err)
	}
	fmt.Println("Server shutdown successfully")
}
