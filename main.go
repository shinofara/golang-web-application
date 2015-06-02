package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/fcgi"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}

func main() {
	goji.Get("/hello/:name", hello)

	listener, _ := net.Listen("tcp", ":9000")
	fcgi.Serve(listener, goji.DefaultMux)
}
