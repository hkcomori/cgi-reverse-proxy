package application

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"strings"
)

type CgiRequest struct {
	requestLine *HttpRequestLine
	headers *HttpHeader
	body *os.File
}

func NewCgiRequest(env *Environ, file *os.File) (*CgiRequest, error) {
	requestLine, err := NewHttpRequestLineFromCgi(env)
	if err != nil {
		return nil, err
	}

	headers, err := NewHttpHeaderFromCgi(env)
	if err != nil {
		return nil, err
	}

	return &CgiRequest{
		requestLine: requestLine,
		headers: headers,
		body: file,
	}, nil
}

func (r *CgiRequest) ToHttpRequest() (*http.Request, error) {
	req, err := http.ReadRequest(r.toReader())
	return req, err
}

func (r *CgiRequest) toReader() *bufio.Reader {
	beforeBody := strings.NewReader(r.stringBeforeBody())
	ioReader := io.MultiReader(beforeBody, r.body)
	return bufio.NewReader(ioReader)
}

func (r *CgiRequest) stringBeforeBody() string {
	return strings.Join([]string{
		r.requestLine.ToString(),
		r.headers.ToString(),
		"",
	}, "\r\n")
}
