package handler

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/http-server-starter-go/app/port"
	"github.com/codecrafters-io/http-server-starter-go/app/response"
	"github.com/codecrafters-io/http-server-starter-go/app/router"
)

type File struct {
	directory string
	matcher   *router.RegexMatcher
}

func (f *File) exists(path string) bool {
	return true
}

func (f *File) read(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

func (f *File) Handle(r port.Request) (port.Response, error) {
	fileName := f.matcher.FindMatches(r.Path())[1]
	fmt.Println("file name: ", fileName)
	filePath := fmt.Sprintf("%s%s", f.directory, fileName)
	fmt.Println("file path: ", filePath)

	if !f.exists(filePath) {
		return response.NewResponse("").StatusCode(404), nil
	}

	content, err := f.read(filePath)
	fmt.Println("content: ", string(content))
	if err != nil {
		return nil, err
	}

	resp := response.NewOctetStream().StatusCode(200).Body(content)
	return resp, nil
}

func NewFile(directory string, matcher *router.RegexMatcher) *File {
	return &File{directory, matcher}
}
