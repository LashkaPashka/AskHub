package createquestion

import (
	"log/slog"
	"net/http"

	"github.com/LashkaPashka/AskHub/internal/Question/http-server/handlers/createquestion/payload"
	"github.com/LashkaPashka/AskHub/internal/Question/lib/converter"
	"github.com/LashkaPashka/AskHub/internal/Question/model"
	"github.com/LashkaPashka/AskHub/pkg/req"
	"github.com/LashkaPashka/AskHub/pkg/res"
)

type Service interface {
	Create(question *model.Question) (success bool, err error)
}

func New(service Service, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[payload.Request](w, r, nil)
		if err != nil {
			http.Error(w, "failed to parse request body", http.StatusBadRequest)
			return
		}

		question := converter.ConvertQuestion(&body)

		success, err := service.Create((*model.Question)(&question))
		if err != nil {
			http.Error(w, "Error: "+err.Error(), http.StatusConflict)
			return
		}

		if !success {
			http.Error(w, "failed to create question", http.StatusConflict)
			return
		}

		res.Encode(w, payload.Response{
			Status: "Question was created successfully!",
		})
	}
}
