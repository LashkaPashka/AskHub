package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/LashkaPashka/AskHub/internal/Answer/http-server/handlers/createanswer"
	"github.com/LashkaPashka/AskHub/internal/Answer/http-server/handlers/deleteanswer"
	"github.com/LashkaPashka/AskHub/internal/Answer/http-server/handlers/getanswer"
	aSv "github.com/LashkaPashka/AskHub/internal/Answer/service"
	aSt "github.com/LashkaPashka/AskHub/internal/Answer/storage/postgresql"

	"github.com/LashkaPashka/AskHub/internal/Question/http-server/handlers/createquestion"
	"github.com/LashkaPashka/AskHub/internal/Question/http-server/handlers/deletequestion"
	"github.com/LashkaPashka/AskHub/internal/Question/http-server/handlers/getquestion"
	"github.com/LashkaPashka/AskHub/internal/Question/http-server/handlers/getquestions"
	qSv "github.com/LashkaPashka/AskHub/internal/Question/service"
	qSt "github.com/LashkaPashka/AskHub/internal/Question/storage/postgresql"
	"github.com/LashkaPashka/AskHub/pkg/config"
	"github.com/gorilla/mux"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	logger := setupLogger(cfg.Env)

	// Init question service
	qStorage := qSt.New(cfg.StoragePath, logger)
	qService := qSv.New(qStorage)

	// Init answer service
	aStorage := aSt.New(cfg.StoragePath, logger)
	aService := aSv.New(aStorage, logger)

	router := mux.NewRouter()

	// API Question
	router.Handle("/questions", createquestion.New(qService, logger)).Methods(http.MethodPost)
	router.Handle("/questions/{id}", getquestion.New(qService)).Methods(http.MethodGet)
	router.Handle("/questions", getquestions.New(qService)).Methods(http.MethodGet)
	router.Handle("/questions/{id}", deletequestion.New(qService, logger)).Methods(http.MethodDelete)

	// API Answer
	router.Handle("/questions/{id}/answers", createanswer.New(aService, logger)).Methods(http.MethodPost)
	router.Handle("/answers/{id}", getanswer.New(aService)).Methods(http.MethodGet)
	router.Handle("/answers/{id}", deleteanswer.New(aService)).Methods(http.MethodDelete)

	logger.Info("starting server", slog.String("address", cfg.Address))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Error("failed to stop server")
		}
	}()

	logger.Info("server started")

	<-done
	logger.Info("stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("failed to stop server")
		return
	}

	logger.Info("server stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
