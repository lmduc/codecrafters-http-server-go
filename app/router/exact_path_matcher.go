package router

import "github.com/codecrafters-io/http-server-starter-go/app/port"

type ExactPathMatcher struct {
	path string
}

func (e *ExactPathMatcher) Match(req port.Request) bool {
	return e.path == req.Path()
}

func NewExactPathMatcher(path string) *ExactPathMatcher {
	return &ExactPathMatcher{path}
}
