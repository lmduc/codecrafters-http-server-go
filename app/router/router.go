package router

import (
	"github.com/codecrafters-io/http-server-starter-go/app/port"
)

type Router struct {
	routeToHandler  map[port.RouteMatcher]port.Handler
	notFoundHandler port.Handler
}

func (r *Router) Handle(req port.Request) (port.Response, error) {
	for routeMatcher, handler := range r.routeToHandler {
		if routeMatcher.Match(req) {
			return handler.Handle(req)
		}
	}

	return r.notFoundHandler.Handle(req)
}

func (r *Router) Register(routeMatcher port.RouteMatcher, handler port.Handler) *Router {
	r.routeToHandler[routeMatcher] = handler
	return r
}

func (r *Router) NotFoundHandler(handler port.Handler) *Router {
	r.notFoundHandler = handler
	return r
}

func NewRouter() *Router {
	return &Router{
		routeToHandler: make(map[port.RouteMatcher]port.Handler),
	}
}
