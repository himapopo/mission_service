package dto

type Response struct {
	Result bool   `json:"result"`
	Error  string `json:"error"`
}
