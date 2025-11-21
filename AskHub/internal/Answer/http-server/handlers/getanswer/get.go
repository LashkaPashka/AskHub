package getanswer

import (
	"net/http"
	"strconv"

	"github.com/LashkaPashka/AskHub/internal/Answer/http-server/handlers/getanswer/payload"
	"github.com/LashkaPashka/AskHub/internal/Answer/model"
	"github.com/LashkaPashka/AskHub/pkg/res"
	"github.com/gorilla/mux"
)

type Service interface {
	GetByID(ID uint) (*model.Answer, error)
}

func New(service Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		strID := mux.Vars(r)["id"]
		intID, err := strconv.Atoi(strID)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		answer, err := service.GetByID(uint(intID))
		if err != nil {
			http.Error(w, "failed to get answer", http.StatusInternalServerError)
			return
		}

		res.Encode(w, payload.Response{
			Answer: answer,
		})
	}
}
