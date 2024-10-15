package main

import (
	"log/slog"
	"os"
	"task2/client"
	"task2/server"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	newServer := server.NewServer(":8080")
	if err := newServer.Start(); err != nil {
		slog.Error("NewServer create failed: %s\n", "Error", err.Error())
		return
	}

	newClient := client.NewClient("http://localhost:8080")
	body, err := newClient.GetVersion()
	if err != nil {
		slog.Error("NewClient create failed: %s\n", "Error", err.Error())
		return
	}
	slog.Info(string(body))
}
