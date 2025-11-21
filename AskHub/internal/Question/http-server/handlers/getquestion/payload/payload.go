package payload

import "github.com/LashkaPashka/AskHub/internal/Question/model"

type Response struct {
	Question *model.Question `json:"question"`
}
