package lambdaapi

type HttpResponse struct {
	StatusCode        int                 `json:"statusCode"`
	Headers           map[string]string   `json:"headers"`
	MultiValueHeaders map[string][]string `json:"multiValueHeaders"`
	IsBase64Encoded   bool                `json:"isBase64Encoded"`
	Body              string              `json:"body"`
}

type IntoHttpResponse interface {
	Into() *HttpResponse
}

func NewHttpResponse() *HttpResponse {
	return &HttpResponse{
		StatusCode:        0,
		Headers:           make(map[string]string),
		MultiValueHeaders: make(map[string][]string),
		IsBase64Encoded:   false,
		Body:              "",
	}
}

func (resp *HttpResponse) WithStatusCode(statusCode int) *HttpResponse {
	resp.StatusCode = statusCode
	return resp
}

func (resp *HttpResponse) WithHeader(key, value string) *HttpResponse {
	resp.Headers[key] = value
	return resp
}
func (resp *HttpResponse) WithBody(value string, base64 bool) *HttpResponse {
	resp.Body = value
	resp.IsBase64Encoded = base64
	return resp
}

func (resp *HttpResponse) Into() *HttpResponse {
	return resp
}
