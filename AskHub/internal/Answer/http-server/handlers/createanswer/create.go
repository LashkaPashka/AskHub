package createanswer

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/LashkaPashka/AskHub/internal/Answer/http-server/handlers/createanswer/payload"
	"github.com/LashkaPashka/AskHub/internal/Answer/lib/converter"
	"github.com/LashkaPashka/AskHub/internal/Answer/model"
	"github.com/LashkaPashka/AskHub/pkg/req"
	"github.com/LashkaPashka/AskHub/pkg/res"
	"github.com/gorilla/mux"
)

type Service interface {
	Create(answer *model.Answer) (bool, error)
}

func New(service Service, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		strID := mux.Vars(r)["id"]
		intID, err := strconv.Atoi(strID)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		logger = logger.With(slog.Int("question_id", intID))

		body, err := req.HandleBody[payload.Request](w, r, logger)
		if err != nil {
			http.Error(w, "failed to parse request body", http.StatusBadRequest)
			return
		}

		answer := converter.ConvertAnswer(intID, body)
		logger = logger.With(slog.Any("answer", answer))

		success, err := service.Create(&answer)
		if err != nil {
			http.Error(w, "failed to create answer", http.StatusInternalServerError)
			return
		}
		if !success {
			http.Error(w, "cannot create answer", http.StatusBadRequest)
			return
		}

		res.Encode(w, &payload.Response{
			Status: "Answer was successfully created",
		})

	}
}
