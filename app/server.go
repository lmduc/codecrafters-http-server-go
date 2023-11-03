package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/codecrafters-io/http-server-starter-go/app/handler"
	"github.com/codecrafters-io/http-server-starter-go/app/request"
	"github.com/codecrafters-io/http-server-starter-go/app/router"
)

var (
	r = router.NewRouter()
)

func prepareRouter(directory string) {
	echoMatcher := router.NewRegexPathMatcher("GET", `/echo/(.+)`)
	getFileMatcher := router.NewRegexPathMatcher("GET", `/files/(.+)`)
	postFileMatcher := router.NewRegexPathMatcher("POST", `/files/(.+)`)

	notFoundHandler := handler.NewNotFound()
	homeHandler := handler.NewHome()
	echoHandler := handler.NewEcho(echoMatcher)
	userAgentHandler := handler.NewUserAgent()
	getFileHandler := handler.NewGetFile(directory, getFileMatcher)
	postFileHandler := handler.NewPostFile(directory, postFileMatcher)

	r.
		NotFoundHandler(notFoundHandler).
		Register(router.NewExactPathMatcher("/"), homeHandler).
		Register(echoMatcher, echoHandler).
		Register(router.NewExactPathMatcher("/user-agent"), userAgentHandler).
		Register(getFileMatcher, getFileHandler).
		Register(postFileMatcher, postFileHandler)
}

func main() {
	fileDirectory := flag.String("directory", "", "Path to the directory")
	flag.Parse()

	prepareRouter(*fileDirectory)

	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		go func() {
			defer conn.Close()

			if err != nil {
				fmt.Println("Error accepting connection: ", err.Error())
				os.Exit(1)
			}

			request := request.NewHTTPRequest()
			if err := request.Read(conn); err != nil {
				fmt.Println("Error when reading request: ", err.Error())
				os.Exit(1)
			}

			response, err := r.Handle(request)

			if err != nil {
				fmt.Println("Error handling request: ", err.Error())
			}

			err = response.Write(conn)

			if err != nil {
				fmt.Println("Error writing response: ", err.Error())
				os.Exit(1)
			}
		}()
	}
}
