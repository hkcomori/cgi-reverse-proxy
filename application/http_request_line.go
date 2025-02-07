package application

import (
	"errors"
	"strings"
)

type HttpRequestLine struct {
	method string
	path string
	proto string
}

func NewHttpRequestLineFromCgi(env *Environ) (*HttpRequestLine, error) {
	if !env.HasAll("REQUEST_METHOD") {
		return nil, errors.New("Request method cannot found")
	}

	if !env.HasAll("SERVER_PROTOCOL") {
		return nil, errors.New("Request method cannot found")
	}

	var path string
	if env.HasAll("PATH_INFO", "QUERY_STRING") {
		query := ""
		if len((*env)["QUERY_STRING"]) > 0 {
			query = "?" + (*env)["QUERY_STRING"]
		}
		path = (*env)["PATH_INFO"] + query
	} else if env.HasAll("SCRIPT_URL", "QUERY_STRING") {
		query := ""
		if len((*env)["QUERY_STRING"]) > 0 {
			query = "?" + (*env)["QUERY_STRING"]
		}
		path = (*env)["SCRIPT_URL"] + query
	} else if env.HasAll("REQUEST_URI") {
		path = (*env)["REQUEST_URI"]
	} else {
		return nil, errors.New("Request path cannot found")
	}

	return &HttpRequestLine{
		method: (*env)["REQUEST_METHOD"],
		path: path,
		proto: (*env)["SERVER_PROTOCOL"],
	}, nil
}

func (r *HttpRequestLine) ToString() string {
	return strings.Join([]string{r.method, r.path, r.proto}, " ")
}
