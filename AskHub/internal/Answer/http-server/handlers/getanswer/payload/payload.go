package payload

import "github.com/LashkaPashka/AskHub/internal/Answer/model"

type Response struct {
	Answer *model.Answer `json:"Answer"`
}
