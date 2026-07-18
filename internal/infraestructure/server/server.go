package server

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/Zumbr3/case-service/internal/adapter/handler"
	"github.com/Zumbr3/case-service/internal/adapter/router"
	"github.com/Zumbr3/case-service/internal/infraestructure/container"
	"github.com/Zumbr3/case-service/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Start(ctx context.Context) error {
	db, err := InitPostgres(ctx)
	if err != nil {
		return err
	}
	defer db.Close()

	repos := container.New(db)
	services := service.NewServices(repos.CaseRepository)
	handlers := handler.NewHandlers(services)
	server := router.Router(handlers)

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

func InitPostgres(ctx context.Context) (*pgxpool.Pool, error) {
	poolConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	db, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(ctx); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
