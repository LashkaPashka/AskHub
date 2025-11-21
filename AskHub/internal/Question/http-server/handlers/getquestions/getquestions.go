package getquestions

import (
	"net/http"

	"github.com/LashkaPashka/AskHub/internal/Question/http-server/handlers/getquestions/payload"
	"github.com/LashkaPashka/AskHub/internal/Question/model"
	"github.com/LashkaPashka/AskHub/pkg/res"
)

type Service interface {
	 GetAll() (questions []model.Question, err error)
}

func New(service Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		questions, err := service.GetAll()
		if err != nil {
			http.Error(w, "Error: "+err.Error(), http.StatusConflict)
			return
		}

		res.Encode(w, payload.Response{
			Questions: questions,
		})
	}
}