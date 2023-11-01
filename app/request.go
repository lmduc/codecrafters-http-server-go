package main

import (
	"net"
	"strings"
)

type Request struct {
	body []byte
}

func (r *Request) Read(conn net.Conn) error {
	r.body = make([]byte, 1024)
	_, err := conn.Read(r.body)
	return err
}

func (r *Request) Path() string {
	firstLine := strings.Split(string(r.body), "\n")[0]
	path := strings.Split(firstLine, " ")[1]
	return path
}
