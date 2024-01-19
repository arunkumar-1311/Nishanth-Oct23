package models

type Message struct {
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
}
