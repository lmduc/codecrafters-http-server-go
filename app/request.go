package main

import "net"

type Request struct {
	body []byte
}

func (r *Request) Read(conn net.Conn) error {
	r.body = make([]byte, 1024)
	_, err := conn.Read(r.body)
	return err
}
