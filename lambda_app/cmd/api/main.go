package main

import (
	"context"
	"fmt"
	"lambda_app/pkg/lambdaapi"
)

func main() {
	app := lambdaapi.NewHttpApiApp()
	app.Register("/hello", []string{lambdaapi.HttpAny}, func(_ context.Context, event *lambdaapi.HttpEvent) (lambdaapi.IntoHttpResponse, error) {
		textBody, _ := event.BodyAsText()
		respBody := fmt.Sprintf("here is waht I got: %s", textBody)
		return lambdaapi.NewHttpResponse().WithStatusCode(200).WithBody(respBody, false), nil
	})
	app.Start()
}
