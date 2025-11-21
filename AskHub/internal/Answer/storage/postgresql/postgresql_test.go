package storage

import (
	"log/slog"
	"os"
	"testing"

	ma "github.com/LashkaPashka/AskHub/internal/Answer/model"
	mq "github.com/LashkaPashka/AskHub/internal/Question/model"
)

var st *Storage

var connStr = "host=localhost user=postgres password=root dbname=askhub port=5432 sslmode=disable"

var testData = mq.Question{
	Text: "What's Go?",
}

var testAnswer = ma.Answer{
	QuestionID: 1,
	UserID: "550e8400-e29b-41d4-a716-446655440000",
	Text: "Go - programming language",
}

func TestMain(m *testing.M) {
	logger := setupLogger("test")
	
	st = New(connStr, logger)

	m.Run()
}

func TestCreate(t *testing.T) {
	success, err := st.Create(&testAnswer)

	if err != nil || !success {
		t.Fatalf("Error: %v", err)
	}
}

func TestGet(t *testing.T) {
	var answerID uint = 2
	answer, err := st.GetByID(answerID)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	
	t.Log(answer)
}

func TestDelete(t *testing.T) {
	var answerID uint = 2
	success, err := st.Delete(answerID)
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
