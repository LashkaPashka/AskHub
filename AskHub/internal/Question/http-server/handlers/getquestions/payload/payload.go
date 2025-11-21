package payload

import "github.com/LashkaPashka/AskHub/internal/Question/model"

type Response struct {
	Questions []model.Question
}