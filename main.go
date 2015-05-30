package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/fcgi"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World %s!", r.URL.Path[1:])
}

func main() {
	l, err := net.Listen("tcp", ":9000")
	if err != nil {
		return
	}
	http.HandleFunc("/", handler)
	fcgi.Serve(l, nil)
}
