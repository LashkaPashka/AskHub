package getquestion

import (
	"net/http"
	"strconv"

	"github.com/LashkaPashka/AskHub/internal/Question/http-server/handlers/getquestion/payload"
	"github.com/LashkaPashka/AskHub/internal/Question/model"
	"github.com/LashkaPashka/AskHub/pkg/res"
	"github.com/gorilla/mux"
)

type Service interface {
	GetByID(ID uint) (question *model.Question, err error)
}

func New(service Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		strID := mux.Vars(r)["id"]

		intID, err := strconv.Atoi(strID)
		if err != nil {
			http.Error(w, "invalid id format", http.StatusBadRequest)
			return
		}

		question, err := service.GetByID(uint(intID))
		if err != nil {
			http.Error(w, "Error: "+err.Error(), http.StatusConflict)
			return
		}

		res.Encode(w, payload.Response{
			Question: question,
		})

	}
}
