package payload

type Request struct {
	UserID string `json:"user_id" validate:"required"`
	Text   string `json:"text" validate:"required"`
}

type Response struct {
	Status string `json:"status"`
}
