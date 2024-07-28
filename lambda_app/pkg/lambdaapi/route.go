package lambdaapi

import "context"

type HttpEventHandler func(context.Context, *HttpEvent) (IntoHttpResponse, error)

type route struct {
	path    string
	methods []string
	handler HttpEventHandler
}

func (route *route) handle(ctx context.Context, event *HttpEvent) (IntoHttpResponse, error) {
	return route.handler(ctx, event)
}

func (route *route) match(event *HttpEvent) bool {
	if event.RequestContext.Http.Path == route.path {
		for _, m := range route.methods {
			if m == HttpAny || event.RequestContext.Http.Method == m {
				return true
			}
		}
	}
	return false
}

type router struct {
	routes []*route
}

func (router router) matchRoute(event *HttpEvent) *route {
	for _, r := range router.routes {
		if r.match(event) {
			return r
		}
	}
	return nil
}
