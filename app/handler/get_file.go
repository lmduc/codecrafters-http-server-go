package handler

import (
	"errors"
	"fmt"
	"os"

	"github.com/codecrafters-io/http-server-starter-go/app/port"
	"github.com/codecrafters-io/http-server-starter-go/app/response"
	"github.com/codecrafters-io/http-server-starter-go/app/router"
)

type GetFile struct {
	directory string
	matcher   *router.RegexMatcher
}

func (f *GetFile) exists(filePath string) bool {
	_, err := os.Stat(filePath)
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func (f *GetFile) read(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

func (f *GetFile) Handle(r port.Request) (port.Response, error) {
	fileName := f.matcher.FindMatches(r.Path())[1]
	filePath := fmt.Sprintf("%s%s", f.directory, fileName)

	if !f.exists(filePath) {
		return notFoundResponse(), nil
	}

	content, err := f.read(filePath)
	fmt.Println("content: ", string(content))
	if err != nil {
		return nil, err
	}

	resp := response.NewOctetStream().StatusCode(200).Body(content)
	return resp, nil
}

func NewGetFile(directory string, matcher *router.RegexMatcher) *GetFile {
	return &GetFile{directory, matcher}
}
