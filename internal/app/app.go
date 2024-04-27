package app

import (
	"classconnect-api/internal/config"
	"classconnect-api/internal/domain"
	httpLayer "classconnect-api/internal/handler/http"
	"classconnect-api/internal/repository/postgres"
	"classconnect-api/pkg/client/postgresql"
	"classconnect-api/pkg/logging"
	"context"
	"errors"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	logger     *slog.Logger
	httpServer *http.Server
}

func New() *App {
	//Init the config
	cfg := config.MustLoad()

	//Init the slog
	logger := logging.InitSlog(cfg.Environment)

	//Init postgres client
	client := postgresql.NewClient(context.Background(), "postgresql://postgres:root@localhost:5432/postgres")

	//Init repositories
	repositories := postgres.NewRepositories(client, logger)

	//Init manager of serviсes
	services := domain.NewServices(repositories)

	//Init the router
	router := httpLayer.NewRouter(services, cfg, logger)

	return &App{
		logger: logger,
		httpServer: &http.Server{
			Addr:    net.JoinHostPort(cfg.HTTPServer.Address, cfg.HTTPServer.Port),
			Handler: router,
		},
	}
}

func (a *App) Run(ctx context.Context) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	//TODO: print logging about start the server
	go func() {
		err := a.httpServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.logger.Error("http server closed unexpectedly", slog.Any("error", err))
			os.Exit(1)
		}
	}()

	<-quit

	a.Stop(ctx)
}

func (a *App) Stop(ctx context.Context) {
	a.logger.Info("shutting down")
	if err := a.httpServer.Shutdown(ctx); err != nil {
		a.logger.Error("an error occurred during server shutdown", slog.Any("error", err))
		os.Exit(1)
	}
}
