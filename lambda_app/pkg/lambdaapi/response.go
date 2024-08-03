package lambdaapi

import "net/http"

type HttpResponse struct {
	StatusCode        int               `json:"statusCode"`
	Headers           map[string]string `json:"headers"`
	MultiValueHeaders http.Header       `json:"multiValueHeaders"`
	IsBase64Encoded   bool              `json:"isBase64Encoded"`
	Body              string            `json:"body"`
}

func NewResponse() *HttpResponse {
	return &HttpResponse{
		StatusCode:        http.StatusOK,
		Headers:           make(map[string]string), // We ignore single value header for now.
		MultiValueHeaders: http.Header{},
		IsBase64Encoded:   false,
		Body:              "",
	}
}

func (resp *HttpResponse) Header() http.Header {
	return resp.MultiValueHeaders
}

func (resp *HttpResponse) Write(data []byte) (int, error) {
	resp.Body += string(data)
	return len(data), nil
}

func (resp *HttpResponse) WriteHeader(statusCode int) {
	resp.StatusCode = statusCode
}
