package dto

type Response struct {
	Status  bool   `json:"status"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}
