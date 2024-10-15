package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"task2/client"
	"task2/server"
	"time"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	srv := server.NewServer(":8080")
	if err := srv.Start(); err != nil {
		fmt.Println(err)
		return
	}

	newClient := client.NewClient("http://localhost:8080")

	_, err := newClient.GetVersion()
	if err != nil {
		slog.Error(err.Error())
		return
	}

	decodedString, err := newClient.PostDecode("aGVsbG8gd29ybGQ=") // hello world
	if err != nil {
		slog.Warn("PostDecode request failed: %s\n", "Error", err.Error())
		return
	}
	slog.Debug("PostDecode request", "text", decodedString)

	status, code, err := newClient.GetHardOp()
	if err != nil {
		slog.Warn("GetHardOp request failed: %s\n", "Error", err.Error())
		return
	}

	slog.Debug("Hard-op request", "status", status, "code", code)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Warn("Server shutdown failed: %s\n", "Error", err.Error())
	}
	slog.Debug("Server shutdown successfully")
}
