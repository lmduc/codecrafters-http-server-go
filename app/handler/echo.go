package handler

import "github.com/codecrafters-io/http-server-starter-go/app/port"

type Echo struct{}

func (e *Echo) Handle(r port.Request) (port.Response, error) {
	return nil, nil
}

func NewEcho() *Echo { return &Echo{} }
