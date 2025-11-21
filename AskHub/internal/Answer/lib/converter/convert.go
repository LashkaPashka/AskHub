package converter

import (
	"github.com/LashkaPashka/AskHub/internal/Answer/http-server/handlers/createanswer/payload"
	"github.com/LashkaPashka/AskHub/internal/Answer/model"
)

func ConvertAnswer(intID int, payload payload.Request) model.Answer {
	return model.Answer{
		QuestionID: uint(intID),
		UserID:     payload.UserID,
		Text:       payload.Text,
	}
}
