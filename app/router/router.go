package router

import (
	"regexp"

	"github.com/codecrafters-io/http-server-starter-go/app/port"
)

type Router struct {
	pathToHandler   map[string]port.Handler
	pathRegex       map[string]*regexp.Regexp
	notFoundHandler port.Handler
}

func (r *Router) Handle(req port.Request) (port.Response, error) {
	for path, handler := range r.pathToHandler {
		if r.pathRegex[path].MatchString(req.Path()) {
			return handler.Handle(req)
		}
	}

	return r.notFoundHandler.Handle(req)
}

func (r *Router) Register(path string, handler port.Handler) *Router {
	r.pathToHandler[path] = handler
	r.pathRegex[path] = regexp.MustCompile(path)
	return r
}

func (r *Router) NotFoundHandler(handler port.Handler) *Router {
	r.notFoundHandler = handler
	return r
}

func NewRouter() *Router { return &Router{} }
