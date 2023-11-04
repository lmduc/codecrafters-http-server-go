package request

import (
	"fmt"
	"net"
	"regexp"
	"strings"
)

var (
	headerRegexp = regexp.MustCompile(`([^:]+):\s+([^\s]*)`)
)

type HTTPRequest struct {
	statusLine []byte
	headers    map[string]string
	body       []byte
}

func (r *HTTPRequest) parseRequest(data []byte) error {
	for i := range data {
		if data[i] == '\n' {
			r.statusLine = data[:i]
			data = data[i+1:]
			break
		}
	}
	fmt.Println("data", string(data))

	var headerLines []byte
	index := strings.Index(string(data), "\r\n\r\n")
	if index != -1 {
		headerLines, r.body = data[:index], data[index+2:]
	} else {
		headerLines, r.body = data, nil
	}

	r.headers = make(map[string]string)

	for _, line := range strings.Split(string(headerLines), "\n") {
		if headerRegexp.MatchString(line) {
			matches := headerRegexp.FindStringSubmatch(line)
			r.headers[matches[1]] = matches[2]
		}
	}

	fmt.Println("body :", string(r.body))

	return nil
}

func (r *HTTPRequest) Read(conn net.Conn) error {
	data := make([]byte, 1024)
	if _, err := conn.Read(data); err != nil {
		return err
	}

	if err := r.parseRequest(data); err != nil {
		return err
	}

	return nil
}

func (r *HTTPRequest) Path() string {
	path := strings.Split(string(r.statusLine), " ")[1]
	return path
}

func (r *HTTPRequest) Header(key string) string {
	return r.headers[key]
}

func (r *HTTPRequest) Body() []byte {
	return r.body
}

func (r *HTTPRequest) Verb() string {
	return strings.Split(string(r.statusLine), " ")[0]
}

func NewHTTPRequest() *HTTPRequest { return &HTTPRequest{} }
