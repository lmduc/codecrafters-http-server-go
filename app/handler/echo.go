package handler

import (
	"github.com/codecrafters-io/http-server-starter-go/app/port"
	"github.com/codecrafters-io/http-server-starter-go/app/response"
	"github.com/codecrafters-io/http-server-starter-go/app/router"
)

type Echo struct {
	matcher *router.RegexMatcher
}

func (e *Echo) Handle(r port.Request) (port.Response, error) {
	match := e.matcher.FindMatch(r.Path())
	response := response.NewTextPlainResponse()
	response.StatusCode(200).Body([]byte(match))
	return response, nil
}

func NewEcho(matcher *router.RegexMatcher) *Echo {
	return &Echo{matcher}
}
