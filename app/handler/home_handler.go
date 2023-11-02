package handler

import (
	"github.com/codecrafters-io/http-server-starter-go/app/port"
	"github.com/codecrafters-io/http-server-starter-go/app/response"
)

type HomeHandler struct {
	response *response.Response
}

func (p *HomeHandler) Handle(_ port.Request) (port.Response, error) {
	return p.response, nil
}

func NewHomeHandler() *HomeHandler {
	resp := response.NewResponse("")
	resp.StatusCode(200)

	return &HomeHandler{resp}
}
