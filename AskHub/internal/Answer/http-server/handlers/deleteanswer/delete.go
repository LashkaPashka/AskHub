package deleteanswer

import (
	"net/http"
	"strconv"

	"github.com/LashkaPashka/AskHub/internal/Answer/http-server/handlers/deleteanswer/payload"
	"github.com/LashkaPashka/AskHub/pkg/res"
	"github.com/gorilla/mux"
)

type Service interface {
	Delete(ID uint) (success bool, err error)
}

func New(service Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		strID := mux.Vars(r)["id"]
		intID, err := strconv.Atoi(strID)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		success, err := service.Delete(uint(intID))
		if err != nil {
			http.Error(w, "failed to create answer", http.StatusInternalServerError)
			return
		}
		if !success {
			http.Error(w, "cannot create answer", http.StatusBadRequest)
			return
		}

		res.Encode(w, payload.Response{
			Status: "Answer was deleted successfully",
		})
	}
}
