package main

import (
	"fmt"
	"net"
	"os"
	"regexp"
)

var (
	echoPathRegex = regexp.MustCompile(`/echo/(.+)`)
)

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

	request := Request{}
	request.Read(conn)

	response := Response{}

	if echoPathRegex.MatchString(request.Path()) {
		text := echoPathRegex.FindString(request.Path())
		response.StatusCode(200).ContentType("text/plain").Body([]byte(text))
	} else if request.Path() == "/" {
		response.StatusCode(200)
	} else {
		response.StatusCode(404)
	}

	response.Write(conn)

	if err != nil {
		fmt.Println("Error writing response: ", err.Error())
		os.Exit(1)
	}
}
