package application

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
)

type Proxy struct {
	scheme string
	address string
    timeoutSeconds int
}

func NewProxy(cfg *Config) (*Proxy, error) {
	parsedUrl, err := url.Parse(cfg.UpstreamAddress)
	if err != nil {
		return nil, errors.New("Cannot parse upstream url: " + cfg.UpstreamAddress)
	}

	var scheme string
	var address string
	switch parsedUrl.Scheme {
	case "unix":
		scheme = "unix"
		address = parsedUrl.Path
	case "http", "tcp":
		scheme = "tcp"
		address = parsedUrl.Host + parsedUrl.Path
	default:
		return nil, errors.New("Unsupported scheme: " + parsedUrl.Scheme)
	}

	return &Proxy{
		scheme: scheme,
		address: address,
        timeoutSeconds: cfg.UpstreamTimeout,
	}, nil
}

func (p *Proxy) RewriteHeader(req *http.Request) *http.Request {
	// header := req.Header
	return req
}

func (p *Proxy) Send(req *http.Request) (*ProxyResponse, error) {
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(network, addr string) (net.Conn, error) {
				return net.Dial(p.scheme, p.address)
			},
		},
        Timeout: time.Duration(p.timeoutSeconds) * time.Second,
	}

	return NewProxyResponseFromRequest(client, req)
}

func (p *Proxy) GetHandler() http.HandlerFunc {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		req := p.RewriteHeader(r)

		res, err := p.Send(req)
		if err != nil {
			log.Println("Proxy failed: ", err)
		}

		for key, values := range res.Header {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}

		w.WriteHeader(res.StatusCode)

		w.Write(res.Body)
	})
}
