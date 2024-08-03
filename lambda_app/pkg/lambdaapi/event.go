package lambdaapi

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/textproto"
	"strings"
)

const (
	HttpGet  = "GET"
	HttpPost = "POST"
	HttpAny  = "*"
)

type HttpEvent struct {
	Version        string            `json:"version"`
	RouteKey       string            `json:"routeKey"`
	RawPath        string            `json:"rawPath"`
	RawQueryString string            `json:"rawQueryString"`
	Cookies        []string          `json:"cookies"`
	Headers        map[string]string `json:"headers"`
	RequestContext struct {
		AccountId    string `json:"accountId"`
		ApiId        string `json:"apiId"`
		DomainName   string `json:"domainName"`
		DomainPrefix string `json:"domainPrefix"`
		Http         struct {
			Method    string `json:"method"`
			Path      string `json:"path"`
			Protocol  string `json:"protocol"`
			SourceIp  string `json:"sourceIp"`
			UserAgent string `json:"userAgent"`
		} `json:"http"`
		RequestId string `json:"requestId"`
		RouteKey  string `json:"routeKey"`
		Stage     string `json:"stage"`
		Time      string `json:"time"`
		TimeEpoch int64  `json:"timeEpoch"`
	} `json:"requestContext"`
	Body            string `json:"body"`
	IsBase64Encoded bool   `json:"isBase64Encoded"`
}

func (event *HttpEvent) AsHttpRequest() (*http.Request, error) {
	requestContext := event.RequestContext

	body, err := event.bodyDecoded()
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(requestContext.Http.Method, event.url(), body)

	for k, v := range event.Headers {
		request.Header.Add(k, v)
	}

	for _, line := range event.Cookies {
		line = textproto.TrimString(line)

		var part string
		for len(line) > 0 { // continue since we have rest
			part, line, _ = strings.Cut(line, ";")
			part = textproto.TrimString(part)
			if part == "" {
				continue
			}
			name, val, _ := strings.Cut(part, "=")
			name = textproto.TrimString(name)

			request.AddCookie(&http.Cookie{Name: name, Value: val})
		}

	}
	return request, err
}

func (event *HttpEvent) url() string {
	baseUrl := fmt.Sprintf("http://%s%s", event.RequestContext.DomainName, event.RawPath)
	if len(event.RawQueryString) > 0 {
		return baseUrl + "?" + event.RawQueryString
	} else {
		return baseUrl
	}
}

func (event *HttpEvent) bodyDecoded() (*bytes.Buffer, error) {
	if event.IsBase64Encoded {
		decoded, err := base64.StdEncoding.DecodeString(event.Body)
		if err != nil {
			return nil, err
		}
		return bytes.NewBuffer(decoded), nil

	}
	return bytes.NewBufferString(event.Body), nil
}

func (event *HttpEvent) BodyAsText() (string, error) {
	decoded, err := event.bodyDecoded()
	if err != nil {
		return "", err
	}
	return decoded.String(), nil
}
