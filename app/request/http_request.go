package request

import (
	"net"
	"regexp"
	"strings"
)

var (
	headerRegexp = regexp.MustCompile(`([^:]+):\s+([^\s]*)`)
)

type HTTPRequest struct {
	body       []byte
	statusLine string
	headers    map[string]string
}

func (r *HTTPRequest) readStatusLine(data []byte) error {
	r.statusLine = strings.Split(string(data), "\n")[0]
	return nil
}

func (r *HTTPRequest) readHeaderLinesAndBody(data []byte) ([]byte, []byte, error) {
	for i := range data {
		if i != 0 && data[i-1] == '\n' && data[i] == '\n' {
			return data[:i-1], data[i+1:], nil
		}
	}
	return data, nil, nil
}

func (r *HTTPRequest) readHeaders(data []byte) error {
	r.headers = make(map[string]string)

	headerLines, body, err := r.readHeaderLinesAndBody(data)
	if err != nil {
		return nil
	}

	r.body = body

	for _, line := range strings.Split(string(headerLines), "\n")[1:] {
		if headerRegexp.MatchString(line) {
			matches := headerRegexp.FindStringSubmatch(line)
			if len(matches) == 3 {
				r.headers[matches[1]] = matches[2]
			}
		}
	}
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

func (r *HTTPRequest) Body() []byte {
	return r.body
}

func NewHTTPRequest() *HTTPRequest { return &HTTPRequest{} }
