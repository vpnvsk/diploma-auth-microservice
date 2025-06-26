package main

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/vpnvsk/amunet_auth_microservices"
	handler_ "github.com/vpnvsk/amunet_auth_microservices/pkg/handler"
	"github.com/vpnvsk/amunet_auth_microservices/pkg/repository"
	service2 "github.com/vpnvsk/amunet_auth_microservices/pkg/service"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	settings := amunet_auth_microservices.NewConfig()
	log := setUpLogger(settings.Env)
	db, err := repository.NewPostgresDb(settings)
	if err != nil {
		panic(err)
	}
	repo := repository.NewRepository(log, db)
	service := service2.NewService(log, repo, settings)
	handler := handler_.NewHandler(log, service, settings)
	srv := new(amunet_auth_microservices.Server)
	go func() {
		if err := srv.Run("9000", handler.InitRoutes()); err != nil {
			errorMessage := fmt.Sprintf("error while running server %s", err.Error())
			panic(errorMessage)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	if err := srv.ShutDown(context.Background()); err != nil {
		log.Error("error while shutting down: %s", err.Error())
	}
}

func setUpLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case "prod":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}
	return log
}
