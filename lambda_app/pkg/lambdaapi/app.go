package lambdaapi

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"log/slog"
)

type HttpApiApp struct {
	router router
}

func NewHttpApiApp() *HttpApiApp {
	return &HttpApiApp{
		router: router{
			routes: make([]*route, 0),
		},
	}
}

func (app *HttpApiApp) handler(ctx context.Context, event *HttpEvent) (*HttpResponse, error) {
	slog.Info("handle event: %v", event)
	route := app.router.matchRoute(event)
	if route != nil {
		resp, err := route.handle(ctx, event)
		if err != nil {
			return nil, err
		}
		return resp.Into(), nil
	}
	return NewHttpResponse().WithStatusCode(404).WithBody("Not Found", false), nil
}

func (app *HttpApiApp) Register(path string, methods []string, handler HttpEventHandler) {
	newRoute := route{
		path:    path,
		methods: methods,
		handler: handler,
	}
	app.router.routes = append(app.router.routes, &newRoute)
}

func (app *HttpApiApp) Start() {
	lambda.Start(app.handler)
}
