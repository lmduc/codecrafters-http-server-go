package handler

import (
	"github.com/codecrafters-io/http-server-starter-go/app/port"
	"github.com/codecrafters-io/http-server-starter-go/app/response"
)

type Home struct {
	response *response.Response
}

func (p *Home) Handle(_ port.Request) (port.Response, error) {
	return p.response, nil
}

func NewHome() *Home {
	resp := response.NewResponse("")
	resp.StatusCode(200)

	return &Home{resp}
}
