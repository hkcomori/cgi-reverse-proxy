package application

import (
	"fmt"
	"net/http"
	"os"
)

type ProxyRequest http.Request

func NewProxyRequestFromCGI(env *Environ, file *os.File) *ProxyRequest {
	cgi, err := NewCgiRequest(env, file)
	if err != nil {
		fmt.Errorf("Invalid CGI request: %s", err)
	}

	req, err := cgi.ToHttpRequest()
	if err != nil {
		fmt.Errorf("Parse request failed: %s", err)
	}

	reqProxy := ProxyRequest(*req)
	return &reqProxy
}
