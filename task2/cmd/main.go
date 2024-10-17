package main

import (
	"log/slog"
	"os"
	"task2/server"
	"task2/test"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	go func() {
		newServer := server.NewServer(":8080")
		if err := newServer.Start(); err != nil {
			slog.Error("NewServer create failed: %s\n", "Error", err.Error())
			return
		}
	}()
	test.RunTests()
}
