package request

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"
)

var (
	headerRegexp = regexp.MustCompile(`([^:]+):\s+([^\s]*)`)
)

type HTTPRequest struct {
	statusLine string
	headers    map[string]string
	body       []byte
}

func (r *HTTPRequest) parseRequest(data []byte) error {
	parts := strings.SplitN(string(data), "\r\n\r\n", 2)
	if len(parts) < 2 {
		return fmt.Errorf("malformed request")
	}

	heads := strings.SplitN(parts[0], "\r\n", 2)
	r.statusLine = heads[0]

	headerLines := ""
	if len(heads) == 2 {
		headerLines = heads[1]
	}

	r.headers = make(map[string]string)

	for _, line := range strings.Split(headerLines, "\n") {
		if headerRegexp.MatchString(line) {
			matches := headerRegexp.FindStringSubmatch(line)
			r.headers[matches[1]] = matches[2]
		}
	}

	r.body = []byte(parts[1])
	if r.headers["Content-Length"] != "" {
		bodyLength, err := strconv.Atoi(r.headers["Content-Length"])
		if err != nil {
			return err
		}

		r.body = r.body[:bodyLength]
	}

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
	path := strings.Split(r.statusLine, " ")[1]
	return path
}

func (r *HTTPRequest) Header(key string) string {
	return r.headers[key]
}

func (r *HTTPRequest) Body() []byte {
	return r.body
}

func (r *HTTPRequest) Verb() string {
	return strings.Split(r.statusLine, " ")[0]
}

func NewHTTPRequest() *HTTPRequest { return &HTTPRequest{} }
