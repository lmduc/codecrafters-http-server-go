package main

import (
	"fmt"
	"net"
	"os"

	"github.com/codecrafters-io/http-server-starter-go/app/handler"
	"github.com/codecrafters-io/http-server-starter-go/app/request"
	"github.com/codecrafters-io/http-server-starter-go/app/router"
)

var (
	echoMatcher = router.NewRegexMatcher(`/echo/(.+)`)

	notFoundHandler  = handler.NewNotFound()
	homeHandler      = handler.NewHome()
	echoHandler      = handler.NewEcho(echoMatcher)
	userAgentHandler = handler.NewUserAgent()

	r = router.NewRouter()
)

func init() {
	r.
		NotFoundHandler(notFoundHandler).
		Register(router.NewExactMatcher("/"), homeHandler).
		Register(echoMatcher, echoHandler).
		Register(router.NewExactMatcher("/user-agent"), userAgentHandler)
}

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer l.Close()

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	request := request.NewHTTPRequest()
	request.Read(conn)

	response, err := r.Handle(request)

	if err != nil {
		fmt.Println("Error handling request: ", err.Error())
	}

	err = response.Write(conn)

	if err != nil {
		fmt.Println("Error writing response: ", err.Error())
		os.Exit(1)
	}
}
