package models

type Response struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
	Data    []Book `json:"data"`
}
