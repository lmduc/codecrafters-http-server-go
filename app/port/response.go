package port

import "net"

type Response interface {
	Write(net.Conn) error
}
