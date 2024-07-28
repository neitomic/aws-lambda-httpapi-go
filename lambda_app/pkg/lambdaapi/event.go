package lambdaapi

import (
	"encoding/base64"
	"encoding/json"
	"errors"
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

func (event *HttpEvent) bodyDecoded() ([]byte, error) {
	if event.IsBase64Encoded {
		return base64.StdEncoding.DecodeString(event.Body)
	}
	return []byte(event.Body), nil
}

func (event *HttpEvent) BodyAsText() (string, error) {
	decoded, err := event.bodyDecoded()
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

func (event *HttpEvent) BodyAsJson(target any, verifyHeader bool) error {
	if verifyHeader {
		value := ""
		if contentType, ok := event.Headers["ContentType"]; ok {
			value = strings.ToLower(contentType)
		}
		if value != "application/json" {
			return errors.New("not a json body")
		}
	}

	decoded, err := event.bodyDecoded()
	if err != nil {
		return err
	}
	return json.Unmarshal(decoded, target)
}
