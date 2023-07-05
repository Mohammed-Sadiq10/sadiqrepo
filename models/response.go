package models

type Response struct {
	StatusCode   int         `json:"statusCode"`
	Status       string      `json:"status"`
	Message      string      `json:"message"`
	ResponseData interface{} `json:"result"`
	Error        interface{} `json:"error"`
	RequestURL   string      `json:"url"`
}
