package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/Zumbr3/case-service/internal/adapter/router"
)

func Start(ctx context.Context) error {
	server := router.Router()

	serverErr := make(chan error, 1)
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			serverErr <- err
		} else {
			serverErr <- nil
		}
	}()

	select {
	case <-ctx.Done():
		slog.Debug("Graceful Shutdown requested")
	case err := <-serverErr:
		if err != nil {
			return err
		}
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		slog.Debug("Server Shutdown Failed", "error", err)
	}

	slog.Debug("Graceful shutdown complete")
	return nil
}
