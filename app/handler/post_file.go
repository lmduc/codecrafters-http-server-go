package handler

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/http-server-starter-go/app/port"
	"github.com/codecrafters-io/http-server-starter-go/app/response"
	"github.com/codecrafters-io/http-server-starter-go/app/router"
)

type PostFile struct {
	directory string
	matcher   *router.RegexPathMatcher
}

func (f *PostFile) write(filePath string, content []byte) error {
	return os.WriteFile(filePath, content, 0644)
}

func (f *PostFile) Handle(r port.Request) (port.Response, error) {
	fmt.Println("inside post")
	fileName := f.matcher.FindMatches(r.Path())[1]
	filePath := fmt.Sprintf("%s%s", f.directory, fileName)

	fmt.Println("body: ", r.Body())
	err := f.write(filePath, r.Body())
	if err != nil {
		return nil, err
	}

	resp := response.NewResponse("").StatusCode(201)
	return resp, nil
}

func NewPostFile(directory string, matcher *router.RegexPathMatcher) *PostFile {
	return &PostFile{directory, matcher}
}
