package handler

import (
	"github.com/codecrafters-io/http-server-starter-go/app/port"
	"github.com/codecrafters-io/http-server-starter-go/app/response"
)

type UserAgent struct{}

func (u *UserAgent) Handle(r port.Request) (port.Response, error) {
	response := response.NewTextPlainResponse()
	response.StatusCode(200).Body([]byte(r.Header("User-Agent")))
	return response, nil
}

func NewUserAgent() *UserAgent { return &UserAgent{} }
