package lambdaapi

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gorilla/mux"
)

type HttpApiApp struct {
	router *mux.Router
}

func NewHttpApiApp() *HttpApiApp {
	return &HttpApiApp{
		router: mux.NewRouter(),
	}
}

type Example http.Handler

func (app *HttpApiApp) handler(ctx context.Context, event *HttpEvent) (*HttpResponse, error) {
	slog.Info("handle event: %v", event)

	request, _ := event.AsHttpRequest()
	response := NewResponse()

	app.router.ServeHTTP(response, request)

	return response, nil
}

func (app *HttpApiApp) HandlerFunc(path string, handler http.HandlerFunc) {
	app.router.HandleFunc(path, handler)
}

func (app *HttpApiApp) Start() {
	lambda.Start(app.handler)
}
