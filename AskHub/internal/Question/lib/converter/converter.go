package converter

import (
	"github.com/LashkaPashka/AskHub/internal/Question/model"
	"github.com/LashkaPashka/AskHub/internal/Question/http-server/handlers/createquestion/payload"
)

func ConvertQuestion(payload *payload.Request) model.Question {
	return model.Question{
		Text: payload.Text,
	}
}