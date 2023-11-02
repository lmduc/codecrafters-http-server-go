package handler

import (
	"github.com/codecrafters-io/http-server-starter-go/app/port"
	"github.com/codecrafters-io/http-server-starter-go/app/response"
)

type NotFoundHandler struct {
	response *response.Response
}

func (p *NotFoundHandler) Handle(_ port.Request) (port.Response, error) {
	return p.response, nil
}

func NewNotFoundHandler() *NotFoundHandler {
	resp := response.NewResponse("")
	resp.StatusCode(404)

	return &NotFoundHandler{resp}
}
