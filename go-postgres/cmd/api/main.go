package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/faizallmaullana/rekrutment-gbu-go/internal/config"
	httpdelivery "github.com/faizallmaullana/rekrutment-gbu-go/internal/delivery/http"
	"github.com/faizallmaullana/rekrutment-gbu-go/internal/infrastructure/postgres"
	"github.com/faizallmaullana/rekrutment-gbu-go/internal/repository"
	"github.com/faizallmaullana/rekrutment-gbu-go/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	settings := config.Load()

	database, err := postgres.NewDatabase(settings.Database.DSN())
	if err != nil {
		logger.Error("connect database", "error", err)
		os.Exit(1)
	}
	sqlDatabase, err := database.DB()
	if err != nil {
		logger.Error("get database connection", "error", err)
		os.Exit(1)
	}
	defer sqlDatabase.Close()

	repository := repository.NewProductRepository(database)
	service := usecase.NewProductService(repository)
	handler := httpdelivery.NewProductHandler(service)

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	router.GET("/healthz", func(context *gin.Context) {
		context.Status(http.StatusNoContent)
	})
	handler.RegisterRoutes(router)
	server := &http.Server{Addr: settings.HTTPAddr, Handler: router}

	go func() {
		logger.Info("http server started", "address", settings.HTTPAddr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("http server stopped", "error", err)
			os.Exit(1)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	shutdownContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(shutdownContext); err != nil {
		logger.Error("shutdown http server", "error", err)
	}
}
