package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Oleska1601/WBOptimizeServer/config"
	"github.com/Oleska1601/WBOptimizeServer/internal/controller/api"
	v1 "github.com/Oleska1601/WBOptimizeServer/internal/controller/api/v1"
	v2 "github.com/Oleska1601/WBOptimizeServer/internal/controller/api/v2"
	v1service "github.com/Oleska1601/WBOptimizeServer/internal/service/v1"
	v2service "github.com/Oleska1601/WBOptimizeServer/internal/service/v2"
	"github.com/Oleska1601/WBOptimizeServer/pkg/logger"
)

func Run(cfg *config.Config) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger, err := logger.New(&cfg.Logger)
	if err != nil {
		logger.Fatal().
			Err(err).
			Str("path", "Run logger.New").
			Msg("failed to init logger")
	}

	service := v1service.New()
	apiV1 := v1.New(service, logger)
	router := api.Register(&cfg.Gin, apiV1)
	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.PortV1)
	server := &http.Server{Addr: addr, Handler: router}
	go func() {
		logger.Info().Str("path", "Run").Str("addr", addr).Msg("start server")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal().
				Err(err).
				Str("path", "Run server.ListenAndServe").
				Msg("failed to process server")
		}
	}()

	service2 := v2service.New()
	apiV2 := v2.New(service2, logger)
	router2 := api.Register(&cfg.Gin, apiV2)
	addr2 := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.PortV2)
	server2 := &http.Server{Addr: addr2, Handler: router2}
	go func() {
		logger.Info().Str("path", "Run").Str("addr", addr2).Msg("start server2")
		if err := server2.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal().
				Err(err).
				Str("path", "Run server2.ListenAndServe").
				Msg("failed to process server 2")
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	shutdownCtx, shutdownCancel := context.WithTimeout(ctx, cfg.Server.ShutdownTimeout)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Error().
			Err(err).
			Str("path", "Run server.Shutdown").
			Msg("failed to shutdown server")
	}

	logger.Info().Msg("shutdown system properly")
}
