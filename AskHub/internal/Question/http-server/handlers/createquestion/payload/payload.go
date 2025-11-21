package payload

type Request struct {
	Text string `json:"text" validate:"required"`
}

type Response struct {
	Status string `json:"status"`
}