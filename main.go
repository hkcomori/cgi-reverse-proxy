package main

import (
	"groxy/application"
	"log"
	"net/http/cgi"
)

func main() {
	cfg, err := application.NewConfig()
	if err != nil {
		log.Printf("Invalid config: %s", err)
	}

	proxy, err := application.NewProxy(cfg)
	if err != nil {
		log.Printf("Cannot open proxy: %s", err)
	}

	cgi.Serve(proxy.GetHandler())
}
