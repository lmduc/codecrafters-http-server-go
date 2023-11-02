package request

import (
	"net"
	"strings"
)

type HTTPRequest struct {
	body []byte
}

func (r *HTTPRequest) Read(conn net.Conn) error {
	r.body = make([]byte, 1024)
	_, err := conn.Read(r.body)
	return err
}

func (r *HTTPRequest) Path() string {
	firstLine := strings.Split(string(r.body), "\n")[0]
	path := strings.Split(firstLine, " ")[1]
	return path
}

func NewHTTPRequest() *HTTPRequest { return &HTTPRequest{} }
