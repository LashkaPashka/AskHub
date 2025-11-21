package deletequestion

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/LashkaPashka/AskHub/internal/Question/http-server/handlers/createquestion/payload"
	"github.com/LashkaPashka/AskHub/pkg/res"
	"github.com/gorilla/mux"
)

type Service interface {
	Delete(ID uint) (success bool, err error)
}

func New(service Service, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		strID := mux.Vars(r)["id"]

		intID, err := strconv.Atoi(strID)
		if err != nil {
			http.Error(w, "invalid id format", http.StatusBadRequest)
			return
		}

		success, err := service.Delete(uint(intID))
		if err != nil {
			http.Error(w, "Error: " + err.Error(), http.StatusConflict)
			return
		}

		if !success {
			http.Error(w, "failed to delete question", http.StatusConflict)
			return
		}

		res.Encode(w, payload.Response{
			Status: "Question was deleted successfully!",
		})
	}
}
