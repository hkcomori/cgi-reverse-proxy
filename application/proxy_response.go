package application

import (
	"io"
	"net/http"
	"strings"

	"groxy/lib"
)

type ProxyResponse struct {
	*http.Response

	// Body represents the response body.
	//
	// The response body is already read from http.Response
	// when the instance is created.
	Body []byte
}

func NewProxyResponse(res *http.Response) (*ProxyResponse, error) {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return &ProxyResponse{
		Response: res,
		Body: body,
	}, nil
}

func NewProxyResponseFromRequest(c *http.Client, req *http.Request) (*ProxyResponse, error) {
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return NewProxyResponse(res)
}

func (pr *ProxyResponse) ToString() string {
	return strings.Join(lib.AppendAll(
		[]string{pr.createStatusLine()},
		pr.createHeaderLines(),
		pr.createBodyLines(),
	), "\r\n")
}

func (pr *ProxyResponse) createStatusLine() string {
	return strings.Join([]string{
		pr.Proto,
		pr.Status,
	}, " ")
}

func (pr *ProxyResponse) createHeaderLines() []string {
	headers := []string{}
	for key, values := range pr.Header {
		headers = append(headers, key + ": " + strings.Join(values, ", "))
	}
	return headers
}

func (pr *ProxyResponse) createBodyLines() []string {
	if len(pr.Body) == 0 {
		return []string{}
	} else {
		return []string{"", string(pr.Body)}
	}
}
