package handler

import (
	"github.com/codecrafters-io/http-server-starter-go/app/port"
	"github.com/codecrafters-io/http-server-starter-go/app/response"
)

type NotFound struct {
	response *response.Response
}

func (p *NotFound) Handle(_ port.Request) (port.Response, error) {
	return p.response, nil
}

func NewNotFound() *NotFound {
	return &NotFound{notFoundResponse()}
}

func notFoundResponse() *response.Response {
	return response.NewResponse("").StatusCode(404)
}
