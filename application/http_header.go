package application

import (
	"net/http"
	"strings"

	"groxy/lib"
)

type HttpHeader http.Header

func NewHttpHeaderFromCgi(env *Environ) (*HttpHeader, error) {
	hasPrefix := func (k string, v string) bool { return strings.HasPrefix(k, "HTTP_") }
	httpEnv := Environ(env.Filter(hasPrefix).All())
	header := make(HttpHeader)
	for key, valueString := range httpEnv {
		trimmedKey := strings.TrimPrefix(key, "HTTP_")
		headerKey := lib.UpperSnakeCaseToTrainCase(trimmedKey)
		values := strings.Split(valueString, ",")
		header[headerKey] = values
	}
	return &header, nil

}

func (h *HttpHeader) ToString() string {
	headerLines := []string{}
	for key, values := range *h {
		valueString := strings.Join(values, ",")
		headerLines = append(headerLines, key + ": " + valueString)
	}
	return strings.Join(headerLines, "\r\n")
}
