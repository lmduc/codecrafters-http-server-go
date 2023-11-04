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
	fmt.Println("status line: ", string(r.statusLine))

	var headerLines []byte
	for i := range data {
		if i > 0 && data[i-1] == '\n' && data[i] == '\n' {
			headerLines = data[:i-1]
			data = data[i+1:]
			break
		}
	}
	r.body = data

	r.headers = make(map[string]string)

	for _, line := range strings.Split(string(headerLines), "\n") {
		if headerRegexp.MatchString(line) {
			matches := headerRegexp.FindStringSubmatch(line)
			r.headers[matches[1]] = matches[2]
		}
	}

	if len(strings.Split(string(data), "\n\n")) > 1 {
		r.body = []byte(strings.Split(string(data), "\n\n")[1])
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
