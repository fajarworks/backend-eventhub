package dto

type Response struct {
	Success bool   `json:"status"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}
