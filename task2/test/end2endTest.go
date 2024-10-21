package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"task2/client"
	"task2/server"
	"time"
)

func main() {
	newServer := server.NewServer(":8081")

	if err := newServer.Start(); err != nil {
		fmt.Println(err)
		return
	}

	newClient := client.NewClient("http://localhost:8081")

	_, err := newClient.GetVersion()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	decodedString, err := newClient.PostDecode("aGVsbG8gd29ybGQ=") // hello world
	if err != nil {
		fmt.Printf("PostDecode request failed: %s\n\n", err.Error())
		return
	}
	fmt.Println("PostDecode request. decodedString :", decodedString)

	status, code, err := newClient.GetHardOp()
	if err != nil {
		fmt.Printf("GetHardOp request failed: %s\n\n", err.Error())
		return
	}
	fmt.Println("GetHardOp request. status :", status, "; code :", code)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	log.Println("Завершение работы сервера...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := newServer.Shutdown(ctx); err != nil {
		log.Fatalf("Ошибка при завершении работы сервера: %s", err)
	}

	log.Println("Сервер успешно завершил работу")
}
