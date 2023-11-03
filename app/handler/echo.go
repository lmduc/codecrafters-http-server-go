package handler

import (
	"github.com/codecrafters-io/http-server-starter-go/app/port"
	"github.com/codecrafters-io/http-server-starter-go/app/response"
	"github.com/codecrafters-io/http-server-starter-go/app/router"
)

type Echo struct {
	matcher *router.RegexPathMatcher
}

func (e *Echo) Handle(r port.Request) (port.Response, error) {
	matches := e.matcher.FindMatches(r.Path())
	response := response.NewTextPlainResponse()
	response.StatusCode(200).Body([]byte(matches[1]))
	return response, nil
}

func NewEcho(matcher *router.RegexPathMatcher) *Echo {
	return &Echo{matcher}
}
