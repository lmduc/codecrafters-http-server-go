package main

import (
	"fmt"
	"net"
)

type Response struct {
	statusCode    int
	contentType   string
	contentLength int
	body          []byte
}

func (r *Response) Write(conn net.Conn) error {
	_, err := conn.Write(
		[]byte(
			fmt.Sprintf(
				"HTTP/1.1 %d OK\r\nContent-Type: %s\r\nContent-Length: %d\r\n\r\n%s",
				r.statusCode,
				r.contentType,
				r.contentLength,
				r.body,
			),
		),
	)

	return err
}

func (r *Response) StatusCode(statusCode int) *Response {
	r.statusCode = statusCode
	return r
}

func (r *Response) ContentType(contentType string) *Response {
	r.contentType = contentType
	return r
}

func (r *Response) Body(body []byte) *Response {
	r.body = body
	r.contentLength = len(body)
	return r
}
