package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/Zumbr3/case-service/internal/infraestructure/server"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	err := server.Start(ctx)
	if err != nil {
		slog.Error("Error when starting server", "error", err)
		os.Exit(1)
	}
}
