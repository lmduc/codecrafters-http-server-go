package port

import "net"

type Request interface {
	Read(net.Conn) error
	Path() string
	Header(string) string
	Body() []byte
	Verb() string
}
