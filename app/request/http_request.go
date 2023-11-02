package request

import (
	"fmt"
	"net"
	"strings"
)

type HTTPRequest struct {
	body       []byte
	statusLine string
	headers    map[string]string
}

func (r *HTTPRequest) readStatusLine(data []byte) error {
	r.statusLine = strings.Split(string(data), "\n")[0]
	fmt.Println("status line: ", r.statusLine)
	return nil
}

func (r *HTTPRequest) readHeaders(data []byte) error {
	fmt.Println("inside read header")
	r.headers = make(map[string]string)
	fmt.Println("number: ", len(strings.Split(string(data), "\n\n")))
	headerLines := strings.Split(string(data), "\n\n")[1]
	fmt.Println("header lines: ", headerLines)
	for _, line := range strings.Split(headerLines, "\n") {
		fmt.Println("line: ", line)
		kvs := strings.Split(line, ": ")
		r.headers[kvs[0]] = kvs[1]
	}
	fmt.Println("headers: ", r.headers)
	return nil
}

func (r *HTTPRequest) Read(conn net.Conn) error {
	data := make([]byte, 1024)
	if _, err := conn.Read(data); err != nil {
		return err
	}

	if err := r.readStatusLine(data); err != nil {
		return err
	}

	if err := r.readHeaders(data); err != nil {
		return err
	}

	return nil
}

func (r *HTTPRequest) Path() string {
	path := strings.Split(r.statusLine, " ")[1]
	return path
}

func (r *HTTPRequest) Header(key string) string {
	return r.headers[key]
}

func NewHTTPRequest() *HTTPRequest { return &HTTPRequest{} }
