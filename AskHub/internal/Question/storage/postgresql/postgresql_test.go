package postgresql

import (
	"log/slog"
	"os"
	"testing"

	"github.com/LashkaPashka/AskHub/internal/Question/model"
)

var st *Storage

var connStr = "host=localhost user=postgres password=root dbname=askhub port=5432 sslmode=disable"

var testData = &model.Question{
	Text: "What's Go?",
}

func TestMain(m *testing.M) {
	logger := setupLogger("test")

	st = New(connStr, logger)

	m.Run()
}

func TestCreate(t *testing.T) {
	success, err := st.Create(testData)
	if err != nil || !success {
		t.Fatalf("Failed crceate, err: %v", err)
	}
}

func TestGetByID(t *testing.T) {
	const ID = 1

	question, err := st.GetByID(ID)
	if err != nil {
		t.Fatalf("Error GetByID, err: %v", err)
	}

	t.Log(question)
}

func TestGetAll(t *testing.T) {
	questions, err := st.GetAll()
	if err != nil {
		t.Fatalf("Error GetAll, err: %v", err)
	}

	t.Log(questions)
}

func TestDelete(t *testing.T) {
	const ID = 1
	
	success, err := st.Delete(ID)
	if err != nil || !success {
		t.Fatalf("Failed create, err: %v", err)
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