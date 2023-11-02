package port

import "net"

type Request interface {
	Read(net.Conn) error
	Path() string
}
