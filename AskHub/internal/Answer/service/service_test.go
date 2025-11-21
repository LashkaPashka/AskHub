package service

import (
	"log/slog"
	"os"
	"testing"

	"github.com/LashkaPashka/AskHub/internal/Answer/model"
	storage "github.com/LashkaPashka/AskHub/internal/Answer/storage/postgresql"
)

var sv *Service

var connStr = "host=localhost user=postgres password=root dbname=askhub port=5432 sslmode=disable"

var answer = model.Answer{
	QuestionID: 1,
	UserID:     "550e8400-e29b-41d4-a716-446655440000",
	Text:       "Hello world!",
}

func TestMain(m *testing.M) {
	logger := setupLogger("test")

	storage := storage.New(connStr, logger)

	sv = New(storage, logger)

	m.Run()
}

func TestCreate(t *testing.T) {
	success, err := sv.Create(&answer)
	if err != nil || !success {
		t.Fatalf("Error: %v", err)
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case "test":
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	}
	return log
}
