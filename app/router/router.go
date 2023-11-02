package router

import (
	"github.com/codecrafters-io/http-server-starter-go/app/port"
)

type Router struct {
	pathToHandler   map[port.PathMatcher]port.Handler
	notFoundHandler port.Handler
}

func (r *Router) Handle(req port.Request) (port.Response, error) {
	for pathMatcher, handler := range r.pathToHandler {
		if pathMatcher.Match(req.Path()) {
			return handler.Handle(req)
		}
	}

	return r.notFoundHandler.Handle(req)
}

func (r *Router) Register(pathMatcher port.PathMatcher, handler port.Handler) *Router {
	r.pathToHandler[pathMatcher] = handler
	return r
}

func (r *Router) NotFoundHandler(handler port.Handler) *Router {
	r.notFoundHandler = handler
	return r
}

func NewRouter() *Router {
	return &Router{
		pathToHandler: make(map[port.PathMatcher]port.Handler),
	}
}
