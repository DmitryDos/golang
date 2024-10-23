package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"task2/server"
	"time"
)

func main() {
	newServer := server.NewServer(":8080")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("Сервер запущен на http://localhost:8080")
		if err := newServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Ошибка при запуске сервера: %s", err)
		}
	}()

	<-stop
	log.Println("Завершение работы сервера...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := newServer.Shutdown(ctx); err != nil {
		log.Fatalf("Ошибка при завершении работы сервера: %s", err)
	}

	log.Println("Сервер успешно завершил работу")
}
