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
	resp := response.NewResponse("")
	resp.StatusCode(404)

	return &NotFound{resp}
}
