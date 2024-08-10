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
	db, err := repository.NewPostgresDb(settings)
	if err != nil {
		panic(err.Error())
	}
	repo := repository.NewRepository(db)
	service := service2.NewService(repo)
	_ = repo
	handler := handler_.NewHandler(service)
	srv := new(amunet_auth_microservices.Server)
	go func() {
		if err := srv.Run("8080", handler.InitRoutes()); err != nil {
			errorMessage := fmt.Sprintf("error while running server %s", err.Error())
			panic(errorMessage)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	if err := srv.ShutDown(context.Background()); err != nil {
		slog.Error("error while shutting down: %s", err.Error())
	}
}
